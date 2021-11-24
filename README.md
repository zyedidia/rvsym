# rvsym

A small RISC-V symbolic execution engine.

# Building

Rvsym requires z3. To install z3, run

```
git clone https://github.com/Z3Prover/z3
cd z3
mkdir build
cd build
cmake -G "Unix Makefiles" ../
make -j4
sudo make install
```

Then to build rvsym run

```
go build ./cmd/rvsym
```

# Usage

To use the makefiles, making sure to set the environment variable `RVSYM_ROOT`
to point to this directory. If you want to test the `rvsoc` example, also
install Yosim and point `YOSIM_ROOT` to that directory.
