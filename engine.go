package rvsym

import (
	"github.com/zyedidia/rvsym/pkg/smt"
)

type Engine struct {
	active      *Machine
	checkpoints []*Checkpoint
	smt         *smt.Solver

	tcs []TestCase
}

func NewEngine(segs []Segment, entrypc uint32) *Engine {
	s := smt.NewSolver()
	mem := NewMemory(nil)

	for _, seg := range segs {
		for i, word := range seg.data {
			addr := seg.addr + uint32(i*4)
			mem.Write32(smt.Int32{C: int32(addr)}, smt.Int32{C: int32(word)}, s)
		}
	}

	machine := NewMachine(int32(entrypc), mem)

	return &Engine{
		active: machine,
		smt:    s,
	}
}

func (e *Engine) Step() bool {
	m := e.active

	m.Exec(e.smt)

	exited := e.HandleExit(m)
	if exited {
		return e.active != nil
	}

	br := m.Status.Br
	switch {
	case m.Status.HasBr && br.cond.Concrete() && br.cond.C:
		m.pc = br.pc
	case m.Status.HasBr && br.cond.Concrete():
		m.pc += 4
	case m.Status.HasBr:
		var cond, alt smt.Bool
		var condpc, altpc int32

		cond, alt = br.cond, br.cond.Not(e.smt)
		condpc, altpc = br.pc, m.pc+4

		e.smt.Push()
		e.smt.Assert(alt)
		res := e.smt.Check()
		e.smt.Pop()
		if res == smt.Sat {
			checkpoint := m.Checkpoint(alt)
			checkpoint.pc = altpc
			e.checkpoints = append(e.checkpoints, checkpoint)
			e.smt.Push()
		}

		m.pc = condpc
		m.AddCond(cond, true, e.smt)
	default:
		m.pc += 4
	}

	if !e.HandleExit(m) {
		m.clearbr()
	}

	return e.active != nil
}

func (e *Engine) HandleExit(m *Machine) bool {
	if !e.HasExit(m) {
		return false
	}

	if len(e.checkpoints) > 0 {
		e.smt.Pop()
		e.active = Restore(e.checkpoints[len(e.checkpoints)-1], e.smt)
		e.checkpoints = e.checkpoints[:len(e.checkpoints)-1]
	} else {
		e.active = nil
	}

	return true
}

func (e *Engine) HasExit(m *Machine) bool {
	if m.Status.Exit != ExitNone {
		switch m.Status.Exit {
		case ExitNormal, ExitFail:
			tc, sat := m.TestCase(e.smt)
			if sat {
				e.tcs = append(e.tcs, tc)
			} else {
				m.Status.Exit = ExitUnsat
			}
		case ExitQuiet:
		}
		return true
	}
	return false
}

func (e *Engine) TestCases() []TestCase {
	return e.tcs
}

func (e *Engine) NumTestCases() int {
	return len(e.tcs)
}
