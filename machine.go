package minisym

import (
	"fmt"

	"github.com/zyedidia/go-z3/st"
	"github.com/zyedidia/go-z3/z3"
)

type Machine struct {
	pc    int
	regs  []st.Int32
	conds []z3.Bool

	mem map[uint32]st.Int32

	ctx *z3.Context
}

func NewMachine(ctx *z3.Context, pc int) *Machine {
	regs := make([]st.Int32, 32)
	for i := range regs {
		regs[i] = st.Int32{C: 0}
	}

	return &Machine{
		pc:    pc,
		regs:  regs,
		conds: make([]z3.Bool, 0),
		mem:   make(map[uint32]st.Int32),
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

type Branch struct {
	pc   int
	cond st.Bool
}

func (m *Machine) Exec(insn uint32) (Branch, bool) {
	op := GetBits(insn, 6, 0).Uint32()

	switch op {
	case OpRarith:
		m.rarith(insn)
	case OpIarith:
		m.iarith(insn)
	case OpBranch:
		return m.branch(insn), true
	case OpLui:
		m.lui(insn)
	case OpAuipc:
		m.auipc(insn)
	case OpJal:
		return m.jal(insn), true
	case OpJalr:
		return m.jalr(insn), true
	case OpLoad:
		m.load(insn)
	case OpStore:
		m.store(insn)
	}
	return Branch{}, false
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
		pc:   m.pc + int(imm),
		cond: cond,
	}
}

func (m *Machine) jal(insn uint32) Branch {
	rd, _, _ := extractRegs(insn)
	imm := extractImm(insn, ImmJ)

	pc := m.pc + int(imm)
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
		// TODO: concretize
		panic("jalr target is symbolic")
	}
	m.WriteReg(rd, st.Int32{C: int32(m.pc + 4)})

	return Branch{
		pc:   int(pc.C),
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

	m.WriteReg(rd, st.Int32{C: int32(m.pc + int(imm))})
}

func (m *Machine) load(insn uint32) {
	rd, rs1, _ := extractRegs(insn)
	imm := extractImm(insn, ImmI)
	funct3 := GetBits(insn, 14, 12).Uint32()

	if funct3 != 0b010 {
		// TODO: lb/lh/lbu/lhu
		panic("unsupported: only lw is implemented")
	}

	rsval := m.regs[rs1]
	if !rsval.IsConcrete() {
		// TODO: concretize
		panic("load address is symbolic")
	}

	addr := uint32(rsval.C) + imm
	if v, ok := m.mem[addr/4]; ok {
		m.WriteReg(rd, v)
	} else {
		panic(fmt.Sprintf("invalid memory access at 0x%x", addr))
	}
}

func (m *Machine) store(insn uint32) {
	_, rs1, rs2 := extractRegs(insn)
	imm := extractImm(insn, ImmS)
	funct3 := GetBits(insn, 14, 12).Uint32()

	if funct3 != 0b010 {
		// TODO: sb/sh
		panic("unsupported: only sw is implemented")
	}

	rsval := m.regs[rs1]
	if !rsval.IsConcrete() {
		// TODO: concretize
		panic("store address is symbolic")
	}

	addr := uint32(rsval.C) + imm
	m.mem[addr/4] = m.regs[rs2]
}
