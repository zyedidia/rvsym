package minisym

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	insns := []uint32{
		0x00500293,
		0x00128463,
		0x00000463,
		0x00528293,
		0x00000063,
	}

	e := NewEngine(insns)

	for i := 0; i < 10; i++ {
		e.Step()
	}

	fmt.Println(len(e.Machines))
	fmt.Println(e.Machines[0].MustReg("x5"))
	fmt.Println(e.Machines[1].MustReg("x5"))
}
