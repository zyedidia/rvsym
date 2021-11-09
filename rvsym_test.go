package rvsym

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestConcrete(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/test.bin")
	if err != nil {
		t.Fatal(err)
	}
	code := LoadCode(data)
	eng := NewEngine(code)

	var exits []Exit
	for i := 0; i < 100; i++ {
		stpa, _ := eng.Step()
		if len(stpa) > 0 {
			exits = append(exits, stpa...)
		}
	}

	for _, e := range exits {
		if e.Status == ExitFail {
			fmt.Printf("--- Universe %d FAILED ---\n", e.Universe)
		} else {
			fmt.Println("---")
		}
		fmt.Printf("Universe %d exited at 0x%x\n", e.Universe, e.Pc)
		fmt.Println("Test case:")
		fmt.Print(eng.UniverseInput(e.Universe))
		fmt.Println("---")
	}

	fmt.Println(eng.machines[0].regs)
	fmt.Println(eng.machines[0].mem)
}
