package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/zyedidia/rvsym"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) <= 0 {
		log.Fatal("no input")
	}

	bin, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	code := rvsym.LoadCode(bin)
	eng := rvsym.NewEngine(code)

	var exits []rvsym.Exit

	for {
		ex, done := eng.Step()
		exits = append(exits, ex...)

		if done {
			break
		}
	}

	for _, e := range exits {
		if e.Status == rvsym.ExitQuiet {
			continue
		} else if e.Status == rvsym.ExitFail {
			fmt.Printf("--- Universe %d FAILED at 0x%x ---\n", e.Universe, e.Pc)
		} else {
			fmt.Printf("--- Universe %d exited at 0x%x ---\n", e.Universe, e.Pc)
		}
		fmt.Println("Test case:")
		fmt.Print(eng.UniverseInput(e.Universe))
		fmt.Println("---")
	}
}
