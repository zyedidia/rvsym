package main

var opts struct {
	Time    *int   `long:"time" description:"Stop execution after a given amount of seconds"`
	Summary bool   `short:"s" long:"summary" description:"Show execution summary"`
	Elf     string `long:"elf" description:"ELF debug information file"`
	Entry   uint32 `long:"entry" default:"4096" description:"Program start address"`
	Profile string `short:"p" long:"profile" description:"Dump profiling information to file"`
	Verbose bool   `short:"V" long:"verbose" description:"Show verbose debug information"`
	Version bool   `short:"v" long:"version" description:"Show version information"`
	Help    bool   `short:"h" long:"help" description:"Show this help message"`
}
