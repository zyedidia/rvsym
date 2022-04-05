# Rvsym: a small RISC-V symbolic execution engine

Rvsym can execute rv32im RISC-V programs symbolically.

# Building

By default, rvsym uses Boolector as the SMT backend. First you need to set up
Boolector:

```
./pkg/smt/setup-boolector.sh
```

Then you can build

```
go build ./cmd/rvsym
```

Alternatively, you may build with dynamic linking to Z3 if it is installed:

```
go build -tags z3,noboolector ./cmd/rvsym
```

In order to build the examples you will also need the RISC-V GNU toolchain.
Pre-built toolchains are available from SiFive
[here](https://www.sifive.com/software). After downloading, unpack the tar
archive to `/opt/riscv` (or if you choose another location, point the `RISCV`
environment variable to your installation).

# Usage

See the `./examples/basic` directory for a number of example programs. Compile with `make`,
and then run the example of your choosing with `rvsym example.elf`.

For example:

```
$ make get_sign.elf
$ rvsym get_sign.elf --summary
--- Test case 0: exit at get_sign.c:20 (0x100a4) ---
a[3:0] -> 0x1
---
--- Test case 1: exit at get_sign.c:20 (0x100a4) ---
a[3:0] -> 0x80000001
---
--- Test case 2: exit at get_sign.c:20 (0x100a4) ---
a[3:0] -> 0x0
---
--- Summary ---
Instructions executed: 148
Total paths: 3
Quiet exits: 0
Unsat exits: 0
Normal exits: 3
Failures: 0
---
```

Rvsym has the following additional options:

```
Usage:
  rvsym [OPTIONS] EXE

Application Options:
      --time=    Stop execution after a given amount of seconds
  -s, --summary  Show execution summary
      --elf=     ELF debug information file
      --entry=   Program start address (default: 4096)
  -p, --profile= Dump profiling information to file
  -V, --verbose  Show verbose debug information
  -v, --version  Show version information
  -h, --help     Show this help message
```

# Notes

Rvsym runs in Linux mode for ELF files, and bare-metal mode for `.hex` or
`.bin` files. In Linux mode, Rvsym sets up the stack properly and provides a
limited set of system calls, shown below. In bare-metal mode, Rvsym does not do
any set-up. In the future bare-metal mode may include support for exceptions
and other privileged CPU features.

The following system calls are provided in Linux mode: `exit`, `open`, `read`,
`write`, `lseek`, `close`, `fstat` (does not provide accurate information),
`brk`. These are a minimal set allowing basic memory allocation and file I/O.
