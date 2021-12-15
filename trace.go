package rvsym

import (
	"time"

	"github.com/zyedidia/rvsym/pkg/smt"
)

type Action struct {
	addr uint32
	time time.Duration
}

type Trace map[Action]smt.Int32

func (t Trace) Add(addr uint32, time time.Duration, val smt.Int32) {
	t[Action{addr, time}] = val
}

func (t Trace) Eq(other Trace, s *smt.Solver) bool {
	s.Push()
	defer s.Pop()

	var asserted bool
	for k, v := range t {
		o, ok := other[k]
		if !ok {
			return false
		}
		if v.Concrete() && o.Concrete() {
			if v.C != o.C {
				return false
			}
			continue
		}

		s.Assert(v.NEqb(o, s))
		asserted = true
	}

	if !asserted {
		return true
	}
	r := s.Check(false)
	return r == smt.Unsat
}
