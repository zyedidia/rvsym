package minisym

import (
	"bytes"
	"fmt"

	"github.com/zyedidia/go-z3/z3"
)

type Engine struct {
	insns    []uint32
	machines []*Machine

	ctx *z3.Context
}

func NewEngine(insns []uint32) *Engine {
	ctx := z3.NewContext(nil)
	return &Engine{
		insns:    insns,
		ctx:      ctx,
		machines: []*Machine{NewMachine(ctx, 0)},
	}
}

func (e *Engine) Step() {
	nmach := len(e.machines)
	for i := 0; i < nmach; i++ {
		m := e.machines[i]
		br, ok := m.Exec(e.insns[m.pc/4])
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
}

func (e *Engine) Context() *z3.Context {
	return e.ctx
}

func (e *Engine) NumUniverses() int {
	return len(e.machines)
}

type RegMap []int32

func (r RegMap) WithName(name string) int32 {
	if reg, ok := RegNums[name]; ok {
		return r[reg]
	}
	panic(fmt.Sprintf("invalid register name %s", name))
}

func (r RegMap) String() string {
	buf := &bytes.Buffer{}
	for i, reg := range r {
		buf.WriteString(fmt.Sprintf("x%d: %d\n", i, reg))
	}
	return buf.String()
}

func (e *Engine) UniverseInput(n int) RegMap {
	m := e.machines[n]

	s := m.MustSolver()
	model := s.Model()

	regmap := make([]int32, 32)
	for i := range m.regs {
		regmap[i] = m.regs[i].Eval(model)
	}
	return regmap
}
