package rvsym

import (
	"sort"

	"github.com/zyedidia/go-z3/st"
)

type Memory struct {
	rdmem *Memory
	mem   map[uint32]st.Int32
}

func NewMemory(rdmem *Memory) *Memory {
	return &Memory{
		rdmem: rdmem,
		mem:   make(map[uint32]st.Int32),
	}
}

func (m *Memory) Keys() []uint32 {
	keysmap := m.keys()
	keys := make([]uint32, 0, len(keysmap))
	for k := range keysmap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func (m *Memory) keys() map[uint32]struct{} {
	if m == nil {
		return nil
	}
	keys := make(map[uint32]struct{})
	for k := range m.mem {
		keys[k] = struct{}{}
	}
	rdkeys := m.rdmem.keys()
	if rdkeys != nil {
		for k := range rdkeys {
			keys[k] = struct{}{}
		}
	}
	return keys
}

func (m *Memory) read(idx uint32) (st.Int32, bool) {
	if m == nil {
		return st.Int32{}, false
	}
	if v, ok := m.mem[idx]; ok {
		return v, true
	}
	v, ok := m.rdmem.read(idx)
	return v, ok
}

func (m *Memory) readz(idx uint32) st.Int32 {
	if m == nil {
		return st.Int32{}
	}
	if v, ok := m.mem[idx]; ok {
		return v
	}
	return m.rdmem.readz(idx)
}

func (m *Memory) Write32(addr uint32, val st.Int32) {
	m.mem[addr>>2] = val
}

func (m *Memory) Write16(addr uint32, val st.Int32) {
	wrb := addr & 0b11
	val = val.ToInt16().ToInt32()
	wrword := val.Lsh(st.Uint64{C: uint64(wrb * 8)})
	wrmask := st.Int32{C: int32(^uint32(0x0000ffff << (wrb * 8)))}
	m.mem[addr>>2] = m.readz(addr >> 2).And(wrmask).Or(wrword)
}

func (m *Memory) Write8(addr uint32, val st.Int32) {
	wrb := addr & 0b11
	val = val.ToInt8().ToInt32()
	wrword := val.Lsh(st.Uint64{C: uint64(wrb * 8)})
	wrmask := st.Int32{C: int32(^uint32(0x00ff << (wrb * 8)))}
	m.mem[addr>>2] = m.readz(addr >> 2).And(wrmask).Or(wrword)
}

func (m *Memory) Read32(addr uint32) (st.Int32, bool) {
	v, ok := m.read(addr >> 2)
	return v, ok
}

func (m *Memory) Read16(addr uint32) (st.Int32, bool) {
	rdb := st.Uint64{C: uint64(addr&0b11) * 8}
	if v, ok := m.read(addr >> 2); ok {
		return v.Rsh(rdb).ToInt16().ToInt32(), true
	}
	return st.Int32{}, false
}

func (m *Memory) Read8(addr uint32) (st.Int32, bool) {
	rdb := st.Uint64{C: uint64(addr&0b11) * 8}
	if v, ok := m.read(addr >> 2); ok {
		return v.Rsh(rdb).ToInt8().ToInt32(), true
	}
	return st.Int32{}, false
}

func (m *Memory) Read16u(addr uint32) (st.Int32, bool) {
	rdb := st.Uint64{C: uint64(addr&0b11) * 8}
	if v, ok := m.read(addr >> 2); ok {
		return v.Rsh(rdb).ToUint16().ToInt32(), true
	}
	return st.Int32{}, false
}

func (m *Memory) Read8u(addr uint32) (st.Int32, bool) {
	rdb := st.Uint64{C: uint64(addr&0b11) * 8}
	if v, ok := m.read(addr >> 2); ok {
		return v.Rsh(rdb).ToUint8().ToInt32(), true
	}
	return st.Int32{}, false
}
