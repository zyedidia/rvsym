package rvsym

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/deadsy/rvda"
	"github.com/zyedidia/rvsym/bits"
	"github.com/zyedidia/rvsym/pkg/smt"
	"github.com/zyedidia/rvsym/rvc"
)

type Machine struct {
	pc   int32
	brk  int32
	regs []smt.Int32
	mem  *Memory
	time time.Duration

	symvals []SymVal
	outputs []Output
	traces  []Trace

	Status Status
}

type Branch struct {
	pc   int32
	cond smt.Bool
}

type Status struct {
	Err   error
	HasBr bool
	Br    Branch
	Exit  ExitStatus
}

func (m *Machine) br(pc int32, cond smt.Bool) {
	m.Status.HasBr = true
	m.Status.Br.pc = pc
	m.Status.Br.cond = cond
}

func (m *Machine) clearbr() {
	m.Status.HasBr = false
}

func (m *Machine) err(err error) {
	m.Status.Err = err
	m.Status.Exit = ExitFail
}

func (m *Machine) exit(exit ExitStatus) {
	m.Status.Exit = exit
}

func NewMachine(pc int32, mem *Memory) *Machine {
	return &Machine{
		pc:     pc,
		regs:   make([]smt.Int32, 32),
		mem:    mem,
		traces: []Trace{make(Trace)},
	}
}

func (m *Machine) WriteReg(reg uint32, val smt.Int32) {
	if reg == 0 {
		// cannot write to register 0
		return
	}
	m.regs[reg] = val
}

func (m *Machine) FetchInsnNoCompression(s *smt.Solver) (uint32, bool, error) {
	word, ok := m.mem.Read32(smt.Int32{C: m.pc}, s)
	if !ok {
		return 0, false, fmt.Errorf("program counter out of bounds")
	} else if !word.Concrete() {
		return 0, false, fmt.Errorf("cannot execute symbolic instruction")
	}
	// fmt.Println(isa.Disassemble(uint(m.pc), uint(word.C)))
	return uint32(word.C), false, nil
}

func (m *Machine) FetchInsnCompression(s *smt.Solver) (uint32, bool, error) {
	lword, ok := m.mem.Read16u(smt.Int32{C: m.pc}, s)
	if !ok {
		return 0, false, fmt.Errorf("program counter out of bounds: %x", m.pc)
	} else if !lword.Concrete() {
		return 0, false, fmt.Errorf("cannot execute symbolic instruction")
	}

	decoded, compressed, illegal := rvc.Decompress(uint32(lword.C))

	if illegal {
		return decoded, compressed, fmt.Errorf("illegal instruction")
	} else if !compressed {
		uword, ok := m.mem.Read16u(smt.Int32{C: m.pc + 2}, s)
		if !ok {
			return 0, false, fmt.Errorf("program counter out of bounds: %x", m.pc)
		} else if !uword.Concrete() {
			return 0, false, fmt.Errorf("cannot execute symbolic instruction")
		}
		decoded = (uint32(uword.C) << 16) | uint32(lword.C)
	}

	return decoded, compressed, nil
}

var isa *rvda.ISA

func init() {
	isa, _ = rvda.New(32, rvda.RV32gc)
}

func (m *Machine) Exec(s *smt.Solver) (isz int32) {
	insn, compressed, err := m.FetchInsnNoCompression(s)

	if compressed {
		isz = 2
	} else {
		isz = 4
	}

	if err != nil {
		m.err(err)
		return
	}

	switch insn {
	case InsnNop:
		return
	case InsnEbreak:
		symnum := m.regs[10] // a0
		if !symnum.Concrete() {
			m.err(fmt.Errorf("symcall number is symbolic"))
		} else {
			m.symcall(int(symnum.C), s)
		}
		return
	case InsnEcall:
		sysnum := m.regs[17] // a7
		if !sysnum.Concrete() {
			m.err(fmt.Errorf("syscall number is symbolic"))
		} else {
			m.syscall(int(sysnum.C), s)
		}
	}

	op := bits.Get(insn, 6, 0)

	switch op {
	case OpRarith:
		m.rarith(insn, s)
	case OpIarith:
		m.iarith(insn, s)
	case OpBranch:
		m.branch(insn, s)
	case OpLui:
		m.lui(insn)
	case OpAuipc:
		m.auipc(insn)
	case OpJal:
		m.jal(insn, isz)
	case OpJalr:
		m.jalr(insn, isz, s)
	case OpLoad:
		m.load(insn, s)
	case OpStore:
		m.store(insn, s)
	case OpAtomic:
		m.atomic(insn, s)
	}

	return
}

