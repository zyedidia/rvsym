package rvsym

import (
	"sort"

	"github.com/zyedidia/rvsym/pkg/z3/st"
	"github.com/zyedidia/rvsym/pkg/z3/z3"
)

type Memory struct {
	rdmem *Memory
	mem   map[uint32]st.Int32
	arrs  []st.ArrayInt32
}

func NewMemory(rdmem *Memory) *Memory {
	var arrs []st.ArrayInt32
	if rdmem != nil {
		arrs = make([]st.ArrayInt32, len(rdmem.arrs))
		copy(arrs, rdmem.arrs)
	}
	return &Memory{
		rdmem: rdmem,
		mem:   make(map[uint32]st.Int32),
		arrs:  arrs,
	}
}

func (m *Memory) AddArray(ctx *z3.Context, name string, base int, length int) {
	m.arrs = append(m.arrs, st.AnyArrayInt32(ctx, name, base, length))
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
	for k := range rdkeys {
		keys[k] = struct{}{}
	}
	return keys
}

func (m *Memory) read(idx st.Uint32, s *z3.Solver) (st.Int32, bool) {
	if m == nil {
		return st.Int32{}, false
	}
	for _, a := range m.arrs {
		if a.InBounds(idx, s) {
			return a.Read(idx, s), true
		}
	}
	// not found in any symbolic array
	if !idx.IsConcrete() {
		return st.Int32{}, false
	}
	if v, ok := m.mem[idx.C]; ok {
		return v, true
	}
	v, ok := m.rdmem.read(idx, s)
	return v, ok
}

func (m *Memory) readz(idx st.Uint32, s *z3.Solver) st.Int32 {
	if m == nil {
		return st.Int32{}
	}
	v, _ := m.read(idx, s)
	return v
}

func (m *Memory) write(idx st.Uint32, val st.Int32, s *z3.Solver) bool {
	for i := range m.arrs {
		if m.arrs[i].InBounds(idx, s) {
			m.arrs[i].Write(idx, val, s)
			return true
		}
	}
	// not found in any symbolic memory
	if !idx.IsConcrete() {
		return false
	}
	m.mem[idx.C] = val
	return true
}

func (m *Memory) Write32(addr st.Uint32, val st.Int32, s *z3.Solver) bool {
	return m.write(addr.Rsh(st.Uint64{C: 2}), val, s)
}

func (m *Memory) Write16(addr st.Uint32, val st.Int32, s *z3.Solver) bool {
	wrb := addr.And(st.Uint32{C: 0b11})
	val = val.ToUint16().ToInt32()
	wrb8 := wrb.Lsh(st.Uint64{C: 3}).ToUint64()
	wrword := val.Lsh(wrb8)
	wrmask := st.Uint32{C: uint32(0x0ffff)}.Lsh(wrb8).Not().ToInt32()
	idx := addr.Rsh(st.Uint64{C: 2})
	return m.write(idx, m.readz(idx, s).And(wrmask).Or(wrword), s)
}

func (m *Memory) Write8(addr st.Uint32, val st.Int32, s *z3.Solver) bool {
	wrb := addr.And(st.Uint32{C: 0b11})
	val = val.ToUint8().ToInt32()
	wrb8 := wrb.Lsh(st.Uint64{C: 3}).ToUint64()
	wrword := val.Lsh(wrb8)
	wrmask := st.Uint32{C: uint32(0x0ff)}.Lsh(wrb8).Not().ToInt32()
	idx := addr.Rsh(st.Uint64{C: 2})
	return m.write(idx, m.readz(idx, s).And(wrmask).Or(wrword), s)
}

func (m *Memory) Read32(addr st.Uint32, s *z3.Solver) (st.Int32, bool) {
	idx := addr.Rsh(st.Uint64{C: 2})
	v, ok := m.read(idx, s)
	return v, ok
}

func (m *Memory) Read16(addr st.Uint32, s *z3.Solver) (st.Int32, bool) {
	rdb := addr.And(st.Uint32{C: 0b11})
	rdb8 := rdb.Lsh(st.Uint64{C: 3}).ToUint64()
	idx := addr.Rsh(st.Uint64{C: 2})
	if v, ok := m.read(idx, s); ok {
		return v.Rsh(rdb8).ToInt16().ToInt32(), true
	}
	return st.Int32{}, false
}

func (m *Memory) Read8(addr st.Uint32, s *z3.Solver) (st.Int32, bool) {
	rdb := addr.And(st.Uint32{C: 0b11})
	rdb8 := rdb.Lsh(st.Uint64{C: 3}).ToUint64()
	idx := addr.Rsh(st.Uint64{C: 2})
	if v, ok := m.read(idx, s); ok {
		return v.Rsh(rdb8).ToInt8().ToInt32(), true
	}
	return st.Int32{}, false
}

func (m *Memory) Read16u(addr st.Uint32, s *z3.Solver) (st.Int32, bool) {
	rdb := addr.And(st.Uint32{C: 0b11})
	rdb8 := rdb.Lsh(st.Uint64{C: 3}).ToUint64()
	idx := addr.Rsh(st.Uint64{C: 2})
	if v, ok := m.read(idx, s); ok {
		return v.Rsh(rdb8).ToUint16().ToInt32(), true
	}
	return st.Int32{}, false
}

func (m *Memory) Read8u(addr st.Uint32, s *z3.Solver) (st.Int32, bool) {
	rdb := addr.And(st.Uint32{C: 0b11})
	rdb8 := rdb.Lsh(st.Uint64{C: 3}).ToUint64()
	idx := addr.Rsh(st.Uint64{C: 2})
	if v, ok := m.read(idx, s); ok {
		return v.Rsh(rdb8).ToUint8().ToInt32(), true
	}
	return st.Int32{}, false
}
