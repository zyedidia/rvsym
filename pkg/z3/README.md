go-z3 provides Go bindings for
the [Z3 SMT solver](https://github.com/Z3Prover/z3).

This repository is a fork of
[aclements/go-z3](https://github.com/aclements/go-z3) that makes some minor
fixes so that the bindings work with a modern version of z3 (4.8.13 at the time
of writing).

Installation
============

First, follow the instructions to
[download and install](https://github.com/Z3Prover/z3/blob/master/README.md)
the Z3 C library. Here is my recommended installation procedure:

```
git clone https://github.com/Z3Prover/z3
cd z3
mkdir build
cd build
cmake -G "Unix Makefiles" ../
make -j4
sudo make install
```

If you installed the C library to a non-default location (such as a
directory under `$HOME`), set the following environment variables:

```sh
# For building:
export CGO_CFLAGS=-I$Z3PREFIX/include CGO_LDFLAGS=-L$Z3PREFIX/lib
# For running binaries (including tests):
export LD_LIBRARY_PATH=$Z3PREFIX/lib
```

Then add go-z3 to your project:

```sh
go get github.com/zyedidia/rvsym/pkg/z3/z3
```

Documentation
=============

See the [godoc](https://godoc.org/github.com/zyedidia/rvsym/pkg/z3/z3).

Example
=======

```go
package main

import (
	"fmt"
	"log"

	"github.com/zyedidia/rvsym/pkg/z3/z3"
)

func main() {
	ctx := z3.NewContext(nil)

	s := z3.NewSolver(ctx)

	z3int := ctx.IntSort()
	x := ctx.Const("x", z3int).(z3.Int)
	y := ctx.Const("y", z3int).(z3.Int)
	z := ctx.Const("z", z3int).(z3.Int)

	zero := ctx.FromInt(0, z3int).(z3.Int)

	s.Assert(x.Add(y, z).GT(ctx.FromInt(4, z3int).(z3.Int)))
	s.Assert(z.GT(zero))

	s.Assert(x.Eq(y).Not())
	s.Assert(y.Eq(z).Not())

	s.Assert(x.Eq(zero).Not())
	s.Assert(y.Eq(zero).Not())
	s.Assert(z.Eq(zero).Not())

	s.Assert(x.Add(y).Eq(ctx.FromInt(-3, z3int).(z3.Int)))

	sat, err := s.Check()
	if err != nil {
		log.Fatal(err)
	}
	if !sat {
		log.Fatal("Unsolvable")
	}

	m := s.Model()
	fmt.Printf("x: %v\n", m.Eval(x, true))
	fmt.Printf("y: %v\n", m.Eval(y, true))
	fmt.Printf("z: %v\n", m.Eval(z, true))
}
```