func (m *Machine) syscall(sysnum int, s *smt.Solver) {
	logger.Println("syscall:", syscalls_riscv[sysnum])
	switch sysnum {
	case SysOpen:
		path := m.regs[10].C
		str, _ := m.readString(uint32(path), 100, s)
		fmt.Println("Open", str)
	case SysExit:
		m.exit(ExitQuiet)
	case SysFstat:
		// fd := m.regs[10].C
		buf := m.regs[11].C
		for i := int32(0); i < 112; i += 4 {
			m.mem.Write32(smt.Int32{C: buf + i}, smt.Int32{C: 0}, s)
		}
		m.regs[10] = smt.Int32{C: 0}
	case SysClose:
	case SysBrk:
		if m.regs[10].C != 0 {
			m.brk = m.regs[10].C
		}
		m.regs[10] = smt.Int32{C: m.brk}
	case SysWrite:
		fd := m.regs[10].C
		buf := m.regs[11].C
		count := m.regs[12].C
		if fd == 1 {
			// stdout
			str := make([]byte, count)
			err := m.readBytes(uint32(buf), str, s)
			if err != nil {
				m.err(err)
			} else {
				fmt.Fprint(os.Stdout, string(str))
			}
			m.regs[10] = smt.Int32{C: int32(len(str))}
		} else {
			m.err(fmt.Errorf("invalid write file descriptor: %d", fd))
		}
	default:
		m.err(fmt.Errorf("unhandled system call: %d", sysnum))
	}
}

