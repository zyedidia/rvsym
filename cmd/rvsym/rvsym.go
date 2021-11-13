package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/zyedidia/rvsym"
)

var summary = flag.Bool("summary", false, "provide a path exploration summary")
var max = flag.Int("max", -1, "Maximum number of machines (-1 for unlimited)")

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
	eng.MaxMachines = *max

	for eng.Step() {
	}

	for i, tc := range eng.TestCases() {
		fmt.Printf("--- Test case %d: %v at 0x%x ---\n", i, tc.Exit, tc.Addr)
		fmt.Print(tc)
		fmt.Println("---")
	}

	if *summary {
		paths := 0
		for _, v := range eng.Stats.Exits {
			paths += v
		}
		fmt.Println("Summary:")
		fmt.Printf("Total paths: %d\n", paths)
		fmt.Printf("Quiet exits: %d\n", eng.Stats.Exits[rvsym.ExitQuiet])
		fmt.Printf("Normal exits: %d\n", eng.Stats.Exits[rvsym.ExitNormal])
		fmt.Printf("Failures: %d\n", eng.Stats.Exits[rvsym.ExitFail])
	}
}
