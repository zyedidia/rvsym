package rvsym

import (
	"fmt"

	"github.com/zyedidia/rvsym/pkg/smt"
)

var symcalls = map[int]EcallFn{
	0: (*Machine).SymPrint,
	1: (*Machine).SymFail,
	2: (*Machine).SymExit,
	3: (*Machine).SymQuietExit,
	4: (*Machine).SymMarkArray,
	5: (*Machine).SymMarkBytes,
}

func (m *Machine) SymPrint(s *smt.Solver) {
	arg := m.regs[11]
	if !arg.Concrete() {
		fmt.Println(arg.S)
	} else {
		fmt.Println(arg.C)
	}
}

func (m *Machine) SymFail(s *smt.Solver) {
	m.exit(ExitFail)
}

func (m *Machine) SymExit(s *smt.Solver) {
	m.exit(ExitNormal)
}

func (m *Machine) SymQuietExit(s *smt.Solver) {
	m.exit(ExitQuiet)
}

func (m *Machine) SymMarkArray(s *smt.Solver) {
	var ptr, nbytes int32
	var ok bool
	if ptr, ok = m.RegConc(11); !ok {
		m.err(fmt.Errorf("array address is symbolic"))
		return
	}
	if nbytes, ok = m.RegConc(12); !ok {
		m.err(fmt.Errorf("array size is symbolic"))
		return
	}
	// mark the region expanded out to the nearest word boundaries
	m.mem.MakeSymAddrRegion(uint32(ptr)/4, (uint32(nbytes)+3)/4, s)

	logger.Printf("Marking symbolic array: %d bytes at 0x%x\n", nbytes, ptr)
}

func (m *Machine) SymMarkBytes(s *smt.Solver) {
	var ptr, nbytes, nameptr int32
	var ok bool

	if ptr, ok = m.RegConc(11); !ok {
		m.err(fmt.Errorf("address is symbolic"))
	}
	if nbytes, ok = m.RegConc(12); !ok {
		m.err(fmt.Errorf("nbytes is symbolic"))
	}
	if nameptr, ok = m.RegConc(13); !ok {
		m.err(fmt.Errorf("name is symbolic"))
	}

	name, err := m.rdstr(uint32(nameptr), s)
	if err != nil {
		m.err(err)
		return
	}
	logger.Printf("Marking symbolic value '%s': %d bytes at 0x%x\n", name, nbytes, ptr)

	m.wrsym(uint32(ptr), uint32(nbytes), string(name), s)
}
