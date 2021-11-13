package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"

	"github.com/zyedidia/rvsym"
)

var summary = flag.Bool("summary", false, "provide a path exploration summary")
var max = flag.Int("max", -1, "Maximum number of machines (-1 for unlimited)")
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

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

	steps := 0
	for eng.Step() {
		steps += eng.NumMachines()
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
		fmt.Println("--- Summary ---")
		fmt.Printf("Instructions executed: %d\n", steps)
		fmt.Printf("Total paths: %d\n", paths)
		fmt.Printf("Quiet exits: %d\n", eng.Stats.Exits[rvsym.ExitQuiet])
		fmt.Printf("Normal exits: %d\n", eng.Stats.Exits[rvsym.ExitNormal])
		fmt.Printf("Failures: %d\n", eng.Stats.Exits[rvsym.ExitFail])
	}
}
