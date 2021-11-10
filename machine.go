package rvsym

import (
	"errors"
	"fmt"

	"github.com/zyedidia/go-z3/st"
	"github.com/zyedidia/go-z3/z3"
)

type Machine struct {
	pc    int32
	regs  []st.Int32
	conds []z3.Bool
	mem   Memory

	done bool

	ctx    *z3.Context
	solver *z3.Solver
}

func NewMachine(ctx *z3.Context, pc int32, mem Memory) *Machine {
	regs := make([]st.Int32, 32)
	for i := range regs {
		regs[i] = st.Int32{C: 0}
	}

	return &Machine{
		pc:    pc,
		regs:  regs,
		conds: make([]z3.Bool, 0),
		mem:   mem,
		ctx:   ctx,
	}
}

func (m *Machine) Copy() *Machine {
	regs := make([]st.Int32, len(m.regs))
	conds := make([]z3.Bool, len(m.conds))
	mem := make(map[uint32]st.Int32)

	copy(regs, m.regs)
	copy(conds, m.conds)

	for k, v := range m.mem {
		mem[k] = v
	}

	return &Machine{
		pc:    m.pc,
		regs:  regs,
		conds: conds,
		mem:   mem,
		ctx:   m.ctx,
	}
}

func (m *Machine) AddCond(cond z3.Bool) {
	m.conds = append(m.conds, cond)
}

func (m *Machine) WriteReg(reg uint32, val st.Int32) {
	if reg == 0 {
		return
	}
	m.regs[reg] = val
}

var ErrUnsat = errors.New("unsatisfiable formula")

func (m *Machine) Solver() (*z3.Solver, error) {
	if m.solver != nil {
		return m.solver, nil
	}

	s := z3.NewSolver(m.ctx)
	for _, c := range m.conds {
		s.Assert(c)
	}
	sat, err := s.Check()
	if err != nil {
		return nil, err
	}
	if !sat {
		return nil, ErrUnsat
	}

	return s, nil
}

type Branch struct {
	pc   int32
	cond st.Bool
}

func (m *Machine) Exec(insn uint32) (br Branch, hasbr bool, ex ExitStatus) {
	switch insn {
	case InsnNop:
		return br, false, ExitNone
	case InsnEcall:
		sysnum := m.regs[10] // a0
		if !sysnum.IsConcrete() {
			// TODO: cleanup
			panic("system call number is symbolic")
		}
		ex := m.symsys(insn, int(sysnum.C))
		if ex != ExitNone {
			ex = m.Exit(ex)
		}
		return br, false, ex
	}

	op := GetBits(insn, 6, 0).Uint32()

	switch op {
	case OpRarith:
		m.rarith(insn)
	case OpIarith:
		m.iarith(insn)
	case OpBranch:
		return m.branch(insn), true, ExitNone
	case OpLui:
		m.lui(insn)
	case OpAuipc:
		m.auipc(insn)
	case OpJal:
		return m.jal(insn), true, ExitNone
	case OpJalr:
		return m.jalr(insn), true, ExitNone
	case OpLoad:
		m.load(insn)
	case OpStore:
		m.store(insn)
	}
	return Branch{}, false, ExitNone
}

func (m *Machine) Exit(ex ExitStatus) ExitStatus {
	m.done = true
	s, err := m.Solver()
	m.solver = s
	if err == ErrUnsat {
		return ExitQuiet
	} else if err != nil {
		return ExitUnsure
	}
	return ex
}

func (m *Machine) symsys(insn uint32, sysnum int) ExitStatus {
	switch sysnum {
	case SysSymbolicRegs:
		for i := range m.regs[1:] {
			m.regs[i+1] = st.AnyInt32(m.ctx, fmt.Sprintf("x%d", i+1))
		}
	case SysSymbolicReg:
		sysarg := m.regs[11] // a1
		if !sysarg.IsConcrete() {
			panic("required symcall argument is symbolic")
		}
		m.regs[sysarg.C] = st.AnyInt32(m.ctx, fmt.Sprintf("x%d", sysarg.C))
	case SysFail:
		return ExitFail
	case SysExit:
		return ExitNormal
	case SysQuietExit:
		return ExitQuiet
	case SysMarkNBytes:
		ptr := m.regs[11]    // a1
		nbytes := m.regs[12] // a2

		if !ptr.IsConcrete() {
			panic("mark address is symbolic")
		}
		if !nbytes.IsConcrete() {
			panic("mark size is symbolic")
		}

		fmt.Printf("INFO: marking %d bytes at 0x%x\n", nbytes.C, ptr.C)

		for i := int32(0); i < nbytes.C/4; i++ {
			idx := i * 4
			m.mem.Write32(uint32(ptr.C+idx), st.AnyInt32(m.ctx, fmt.Sprintf("0x%x[%d]", ptr, idx)))
		}
		left := nbytes.C % 4
		for i := int32(0); i < left; i++ {
			idx := nbytes.C/4 + i
			m.mem.Write8(uint32(ptr.C+idx), st.AnyInt32(m.ctx, fmt.Sprintf("0x%x[%d]", ptr, idx)))
		}
	}
	return ExitNone
}

