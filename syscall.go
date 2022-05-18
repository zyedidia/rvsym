package rvsym

import (
	"fmt"
	"io/fs"
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
	if fd, ok := m.RegConc(Ra0); !ok {
		m.err(fmt.Errorf("symbolic fd"))
		return
	} else if buf, ok := m.RegConc(Ra1); !ok {
		m.err(fmt.Errorf("symbolic buf"))
		return
	} else if count, ok := m.RegConc(Ra2); !ok {
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
				m.WriteReg(Ra0, smt.Int32{C: -1})
			} else {
				m.WriteReg(Ra0, smt.Int32{C: int32(n)})
			}
		}
	}
}

func (m *Machine) SysClose(s *smt.Solver) {
	if fd, ok := m.RegConc(Ra0); !ok {
		m.err(fmt.Errorf("symbolic fd"))
		return
	} else {
		if f, ok := m.sys.fdtbl.files[int(fd)]; ok && fd >= 3 {
			f.Close()
		}
		delete(m.sys.fdtbl.files, int(fd))
	}
}

func (m *Machine) SysFstat(s *smt.Solver) {
	if buf, ok := m.RegConc(Ra1); !ok {
		m.err(fmt.Errorf("symbolic fd"))
		return
	} else {
		for i := int32(0); i < 112; i += 4 {
			m.mem.WriteWord(smt.Int32{C: buf + i}, smt.Int32{C: 0}, s)
		}
		m.WriteReg(Ra0, smt.Int32{C: 0})
	}
}

func (m *Machine) SysLseek(s *smt.Solver) {
	if fd, ok := m.RegConc(Ra0); !ok {
		m.err(fmt.Errorf("symbolic fd"))
		return
	} else if offset, ok := m.RegConc(Ra1); !ok {
		m.err(fmt.Errorf("symbolic offset"))
		return
	} else if whence, ok := m.RegConc(Ra2); !ok {
		m.err(fmt.Errorf("symbolic whence"))
		return
	} else {
		if f, ok := m.sys.fdtbl.files[int(fd)]; !ok {
			m.err(fmt.Errorf("invalid fd %d", fd))
			return
		} else {
			ret, err := f.Seek(int64(offset), int(whence))
			if err != nil {
				m.WriteReg(Ra0, smt.Int32{C: -1})
			} else {
				m.WriteReg(Ra0, smt.Int32{C: int32(ret)})
			}
		}
	}
}

func (m *Machine) SysOpen(s *smt.Solver) {
	if pathname, ok := m.RegConc(Ra0); !ok {
		m.err(fmt.Errorf("symbolic path name"))
		return
	} else if flags, ok := m.RegConc(Ra1); !ok {
		m.err(fmt.Errorf("symbolic flags"))
		return
	} else if mode, ok := m.RegConc(Ra2); !ok {
		m.err(fmt.Errorf("symbolic mode"))
		return
	} else {
		path, err := m.rdstr(uint32(pathname), s)
		if err != nil {
			m.err(err)
			return
		}
		file, err := os.OpenFile(string(path), int(flags), fs.FileMode(mode))
		if err != nil {
			m.WriteReg(Ra0, smt.Int32{C: -1})
			return
		}

		m.sys.fdtbl.files[int(file.Fd())] = file
		m.WriteReg(Ra0, smt.Int32{C: int32(file.Fd())})
	}
}

func (m *Machine) SysRead(s *smt.Solver) {
	if fd, ok := m.RegConc(Ra0); !ok {
		m.err(fmt.Errorf("symbolic fd"))
		return
	} else if buf, ok := m.RegConc(Ra1); !ok {
		m.err(fmt.Errorf("symbolic buf"))
		return
	} else if count, ok := m.RegConc(Ra2); !ok {
		m.err(fmt.Errorf("symbolic count"))
		return
	} else {
		if f, ok := m.sys.fdtbl.files[int(fd)]; !ok {
			m.err(fmt.Errorf("invalid fd %d", fd))
			return
		} else {
			b := make([]byte, count)
			n, err := f.Read(b)
			if err != nil {
				m.WriteReg(Ra0, smt.Int32{C: -1})
				return
			}
			ok := m.mem.WriteBytes(uint32(buf), b[:n], s)
			if !ok {
				m.err(fmt.Errorf("invalid memory access during read syscall"))
				return
			}
			m.WriteReg(Ra0, smt.Int32{C: int32(n)})
		}
	}
}

func (m *Machine) SysBrk(s *smt.Solver) {
	if addr, ok := m.RegConc(Ra0); !ok {
		m.err(fmt.Errorf("symbolic addr"))
		return
	} else {
		if addr != 0 {
			m.sys.brk = uint32(addr)
		}
		m.WriteReg(Ra0, smt.Int32{C: int32(m.sys.brk)})
	}
}
