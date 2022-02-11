# Rvsym: a small RISC-V symbolic execution engine

Rvsym can execute bare-metal rv32im RISC-V programs symbolically.

# Building

```
go build ./cmd/rvsym
```

This build uses Boolector as the SMT backend and statically links. Alternatively,
you may build with dynamic linking to Z3 if it is installed:

```
go build -tags z3,noboolector ./cmd/rvsym
```

The makefiles for building examples try to locate your rvsym install directory
using git. If you didn't download using git, or you are trying to build an
example outside of the rvsym repository, you can set the `RVSYM_ROOT`
environment variable to point to your installation (if you want to use the
included library and build rules in `lib/`).

# Usage

See the `./examples/basic` directory for a number of example programs. Compile with `make`,
and then run the example of your choosing with `rvsym example.hex`.

For example:

```
$ make get_sign.hex
$ rvsym get_sign.hex --summary
--- Test case 0: exit at 0x1220 ---
a[3:0] -> 0x1
---
--- Test case 1: exit at 0x1220 ---
a[3:0] -> 0x80000001
---
--- Test case 2: exit at 0x1220 ---
a[3:0] -> 0x0
---
--- Summary ---
Instructions executed: 29
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
  rvsym [OPTIONS] BIN/HEX

Application Options:
      --time=    Stop execution after a given amount of seconds
  -s, --summary  Show execution summary
      --dump     Dump smt2 requests for generated test cases
      --elf=     ELF debug information file
      --entry=   Program start address (default: 4096)
  -p, --profile= Dump profiling information to file
  -V, --verbose  Show verbose debug information
  -v, --version  Show version information
  -h, --help     Show this help message
```
