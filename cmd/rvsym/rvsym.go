package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/signal"
	"runtime/pprof"

	"github.com/go-errors/errors"
	"github.com/zyedidia/rvsym"
	"github.com/zyedidia/rvsym/bininfo"
)

var summary = flag.Bool("summary", false, "provide a path exploration summary")
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var hex = flag.Bool("hex", false, "print test cases in hex")
var forks = flag.Int("forks", math.MaxInt, "maximum number of forks")

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
	eng := rvsym.NewEngine(code, *forks)

	var dwarf io.ReaderAt
	if len(args) > 1 {
		dwarf, _ = os.Open(args[1])
	} else {
		dwarf = bytes.NewReader(bin)
	}
	binfo, _ := bininfo.Read(dwarf)

	defer func() {
		if err := recover(); err != nil {
			if *summary {
				fmt.Print(eng.Summary())
			}
			fmt.Printf("%s\n%v\n", "a fatal error occurred", errors.Wrap(err, 2).ErrorStack())
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			if *summary {
				fmt.Print(eng.Summary())
			}
			os.Exit(0)
		}
	}()

	var last int
	for eng.Step() {
		last = showtc(eng, binfo, last)
	}
	showtc(eng, binfo, last)

	if *summary {
		fmt.Print(eng.Summary())
	}
}

func showtc(eng *rvsym.Engine, binfo *bininfo.BinFile, last int) int {
	if len(eng.TestCases()) != last {
		last = len(eng.TestCases())
		tc := eng.TestCases()[last-1]
		fmt.Printf("--- Test case %d: %v at %s ---\n", last-1, tc.Exit, binfo.PosStr(uint64(tc.Addr)))
		fmt.Print(tc.String(*hex))
		fmt.Println("---")
	}
	return last
}
