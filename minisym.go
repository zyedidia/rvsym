package minisym

import (
	"fmt"

	"github.com/zyedidia/go-z3/st"
	"github.com/zyedidia/go-z3/z3"
)

type Machine struct {
	Regs []st.Int32

	ctx *z3.Context
}

func NewMachine() *Machine {
	ctx := z3.NewContext(nil)

	regs := make([]st.Int32, 32)

	regs[0] = st.Int32{C: 0}
	for i := range regs[1:] {
		regs[i+1] = st.AnyInt32(ctx, fmt.Sprintf("x%d", i+1))
	}

	return &Machine{
		Regs: regs,
		ctx:  ctx,
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

func (m *Machine) Exec(insn uint32) {
	op := GetBits(insn, 6, 0).Uint32()

	switch op {
	case OpRarith:
		m.execRarith(insn, op)
	case OpIarith:
		m.execIarith(insn, op)
	}
}

func (m *Machine) execIarith(insn, op uint32) {
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

func (m *Machine) execRarith(insn, op uint32) {
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

func (m *Machine) Model() *z3.Model {
	s := z3.NewSolver(m.ctx)
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
