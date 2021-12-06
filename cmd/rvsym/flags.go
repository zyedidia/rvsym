package main

var opts struct {
	Entry   uint32 `long:"entry" default:"4096" description:"Address of program entrypoint"`
	Profile string `short:"p" long:"profile" description:"Dump profiling information to file"`
	Verbose bool   `short:"V" long:"verbose" description:"Show verbose debug information"`
	Version bool   `short:"v" long:"version" description:"Show version information"`
	Help    bool   `short:"h" long:"help" description:"Show this help message"`
}
