package rvsym

import (
	"bytes"
	"debug/elf"
	"fmt"
	"io"

	"github.com/marcinbor85/gohex"
)

type Segment struct {
	addr uint32
	data []byte
}

type Loader interface {
	Load(data []byte) (segs []Segment, entry uint32, err error)
}

type RawLoader struct {
	Entry uint32
}

func (l *RawLoader) Load(data []byte) ([]Segment, uint32, error) {
	seg := Segment{
		addr: l.Entry,
		data: data,
	}
	return []Segment{seg}, l.Entry, nil
}

type IHexLoader struct {
	Entry uint32
}

func (l *IHexLoader) Load(data []byte) ([]Segment, uint32, error) {
	mem := gohex.NewMemory()
	if err := mem.ParseIntelHex(bytes.NewReader(data)); err != nil {
		return nil, 0, err
	}
	hexsegs := mem.GetDataSegments()
	segs := make([]Segment, len(hexsegs))
	for i, segment := range hexsegs {
		segs[i] = Segment{
			addr: segment.Address,
			data: segment.Data,
		}
	}
	return segs, l.Entry, nil
}

type ElfLoader struct{}

func (l *ElfLoader) Load(data []byte) ([]Segment, uint32, error) {
	r := bytes.NewReader(data)
	f, err := elf.NewFile(r)
	if err != nil {
		return nil, 0, err
	}
	defer f.Close()

	if f.Type != elf.ET_EXEC {
		return nil, 0, fmt.Errorf("invalid elf file type: %v", f.Type)
	}

	segs := make([]Segment, 0, len(f.Progs))
	for _, p := range f.Progs {
		if p.Type == elf.PT_LOAD {
			data := make([]byte, p.Memsz)
			fdata := make([]byte, p.Filesz)
			n, err := p.ReadAt(fdata, 0)
			if err != nil && err != io.EOF {
				return nil, 0, err
			}
			copy(data, fdata[:n])

			segs = append(segs, Segment{
				addr: uint32(p.Vaddr),
				data: data,
			})
		}
	}

	return segs, uint32(f.Entry), nil
}
