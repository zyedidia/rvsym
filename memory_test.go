package rvsym

import (
	"fmt"
	"testing"

	"github.com/zyedidia/go-z3/st"
)

func TestMem(t *testing.T) {
	mem := make(Memory)
	mem.Write32(0, st.Int32{C: 0x0A0B0C0D})
	// mem.Write8(0, st.Int32{C: 0x0ff})
	// mem.Write8(1, st.Int32{C: 0x0ff})
	// mem.Write16(2, st.Int32{C: 0x0ffff})

	if v, ok := mem.Read8(0); ok {
		fmt.Printf("%x\n", v.C)
	}
	if v, ok := mem.Read8(1); ok {
		fmt.Printf("%x\n", v.C)
	}
	if v, ok := mem.Read8(2); ok {
		fmt.Printf("%x\n", v.C)
	}
	if v, ok := mem.Read8(3); ok {
		fmt.Printf("%x\n", v.C)
	}

	// if v, ok := mem.Read32(0); !ok {
	// 	t.Fatal("bad out of bounds")
	// } else if uint32(v.C) != 0xffffffff {
	// 	t.Fatal("bad", v)
	// }
}
