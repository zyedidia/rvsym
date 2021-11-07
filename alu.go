package rvsym

import "github.com/zyedidia/go-z3/st"

func alu(a, b st.Int32, op uint32, sub, arishift bool) st.Int32 {
	switch op {
	case AluAdd:
		if sub {
			return a.Sub(b)
		} else {
			return a.Add(b)
		}
	case AluShl:
		return a.Lsh(b.ToUint64())
	case AluSlt:
		return a.LT(b).ToInt32()
	case AluSltu:
		return a.ToUint32().LT(b.ToUint32()).ToInt32()
	case AluXor:
		return a.Xor(b)
	case AluShr:
		if arishift {
			return a.Rsh(b.ToUint64())
		} else {
			return a.ToUint32().Rsh(b.ToUint64()).ToInt32()
		}
	case AluOr:
		return a.Or(b)
	case AluAnd:
		return a.And(b)
	}
	panic("unreachable")
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
