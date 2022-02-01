package bits

type Vec struct {
	B  uint32
	Sz int
}

func Join(vecs ...Vec) uint32 {
	var x uint32
	start := 0
	for i := len(vecs) - 1; i >= 0; i-- {
		v := vecs[i]
		x = Set(x, start+v.Sz-1, start, v.B)
		start += v.Sz
	}
	return x
}

func GetBit(x uint32, bit int) uint32 {
	return (x >> bit) & 1
}

func Repeat(bit uint32, n int) uint32 {
	if bit == 0 {
		return 0
	}
	return 0xffffffff & Mask(n)
}

func Mask(nbits int) uint32 {
	if nbits == 32 {
		return ^uint32(0)
	}
	return (1 << nbits) - 1
}

func Get(x uint32, ub, lb int) uint32 {
	return (x >> lb) & Mask(ub-lb+1)
}

func Clear(x uint32, ub, lb int) uint32 {
	mask := Mask(ub - lb + 1)
	return x & ^(mask << lb)
}

func Set(x uint32, ub, lb int, v uint32) uint32 {
	return Clear(x, ub, lb) | (v << lb)
}

func RemapBit(i uint32, from, to int) uint32 {
	return GetBit(i, from) << to
}

func Remap(i uint32, fromub, fromlb, toub, tolb int) uint32 {
	return Get(i, fromub, fromlb) << tolb
}

func Sext(x uint32, width int) uint32 {
	n := 31 - (width - 1)
	return uint32(int32(x<<n) >> n)
}
