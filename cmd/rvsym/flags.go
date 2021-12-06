package main

var opts struct {
	Profile string `short:"p" long:"profile" description:"Dump profiling information to file"`
	Verbose bool   `short:"V" long:"verbose" description:"Show verbose debug information"`
	Version bool   `short:"v" long:"version" description:"Show version information"`
	Help    bool   `short:"h" long:"help" description:"Show this help message"`
}