func (m *Machine) symcall(symnum int, s *smt.Solver) {
	mustconc := func(val smt.Int32, errmsg string) (uint32, bool) {
		if !val.Concrete() {
			m.err(fmt.Errorf(errmsg))
			return 0, false
		}
		return uint32(val.C), true
	}

	switch symnum {
	case SymPrint:
		arg := m.regs[11] // a1
		if !arg.Concrete() {
			fmt.Println(arg.S)
		} else {
			fmt.Println(arg.C)
		}
	case SymElapseNs:
		if arg, ok := mustconc(m.regs[11], "time elapsed is symbolic"); !ok {
			return
		} else {
			m.time += time.Nanosecond * time.Duration(arg)
		}
	case SymSnapshot:
		m.regs[10] = smt.Int32{C: int32(len(m.traces) - 1)}
	case SymTraceReset:
		m.time = time.Duration(0)
		m.traces = append(m.traces, make(Trace))
	case SymSnapshotEq:
		s1 := m.regs[11].C
		s2 := m.regs[12].C
		eq := m.traces[s1].Eq(m.traces[s2], s)
		if eq {
			m.regs[10] = smt.Int32{C: 1}
		} else {
			m.regs[10] = smt.Int32{C: 0}
		}
	case SymFail:
		m.exit(ExitFail)
	case SymExit:
		m.exit(ExitNormal)
	case SymQuietExit:
		m.exit(ExitQuiet)
	case SymMarkArray:
		var ptr, nbytes uint32
		var ok bool
		if ptr, ok = mustconc(m.regs[11], "array address is symbolic"); !ok {
			return
		}
		if nbytes, ok = mustconc(m.regs[12], "array size is symbolic"); !ok {
			return
		}
		m.mem.AddArray(s, ptr/4, (nbytes+3)/4)

		logger.Printf("Marking symbolic array: %d bytes at 0x%x\n", nbytes, ptr)
	case SymMarkOutput:
		var ptr, nbytes, nameptr uint32
		var ok bool
		if ptr, ok = mustconc(m.regs[11], "value address is symbolic"); !ok {
			return
		}
		if nbytes, ok = mustconc(m.regs[12], "value size is symbolic"); !ok {
			return
		}
		if nameptr, ok = mustconc(m.regs[13], "value name address is symbolic"); !ok {
			return
		}

		name, err := m.readString(nameptr, nbytes, s)
		if err != nil {
			m.err(err)
			return
		}

		logger.Printf("Marking output '%s': %d bytes at 0x%x\n", name, nbytes, ptr)

		m.outputs = append(m.outputs, Output{
			base: ptr,
			size: nbytes,
			name: name,
		})
	case SymMarkBytes:
		var ptr, nbytes, nameptr uint32
		var ok bool
		if ptr, ok = mustconc(m.regs[11], "value address is symbolic"); !ok {
			return
		}
		if nbytes, ok = mustconc(m.regs[12], "value size is symbolic"); !ok {
			return
		}
		if nameptr, ok = mustconc(m.regs[13], "value name address is symbolic"); !ok {
			return
		}

		name, err := m.readString(nameptr, nbytes, s)
		if err != nil {
			m.err(err)
			return
		}
		logger.Printf("Marking symbolic value '%s': %d bytes at 0x%x\n", name, nbytes, ptr)

		markUnaligned := func(base, idx, length uint32) uint32 {
			for i := idx; i < idx+length; i++ {
				i32 := s.AnyInt32()
				s.Assert(i32.And(smt.Int32{C: ^0x0ff}, s).Eqz(s))
				m.markSym(i32, fmt.Sprintf("%s[%d]", name, i-base))
				m.mem.Write8(smt.Int32{C: int32(i)}, i32, s)
			}
			return length
		}
		markAligned := func(base, idx, length uint32) uint32 {
			for i := idx; i < idx+length; i += 4 {
				i32 := s.AnyInt32()
				m.markSym(i32, fmt.Sprintf("%s[%d:%d]", name, i+3-base, i-base))
				m.mem.Write32(smt.Int32{C: int32(i)}, i32, s)
			}
			return length
		}

		if ptr%4 == 0 && nbytes%4 == 0 {
			markAligned(ptr, ptr, nbytes)
		} else if ptr%4 == 0 && nbytes > 4 {
			written := markAligned(ptr, ptr, nbytes-(nbytes%4))
			markUnaligned(ptr, ptr+written, nbytes%4)
		} else {
			// TODO: can optimize further when ptr % 4 != 0 but nbytes >= 4
			markUnaligned(ptr, ptr, nbytes)
		}
	}
}

func (m *Machine) markSym(val smt.Int32, name string) {
	m.symvals = append(m.symvals, SymVal{val, name})
}

func (m *Machine) readBytes(ptr uint32, p []byte, s *smt.Solver) error {
	for i := uint32(0); i < uint32(len(p)); i++ {
		if b, ok := m.mem.Read8u(smt.Int32{C: int32(ptr + i)}, s); !ok {
			return fmt.Errorf("out of bounds string")
		} else if !b.Concrete() {
			return fmt.Errorf("symbolic byte in string")
		} else {
			p[i] = byte(b.C)
		}
	}
	return nil
}

func (m *Machine) readString(ptr uint32, size uint32, s *smt.Solver) (string, error) {
	buf := &bytes.Buffer{}
	for i := uint32(0); ; i++ {
		if b, ok := m.mem.Read8u(smt.Int32{C: int32(ptr + i)}, s); !ok {
			return "", fmt.Errorf("out of bounds string")
		} else if !b.Concrete() {
			return "", fmt.Errorf("symbolic byte in string")
		} else if b.C == 0 {
			break
		} else {
			buf.WriteByte(byte(b.C))
		}
	}
	return buf.String(), nil
}

func (m *Machine) rarith(insn uint32, s *smt.Solver) {
	op := (funct7(insn) << 3) | funct3(insn)
	m.WriteReg(rd(insn), m.alu(m.regs[rs1(insn)], m.regs[rs2(insn)], AluOp(op), s))
}

