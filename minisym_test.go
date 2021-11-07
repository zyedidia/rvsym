package minisym

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	insns := []uint32{
		0x00512093,
	}

	m := NewMachine()

	for _, i := range insns {
		m.Exec(i)
	}

	fmt.Println(m.MustReg("x1"))

	regs := m.Assignment()
	for i, r := range regs {
		fmt.Printf("%d: %d\n", i, r)
	}
}
