package rvsym

import (
	"encoding/binary"
	"fmt"

	"github.com/deadsy/rvda"
	"github.com/zyedidia/rvsym/bits"
	"github.com/zyedidia/rvsym/pkg/smt"
	"github.com/zyedidia/rvsym/rvc"
)

type Machine struct {
	mstate

	Status Status
}

type cache struct {
	base uint32
	data []byte
}

type mstate struct {
	pc      int32
	regs    []smt.Int32
	mem     *Memory
	symvals []SymVal

	sys *SysState

	icache cache
}

type SymVal struct {
	val  smt.Int32
	name string
}

func (m *Machine) markSym(val smt.Int32, name string) {
	m.symvals = append(m.symvals, SymVal{val, name})
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

func (m *Machine) exit(stat ExitStatus) {
	m.Status.Exit = stat
}

func NewMachine(pc int32, mem *Memory) *Machine {
	return &Machine{
		mstate: mstate{
			pc:   pc,
			regs: make([]smt.Int32, 32),
			mem:  mem,
			sys:  NewSysState(0x100000),
		},
	}
}

func (m *Machine) WriteReg(reg uint32, val smt.Int32) {
	if reg == 0 {
		return
	}
	m.regs[reg] = val
}

func (m *Machine) Reg(reg uint32) smt.Int32 {
	return m.regs[reg]
}

func (m *Machine) RegConc(reg uint32) (int32, bool) {
	r := m.regs[reg]
	if r.Concrete() {
		return r.C, true
	}
	return 0, false
}

var isa *rvda.ISA

func init() {
	isa, _ = rvda.New(32, rvda.RV32gc)
}

func (m *Machine) fetchC(s *smt.Solver) (uint32, bool, error) {
	if uint32(m.pc) >= m.icache.base && uint32(m.pc)+3 < m.icache.base+uint32(len(m.icache.data)) {
		// pc in icache
		insn := binary.LittleEndian.Uint16(m.icache.data[uint32(m.pc)-m.icache.base:])
		decoded, compressed, illegal := rvc.Decompress(uint32(insn))

		if illegal {
			return decoded, compressed, fmt.Errorf("illegal instruction")
		} else if !compressed {
			insn := binary.LittleEndian.Uint32(m.icache.data[uint32(m.pc)-m.icache.base:])
			return insn, false, nil
		}
		return decoded, true, nil
	}

	// not in cache -- read 16-bit value and following 16-bit if not compressed
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

func (m *Machine) fetch(s *smt.Solver) (uint32, bool, error) {
	if uint32(m.pc) >= m.icache.base && uint32(m.pc)+3 < m.icache.base+uint32(len(m.icache.data)) {
		// pc in icache
		insn := binary.LittleEndian.Uint32(m.icache.data[uint32(m.pc)-m.icache.base:])
		return insn, false, nil
	}

	word, ok := m.mem.ReadWord(smt.Int32{C: m.pc}, s)
	if !ok {
		return 0, false, fmt.Errorf("program counter out of bounds")
	} else if !word.Concrete() {
		return 0, false, fmt.Errorf("cannot execute symbolic instruction")
	}
	return uint32(word.C), false, nil
}

func (m *Machine) Exec(s *smt.Solver) (isz int32) {
	insn, compressed, err := m.fetch(s)

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
		// a0
		m.ecall(10, symcalls, s)
		return
	case InsnEcall:
		// a7
		m.ecall(17, syscalls, s)
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

	return isz
}

func (m *Machine) ecall(reg uint32, ecalls map[int]EcallFn, s *smt.Solver) {
	if num, ok := m.RegConc(reg); ok {
		if fn, ok := ecalls[int(num)]; ok {
			fn(m, s)
		} else {
			m.err(fmt.Errorf("invalid env call %d\n", num))
		}
	} else {
		m.err(fmt.Errorf("env call number is symbolic"))
	}
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
	pc, ok := m.RegConc(rs1(insn))
	if !ok {
		m.err(fmt.Errorf("jalr target is symbolic"))
		return
	}

	m.WriteReg(rd(insn), smt.Int32{C: m.pc + isz})
	m.br(pc+imm, smt.Bool{C: true})
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
		rdval, valid = m.mem.ReadWord(addr, s)
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
		m.mem.WriteWord(addr, stval, s)
	}
}

func (m *Machine) atomic(insn uint32, s *smt.Solver) {
	m.err(fmt.Errorf("unimplemented: atomic instruction"))
}

func (m *Machine) rdbytes(ptr uint32, p []byte, s *smt.Solver) error {
	ok := m.mem.ReadBytes(ptr, p, s)
	if !ok {
		return fmt.Errorf("invalid attempt to read %d bytes", len(p))
	}
	return nil
}

func (m *Machine) rdstr(ptr uint32, s *smt.Solver) ([]byte, error) {
	var buf [1]byte
	var result = make([]byte, 0)
	for {
		ok := m.mem.ReadBytes(ptr, buf[:], s)
		if !ok {
			return nil, fmt.Errorf("invalid attempt to read string")
		}
		if buf[0] == 0 {
			break
		}
		result = append(result, buf[0])
		ptr++
	}
	return result, nil
}

func (m *Machine) wrsym(addr uint32, length uint32, name string, s *smt.Solver) {
	i := addr
	for i < addr+length {
		i32 := s.AnyInt32()
		if i%4 == 0 && length-i >= 4 {
			m.markSym(i32, fmt.Sprintf("%s[%d:%d]", name, i+3-addr, i-addr))
			m.mem.WriteWord(smt.Int32{C: int32(i)}, i32, s)
			i += 4
		} else {
			s.Assert(i32.And(smt.Int32{C: ^0x0ff}, s).Eqz(s))
			m.markSym(i32, fmt.Sprintf("%s[%d]", name, i-addr))
			m.mem.Write8(smt.Int32{C: int32(i)}, i32, s)
			i++
		}
	}
}

func concretize(val smt.Int32, s *smt.Solver) (int32, bool) {
	if val.Concrete() {
		return val.C, true
	}

	res := s.Check(true)
	if res != smt.Sat {
		return 0, false
	}

	model := s.Model()
	cval := model.Eval(val)
	s.Assert(val.Eqb(smt.Int32{C: cval}, s))

	logger.Printf("concretizing value -> %x\n", cval)
	return cval, true
}
