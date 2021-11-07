package minisym

import (
	"fmt"

	"github.com/zyedidia/go-z3/st"
	"github.com/zyedidia/go-z3/z3"
)

type Engine struct {
	Machines []*Machine
}

func NewEngine(insns []uint32) *Engine {
	return &Engine{
		Machines: []*Machine{NewMachine(insns, 0)},
	}
}

func (e *Engine) Step() {
	for i := 0; i < len(e.Machines); i++ {
		m := e.Machines[i]
		f := m.Exec(m.Insns[m.PC/4])
		if f != nil {
			if f.Cond.IsConcrete() {
				if f.Cond.C {
					m.PC = f.PC
				} else {
					m.PC += 4
				}
			} else {
				copied := &Machine{
					Insns: m.Insns,
					PC:    m.PC + 4,
					ctx:   m.ctx,

					Regs:  make([]st.Int32, len(m.Regs)),
					conds: make([]z3.Bool, len(m.conds)),
				}
				copy(copied.Regs, m.Regs)
				copy(copied.conds, m.conds)

				e.Machines = append(e.Machines, copied)

				m.conds = append(m.conds, f.Cond.S)
				m.PC = f.PC
			}
		} else {
			m.PC += 4
		}
	}
}

type Machine struct {
	Insns []uint32
	PC    int

	Regs []st.Int32

	ctx *z3.Context

	conds []z3.Bool
}

type Fork struct {
	PC   int
	Cond st.Bool
}

func NewMachine(insns []uint32, start int) *Machine {
	ctx := z3.NewContext(nil)

	regs := make([]st.Int32, 32)

	regs[0] = st.Int32{C: 0}
	for i := range regs[1:] {
		regs[i+1] = st.AnyInt32(ctx, fmt.Sprintf("x%d", i+1))
	}

	return &Machine{
		Insns: insns,
		PC:    start,
		Regs:  regs,
		ctx:   ctx,
	}
}

func (m *Machine) Reg(name string) (st.Int32, error) {
	if r, ok := RegNums[name]; ok {
		return m.Regs[r], nil
	}
	return st.Int32{C: 0}, fmt.Errorf("invalid register %s", name)
}

func (m *Machine) MustReg(name string) st.Int32 {
	r, err := m.Reg(name)
	if err != nil {
		panic(err)
	}
	return r
}

func (m *Machine) Exec(insn uint32) *Fork {
	op := GetBits(insn, 6, 0).Uint32()

	switch op {
	case OpRarith:
		m.execRarith(insn)
	case OpIarith:
		m.execIarith(insn)
	case OpBranch:
		return m.execBranch(insn)
	case OpLui:
	case OpAuipc:
	case OpJal:
	case OpJalr:
	case OpLoad:
	case OpStore:
	}
	return nil
}

func (m *Machine) execBranch(insn uint32) *Fork {
	funct3 := GetBits(insn, 14, 12).Uint32()
	rs1 := GetBits(insn, 19, 15).Uint32()
	rs2 := GetBits(insn, 24, 20).Uint32()
	imm := int32(CatBits(
		Bits{GetBits(insn, 31, 31).Uint32(), 20},
		GetBits(insn, 31, 31),
		GetBits(insn, 7, 7),
		GetBits(insn, 30, 25),
		GetBits(insn, 11, 8),
		Bits{0, 1},
	).Uint32())

	var aluop uint32
	switch funct3 {
	case 0b000, 0b001:
		aluop = AluXor
	case 0b100, 0b101:
		aluop = AluSlt
	case 0b110, 0b111:
		aluop = AluSltu
	}

	var nextpc uint32
	switch funct3 {
	case 0b000, 0b101, 0b111:
		nextpc = NextPCZ
	case 0b001, 0b100, 0b110:
		nextpc = NextPCNZ
	}

	var result st.Int32
	switch aluop {
	case AluXor:
		result = m.Regs[rs1].Xor(m.Regs[rs2])
	case AluSlt:
		result = m.Regs[rs1].LT(m.Regs[rs2]).ToInt32()
	case AluSltu:
		result = m.Regs[rs1].ToUint32().LT(m.Regs[rs2].ToUint32()).ToInt32()
	}

	switch nextpc {
	case NextPCZ:
		return &Fork{
			PC:   m.PC + int(imm),
			Cond: result.Eq(st.Int32{C: 0}),
		}
	case NextPCNZ:
		return &Fork{
			PC:   m.PC + int(imm),
			Cond: result.NE(st.Int32{C: 0}),
		}
	}
	return nil
}

