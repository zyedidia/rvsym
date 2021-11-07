package minisym

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestConcrete(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/riscvtest.o")
	if err != nil {
		t.Fatal(err)
	}
	code := LoadCode(data)

	eng := NewEngine(code)

	for i := 0; i < 100; i++ {
		eng.Step()
	}

	fmt.Println(eng.machines[0].regs)
	fmt.Println(eng.machines[0].mem)
}
