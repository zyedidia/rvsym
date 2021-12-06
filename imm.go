package rvsym

import "github.com/zyedidia/rvsym/bits"

type ImmType int

const (
	ImmTypeI ImmType = iota
	ImmTypeS
	ImmTypeB
	ImmTypeJ
	ImmTypeU
)

func extractImm(insn uint32, typ ImmType) int32 {
	switch typ {
	case ImmTypeI:
		return int32(bits.Sext(bits.Remap(insn, 31, 20, 11, 0), 12))
	case ImmTypeS:
		return int32(bits.Sext(
			bits.Remap(insn, 11, 7, 4, 0)|
				bits.Remap(insn, 31, 25, 11, 5),
			12,
		))
	case ImmTypeB:
		return int32(bits.Sext(
			bits.RemapBit(insn, 7, 11)|
				bits.Remap(insn, 11, 8, 4, 1)|
				bits.Remap(insn, 30, 25, 10, 5)|
				bits.RemapBit(insn, 31, 12),
			13,
		))
	case ImmTypeJ:
		return int32(bits.Sext(
			bits.RemapBit(insn, 31, 20)|
				bits.Remap(insn, 30, 21, 10, 1)|
				bits.RemapBit(insn, 20, 11)|
				bits.Remap(insn, 19, 12, 19, 12),
			21,
		))
	case ImmTypeU:
		return int32(insn & ^bits.Mask(12))
	}
	panic("invalid immediate type")
}
