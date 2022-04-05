package rvsym

import (
	"github.com/zyedidia/rvsym/pkg/smt"
)

type AluOp int

const (
	AluAdd    = 0b0000000000
	AluSub    = 0b0100000000
	AluSlt    = 0b0000000010
	AluSltu   = 0b0000000011
	AluXor    = 0b0000000100
	AluSll    = 0b0000000001
	AluSrl    = 0b0000000101
	AluSra    = 0b0100000101
	AluOr     = 0b0000000110
	AluAnd    = 0b0000000111
	AluMul    = 0b0000001000
	AluMulH   = 0b0000001001
	AluMulHSU = 0b0000001010
	AluMulHU  = 0b0000001011
	AluDiv    = 0b0000001100
	AluDivU   = 0b0000001101
	AluRem    = 0b0000001110
	AluRemU   = 0b0000001111
)

func (m *Machine) alu(a, b smt.Int32, op AluOp, s *smt.Solver) smt.Int32 {
	switch op {
	case AluAdd:
		return a.Add(b, s)
	case AluSub:
		return a.Sub(b, s)
	case AluSlt:
		return a.Slt(b, s)
	case AluSltu:
		return a.Ult(b, s)
	case AluXor:
		return a.Xor(b, s)
	case AluSll:
		return a.Sll(b, s)
	case AluSrl:
		return a.Srl(b, s)
	case AluSra:
		return a.Sra(b, s)
	case AluOr:
		return a.Or(b, s)
	case AluAnd:
		return a.And(b, s)
	case AluMul:
		return a.Mul(b, s)
	case AluMulH:
		return a.ToInt64s(s).Mul(b.ToInt64s(s), s).Upper32(s)
	case AluMulHSU:
		return a.ToInt64s(s).Mul(b.ToInt64z(s), s).Upper32(s)
	case AluMulHU:
		return a.ToInt64z(s).Mul(b.ToInt64z(s), s).Upper32(s)
	case AluDiv:
		return a.Div(b, s)
	case AluDivU:
		return a.Divu(b, s)
	case AluRem:
		return a.Rem(b, s)
	case AluRemU:
		return a.Remu(b, s)
	}
	panic("invalid alu op")
}
