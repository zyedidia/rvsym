package rvsym

import "math/rand"

type boolgen struct {
	src       rand.Source
	cache     int64
	remaining int
}

func (b *boolgen) Bool() bool {
	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}

var gen boolgen

func init() {
	gen = boolgen{
		src: rand.NewSource(42),
	}
}

func randbool() bool {
	return gen.Bool()
}
