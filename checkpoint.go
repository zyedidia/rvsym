package rvsym

import (
	"fmt"

	"github.com/zyedidia/rvsym/pkg/smt"
)

type Checkpoint struct {
	mstate

	cond smt.Bool
}

func Restore(cp *Checkpoint, s *smt.Solver) *Machine {
	s.Assert(cp.cond)
	return &Machine{
		mstate: cp.mstate,
		icache: cache{
			data: make([]byte, 0, 1024),
		},
	}
	return nil
}

func (m *Machine) Checkpoint(cond smt.Bool) *Checkpoint {
	return nil
}

func (m *Machine) AddCond(cond smt.Bool, checksat bool, s *smt.Solver) {
	s.Assert(cond)

	if checksat {
		res := s.Check(false)
		if res == smt.Unsat {
			m.exit(ExitUnsat)
		} else if res == smt.Unknown {
			m.err(fmt.Errorf("smt solver returned unknown"))
		}
	}
}

func (m *Machine) TestCase(s *smt.Solver) (TestCase, bool) {
	res := s.Check(true)
	if res != smt.Sat {
		return TestCase{}, false
	}
	model := s.Model()
	vars := make([]Assignment, len(m.symvals))
	for i, v := range m.symvals {
		vars[i] = Assignment{
			Name: v.name,
			Val:  model.Eval(v.val),
		}
	}
	dump := s.String()
	return TestCase{
		Assignments: vars,
		Pc:          m.pc,
		Exit:        m.Status.Exit,
		Err:         m.Status.Err,
		Dump:        dump,
	}, true
}
