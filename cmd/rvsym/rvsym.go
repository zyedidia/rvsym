package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"strings"
	"time"

	"github.com/go-errors/errors"
	"github.com/jessevdk/go-flags"
	"github.com/zyedidia/rvsym"
	"github.com/zyedidia/rvsym/addr2line"
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
	flagparser.Usage = "[OPTIONS] EXE"
	args, err := flagparser.Parse()
	if err != nil {
		fatal(err)
	}

	if opts.Version {
		fmt.Println("rvsym version", Version)
		os.Exit(0)
	}

	if opts.Verbose {
		logger := log.New(os.Stdout, "INFO: ", 0)
		rvsym.SetLogger(logger)
	}

	if len(args) <= 0 || opts.Help {
		flagparser.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	if opts.Profile != "" {
		f, err := os.Create(opts.Profile)
		must("profile", err)
		defer f.Close() // error handling omitted for example
		err = pprof.StartCPUProfile(f)
		must("profile", err)
		defer pprof.StopCPUProfile()
	}

	bin, err := ioutil.ReadFile(args[0])
	must("read", err)

	mode := rvsym.EmuBareMetal

	var loader rvsym.Loader
	if strings.HasSuffix(args[0], ".hex") {
		loader = &rvsym.IHexLoader{
			Entry: opts.Entry,
		}
	} else if strings.HasSuffix(args[0], ".bin") {
		loader = &rvsym.RawLoader{
			Entry: opts.Entry,
		}
	} else {
		loader = &rvsym.ElfLoader{}
		opts.Elf = args[0]
		mode = rvsym.EmuLinux
	}

	converter := &addr2line.Converter{
		Elf: opts.Elf,
	}

	segs, entry, err := loader.Load(bin)
	must("load", err)

	if opts.Func != "" {
		entry, err = converter.FuncToAddr(opts.Func)
		must("elf lookup", err)
		mode = rvsym.EmuUnderConstrained
	}

	eng := rvsym.NewEngine(segs, entry, mode)

	showtc := func(eng *rvsym.Engine, last int) int {
		if eng.NumTestCases() != last {
			last = eng.NumTestCases()
			tc := eng.TestCases()[last-1]

			str := fmt.Sprintf("0x%x", uint32(tc.Pc))
			if opts.Elf != "" {
				pos, err := converter.Lookup(uint32(tc.Pc))
				if err == nil {
					str = fmt.Sprintf("%s (%s)", pos.String(), str)
				}
			}

			fmt.Printf("--- Test case %d: %v at %s ---\n", last-1, tc.Exit, str)
			fmt.Print(tc.String())
			fmt.Println("---")
		}
		return last
	}

	showSummary := func(eng *rvsym.Engine) {
		if opts.Summary {
			fmt.Print(eng.Summary())
		}
	}

	defer func() {
		if err := recover(); err != nil {
			showSummary(eng)
			fmt.Printf("%s\n%v\n", "a fatal error occurred", errors.Wrap(err, 2).ErrorStack())
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			showSummary(eng)
			os.Exit(0)
		}
	}()

	var timer *time.Timer
	if opts.Time != nil {
		timeout := time.Second * time.Duration(*opts.Time)
		timer = time.NewTimer(timeout)
	} else {
		timer = time.NewTimer(0)
		timer.Stop()
	}

	var last int

loop:
	for eng.Step() {
		select {
		case <-timer.C:
			log.Println("Execution timed out")
			break loop
		default:
			last = showtc(eng, last)
		}
	}
	showtc(eng, last)

	showSummary(eng)
}
