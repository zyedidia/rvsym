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
	fmt.Println("TODO: SymPrint")
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
	fmt.Println("TODO: SymMarkArray")
}
func (m *Machine) SymMarkBytes(s *smt.Solver) {
	fmt.Println("TODO: SymMarkBytes")
}
