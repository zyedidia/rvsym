package rvsym

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
	if len(data)%4 != 0 {
		return nil, 0, fmt.Errorf("load: not a multiple of 4")
	}

	seg := Segment{
		addr: b.Entry,
		data: make([]uint32, 0, len(data)/4),
	}

	i := 0
	for i < len(data) {
		var ui uint32
		buf := bytes.NewReader(data[i:])
		binary.Read(buf, binary.LittleEndian, &ui)
		seg.data = append(seg.data, ui)
		i += int(buf.Size()) - buf.Len()
	}
	return []Segment{seg}, 0, nil

}
