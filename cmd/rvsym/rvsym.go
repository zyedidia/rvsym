package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/zyedidia/rvsym"
)

func fatal(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}

func must(desc string, err error) {
	if err != nil {
		fatal(desc, ":", err)
	}
}

func main() {
	flagparser := flags.NewParser(&opts, flags.PassDoubleDash|flags.PrintErrors)
	flagparser.Usage = "[OPTIONS] BINFILE"
	args, err := flagparser.Parse()
	if err != nil {
		os.Exit(1)
	}

	if opts.Version {
		fmt.Println("rvsym version", Version)
		os.Exit(0)
	}

	if len(args) <= 0 || opts.Help {
		flagparser.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	bin, err := ioutil.ReadFile(args[0])
	must("read", err)

	loader := &rvsym.BinaryLoader{
		Entry: 0x1000,
	}
	segs, entry, err := loader.Load(bin)
	must("load", err)

	eng := rvsym.NewEngine(segs, entry)

	var last int
	for eng.Step() {
		last = showtc(eng, last)
	}
	showtc(eng, last)
}

func showtc(eng *rvsym.Engine, last int) int {
	if eng.NumTestCases() != last {
		last = eng.NumTestCases()
		tc := eng.TestCases()[last-1]
		fmt.Printf("--- Test case %d: %v at 0x%x ---\n", last-1, tc.Exit, uint64(tc.Pc))
		fmt.Print(tc.String(false))
		fmt.Println("---")
	}
	return last
}
