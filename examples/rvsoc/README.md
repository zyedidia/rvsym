# rvsoc

A small RISC-V System-on-Chip. Currently features a multicycle processor
supporting most of the unprivileged RV32i spec.

# Requirements

You can skip some requirements if you are only doing simulation.

* Yosys (simulation, synthesis)
* NextPNR-ECP5 (synthesis)
* Project Trellis toolchain (synthesis bitstream generation)
* Verilator (linting)
* GTKWave (waveform simulation)
* g++ (simulation)

# Usage

All make commmands are run from `rtl/`.

Linting:

```
make lint
```

Simulation:

```shell
make soc_top_sim.cc # build C++ simulator
make tb             # build testbench using soc_tb.cc and soc_top_sim.cc
make test           # run testbench
make waveform       # open generated waveform (requires GTKWave)
```

To run a particular test, first edit `rtl/soc/ram.sv` to select the appropriate
`.mem` file to load into memory. Then re-run the testbench: `make clean && make test`.

Synthesis:

```shell
make synth # synthesize FPGA bitstream
make prog  # send bitstream to FPGA (must be plugged in)
```

# TODO:

* Support the `Zicsr` control/status registers extension.
* Support M-mode from the privilege spec.
* Improve test infrastructure

# RISC-V Spec

* [Unprivileged spec](https://github.com/riscv/riscv-isa-manual/releases/download/Ratified-IMAFDQC/riscv-spec-20191213.pdf)
* [Privileged spec](https://github.com/riscv/riscv-isa-manual/releases/download/Ratified-IMFDQC-and-Priv-v1.11/riscv-privileged-20190608.pdf)
