package rvsym

import (
	"bytes"
	"fmt"

	"github.com/zyedidia/rvsym/bits"
	"github.com/zyedidia/rvsym/pkg/smt"
)

type Machine struct {
	pc   int32
	regs []smt.Int32
	mem  *Memory

	Status Status

	symvals []SymVal
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
		pc:   pc,
		regs: make([]smt.Int32, 32),
		mem:  mem,
	}
}

func (m *Machine) WriteReg(reg uint32, val smt.Int32) {
	if reg == 0 {
		// cannot write to register 0
		return
	}
	m.regs[reg] = val
}

func (m *Machine) FetchInsn(s *smt.Solver) (uint32, error) {
	word, ok := m.mem.Read32(smt.Int32{C: m.pc}, s)
	if !ok {
		return 0, fmt.Errorf("program counter out of bounds")
	} else if !word.Concrete() {
		return 0, fmt.Errorf("cannot execute symbolic instruction")
	}
	return uint32(word.C), nil
}

func (m *Machine) Exec(s *smt.Solver) {
	insn, err := m.FetchInsn(s)
	if err != nil {
		m.err(err)
		return
	}

	switch insn {
	case InsnNop:
		return
	case InsnEcall:
		symnum := m.regs[10] // a0
		if !symnum.Concrete() {
			m.err(fmt.Errorf("symcall number is symbolic"))
		} else {
			m.symcall(insn, int(symnum.C), s)
		}
		return
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
		m.jal(insn)
	case OpJalr:
		m.jalr(insn, s)
	case OpLoad:
		m.load(insn, s)
	case OpStore:
		m.store(insn, s)
	}
}

func (m *Machine) symcall(insn uint32, symnum int, s *smt.Solver) {
	mustconc := func(val smt.Int32, errmsg string) (int32, bool) {
		if !val.Concrete() {
			m.err(fmt.Errorf(errmsg))
			return 0, false
		}
		return val.C, true
	}

	switch symnum {
	case SymPrint:
		arg := m.regs[11] // a1
		if !arg.Concrete() {
			fmt.Println(arg.S)
		} else {
			fmt.Println(arg.C)
		}
	case SymFail:
		m.exit(ExitFail)
	case SymExit:
		m.exit(ExitNormal)
	case SymQuietExit:
		m.exit(ExitQuiet)
	case SymMarkArray:
		var ptr, nbytes int32
		var ok bool
		if ptr, ok = mustconc(m.regs[11], "array address is symbolic"); !ok {
			return
		}
		if nbytes, ok = mustconc(m.regs[12], "array size is symbolic"); !ok {
			return
		}
		m.mem.AddArray(s, uint32(ptr/4), uint32((nbytes+3)/4))
	case SymMarkBytes:
		var ptr, nbytes, nameptr int32
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

		markUnaligned := func(base, idx, length int32) int32 {
			for i := idx; i < idx+length; i++ {
				i32 := s.AnyInt32()
				s.Assert(i32.And(smt.Int32{C: ^0x0ff}, s).Eqz(s))
				m.markSym(i32, fmt.Sprintf("%s[%d]", name, uint32(i-base)))
				m.mem.Write8(smt.Int32{C: i}, i32, s)
			}
			return length
		}
		markAligned := func(base, idx, length int32) int32 {
			for i := idx; i < idx+length; i += 4 {
				i32 := s.AnyInt32()
				m.markSym(i32, fmt.Sprintf("%s[%d:%d]", name, uint32(i+3-base), uint32(i-base)))
				m.mem.Write32(smt.Int32{C: i}, i32, s)
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

func (m *Machine) readString(ptr int32, size int32, s *smt.Solver) (string, error) {
	buf := &bytes.Buffer{}
	for i := int32(0); ; i++ {
		if b, ok := m.mem.Read8u(smt.Int32{C: ptr + i}, s); !ok {
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

func (m *Machine) jal(insn uint32) {
	imm := extractImm(insn, ImmTypeJ)
	pc := m.pc + imm
	m.WriteReg(rd(insn), smt.Int32{C: m.pc + 4})
	m.br(pc, smt.Bool{C: true})
}

func (m *Machine) jalr(insn uint32, s *smt.Solver) {
	imm := extractImm(insn, ImmTypeI)
	pc := m.regs[rs1(insn)].Add(smt.Int32{C: imm}, s)
	if !pc.Concrete() {
		m.err(fmt.Errorf("jalr target is symbolic"))
		return
	}

	m.WriteReg(rd(insn), smt.Int32{C: m.pc + 4})
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
		m.mem.Write8(addr, stval, s)
	case ExtHalf:
		m.mem.Write16(addr, stval, s)
	case ExtWord:
		m.mem.Write32(addr, stval, s)
	}
}

type SymVal struct {
	val  smt.Int32
	name string
}

type Checkpoint struct {
	pc   int32
	regs []smt.Int32
	mem  *Memory
	vals []SymVal

	cond smt.Bool
}

func Restore(cp *Checkpoint, s *smt.Solver) *Machine {
	s.Assert(cp.cond)
	return &Machine{
		pc:      cp.pc,
		regs:    cp.regs,
		mem:     cp.mem,
		symvals: cp.vals,
	}
}

func (m *Machine) Checkpoint(cond smt.Bool) *Checkpoint {
	cp := &Checkpoint{
		regs: make([]smt.Int32, len(m.regs)),
		mem:  NewMemory(m.mem),
		vals: make([]SymVal, len(m.symvals)),
		pc:   m.pc,
		cond: cond,
	}

	copy(cp.regs, m.regs)
	copy(cp.vals, m.symvals)
	// duplicate memory because current m.mem must become read-only
	m.mem = NewMemory(m.mem)

	return cp
}

func (m *Machine) AddCond(cond smt.Bool, checksat bool, s *smt.Solver) {
	s.Assert(cond)

	if checksat {
		res := s.Check()
		if res == smt.Unsat {
			m.exit(ExitUnsat)
		} else if res == smt.Unknown {
			m.err(fmt.Errorf("smt solver returned unknown"))
		}
	}
}

func (m *Machine) TestCase(s *smt.Solver) (TestCase, bool) {
	res := s.Check()
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
	return TestCase{
		Assignments: vars,
		Pc:          m.pc,
		Exit:        m.Status.Exit,
		Err:         m.Status.Err,
	}, true
}
