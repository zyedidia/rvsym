package minisym

type Bits struct {
	val uint32
	n   int
}

func GetBits(i uint32, top, bot int) Bits {
	return Bits{i >> bot, top - bot + 1}
}

func CatBits(bs ...Bits) Bits {
	var b Bits
	for i := range bs {
		b = b.Cat(bs[i])
	}
	return b
}

func (b1 Bits) Cat(b2 Bits) Bits {
	return Bits{
		val: (b1.val << b2.n) | b2.Uint32(),
		n:   b1.n + b2.n,
	}
}

func (b Bits) Uint32() uint32 {
	return b.val & ^((^uint32(0) >> b.n) << b.n)
}
