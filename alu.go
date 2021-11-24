package rvsym

import (
	"github.com/zyedidia/go-z3/st"
)

func alu(a, b st.Int32, op uint32, sub, sharith, malu bool) st.Int32 {
	if malu {
		maluRet := st.Int32{Secret: false}
		switch op {
		case MAluMul:
			//return a.Mul(b)
			maluRet = a.Mul(b)
		case MAluMulH:
			//return a.ToInt64().Mul(b.ToInt64()).Upper32()
			maluRet = a.ToInt64().Mul(b.ToInt64()).Upper32()
		case MAluMulHSU:
			//return a.ToInt64().ToUint64().Mul(b.ToUint32().ToUint64()).Upper32()
			maluRet = a.ToInt64().ToUint64().Mul(b.ToUint32().ToUint64()).Upper32()
		case MAluMulHU:
			//return a.ToUint32().ToUint64().Mul(b.ToUint32().ToUint64()).Upper32()
			maluRet = a.ToUint32().ToUint64().Mul(b.ToUint32().ToUint64()).Upper32()
		case MAluDiv:
			//return a.Quo(b)
			maluRet = a.Quo(b)
		case MAluDivU:
			//return a.ToUint32().Quo(b.ToUint32()).ToInt32()
			maluRet = a.ToUint32().Quo(b.ToUint32()).ToInt32()
		case MAluRem:
			//return a.Rem(b)
			maluRet = a.Rem(b)
		case MAluRemU:
			//return a.ToUint32().Rem(b.ToUint32()).ToInt32()
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
			//return a.Sub(b)
			ret = a.Sub(b)
		} else {
			//return a.Add(b)
			ret = a.Add(b)
		}
	case AluShl:
		//return a.Lsh(b.ToUint64())
		ret = a.Lsh(b.ToUint64())
	case AluSlt:
		//return a.LT(b).ToInt32()
		ret = a.LT(b).ToInt32()
	case AluSltu:
		//return a.ToUint32().LT(b.ToUint32()).ToInt32()
		ret = a.ToUint32().LT(b.ToUint32()).ToInt32()
	case AluXor:
		//return a.Xor(b)
		ret = a.Xor(b)
	case AluShr:
		if sharith {
			//return a.Rsh(b.ToUint64())
			ret = a.Rsh(b.ToUint64())
		} else {
			//return a.ToUint32().Rsh(b.ToUint64()).ToInt32()
			ret = a.ToUint32().Rsh(b.ToUint64()).ToInt32()
		}
	case AluOr:
		//return a.Or(b)
		ret = a.Or(b)
	case AluAnd:
		//return a.And(b)
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
