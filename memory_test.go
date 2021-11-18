package rvsym

import (
	"testing"

	"github.com/zyedidia/go-z3/st"
)

func TestMem(t *testing.T) {
	mem := NewMemory(nil)
	mem.Write32(0, st.Int32{C: 0x0A0B0C0D})

	if v, ok := mem.Read8(0); !ok || v.C != 0xd {
		t.Fatal("fail 0")
	}
	if v, ok := mem.Read8(1); !ok || v.C != 0xc {
		t.Fatal("fail 1")
	}
	if v, ok := mem.Read8(2); !ok || v.C != 0xb {
		t.Fatal("fail 2")
	}
	if v, ok := mem.Read8(3); !ok || v.C != 0xa {
		t.Fatal("fail 3")
	}
}
