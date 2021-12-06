package bits

func GetBit(x uint32, bit int) uint32 {
	return (x >> bit) & 1
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
	return Clear(x, lb, ub) | (v << lb)
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
