package rvsym

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/zyedidia/go-z3/st"
	"github.com/zyedidia/go-z3/z3"
)

// A Machine stores the state for one execution path of the program.
type Machine struct {
	pc    int32
	regs  []st.Int32
	conds []z3.Bool
	mem   Memory

	ctx    *z3.Context
	solver *z3.Solver

	Status Status

	marked []Mark
}

type Mark struct {
	val  st.Int32
	name string
}

type Status struct {
	Err   error
	Br    Branch
	HasBr bool
	Exit  ExitStatus
}

func (s *Status) SetBranch(br Branch) {
	s.Br = br
	s.HasBr = true
}

func (s *Status) ClearBranch() {
	s.HasBr = false
}

// NewMachine returns a new machine, with the given initial memory and program
// counter.
func NewMachine(ctx *z3.Context, pc int32, mem Memory) *Machine {
	regs := make([]st.Int32, 32)
	for i := range regs {
		regs[i] = st.Int32{C: 0}
	}

	return &Machine{
		pc:     pc,
		regs:   regs,
		mem:    mem,
		ctx:    ctx,
		solver: z3.NewSolver(ctx),
	}
}

// Copy this machine.
func (m *Machine) Copy() *Machine {
	regs := make([]st.Int32, len(m.regs))
	conds := make([]z3.Bool, len(m.conds))
	mem := make(map[uint32]st.Int32)
	marked := make([]Mark, len(m.marked))

	copy(regs, m.regs)
	copy(conds, m.conds)
	copy(marked, m.marked)

	for k, v := range m.mem {
		mem[k] = v
	}

	solver := z3.NewSolver(m.ctx)
	for _, cond := range m.conds {
		solver.Assert(cond)
	}

	return &Machine{
		pc:     m.pc,
		regs:   regs,
		conds:  conds,
		mem:    mem,
		ctx:    m.ctx,
		marked: marked,
		solver: solver,
	}
}

// AddCond adds a condition to the list of path constraints of this machine.
func (m *Machine) AddCond(cond z3.Bool, checksat bool) {
	m.conds = append(m.conds, cond)

	m.solver.Assert(cond)

	if checksat {
		_, err := m.Solver()
		if err == ErrUnsat {
			m.exit(ExitUnsat)
		} else if err != nil {
			m.Status.Err = err
		}
	}
}

// WriteReg writes the given value to the given register index.
func (m *Machine) WriteReg(reg uint32, val st.Int32) {
	if reg == 0 {
		// cannot write to register 0
		return
	}
	m.regs[reg] = val
}

type SolverErr byte

var ErrUnsat SolverErr = 0

func (s SolverErr) Error() string {
	return "solver error: unsatisfiable formula"
}

// Generate a solver
func (m *Machine) Solver() (*z3.Solver, error) {
	s := m.solver
	sat, err := s.Check()
	if err != nil {
		return nil, err
	} else if !sat {
		return nil, ErrUnsat
	}
	return s, nil
}

// Branch: jumps to PC if cond is true.
type Branch struct {
	pc   int32
	cond st.Bool
}

func (m *Machine) Exec(insn uint32) {
	switch insn {
	case InsnNop:
		return
	case InsnEcall:
		symnum := m.regs[10] // a0
		if !symnum.IsConcrete() {
			m.Status.Err = fmt.Errorf("symcall number is symbolic")
		} else {
			m.symcall(insn, int(symnum.C))
		}
		return
	}

	op := GetBits(insn, 6, 0).Uint32()

	switch op {
	case OpRarith:
		m.rarith(insn)
	case OpIarith:
		m.iarith(insn)
	case OpBranch:
		m.branch(insn)
	case OpLui:
		m.lui(insn)
	case OpAuipc:
		m.auipc(insn)
	case OpJal:
		m.jal(insn)
	case OpJalr:
		m.jalr(insn)
	case OpLoad:
		m.load(insn)
	case OpStore:
		m.store(insn)
	}
}

func (m *Machine) mark(val st.Int32, name string) {
	m.marked = append(m.marked, Mark{val, name})
}

