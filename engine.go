package rvsym

import (
	"bytes"
	"fmt"

	"github.com/zyedidia/rvsym/pkg/z3/st"
	"github.com/zyedidia/rvsym/pkg/z3/z3"
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

	Stats    Stats
	MaxForks int

	errs map[int32]bool

	solver *z3.Solver
	ctx    *z3.Context
}

type Stats struct {
	Exits map[ExitStatus]int
	Steps int
	Forks int
}

func NewEngine(insns []uint32, maxforks int) *Engine {
	ctx := z3.NewContext(nil)
	sol := z3.NewSolver(ctx)
	mem := NewMemory(nil)

	for i, ins := range insns {
		mem.Write32(st.Uint32{C: uint32(i * 4)}, st.Int32{C: int32(ins)}, sol)
	}

	return &Engine{
		insns:  insns,
		ctx:    ctx,
		solver: sol,
		active: NewMachine(ctx, sol, 0, mem),
		errs:   make(map[int32]bool),
		Stats: Stats{
			Exits: make(map[ExitStatus]int),
		},
		MaxForks: maxforks,
	}
}

func (e *Engine) Step() bool {
	e.Stats.Steps++
	m := e.active

	m.Exec(e.insns[m.pc/4])

	if e.HandleExit(m) {
	} else if m.Status.HasBr {
		br := m.Status.Br
		if br.cond.IsConcrete() {
			if br.cond.C {
				m.pc = br.pc
			} else {
				m.pc += 4
			}
		} else {
			var cond, alt z3.Bool
			var condpc, altpc int32
			if randbool() {
				cond, alt = br.cond.S, br.cond.S.Not()
				condpc, altpc = br.pc, m.pc+4
			} else {
				alt, cond = br.cond.S, br.cond.S.Not()
				altpc, condpc = br.pc, m.pc+4
			}

			if e.Stats.Forks < e.MaxForks {
				e.solver.Push()
				e.solver.Assert(alt)
				sat, err := e.solver.Check()
				e.solver.Pop()
				if sat || err != nil {
					checkpoint := m.Checkpoint(alt)
					checkpoint.pc = altpc
					e.checkpoints = append(e.checkpoints, checkpoint)
					e.solver.Push()
					e.Stats.Forks++
				}
			}

			m.pc = condpc
			m.AddCond(cond, true)
		}

		if !e.HandleExit(m) {
			m.Status.ClearBranch()
		}
	} else {
		m.pc += 4
	}

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
		if m.Status.Exit == ExitMem {
			if !e.errs[m.pc] {
				m.Status.Exit = ExitFail
				e.errs[m.pc] = true
			} else {
				m.Status.Exit = ExitQuiet
			}
		}

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