func (m *Machine) iarith(insn uint32, s *smt.Solver) {
	op := funct3(insn)
	var imm int32
	if op == AluSrl {
		imm = int32(shamt(insn))
		op |= funct7(insn) << 3
	} else {
		imm = extractImm(insn, ImmTypeI)
	}
	m.WriteReg(rd(insn), m.alu(m.regs[rs1(insn)], smt.Int32{C: imm}, AluOp(op), s))
}

func (m *Machine) branch(insn uint32, s *smt.Solver) {
	imm := extractImm(insn, ImmTypeB)

	const (
		beq  = 0b000
		bne  = 0b001
		blt  = 0b100
		bge  = 0b101
		bltu = 0b110
		bgeu = 0b111
	)

	r1, r2 := m.regs[rs1(insn)], m.regs[rs2(insn)]

	var cond smt.Bool
	switch funct3(insn) {
	case beq:
		cond = r1.Eqb(r2, s)
	case bne:
		cond = r1.NEqb(r2, s)
	case blt:
		cond = r1.Sltb(r2, s)
	case bge:
		cond = r1.Sgeb(r2, s)
	case bltu:
		cond = r1.Ultb(r2, s)
	case bgeu:
		cond = r1.Ugeb(r2, s)
	}

	m.br(m.pc+imm, cond)
}

func (m *Machine) lui(insn uint32) {
	m.WriteReg(rd(insn), smt.Int32{C: extractImm(insn, ImmTypeU)})
}

func (m *Machine) auipc(insn uint32) {
	m.WriteReg(rd(insn), smt.Int32{C: m.pc + extractImm(insn, ImmTypeU)})
}

func (m *Machine) jal(insn uint32, isz int32) {
	imm := extractImm(insn, ImmTypeJ)
	pc := m.pc + imm
	m.WriteReg(rd(insn), smt.Int32{C: m.pc + isz})
	m.br(pc, smt.Bool{C: true})
}

func (m *Machine) jalr(insn uint32, isz int32, s *smt.Solver) {
	imm := extractImm(insn, ImmTypeI)
	pc := m.regs[rs1(insn)].Add(smt.Int32{C: imm}, s)
	if !pc.Concrete() {
		m.err(fmt.Errorf("jalr target is symbolic"))
		return
	}

	m.WriteReg(rd(insn), smt.Int32{C: m.pc + isz})
	m.br(pc.C, smt.Bool{C: true})
}

func (m *Machine) load(insn uint32, s *smt.Solver) {
	imm := extractImm(insn, ImmTypeI)
	addr := m.regs[rs1(insn)].Add(smt.Int32{C: imm}, s)

	var rdval smt.Int32
	var valid bool
	switch funct3(insn) {
	case ExtByte:
		rdval, valid = m.mem.Read8(addr, s)
	case ExtHalf:
		rdval, valid = m.mem.Read16(addr, s)
	case ExtWord:
		rdval, valid = m.mem.Read32(addr, s)
	case ExtByteU:
		rdval, valid = m.mem.Read8u(addr, s)
	case ExtHalfU:
		rdval, valid = m.mem.Read16u(addr, s)
	}

	if !valid {
		if addr.Concrete() {
			m.err(fmt.Errorf("out of bounds access at 0x%x", addr.C))
		} else {
			m.err(fmt.Errorf("symbolic out of bounds access"))
		}
		return
	}

	m.WriteReg(rd(insn), rdval)
}

func (m *Machine) store(insn uint32, s *smt.Solver) {
	imm := extractImm(insn, ImmTypeS)

	addr := m.regs[rs1(insn)].Add(smt.Int32{C: imm}, s)
	stval := m.regs[rs2(insn)]

	switch funct3(insn) {
	case ExtByte:
		stval = stval.ToInt8(s).ToInt32z(s)
		m.mem.Write8(addr, stval, s)
	case ExtHalf:
		stval = stval.ToInt16(s).ToInt32z(s)
		m.mem.Write16(addr, stval, s)
	case ExtWord:
		m.mem.Write32(addr, stval, s)
	}

	if addr.Concrete() {
		addrc := uint32(addr.C)
		for _, o := range m.outputs {
			if addrc >= o.base && addrc < o.base+o.size {
				logger.Printf("write %v to '%s' at time %v\n", stval, o.name, m.time)
				m.traces[len(m.traces)-1].Add(addrc, m.time, stval)
			}
		}
	}
}

