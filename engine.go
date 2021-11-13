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

func (tc TestCase) String() string {
	buf := &bytes.Buffer{}
	if tc.Err != nil {
		buf.WriteString(tc.Err.Error())
		buf.WriteByte('\n')
	}
	for _, a := range tc.Assignments {
		buf.WriteString(fmt.Sprintf("%s -> %v\n", a.Name, a.Val))
	}
	return buf.String()
}

type Engine struct {
	insns    []uint32
	machines []*Machine

	paths []TestCase

	Stats       Stats
	MaxMachines int

	ctx *z3.Context
}

type Stats struct {
	Exits map[ExitStatus]int
}

func NewEngine(insns []uint32) *Engine {
	ctx := z3.NewContext(nil)
	mem := make(Memory)

	for i, ins := range insns {
		mem.Write32(uint32(i*4), st.Int32{C: int32(ins)})
	}

	return &Engine{
		insns:       insns,
		ctx:         ctx,
		machines:    []*Machine{NewMachine(ctx, 0, mem)},
		MaxMachines: -1,
		Stats: Stats{
			Exits: make(map[ExitStatus]int),
		},
	}
}

func (e *Engine) Step() bool {
	for i := 0; i < len(e.machines); {
		m := e.machines[i]

		m.Exec(e.insns[m.pc/4])
		if m.Status.Err != nil {
			m.Status.Exit = ExitFail
		}

		if e.HandleExit(i) {
			continue
		} else if m.Status.HasBr {
			br := m.Status.Br
			if br.cond.IsConcrete() {
				if br.cond.C {
					m.pc = br.pc
				} else {
					m.pc += 4
				}
			} else {
				if e.MaxMachines == -1 || len(e.machines) < e.MaxMachines {
					copied := m.Copy()
					copied.pc += 4
					copied.AddCond(br.cond.S.Not(), true)
					if !e.HasExit(copied) {
						e.machines = append(e.machines, copied)
					}
				}

				m.pc = br.pc
				m.AddCond(br.cond.S, true)

				if e.HandleExit(i) {
					continue
				}
			}
			m.Status.ClearBranch()
		} else {
			m.pc += 4
		}
		i++
	}

	return len(e.machines) != 0
}

func (e *Engine) HandleExit(machidx int) bool {
	m := e.machines[machidx]

	if e.HasExit(m) {
		e.machines[machidx] = e.machines[len(e.machines)-1]
		e.machines[len(e.machines)-1] = nil
		e.machines = e.machines[:len(e.machines)-1]

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
			}
		case ExitFail:
			tc, err := m.TestCase()
			if err != ErrUnsat {
				e.paths = append(e.paths, tc)
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