func (m *Machine) symcall(insn uint32, sysnum int) {
	switch sysnum {
	case SymSymbolicRegs:
		for i := range m.regs[1:] {
			name := fmt.Sprintf("x%d", i+1)
			m.regs[i+1] = st.AnyInt32(m.ctx, name)
			m.mark(m.regs[i+1], name)
		}
	case SymSymbolicReg:
		sysarg := m.regs[11] // a1
		if !sysarg.IsConcrete() {
			m.Status.Err = fmt.Errorf("required symcall argument is symbolic")
			return
		}
		name := fmt.Sprintf("x%d", sysarg.C)
		m.regs[sysarg.C] = st.AnyInt32(m.ctx, name)
		m.mark(m.regs[sysarg.C], name)
	case SymFail:
		m.exit(ExitFail)
	case SymExit:
		m.exit(ExitNormal)
	case SymQuietExit:
		m.exit(ExitQuiet)
	case SymMarkNBytes:
		ptr := m.regs[11]     // a1
		nbytes := m.regs[12]  // a2
		nameptr := m.regs[13] // a3

		if !ptr.IsConcrete() {
			m.Status.Err = fmt.Errorf("mark address is symbolic")
			return
		}
		if !nbytes.IsConcrete() {
			m.Status.Err = fmt.Errorf("mark size is symbolic")
			return
		}
		if !nameptr.IsConcrete() {
			m.Status.Err = fmt.Errorf("mark name address is symbolic")
			return
		}

		nameb := &bytes.Buffer{}
		for i := int32(0); ; i++ {
			b, ok := m.mem.Read8u(uint32(nameptr.C + i))
			if !ok || !b.IsConcrete() {
				m.Status.Err = fmt.Errorf("out of bounds name while marking bytes")
				return
			}
			if b.C == 0 {
				break
			}
			nameb.WriteByte(byte(b.C))
		}
		name := nameb.String()

		fmt.Printf("INFO: marking %d bytes at 0x%x\n", nbytes.C, ptr.C)

		for i := int32(0); i < nbytes.C/4; i++ {
			idx := i * 4
			i32 := st.AnyInt32(m.ctx, fmt.Sprintf("0x%x", ptr.C+idx))
			m.mark(i32, fmt.Sprintf("%s[%d:%d]", name, idx, idx+3))
			m.mem.Write32(uint32(ptr.C+idx), i32)
		}
		left := nbytes.C % 4
		for i := int32(0); i < left; i++ {
			idx := nbytes.C - left + i
			i32 := st.AnyInt32(m.ctx, fmt.Sprintf("0x%x", ptr.C+idx))
			m.mark(i32, fmt.Sprintf("%s[%d]", name, idx))
			m.mem.Write8(uint32(ptr.C+idx), i32)
		}
	case SymDump:
		fmt.Print(m.String())
	}
}

func (m *Machine) rarith(insn uint32) {
	aluop := GetBits(insn, 14, 12).Uint32() // funct3
	rd, rs1, rs2 := extractRegs(insn)
	funct7 := GetBits(insn, 31, 25).Uint32()

	modify := funct7 == 0b0100000
	muldiv := funct7 == 0b0000001

	m.WriteReg(rd, alu(m.regs[rs1], m.regs[rs2], aluop, modify, modify, muldiv))
}

func (m *Machine) iarith(insn uint32) {
	aluop := GetBits(insn, 14, 12).Uint32() // funct3
	modify := GetBits(insn, 30, 30).Uint32() != 0
	imm := st.Int32{C: int32(extractImm(insn, ImmI))}
	rd, rs1, _ := extractRegs(insn)

	m.WriteReg(rd, alu(m.regs[rs1], imm, aluop, false, modify, false))
}

func (m *Machine) branch(insn uint32) {
	funct3 := GetBits(insn, 14, 12).Uint32()
	_, rs1, rs2 := extractRegs(insn)
	imm := int32(extractImm(insn, ImmB))

	var aluop uint32
	switch funct3 {
	case 0b000, 0b001:
		aluop = AluXor
	case 0b100, 0b101:
		aluop = AluSlt
	case 0b110, 0b111:
		aluop = AluSltu
	}

	result := alu(m.regs[rs1], m.regs[rs2], aluop, false, false, false)

	var cond st.Bool
	switch funct3 {
	case 0b000, 0b101, 0b111:
		// branch if alu result is zero
		cond = result.Eq(st.Int32{C: 0})
	case 0b001, 0b100, 0b110:
		// branch if alu result is non-zero
		cond = result.NE(st.Int32{C: 0})
	}

	m.Status.SetBranch(Branch{
		pc:   m.pc + int32(imm),
		cond: cond,
	})
}

func (m *Machine) jal(insn uint32) {
	rd, _, _ := extractRegs(insn)
	imm := extractImm(insn, ImmJ)

	pc := m.pc + int32(imm)
	m.WriteReg(rd, st.Int32{C: m.pc + 4})

	m.Status.SetBranch(Branch{
		pc:   pc,
		cond: st.Bool{C: true},
	})
}

func (m *Machine) jalr(insn uint32) {
	rd, rs1, _ := extractRegs(insn)
	imm := extractImm(insn, ImmI)

	pc := m.regs[rs1].ToInt32().Add(st.Int32{C: int32(imm)})

	if !pc.IsConcrete() {
		if c, ok := m.concretize(pc); ok {
			pc = st.Int32{C: c}
		} else {
			return
		}
	}
	m.WriteReg(rd, st.Int32{C: m.pc + 4})

	m.Status.SetBranch(Branch{
		pc:   pc.C,
		cond: st.Bool{C: true},
	})
}

