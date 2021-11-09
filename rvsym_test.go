package rvsym

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestConcrete(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/riscvtest.bin")
	if err != nil {
		t.Fatal(err)
	}
	code := LoadCode(data)
	eng := NewEngine(code)

	var panics []Panic
	for i := 0; i < 100; i++ {
		stpa := eng.Step()
		if len(stpa) > 0 {
			panics = append(panics, stpa...)
		}
	}

	for _, p := range panics {
		fmt.Printf("Universe %d panicked at 0x%x\n", p.Universe, p.Pc)
		fmt.Println(eng.UniverseInput(p.Universe))
	}

	fmt.Println(eng.machines[0].regs)
	fmt.Println(eng.machines[0].mem)
}
