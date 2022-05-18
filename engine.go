package rvsym

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/zyedidia/rvsym/pkg/smt"
)

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
		words := toWords(seg.data)
		for i, word := range words {
			addr := int32(seg.addr + uint32(i*4))
			mem.WriteWord(smt.Int32{C: addr}, smt.Int32{C: int32(word)}, s)
		}
	}

	machine := NewMachine(int32(entry), mem)

	if len(segs) > 0 {
		machine.icache = cache{
			base: segs[0].addr,
			data: segs[0].data,
		}
	}

	switch mode {
	case EmuLinux:
		// linux mode: use an initial stack pointer and point it at argc (0).
		sp := int32(0x7ffff00)
		machine.regs[Rsp] = smt.Int32{C: sp}
		mem.WriteWord(smt.Int32{C: sp}, smt.Int32{C: 0}, s)
	case EmuUnderConstrained:
		for i := range machine.regs {
			if i == Rzero || i == Rsp {
				continue
			}
			sym := s.AnyInt32()
			// machine.markSym(sym, RegNames[i])
			machine.regs[i] = sym
		}
		sp := int32(0x7ffff00)
		machine.regs[Rsp] = smt.Int32{C: sp}
	}
	machine.mode = mode

	return &Engine{
		active: machine,
		smt:    s,
		Stats: Stats{
			Exits: make(map[ExitStatus]int),
		},
	}
}

func (e *Engine) Step() bool {
	e.Stats.Steps++
	m := e.active

	isz := m.Exec(e.smt)

	exited := e.handleExit(m)
	if exited {
		return e.active != nil
	}

	br := m.Status.Br
	switch {
	case m.Status.HasBr && br.cond.Concrete() && br.cond.C:
		m.pc = br.pc
	case m.Status.HasBr && br.cond.Concrete():
		m.pc += isz
	case m.Status.HasBr:
		var cond, alt smt.Bool
		var condpc, altpc int32

		if randbool() {
			cond, alt = br.cond, br.cond.Not(e.smt)
			condpc, altpc = br.pc, m.pc+isz
		} else {
			alt, cond = br.cond, br.cond.Not(e.smt)
			altpc, condpc = br.pc, m.pc+isz
		}

		e.smt.Push()
		e.smt.Assert(alt)
		res := e.smt.Check(false)
		if res == smt.Sat {
			// alt is sat, so we take that branch and checkpoint cond to return
			// to later
			m.pc = altpc

			checkpoint := m.Checkpoint(cond)
			checkpoint.pc = condpc
			e.checkpoints = append(e.checkpoints, checkpoint)
		} else {
			// alt was unsat so we go directly to cond
			e.smt.Pop()
			m.pc = condpc
			m.AddCond(cond, true, e.smt)
		}
	default:
		m.pc += isz
	}

	if !e.handleExit(m) {
		m.clearbr()
	}

	return e.active != nil
}

func (e *Engine) handleExit(m *Machine) bool {
	if !e.hasExit(m) {
		return false
	}

	if len(e.checkpoints) > 0 {
		e.smt.Pop()
		e.active = Restore(e.checkpoints[len(e.checkpoints)-1], e.smt)
		e.checkpoints = e.checkpoints[:len(e.checkpoints)-1]
	} else {
		e.active = nil
	}
	e.Stats.Exits[m.Status.Exit]++

	return true
}

func (e *Engine) hasExit(m *Machine) bool {
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

func toWords(data []byte) []uint32 {
	words := make([]uint32, 0, len(data)/4)
	for len(data) > 0 {
		if len(data) >= 4 {
			words = append(words, binary.LittleEndian.Uint32(data))
			data = data[4:]
		} else {
			if len(data) == 3 {
				words = append(words, uint32(data[0])|uint32(data[1])<<8|uint32(data[2])<<16)
			} else if len(data) == 2 {
				words = append(words, uint32(data[0])|uint32(data[1])<<8)
			} else {
				words = append(words, uint32(data[0]))
			}
			data = data[len(data):]
		}
	}
	return words
}
