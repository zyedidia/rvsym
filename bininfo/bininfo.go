package bininfo

import (
	"debug/dwarf"
	"debug/elf"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrInvalidElfType = errors.New("invalid elf type")
)

type Position struct {
	File string
	Line int
}

type BinFile struct {
	pie       bool
	positions map[uint64]Position
}

// Read creates a new BinFile from an io.ReaderAt.
func Read(r io.ReaderAt) (*BinFile, error) {
	f, err := elf.NewFile(r)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b := &BinFile{
		pie:       false,
		positions: make(map[uint64]Position),
	}

	if f.Type == elf.ET_DYN {
		b.pie = true
	} else if f.Type != elf.ET_EXEC {
		return nil, ErrInvalidElfType
	}

	// Get the vaddr of the first loadable segment. I'm not sure if this is the
	// right way to find the vaddr offset but it seems to work and I couldn't
	// find any documentation about this.
	var vaddr uint64
	if b.pie {
		for _, p := range f.Progs {
			if p.Type == elf.PT_LOAD {
				vaddr = p.Vaddr
				break
			}
		}
	}

	err = b.buildAddrCache(f, vaddr)

	return b, err
}

func (b *BinFile) buildAddrCache(f *elf.File, offset uint64) error {
	dw, err := f.DWARF()
	if err != nil {
		return err
	}

	r := dw.Reader()

	var filetable []string
	for {
		e, err := r.Next()
		if err != nil {
			return err
		}
		if e == nil {
			break
		}

		if e.Tag == dwarf.TagInlinedSubroutine {
			callfile, okFile := e.Val(dwarf.AttrCallFile).(int64)
			callline, okLine := e.Val(dwarf.AttrCallLine).(int64)
			lowpc, okPC := e.Val(dwarf.AttrLowpc).(uint64)
			if okFile && okLine && okPC && callfile < int64(len(filetable)) {
				b.addLineCacheEntry(
					filetable[callfile],
					int(callline),
					lowpc-offset,
				)
			}
		} else if e.Tag == dwarf.TagCompileUnit {
			filetable = make([]string, 0)

			lr, err := dw.LineReader(e)
			if err != nil || lr == nil {
				continue
			}
			var entry dwarf.LineEntry

			for {
				err = lr.Next(&entry)
				if err == io.EOF {
					break
				} else if err != nil || !entry.IsStmt {
					continue
				}

				var file string
				if entry.File == nil {
					file = "<unknown>"
				} else {
					file = entry.File.Name
				}
				filetable = append(filetable, file)

				b.addLineCacheEntry(
					file,
					entry.Line,
					entry.Address-offset,
				)
			}
		}
	}
	return nil
}

func (b *BinFile) addLineCacheEntry(file string, line int, addr uint64) {
	if _, ok := b.positions[uint64(addr)]; !ok {
		b.positions[uint64(addr)] = Position{
			File: file,
			Line: line,
		}
	}
}

func (b *BinFile) Pos(addr uint64) (Position, bool) {
	if b == nil {
		return Position{}, false
	}
	p, ok := b.positions[addr]

	if ok && strings.Contains(p.File, "rvsym.h") {
		// hack: sometimes dwarf marks positions as inside rvsym.h, so we retry
		// on the previous instruction which dwarf has usually pointed to the
		// calling line number (in the user's code).
		return b.Pos(addr - 4)
	}

	return p, ok
}

func (b *BinFile) PosStr(addr uint64) string {
	p, ok := b.Pos(addr)
	if !ok {
		return fmt.Sprintf("0x%x", addr)
	}
	var path string
	current, err := os.Getwd()
	if err != nil {
		path = p.File
	} else if rel, err := filepath.Rel(current, p.File); err != nil {
		path = p.File
	} else {
		path = rel
	}
	return fmt.Sprintf("%s:%d", path, p.Line)
}
