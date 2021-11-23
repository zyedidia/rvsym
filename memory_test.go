package rvsym

import (
	"testing"

	"github.com/zyedidia/rvsym/pkg/z3/st"
)

func TestMem(t *testing.T) {
	mem := NewMemory(nil)
	mem.Write32(st.Uint32{C: 0}, st.Int32{C: 0x0A0B0C0D}, nil)

	if v, ok := mem.Read8(st.Uint32{C: 0}, nil); !ok || v.C != 0xd {
		t.Fatal("fail 0")
	}
	if v, ok := mem.Read8(st.Uint32{C: 1}, nil); !ok || v.C != 0xc {
		t.Fatal("fail 1")
	}
	if v, ok := mem.Read8(st.Uint32{C: 2}, nil); !ok || v.C != 0xb {
		t.Fatal("fail 2")
	}
	if v, ok := mem.Read8(st.Uint32{C: 3}, nil); !ok || v.C != 0xa {
		t.Fatal("fail 3")
	}
}
