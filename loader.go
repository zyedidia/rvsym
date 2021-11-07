package minisym

import (
	"bytes"
	"encoding/binary"
)

func LoadCode(data []byte) []uint32 {
	insns := make([]uint32, 0, len(data)/4)

	i := 0
	for i < len(data) {
		var ui uint32
		buf := bytes.NewReader(data[i:])
		binary.Read(buf, binary.LittleEndian, &ui)
		insns = append(insns, ui)
		i += int(buf.Size()) - buf.Len()
	}
	return insns
}
