package rvsym

import (
	"bytes"
	"fmt"

	"github.com/zyedidia/go-z3/st"
	"github.com/zyedidia/go-z3/z3"
)

type Assignment struct {
	Name string
	Val  int32
}

type TestCase struct {
	Assignments []Assignment
	Exit        ExitStatus
	Addr        int32
	Err         error // possible error this test case causes
}

func (tc TestCase) String(hex bool) string {
	buf := &bytes.Buffer{}
	if tc.Err != nil {
		buf.WriteString(tc.Err.Error())
		buf.WriteByte('\n')
	}
	for _, a := range tc.Assignments {
		if hex {
			buf.WriteString(fmt.Sprintf("%s -> 0x%x\n", a.Name, uint32(a.Val)))
		} else {
			buf.WriteString(fmt.Sprintf("%s -> %d\n", a.Name, a.Val))
		}
	}
	return buf.String()
}

type Engine struct {
	insns       []uint32
	checkpoints []*Checkpoint
	active      *Machine

	paths []TestCase

	Stats       Stats
	MaxMachines int

	solver *z3.Solver
	ctx    *z3.Context
}

type Stats struct {
	Exits map[ExitStatus]int
	Steps int
}

func NewEngine(insns []uint32) *Engine {
	ctx := z3.NewContext(nil)
	sol := z3.NewSolver(ctx)
	mem := NewMemory(nil)

	for i, ins := range insns {
		mem.Write32(st.Uint32{C: uint32(i * 4)}, st.Int32{C: int32(ins)}, sol)
	}

	return &Engine{
		insns:       insns,
		ctx:         ctx,
		solver:      sol,
		active:      NewMachine(ctx, sol, 0, mem),
		MaxMachines: -1,
		Stats: Stats{
			Exits: make(map[ExitStatus]int),
		},
	}
}

func (e *Engine) Step() bool {
	m := e.active

	m.Exec(e.insns[m.pc/4])
	if m.Status.Err != nil {
		m.Status.Exit = ExitFail
	}

	if e.HandleExit(m) {
		return e.active != nil
	} else if m.Status.HasBr {
		br := m.Status.Br
		if br.cond.IsConcrete() {
			if br.cond.C {
				m.pc = br.pc
			} else {
				m.pc += 4
			}
		} else {
			nobr := br.cond.S.Not()

			e.solver.Push()
			e.solver.Assert(nobr)
			sat, err := e.solver.Check()
			e.solver.Pop()
			if sat || err != nil {
				checkpoint := m.Checkpoint(nobr)
				checkpoint.pc += 4
				e.checkpoints = append(e.checkpoints, checkpoint)
				e.solver.Push()
			}

			m.pc = br.pc
			m.AddCond(br.cond.S, true)
		}

		if e.HandleExit(m) {
			return e.active != nil
		}
		m.Status.ClearBranch()
	} else {
		m.pc += 4
	}
	e.Stats.Steps++

	return e.active != nil
}

func (e *Engine) HandleExit(m *Machine) bool {
	if e.HasExit(m) {
		e.active = nil
		if len(e.checkpoints) > 0 {
			e.solver.Pop()
			e.active = Restore(e.checkpoints[len(e.checkpoints)-1], e.ctx, e.solver)
			e.checkpoints = e.checkpoints[:len(e.checkpoints)-1]
		}
		e.Stats.Exits[m.Status.Exit]++
		return true
	}
	return false
}

func (e *Engine) HasExit(m *Machine) bool {
	if m.Status.Exit != ExitNone {
		switch m.Status.Exit {
		case ExitNormal:
			tc, err := m.TestCase()
			if err != ErrUnsat {
				e.paths = append(e.paths, tc)
			} else {
				m.Status.Exit = ExitUnsat
			}
		case ExitFail:
			tc, err := m.TestCase()
			if err != ErrUnsat {
				e.paths = append(e.paths, tc)
				fmt.Println("INFO: found failure")
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
	return e.paths
}

func (e *Engine) Summary() string {
	buf := &bytes.Buffer{}
	paths := 0
	for _, v := range e.Stats.Exits {
		paths += v
	}
	fmt.Fprintln(buf, "--- Summary ---")
	fmt.Fprintf(buf, "Instructions executed: %d\n", e.Stats.Steps)
	fmt.Fprintf(buf, "Total paths: %d\n", paths)
	fmt.Fprintf(buf, "Quiet exits: %d\n", e.Stats.Exits[ExitQuiet])
	fmt.Fprintf(buf, "Unsat exits: %d\n", e.Stats.Exits[ExitUnsat])
	fmt.Fprintf(buf, "Normal exits: %d\n", e.Stats.Exits[ExitNormal])
	fmt.Fprintf(buf, "Failures: %d\n", e.Stats.Exits[ExitFail])
	fmt.Fprintln(buf, "---")
	return buf.String()
}
