package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"

	"github.com/zyedidia/rvsym"
	"github.com/zyedidia/rvsym/bininfo"
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

	var dwarf io.ReaderAt
	if len(args) > 1 {
		dwarf, _ = os.Open(args[1])
	} else {
		dwarf = bytes.NewReader(bin)
	}
	binfo, _ := bininfo.Read(dwarf)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			for i, tc := range eng.TestCases() {
				fmt.Printf("--- Test case %d: %v at %s ---\n", i, tc.Exit, binfo.PosStr(uint64(tc.Addr)))
				fmt.Print(tc)
				fmt.Println("---")
			}

			fmt.Print(eng.Summary())
			os.Exit(1)
		}
	}()

	for eng.Step() {
	}

	for i, tc := range eng.TestCases() {
		fmt.Printf("--- Test case %d: %v at %s ---\n", i, tc.Exit, binfo.PosStr(uint64(tc.Addr)))
		fmt.Print(tc)
		fmt.Println("---")
	}

	if *summary {
		fmt.Print(eng.Summary())
	}
}
