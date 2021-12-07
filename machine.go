package rvsym

import (
	"bytes"
	"fmt"

	"github.com/zyedidia/rvsym/pkg/z3/st"
	"github.com/zyedidia/rvsym/pkg/z3/z3"
)

// A Machine stores the state for one execution path of the program.
type Machine struct {
	pc   int32
	regs []st.Int32
	mem  *Memory

	ctx    *z3.Context
	solver *z3.Solver

	Status Status

	marked []Mark

	stores []st.Int32
	secretStores []st.Int32

	foundAddrs []st.Int32

}

type Checkpoint struct {
	pc     int32
	regs   []st.Int32
	mem    *Memory
	marked []Mark

	cond z3.Bool
}

type Mark struct {
	val  st.Int32
	name string
}

type Secret struct {
	val st.Int32
	name string
}

var secretCounter int = 0
var secretMax int = 1

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
func NewMachine(ctx *z3.Context, solver *z3.Solver, pc int32, mem *Memory) *Machine {
	regs := make([]st.Int32, 32)
	for i := range regs {
		regs[i] = st.Int32{C: 0}
	}

	return &Machine{
		pc:     pc,
		regs:   regs,
		mem:    mem,
		ctx:    ctx,
		solver: solver,
	}
}

func Restore(cp *Checkpoint, ctx *z3.Context, solver *z3.Solver) *Machine {
	solver.Assert(cp.cond)
	return &Machine{
		pc:     cp.pc,
		regs:   cp.regs,
		mem:    cp.mem,
		marked: cp.marked,
		ctx:    ctx,
		solver: solver,
	}
}

func (m *Machine) Checkpoint(cond z3.Bool) *Checkpoint {
	cpregs := make([]st.Int32, len(m.regs))
	cpmem := NewMemory(m.mem)
	cpmarked := make([]Mark, len(m.marked))
	m.mem = NewMemory(m.mem)

	copy(cpregs, m.regs)
	copy(cpmarked, m.marked)

	return &Checkpoint{
		cond:   cond,
		regs:   cpregs,
		mem:    cpmem,
		marked: cpmarked,
		pc:     m.pc,
	}
}

