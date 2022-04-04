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
	}
	return nil
}

func (m *Machine) Checkpoint(cond smt.Bool) *Checkpoint {
	cp := &Checkpoint{
		mstate: mstate{
			regs:    make([]smt.Int32, len(m.regs)),
			mem:     m.mem.Copy(),
			symvals: make([]SymVal, len(m.symvals)),
			pc:      m.pc,
			sys:     m.sys.Copy(),
		},
		cond: cond,
	}

	copy(cp.regs, m.regs)
	copy(cp.symvals, m.symvals)

	return cp
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
	return TestCase{
		Assignments: vars,
		Pc:          m.pc,
		Exit:        m.Status.Exit,
		Err:         m.Status.Err,
	}, true
}