func (m *Machine) lui(insn uint32) {
	rd, _, _ := extractRegs(insn)
	imm := int32(extractImm(insn, ImmU))

	m.WriteReg(rd, st.Int32{C: imm})
}

func (m *Machine) auipc(insn uint32) {
	rd, _, _ := extractRegs(insn)
	imm := extractImm(insn, ImmU)

	m.WriteReg(rd, st.Int32{C: m.pc + int32(imm)})
}

func (m *Machine) load(insn uint32) {
	rd, rs1, _ := extractRegs(insn)
	imm := extractImm(insn, ImmI)
	funct3 := GetBits(insn, 14, 12).Uint32()

	rsval := m.regs[rs1]
	if !rsval.IsConcrete() {
		if c, ok := m.concretize(rsval); ok {
			rsval = st.Int32{C: c}
		} else {
			return
		}
	}
	addr := uint32(int32(rsval.C) + int32(imm))

	var rdval st.Int32
	var valid bool
	switch funct3 {
	case ExtByte:
		rdval, valid = m.mem.Read8(addr)
	case ExtHalf:
		rdval, valid = m.mem.Read16(addr)
	case ExtWord:
		rdval, valid = m.mem.Read32(addr)
	case ExtByteU:
		rdval, valid = m.mem.Read8u(addr)
	case ExtHalfU:
		rdval, valid = m.mem.Read16u(addr)
	default:
		m.Status.Err = fmt.Errorf("invalid load instruction")
		return
	}

	if !valid {
		m.Status.Err = fmt.Errorf("invalid memory access at 0x%x", addr)
		return
	}

	m.WriteReg(rd, rdval)
}

func (m *Machine) store(insn uint32) {
	_, rs1, rs2 := extractRegs(insn)
	imm := extractImm(insn, ImmS)
	funct3 := GetBits(insn, 14, 12).Uint32()

	rsval := m.regs[rs1]
	if !rsval.IsConcrete() {
		if c, ok := m.concretize(rsval); !ok {
			return
		} else {
			rsval = st.Int32{C: c}
		}
	}

	addr := uint32(int32(rsval.C) + int32(imm))

	stval := m.regs[rs2]

	switch funct3 {
	case ExtByte:
		m.mem.Write8(addr, stval)
	case ExtHalf:
		m.mem.Write16(addr, stval)
	case ExtWord:
		m.mem.Write32(addr, stval)
	default:
		m.Status.Err = fmt.Errorf("invalid store instruction")
	}
}

func extractRegs(insn uint32) (rd, rs1, rs2 uint32) {
	rd = GetBits(insn, 11, 7).Uint32()
	rs1 = GetBits(insn, 19, 15).Uint32()
	rs2 = GetBits(insn, 24, 20).Uint32()
	return rd, rs1, rs2
}

func (m *Machine) concretize(val st.Int32) (int32, bool) {
	if val.IsConcrete() {
		return val.C, true
	}

	s, err := m.Solver()
	if err == ErrUnsat {
		m.exit(ExitUnsat)
	} else if err != nil {
		m.Status.Err = err
	} else {
		model := s.Model()
		concrete := val.Eval(model)
		m.AddCond(val.Eq(st.Int32{C: concrete}).S, false)
		return concrete, true
	}
	return 0, false
}

func (m *Machine) exit(ex ExitStatus) {
	m.Status.Exit = ex
}

func (m *Machine) TestCase() (TestCase, error) {
	s, err := m.Solver()
	if err == ErrUnsat {
		return TestCase{}, err
	} else if err != nil {
		m.Status.Err = err
	}
	model := s.Model()
	vars := make([]Assignment, len(m.marked))
	for i, v := range m.marked {
		vars[i] = Assignment{
			Name: v.name,
			Val:  v.val.Eval(model),
		}
	}
	return TestCase{
		Assignments: vars,
		Addr:        m.pc,
		Exit:        m.Status.Exit,
		Err:         m.Status.Err,
	}, nil
}

func (m *Machine) String() string {
	buf := &bytes.Buffer{}
	for i, reg := range m.regs {
		buf.WriteString(fmt.Sprintf("x%x: ", i))
		if reg.IsConcrete() {
			buf.WriteString(fmt.Sprintf("%d", reg.C))
		} else {
			buf.WriteString(fmt.Sprintf("%v", reg.S))
		}
		buf.WriteByte('\n')
	}

	keys := make([]uint32, len(m.mem))
	i := 0
	for k := range m.mem {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, addr := range keys {
		val := m.mem[addr]
		buf.WriteString(fmt.Sprintf("0x%x: ", addr*4))
		if val.IsConcrete() {
			buf.WriteString(fmt.Sprintf("%d", val.C))
		} else {
			buf.WriteString(fmt.Sprintf("%v", val.S))
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}