// AddCond adds a condition to the list of path constraints of this machine.
func (m *Machine) AddCond(cond z3.Bool, checksat bool) {
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

func check(s *z3.Solver) error {
	sat, err := s.Check()
	if err != nil {
		return err
	} else if !sat {
		return ErrUnsat
	}
	return nil
}

// Generate a solver
func (m *Machine) Solver() (*z3.Solver, error) {
	return m.solver, check(m.solver)
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
	fmt.Printf("sysnum: %d\n", sysnum);
	switch sysnum {
	case SymSymbolicRegs:
		for i := range m.regs[1:] {
			name := fmt.Sprintf("x%d", i+1)
			m.regs[i+1] = st.AnyInt32(m.ctx, name)
			m.mark(m.regs[i+1], name)
		}
	case SymPrint:
		sysarg := m.regs[11] // a1
		if !sysarg.IsConcrete() {
			m.Status.Err = fmt.Errorf("required symcall argument is symbolic")
			return
		}
		fmt.Println(sysarg)
	case SymFail:
		m.exit(ExitFail)
	case SymExit:
		m.exit(ExitNormal)
	case SymQuietExit:
		m.exit(ExitQuiet)
	case SymMarkArray:
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
			b, ok := m.mem.Read8u(st.Uint32{C: uint32(nameptr.C + i)}, m.solver)
			if !ok || !b.IsConcrete() {
				m.Status.Err = fmt.Errorf("out of bounds name while marking array")
				return
			}
			if b.C == 0 {
				break
			}
			nameb.WriteByte(byte(b.C))
		}
		name := nameb.String()

		fmt.Printf("INFO: marking '%s' as an array: %d bytes at 0x%x\n", name, nbytes.C, ptr.C)

		m.mem.AddArray(m.ctx, name, int(ptr.C/4), int((nbytes.C+3)/4))
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
			b, ok := m.mem.Read8u(st.Uint32{C: uint32(nameptr.C + i)}, m.solver)
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

		fmt.Printf("INFO: marking '%s': %d bytes at 0x%x\n", name, nbytes.C, ptr.C)

		for i := int32(0); i < nbytes.C/4; i++ {
			idx := i * 4
			i32 := st.AnyInt32(m.ctx, fmt.Sprintf("0x%x", ptr.C+idx))
			m.mark(i32, fmt.Sprintf("%s[%d:%d]", name, idx, idx+3))
			m.mem.Write32(st.Uint32{C: uint32(ptr.C + idx)}, i32, m.solver)
		}
		left := nbytes.C % 4
		for i := int32(0); i < left; i++ {
			idx := nbytes.C - left + i
			i32 := st.AnyInt32(m.ctx, fmt.Sprintf("0x%x", ptr.C+idx))
			m.mark(i32, fmt.Sprintf("%s[%d]", name, idx))
			m.mem.Write8(st.Uint32{C: uint32(ptr.C + idx)}, i32, m.solver)
		}
	case SymDump:
		fmt.Print(m.String())
	case SymMarkNSecret:
		fmt.Printf("Secret data marked\n")
		ptr := m.regs[11]     // a1
		nbytes := m.regs[12]  // a2
		nameptr := m.regs[13] // a3

		if !ptr.IsConcrete() {
			m.Status.Err = fmt.Errorf("secret address is symbolic")
			return
		}
		if !nbytes.IsConcrete() {
			m.Status.Err = fmt.Errorf("secret size is symbolic")
			return
		}
		if !nameptr.IsConcrete() {
			m.Status.Err = fmt.Errorf("secret name address is symbolic")
			return
		}
		nameb := &bytes.Buffer{}
		for i := int32(0); ; i++ {
			b, ok := m.mem.Read8u(st.Uint32{C: uint32(nameptr.C + i)}, m.solver)
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

		fmt.Printf("INFO: marking %d bytes %s secret at 0x%x\n", nbytes.C, name, ptr.C)
		for i := int32(0); i < nbytes.C/4; i++ {
			idx := i * 4
			i32, _ := m.mem.Read32(st.Uint32{C: uint32(ptr.C+idx)}, m.solver)
			i32.Secret = true
			m.mem.Write32(st.Uint32{C: uint32(ptr.C+idx)}, i32, m.solver)

			//m.secret(i32, fmt.Sprintf("%s[%d:%d]", name, idx, idx+3))

			if i32.IsConcrete() {
				fmt.Printf("Secret data: %d\n", i32.C)
			}
		}
		left := nbytes.C % 4
		for i := int32(0); i < left; i++ {
			idx := nbytes.C - left + i
			i32, _ := m.mem.Read8(st.Uint32{C: uint32(ptr.C+idx)}, m.solver)
			i32.Secret = true
			m.mem.Write8(st.Uint32{C: uint32(ptr.C+idx)}, i32, m.solver)
			//m.secret(i32, fmt.Sprintf("%s[%d]", name, idx))
		}
	}
	fmt.Printf("Exiting symcall\n");
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
	rd, rs1, shamt := extractRegs(insn)
	var imm st.Int32
	if aluop == AluShr {
		imm = st.Int32{C: int32(shamt)}
	} else {
		imm = st.Int32{C: int32(extractImm(insn, ImmI))}
	}

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

	/*
	rsval := m.regs[rs1]

	if !rsval.IsConcrete() {
		fmt.Printf("concretizing load\n")
		if c, ok := m.concretize(rsval); ok {
			rsval = st.Int32{C: c}
		} else {
			return
		}
	}
	*/
	//addr := (st.Int32{C: (rsval.C) + int32(imm)}).ToUint32()
	addr := m.regs[rs1].Add(st.Int32{C: int32(imm)}).ToUint32()
	//cAddr, _ := m.concretize(addr.ToInt32())
	//addr = st.Uint32{C: uint32(cAddr)}

	//fmt.Print("CONCRETE load at ")
	//fmt.Printf("%x\n", addr.C)


	var rdval st.Int32
	var valid bool
	switch funct3 {
	case ExtByte:
		//fmt.Printf("load -- Read8\n")
		rdval, valid = m.mem.Read8(addr, m.solver)
	case ExtHalf:
		rdval, valid = m.mem.Read16(addr, m.solver)
		//fmt.Printf("load -- Read16\n")
	case ExtWord:
		rdval, valid = m.mem.Read32(addr, m.solver)
		//fmt.Printf("load -- Read32\n")
	case ExtByteU:
		rdval, valid = m.mem.Read8u(addr, m.solver)
		//fmt.Printf("load -- Read8u\n")
	case ExtHalfU:
		rdval, valid = m.mem.Read16u(addr, m.solver)
		//fmt.Printf("load -- Read16u\n")
	default:
		m.Status.Err = fmt.Errorf("invalid load instruction")
		return
	}

	if !valid {
		if !addr.IsConcrete() {
			m.Status.Err = fmt.Errorf("invalid memory access possible at symbolic memory address")
		} else {
			m.Status.Err = fmt.Errorf("invalid memory access at 0x%x", addr.C)
		}
		m.Status.Exit = ExitMem
		return
	}

	m.WriteReg(rd, rdval)
}

func (m *Machine) store(insn uint32) {
	_, rs1, rs2 := extractRegs(insn)
	imm := extractImm(insn, ImmS)
	funct3 := GetBits(insn, 14, 12).Uint32()

	//fmt.Printf("rs1: %d, rs2: %d\n", rs1, rs2);
	//addr := m.regs[rs1].Add(st.Int32{C: int32(imm)}).ToUint32()

	rsval := m.regs[rs1]
	stval := m.regs[rs2]


	addrsym := rsval.Add(st.Int32{C: int32(imm)})

	/*
	fmt.Print("INFO: store at ")
	if addrsym.IsConcrete() {
		fmt.Printf("%x\n", addrsym.C)
	} else {
		fmt.Printf("%v\n", addrsym.S)
	}
	*/

	/*
	if stval.Secret {
		fmt.Printf("NOTE: rsval is secret!\n")
	}
	if stval.IsConcrete() {
		fmt.Printf("Storing data: %d\n", stval.C)
	} else {
		fmt.Printf("Store is symbolic\n")
	}
	*/

	loopLen := 0
	if stval.Secret {
		loopLen = len(m.stores)
	} else if !stval.Secret && !stval.IsConcrete() { // Only do silent stores
		loopLen = len(m.secretStores)	         // check if storing secret value, 
	}					         // or if storing non-secret, 
						         // symbolic value (aka attacker
						         // controlled data)

	for i := loopLen - 1; i >= 0; i-- {

		var a st.Int32
		if stval.Secret {
			a = m.stores[i]
			//fmt.Printf("storing secret, checking symbolic addresses\n")
		} else {
			a = m.secretStores[i]
			//fmt.Printf("storing symbolic, checking secret addresses\n")
		}

		//fmt.Printf("Checking secret addresses\n")

		cond := a.Eq(addrsym)

		if !cond.IsConcrete() {
			m.solver.Push()
			m.solver.Assert(cond.S)
		}
		if !cond.IsConcrete() || cond.C {
			sat, err := m.solver.Check()
			if sat && err == nil {
				found := false
				for k := 0; k < len(m.foundAddrs); k++ {
					addrFound := m.foundAddrs[k]
					condition := addrsym.Eq(addrFound)

					if !condition.IsConcrete() {
						continue
					}
					if condition.C {
						found = true
					}
				}

				if !found {
					m.foundAddrs = append(m.foundAddrs, addrsym)
				model := m.solver.Model()
				sameaddr := addrsym.Eval(model)

				fmt.Printf("Stores to same addresses %x\n", sameaddr)

				secretCounter++
				if secretCounter == secretMax {

					m.Status.Err = fmt.Errorf("Stores to same addresses %x\n", sameaddr)
					m.Status.Exit = ExitFail
					return
				}
				}
			}
		}
		if !cond.IsConcrete() {
			m.solver.Pop()
		}
	}

	addr := m.regs[rs1].Add(st.Int32{C: int32(imm)}).ToUint32()


	if stval.Secret { // secretStores and stores are disjoint
			  // because we don't want secret overwriting
			  // secret stores to be an error
		//addrsym.Secret = true
		m.secretStores = append(m.secretStores, addrsym)
		//m.secretStores = append(m.secretStores, addr)
	} else if !stval.Secret && !stval.IsConcrete() {
		m.stores = append(m.stores, addrsym)
		//m.stores = append(m.stores, addr)
	}




	if !stval.Secret && stval.IsConcrete() {
		for j := 0; j < len(m.secretStores); j++ {
			addrSecret := m.secretStores[j]
			if addrSecret.IsConcrete() {
				if addrSecret.C == addrsym.C {
					m.secretStores = append(m.secretStores[:j], m.secretStores[j+1:]...)
					//fmt.Print("Removed secret address\n")
					// Remove secret stores overwritten by non-secret, non-attacker controlled data
				}
			}
		}
	}

	if !stval.IsConcrete() {
		for j := 0; j < len(m.stores); j++ {
			addrSymbolic := m.stores[j]
			if addrSymbolic.IsConcrete() {
				if addrSymbolic.C == addrsym.C {
					m.stores = append(m.stores[:j], m.stores[j+1:]...)
					//fmt.Print("Removed symbolic address\n")
					// Remove secret stores overwritten by non-secret, non-attacker controlled data
				}
			}
		}
	}

	/*
	if !rsval.IsConcrete() {
		if c, ok := m.concretize(rsval); !ok {
			return
		} else {
			rsval = st.Int32{C: c}
		}
	}
	*/

	//addr := (st.Int32{C: (rsval.C) + int32(imm)}).ToUint32()

	//addr := m.regs[rs1].Add(st.Int32{C: int32(imm)}).ToUint32()
	//cAddr, _ := m.concretize(addr.ToInt32())
	//addr = st.Uint32{C: uint32(cAddr)}

	//fmt.Print("CONCRETE store at ")
	//fmt.Printf("%x\n", addr.C)

	switch funct3 {
	case ExtByte:
		m.mem.Write8(addr, stval, m.solver)
	case ExtHalf:
		m.mem.Write16(addr, stval, m.solver)
	case ExtWord:
		m.mem.Write32(addr, stval, m.solver)
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

func concretize(val st.Int32, s *z3.Solver) (int32, error) {
	if val.IsConcrete() {
		return val.C, nil
	}

	err := check(s)
	if err != nil {
		return 0, err
	}
	model := s.Model()
	concrete := val.Eval(model)
	s.Assert(val.Eq(st.Int32{C: concrete}).S)
	fmt.Printf("INFO: concretizing value -> %d\n", concrete)
	return concrete, nil
}

func (m *Machine) concretize(val st.Int32) (int32, bool) {
	concrete, err := concretize(val, m.solver)
	if err == ErrUnsat {
		m.exit(ExitUnsat)
	} else if err != nil {
		m.Status.Err = err
		m.exit(ExitFail)
	} else {
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

	keys := m.mem.Keys()

	for _, addr := range keys {
		val := m.mem.readz(st.Uint32{C: addr}, m.solver)
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