func (m *Machine) atomic(insn uint32, s *smt.Solver) {
	funct5 := funct7(insn) >> 2

	const (
		lr   = 0b00010
		sc   = 0b00011
		swap = 0b00001
		add  = 0b00000
		xor  = 0b00100
		and  = 0b01100
		or   = 0b01000
		min  = 0b10000
		max  = 0b10100
		minu = 0b11000
		maxu = 0b11100
	)

	addr := m.regs[rs1(insn)]

	if funct5 == sc {
		m.mem.Write32(addr, m.regs[rs2(insn)], s)
		return
	}

	rdval, valid := m.mem.Read32(addr, s)
	if !valid {
		m.err(fmt.Errorf("invalid memory read at %v", addr))
		return
	}

	m.WriteReg(rd(insn), rdval)

	if funct5 == lr {
		return
	}

	a := rdval
	b := m.regs[rs2(insn)]
	var result smt.Int32
	switch funct5 {
	case add:
		result = a.Add(b, s)
	case xor:
		result = a.Xor(b, s)
	case and:
		result = a.And(b, s)
	case or:
		result = a.Or(b, s)
	case swap:
		fallthrough
	case min:
		fallthrough
	case max:
		fallthrough
	case minu:
		fallthrough
	case maxu:
		m.err(fmt.Errorf("atomic operation unimplemented"))
		return
	}

	m.mem.Write32(addr, result, s)
}

type Output struct {
	base uint32
	size uint32
	name string
}

type SymVal struct {
	val  smt.Int32
	name string
}

type Checkpoint struct {
	pc     int32
	brk    int32
	regs   []smt.Int32
	mem    *Memory
	vals   []SymVal
	outs   []Output
	traces []Trace

	cond smt.Bool
}

func Restore(cp *Checkpoint, s *smt.Solver) *Machine {
	s.Assert(cp.cond)
	return &Machine{
		pc:      cp.pc,
		brk:     cp.brk,
		regs:    cp.regs,
		mem:     cp.mem,
		symvals: cp.vals,
		outputs: cp.outs,
		traces:  cp.traces,
	}
}

func (m *Machine) Checkpoint(cond smt.Bool) *Checkpoint {
	cp := &Checkpoint{
		regs:   make([]smt.Int32, len(m.regs)),
		mem:    NewMemory(m.mem),
		vals:   make([]SymVal, len(m.symvals)),
		pc:     m.pc,
		brk:    m.brk,
		cond:   cond,
		outs:   make([]Output, len(m.outputs)),
		traces: make([]Trace, len(m.traces)),
	}

	copy(cp.regs, m.regs)
	copy(cp.vals, m.symvals)
	copy(cp.outs, m.outputs)
	copy(cp.traces, m.traces)
	// duplicate memory because current m.mem must become read-only
	m.mem = NewMemory(m.mem)

	return cp
}

func (m *Machine) AddCond(cond smt.Bool, checksat bool, s *smt.Solver) {
	s.Assert(cond)

	if checksat {
		res := s.Check(false)
		if res == smt.Unsat {
			m.exit(ExitUnsat)
		} else if res == smt.Unknown {
			m.err(fmt.Errorf("smt solver returned unknown"))
		}
	}
}

func (m *Machine) TestCase(s *smt.Solver) (TestCase, bool) {
	res := s.Check(true)
	if res != smt.Sat {
		return TestCase{}, false
	}
	model := s.Model()
	vars := make([]Assignment, len(m.symvals))
	for i, v := range m.symvals {
		vars[i] = Assignment{
			Name: v.name,
			Val:  model.Eval(v.val),
		}
	}
	dump := s.String()
	return TestCase{
		Assignments: vars,
		Pc:          m.pc,
		Exit:        m.Status.Exit,
		Err:         m.Status.Err,
		Dump:        dump,
	}, true
}
