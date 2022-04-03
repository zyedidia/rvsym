package rvsym

import "github.com/zyedidia/rvsym/pkg/smt"

type Engine struct {
	active      *Machine
	checkpoints []*Checkpoint
	smt         *smt.Solver

	tcs []TestCase

	Stats Stats
}

type Stats struct {
	Exits map[ExitStatus]int
	Steps int
	Forks int
}

func NewEngine(segs []Segment, entry uint32, mode EmuMode) *Engine {
	s := smt.NewSolver()
	mem := NewMemory()

	for _, seg := range segs {
		// TODO
	}

	machine := NewMachine(int32(entry), mem)

	if mode == EmuLinux {
		sp := int32(0x7ffff00)
		machine.regs[2] = smt.Int32{C: sp}
		mem.WriteWord(smt.Int32{C: sp}, smt.Int32{C: 0}, s)
	}

	return &Engine{
		active: machine,
		smt:    s,
		Stats: Stats{
			Exits: make(map[ExitStatus]int),
		},
	}
}

func (e *Engine) Step() bool {
	return false
}

func (e *Engine) handleExit(m *Machine) bool {
	return false
}

func (e *Engine) hasExit(m *Machine) bool {
	return false
}

func (e *Engine) TestCases() []TestCase {
	return nil
}

func (e *Engine) NumTestCases() int {
	return 0
}

func (e *Engine) Summary() string {
	return ""
}
