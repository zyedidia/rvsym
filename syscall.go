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

type SysState struct {
	fdtbl FdTable
	brk   uint32
}

func NewSysState(brk uint32) *SysState {
	return &SysState{
		fdtbl: NewFdTable(),
		brk:   brk,
	}
}

func (s *SysState) Copy() *SysState {
	return &SysState{
		fdtbl: s.fdtbl.Copy(),
		brk:   s.brk,
	}
}

type FdTable struct {
	files map[int]*os.File
}

func NewFdTable() FdTable {
	return FdTable{
		files: map[int]*os.File{
			0: os.Stdin,
			1: os.Stdout,
			2: os.Stderr,
		},
	}
}

func (t FdTable) Copy() FdTable {
	files := make(map[int]*os.File)
	for k, v := range t.files {
		files[k] = v
	}
	return FdTable{
		files: files,
	}
}

func (m *Machine) SysExit(s *smt.Solver) {
	m.SymQuietExit(s)
}

func (m *Machine) SysWrite(s *smt.Solver) {
	if fd, ok := m.RegConc(10); !ok {
		m.err(fmt.Errorf("symbolic fd"))
		return
	} else if buf, ok := m.RegConc(11); !ok {
		m.err(fmt.Errorf("symbolic buf"))
		return
	} else if count, ok := m.RegConc(12); !ok {
		m.err(fmt.Errorf("symbolic count"))
		return
	} else {
		if f, ok := m.sys.fdtbl.files[int(fd)]; !ok {
			m.err(fmt.Errorf("invalid fd %d", fd))
			return
		} else {
			b := make([]byte, count)
			err := m.rdbytes(uint32(buf), b, s)
			if err != nil {
				m.err(err)
				return
			}
			n, err := f.Write(b)
			if err != nil {
				m.WriteReg(10, smt.Int32{C: -1})
			} else {
				m.WriteReg(10, smt.Int32{C: int32(n)})
			}
		}
	}
}

func (m *Machine) SysClose(s *smt.Solver) {
	if fd, ok := m.RegConc(10); !ok {
		m.err(fmt.Errorf("symbolic fd"))
		return
	} else {
		if f, ok := m.sys.fdtbl.files[int(fd)]; ok {
			f.Close()
		}
		delete(m.sys.fdtbl.files, int(fd))
	}
}

func (m *Machine) SysFstat(s *smt.Solver) {
	if buf, ok := m.RegConc(11); !ok {
		m.err(fmt.Errorf("symbolic fd"))
		return
	} else {
		for i := int32(0); i < 112; i += 4 {
			m.mem.WriteWord(smt.Int32{C: buf + i}, smt.Int32{C: 0}, s)
		}
		m.WriteReg(10, smt.Int32{C: 0})
	}
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
	if addr, ok := m.RegConc(10); !ok {
		m.err(fmt.Errorf("symbolic addr"))
		return
	} else {
		if addr != 0 {
			m.sys.brk = uint32(addr)
		}
		m.WriteReg(10, smt.Int32{C: int32(m.sys.brk)})
	}
}
