package rvc

import (
	"testing"
)

func TestDecompress(t *testing.T) {
	insn := uint32(0xbfe9)
	dinsn, _, _ := Decompress(insn)
	if dinsn != 0xfdbff06f {
		t.Fatal()
	}
}