func (m *Machine) execIarith(insn uint32) {
	aluop := GetBits(insn, 14, 12).Uint32()
	imm := GetBits(insn, 31, 20).Uint32()
	rd := GetBits(insn, 11, 7).Uint32()
	rs1 := GetBits(insn, 19, 15).Uint32()
	modify := GetBits(insn, 30, 30).Uint32() != 0

	i32imm := st.Int32{C: int32(imm)}
	u32imm := st.Uint32{C: imm}
	u64imm := st.Uint64{C: uint64(imm)}

	switch aluop {
	case AluAdd:
		m.Regs[rd] = m.Regs[rs1].Add(i32imm)
	case AluSlt:
		m.Regs[rd] = m.Regs[rs1].LT(i32imm).ToInt32()
	case AluSltu:
		m.Regs[rd] = m.Regs[rs1].ToUint32().LT(u32imm).ToInt32()
	case AluXor:
		m.Regs[rd] = m.Regs[rs1].Xor(i32imm)
	case AluOr:
		m.Regs[rd] = m.Regs[rs1].Or(i32imm)
	case AluAnd:
		m.Regs[rd] = m.Regs[rs1].And(i32imm)
	case AluShl:
		m.Regs[rd] = m.Regs[rs1].Lsh(u64imm)
	case AluShr:
		if modify {
			panic("unimplemented SRAI")
		} else {
			m.Regs[rd] = m.Regs[rs1].Lsh(u64imm)
		}
	}
}

func (m *Machine) execRarith(insn uint32) {
	aluop := GetBits(insn, 14, 12).Uint32() // funct3
	modify := GetBits(insn, 30, 30).Uint32() != 0
	rd := GetBits(insn, 11, 7).Uint32()
	rs1 := GetBits(insn, 19, 15).Uint32()
	rs2 := GetBits(insn, 24, 20).Uint32()

	switch aluop {
	case AluAdd:
		if modify {
			// sub
			m.Regs[rd] = m.Regs[rs1].Sub(m.Regs[rs2])
		} else {
			// add
			m.Regs[rd] = m.Regs[rs1].Add(m.Regs[rs2])
		}
	case AluShl:
		m.Regs[rd] = m.Regs[rs1].Lsh(m.Regs[rs2].ToUint64())
	case AluSlt:
		m.Regs[rd] = m.Regs[rs1].LT(m.Regs[rs2]).ToInt32()
	case AluSltu:
		m.Regs[rd] = m.Regs[rs1].ToUint32().LT(m.Regs[rs2].ToUint32()).ToInt32()
	case AluXor:
		m.Regs[rd] = m.Regs[rs1].Xor(m.Regs[rs2])
	case AluShr:
		if modify {
			panic("unimplemented SRA")
		} else {
			m.Regs[rd] = m.Regs[rs1].Rsh(m.Regs[rs2].ToUint64())
		}
	case AluOr:
		m.Regs[rd] = m.Regs[rs1].Or(m.Regs[rs2])
	case AluAnd:
		m.Regs[rd] = m.Regs[rs1].And(m.Regs[rs2])
	}
}

func (m *Machine) concretize(val st.Int32) int32 {
	return val.Eval(m.Model())
}

func (m *Machine) Context() *z3.Context {
	return m.ctx
}

func (m *Machine) Model() *z3.Model {
	s := z3.NewSolver(m.ctx)

	for _, c := range m.conds {
		s.Assert(c)
	}

	sat, err := s.Check()
	if err != nil {
		panic(err)
	}
	if !sat {
		panic(fmt.Sprintf("concretization unsatisfiable"))
	}
	return s.Model()
}

func (m *Machine) Assignment() []int32 {
	model := m.Model()

	regs := make([]int32, 32)
	for i := range regs {
		regs[i] = m.Regs[i].Eval(model)
	}
	return regs
}
