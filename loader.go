package rvsym

import (
	"bytes"
	"debug/elf"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/marcinbor85/gohex"
)

type Segment struct {
	addr uint32
	data []uint32
}

type Loader interface {
	Load(data []byte) (segs []Segment, entrypc uint32, err error)
}

// A BinaryLoader loads a single blob of data, maps it at address 0, and uses an
// entry point of 0.
type BinaryLoader struct {
	Entry uint32
}

func (b *BinaryLoader) Load(data []byte) ([]Segment, uint32, error) {
	seg := Segment{
		addr: b.Entry,
	}

	seg.data = toWords(data)
	return []Segment{seg}, b.Entry, nil
}

type IntelHexLoader struct {
	Entry uint32
}

func (l *IntelHexLoader) Load(data []byte) ([]Segment, uint32, error) {
	mem := gohex.NewMemory()
	if err := mem.ParseIntelHex(bytes.NewReader(data)); err != nil {
		return nil, 0, err
	}
	hexsegs := mem.GetDataSegments()
	segs := make([]Segment, len(hexsegs))
	for i, segment := range hexsegs {
		segs[i] = Segment{
			addr: segment.Address,
			data: toWords(segment.Data),
		}
	}
	return segs, l.Entry, nil
}

type ElfLoader struct{}

func (e *ElfLoader) Load(data []byte) ([]Segment, uint32, error) {
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
				data: toWords(data),
			})
		}
	}

	return segs, uint32(f.Entry), nil
}

func toWords(data []byte) []uint32 {
	words := make([]uint32, 0, len(data)/4)
	i := 0
	for i < len(data) {
		var ui uint32
		buf := bytes.NewReader(data[i:])
		binary.Read(buf, binary.LittleEndian, &ui)
		words = append(words, ui)
		i += int(buf.Size()) - buf.Len()
	}
	return words
}
