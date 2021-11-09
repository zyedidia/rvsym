package rvsym

import (
	"bytes"
	"fmt"

	"github.com/zyedidia/go-z3/st"
	"github.com/zyedidia/go-z3/z3"
)

type Engine struct {
	insns    []uint32
	machines []*Machine

	ctx *z3.Context
}

func NewEngine(insns []uint32) *Engine {
	ctx := z3.NewContext(nil)
	mem := make(Memory)

	for i, ins := range insns {
		mem.Write32(uint32(i*4), st.Int32{C: int32(ins)})
	}

	return &Engine{
		insns:    insns,
		ctx:      ctx,
		machines: []*Machine{NewMachine(ctx, 0, mem)},
	}
}

type Exit struct {
	Pc       int32
	Universe int
	Status   ExitStatus
}

func (e *Engine) Step() ([]Exit, bool) {
	var exits []Exit

	nmach := len(e.machines)
	done := 0
	for i := 0; i < nmach; i++ {
		m := e.machines[i]

		if m.done {
			done++
			continue
		}

		br, ok, exit := m.Exec(e.insns[m.pc/4])
		if exit != ExitNone {
			exits = append(exits, Exit{Pc: m.pc, Universe: i, Status: exit})
		}
		if ok && br.cond.IsConcrete() {
			if br.cond.C {
				m.pc = br.pc
			} else {
				m.pc += 4
			}
		} else if ok {
			copied := m.Copy()
			copied.pc += 4
			copied.AddCond(br.cond.S.Not())
			e.machines = append(e.machines, copied)

			m.pc = br.pc
			m.AddCond(br.cond.S)
		} else {
			m.pc += 4
		}
	}
	return exits, done == len(e.machines)
}

func (e *Engine) Context() *z3.Context {
	return e.ctx
}

func (e *Engine) NumUniverses() int {
	return len(e.machines)
}

type TestVal struct {
	Name  string
	Value int32
}

type TestCase []TestVal

func (t TestCase) String() string {
	buf := &bytes.Buffer{}
	for _, val := range t {
		buf.WriteString(fmt.Sprintf("%s -> %d\n", val.Name, val.Value))
	}
	return buf.String()
}

func (e *Engine) UniverseInput(n int) TestCase {
	m := e.machines[n]

	s, err := m.Solver()
	if err != nil {
		return nil
	}
	model := s.Model()

	testcase := make(TestCase, 0)
	for i := range m.regs {
		if !m.regs[i].IsConcrete() {
			testcase = append(testcase, TestVal{
				Name:  fmt.Sprintf("x%d", i),
				Value: m.regs[i].Eval(model),
			})
		}
	}

	for addr, val := range m.mem {
		if !val.IsConcrete() {
			testcase = append(testcase, TestVal{
				Name:  fmt.Sprintf("0x%x", addr),
				Value: val.Eval(model),
			})
		}
	}

	return testcase
}
