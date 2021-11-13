package rvsym

import (
	"testing"

	"github.com/zyedidia/go-z3/st"
)

func TestMem(t *testing.T) {
	mem := make(Memory)
	mem.Write32(0, st.Int32{C: 0})
	mem.Write8(0, st.Int32{C: 0x0ff})
	mem.Write8(1, st.Int32{C: 0x0ff})
	mem.Write16(2, st.Int32{C: 0x0ffff})

	if v, ok := mem.Read32(0); !ok {
		t.Fatal("bad out of bounds")
	} else if uint32(v.C) != 0xffffffff {
		t.Fatal("bad", v)
	}
}