func (m *Machine) concretize(val st.Int32) int32 {
	if val.IsConcrete() {
		return val.C
	}

	fmt.Println("INFO: concretizing")

	s, err := m.Solver()
	if err != nil {
		m.Exit(ExitConcretize)
	}
	model := s.Model()

	concrete := val.Eval(model)
	m.AddCond(val.Eq(st.Int32{C: concrete}).S)
	return concrete
}

func extractRegs(insn uint32) (rd, rs1, rs2 uint32) {
	rd = GetBits(insn, 11, 7).Uint32()
	rs1 = GetBits(insn, 19, 15).Uint32()
	rs2 = GetBits(insn, 24, 20).Uint32()
	return rd, rs1, rs2
}

func (m *Machine) rarith(insn uint32) {
	aluop := GetBits(insn, 14, 12).Uint32() // funct3
	modify := GetBits(insn, 30, 30).Uint32() != 0
	rd, rs1, rs2 := extractRegs(insn)

	m.WriteReg(rd, alu(m.regs[rs1], m.regs[rs2], aluop, modify, modify))
}

func (m *Machine) iarith(insn uint32) {
	aluop := GetBits(insn, 14, 12).Uint32()
	modify := GetBits(insn, 30, 30).Uint32() != 0
	imm := st.Int32{C: int32(extractImm(insn, ImmI))}
	rd, rs1, _ := extractRegs(insn)

	m.WriteReg(rd, alu(m.regs[rs1], imm, aluop, false, modify))
}

func (m *Machine) branch(insn uint32) Branch {
	funct3 := GetBits(insn, 14, 12).Uint32()
	_, rs1, rs2 := extractRegs(insn)
	imm := int32(extractImm(insn, ImmB))

	var aluop uint32
	switch funct3 {
	case 0b000, 0b001:
		aluop = AluXor
	case 0b100, 0b101:
		aluop = AluSlt
	case 0b110, 0b111:
		aluop = AluSltu
	}

	result := alu(m.regs[rs1], m.regs[rs2], aluop, false, false)

	var cond st.Bool
	switch funct3 {
	case 0b000, 0b101, 0b111:
		// branch if alu result is zero
		cond = result.Eq(st.Int32{C: 0})
	case 0b001, 0b100, 0b110:
		// branch if alu result is non-zero
		cond = result.NE(st.Int32{C: 0})
	}

	return Branch{
		pc:   m.pc + int32(imm),
		cond: cond,
	}
}

func (m *Machine) jal(insn uint32) Branch {
	rd, _, _ := extractRegs(insn)
	imm := extractImm(insn, ImmJ)

	pc := m.pc + int32(imm)
	m.WriteReg(rd, st.Int32{C: int32(m.pc + 4)})

	return Branch{
		pc:   pc,
		cond: st.Bool{C: true},
	}
}

func (m *Machine) jalr(insn uint32) Branch {
	rd, rs1, _ := extractRegs(insn)
	imm := extractImm(insn, ImmI)

	pc := m.regs[rs1].Add(st.Int32{C: int32(imm)})

	if !pc.IsConcrete() {
		pc = st.Int32{C: m.concretize(pc)}
	}
	m.WriteReg(rd, st.Int32{C: int32(m.pc + 4)})

	return Branch{
		pc:   int32(pc.C),
		cond: st.Bool{C: true},
	}
}

func (m *Machine) lui(insn uint32) {
	rd, _, _ := extractRegs(insn)
	imm := extractImm(insn, ImmU)

	m.WriteReg(rd, st.Int32{C: int32(imm)})
}

func (m *Machine) auipc(insn uint32) {
	rd, _, _ := extractRegs(insn)
	imm := extractImm(insn, ImmU)

	m.WriteReg(rd, st.Int32{C: m.pc + int32(imm)})
}

func (m *Machine) load(insn uint32) {
	rd, rs1, _ := extractRegs(insn)
	imm := extractImm(insn, ImmI)
	funct3 := GetBits(insn, 14, 12).Uint32()

	rsval := m.regs[rs1]
	if !rsval.IsConcrete() {
		rsval = st.Int32{C: m.concretize(rsval)}
	}
	addr := uint32(rsval.C) + imm

	var rdval st.Int32
	switch funct3 {
	case ExtByte:
		rdval = m.mem.Read8(addr)
	case ExtHalf:
		rdval = m.mem.Read16(addr)
	case ExtWord:
		rdval = m.mem.Read32(addr)
	case ExtByteU:
		rdval = m.mem.Read8u(addr)
	case ExtHalfU:
		rdval = m.mem.Read16u(addr)
	default:
		panic("invalid load instruction")
	}

	m.WriteReg(rd, rdval)
}

func (m *Machine) store(insn uint32) {
	_, rs1, rs2 := extractRegs(insn)
	imm := extractImm(insn, ImmS)
	funct3 := GetBits(insn, 14, 12).Uint32()

	rsval := m.regs[rs1]
	if !rsval.IsConcrete() {
		rsval = st.Int32{C: m.concretize(rsval)}
	}

	addr := uint32(rsval.C) + imm

	stval := m.regs[rs2]

	switch funct3 {
	case ExtByte:
		m.mem.Write8(addr, stval)
	case ExtHalf:
		m.mem.Write16(addr, stval)
	case ExtWord:
		m.mem.Write32(addr, stval)
	default:
		panic("invalid store instruction")
	}
}
