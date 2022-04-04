package rvsym

import (
	"fmt"
	"os"

	"github.com/zyedidia/rvsym/pkg/smt"
)

type EcallFn func(*Machine, *smt.Solver)

var syscalls = map[int]EcallFn{
	93:   (*Machine).SysExit,
	64:   (*Machine).SysWrite,
	57:   (*Machine).SysClose,
	80:   (*Machine).SysFstat,
	62:   (*Machine).SysLseek,
	1024: (*Machine).SysOpen,
	63:   (*Machine).SysRead,
	214:  (*Machine).SysBrk,
}

type FdTable struct {
	files map[int]*os.File
}

func NewFdTable() *FdTable {
	return &FdTable{
		files: map[int]*os.File{
			0: os.Stdin,
			1: os.Stdout,
			2: os.Stderr,
		},
	}
}

func (t *FdTable) Copy() *FdTable {
	files := make(map[int]*os.File)
	for k, v := range t.files {
		files[k] = v
	}
	return &FdTable{
		files: files,
	}
}

func (m *Machine) SysExit(s *smt.Solver) {
	m.SymQuietExit(s)
}

func (m *Machine) SysWrite(s *smt.Solver) {
	fmt.Println("TODO: SysWrite")
}

func (m *Machine) SysClose(s *smt.Solver) {
	// if fd, ok := m.RegConc(10); !ok {
	// 	m.err(fmt.Errorf("symbolic fd"))
	// 	return
	// } else {
	// 	if f, ok := m.fdtbl.files[fd]; ok {
	// 		f.Close()
	// 	}
	// 	delete(m.fdtbl.files, fd)
	// }
}

func (m *Machine) SysFstat(s *smt.Solver) {
	fmt.Println("TODO: SysFstat")
}

func (m *Machine) SysLseek(s *smt.Solver) {
	fmt.Println("TODO: SysLseek")
}

func (m *Machine) SysOpen(s *smt.Solver) {
	fmt.Println("TODO: SysOpen")
}

func (m *Machine) SysRead(s *smt.Solver) {
	fmt.Println("TODO: SysRead")
}

func (m *Machine) SysBrk(s *smt.Solver) {
	fmt.Println("TODO: SysBrk")
}
