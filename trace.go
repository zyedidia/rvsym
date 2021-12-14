package rvsym

import (
	"time"

	"github.com/zyedidia/rvsym/pkg/smt"
)

type Action struct {
	addr uint32
	val  smt.Int32
	time time.Duration
}

type Trace []Action

func (t *Trace) Append(a Action) {
	*t = append(*t, a)
}

func (t Trace) Eq(other Trace) bool {
	for i, a := range t {
		if a != other[i] {
			return false
		}
	}
	return true
}
