package rvsym

import (
	"github.com/zyedidia/rvsym/pkg/z3/st"
)

func alu(a, b st.Int32, op uint32, sub, sharith, malu bool) st.Int32 {
	if malu {
		maluRet := st.Int32{Secret: false}
		switch op {
		case MAluMul:
			maluRet = a.Mul(b)
		case MAluMulH:
			maluRet = a.ToInt64().Mul(b.ToInt64()).Upper32()
		case MAluMulHSU:
			maluRet = a.ToInt64().ToUint64().Mul(b.ToUint32().ToUint64()).Upper32()
		case MAluMulHU:
			maluRet = a.ToUint32().ToUint64().Mul(b.ToUint32().ToUint64()).Upper32()
		case MAluDiv:
			maluRet = a.Quo(b)
		case MAluDivU:
			maluRet = a.ToUint32().Quo(b.ToUint32()).ToInt32()
		case MAluRem:
			maluRet = a.Rem(b)
		case MAluRemU:
			maluRet = a.ToUint32().Rem(b.ToUint32()).ToInt32()
		}
		if (a.Secret || b.Secret) {
			maluRet.Secret = true
		}
		return maluRet
	}
	ret := st.Int32{}
	switch op {
	case AluAdd:
		if sub {
			ret = a.Sub(b)
		} else {
			ret = a.Add(b)
		}
	case AluShl:
		ret = a.Lsh(b.ToUint64())
	case AluSlt:
		ret = a.LT(b).ToInt32()
	case AluSltu:
		ret = a.ToUint32().LT(b.ToUint32()).ToInt32()
	case AluXor:
		ret = a.Xor(b)
	case AluShr:
		if sharith {
			ret = a.Rsh(b.ToUint64())
		} else {
			ret = a.ToUint32().Rsh(b.ToUint64()).ToInt32()
		}
	case AluOr:
		ret = a.Or(b)
	case AluAnd:
		ret = a.And(b)
	default:
		panic("unreachable")
	}
	if (a.Secret || b.Secret) {
		ret.Secret = true
	}

	return ret
}

func extractImm(insn uint32, typ ImmType) uint32 {
	switch typ {
	case ImmI:
		return CatBits(
			RepeatBit(insn, 31, 20),
			GetBits(insn, 31, 20),
		).Uint32()
	case ImmS:
		return CatBits(
			RepeatBit(insn, 31, 20),
			GetBits(insn, 31, 25),
			GetBits(insn, 11, 7),
		).Uint32()
	case ImmB:
		return CatBits(
			RepeatBit(insn, 31, 20),
			GetBits(insn, 7, 7),
			GetBits(insn, 30, 25),
			GetBits(insn, 11, 8),
			Bits{0, 1},
		).Uint32()
	case ImmJ:
		return CatBits(
			RepeatBit(insn, 31, 12),
			GetBits(insn, 19, 12),
			GetBits(insn, 20, 20),
			GetBits(insn, 30, 21),
			Bits{0, 1},
		).Uint32()
	case ImmU:
		return CatBits(
			GetBits(insn, 31, 12),
			Bits{0, 12},
		).Uint32()
	}
	return 0
}
