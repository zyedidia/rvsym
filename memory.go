package rvsym

import (
	"sort"

	"github.com/zyedidia/rvsym/pkg/smt"
)

// A Memory represents the address space of the symbolic process. It tracks a
// concrete memory space as a map of concrete addresses to concolic values. It
// also tracks regions of memory which can be indexed with fully symbolic
// addresses. These regions of the address space must be registered ahead of
// time, and during reads/writes can also be bounds checked if desired.
type Memory struct {
	rdmem *Memory
	mem   map[int32]smt.Int32
	arrs  []smt.ArrayInt32
	valid map[int32]struct{}
}

// NewMemory creates a new memory. The rdmem argument provides an existing
// address space that can contain pre-populated values that will only be read
// by this memory. This avoids unnecessarily copying data that will only ever
// be read when copying an address space.
func NewMemory(rdmem *Memory) *Memory {
	var arrs []smt.ArrayInt32
	valid := make(map[int32]struct{})
	if rdmem != nil {
		arrs = make([]smt.ArrayInt32, len(rdmem.arrs))
		copy(arrs, rdmem.arrs)
		for k, v := range rdmem.valid {
			valid[k] = v
		}
	}
	return &Memory{
		rdmem: rdmem,
		mem:   make(map[int32]smt.Int32),
		arrs:  arrs,
		valid: valid,
	}
}

type tmpval struct {
	idx uint32
	val smt.Int32
}

// AddArray registers the region [base,base+length) as a symbolic address
// region. Note that base and length are word addresses.
func (m *Memory) AddArray(s *smt.Solver, base uint32, length uint32) {
	vals := make([]tmpval, 0, length)
	for i := base; i < base+length; i++ {
		v, ok := m.mem[int32(i)]
		if ok {
			vals = append(vals, tmpval{
				idx: i,
				val: v,
			})
		}
	}
	m.arrs = append(m.arrs, s.AnyArrayInt32(base, length))
	for _, v := range vals {
		m.write(smt.Int32{C: int32(v.idx)}, v.val, s)
	}
}

// Keys returns all concrete addresses with values in sorted order.
func (m *Memory) Keys() []int32 {
	keysmap := m.keys()
	keys := make([]int32, 0, len(keysmap))
	for k := range keysmap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func (m *Memory) keys() map[int32]struct{} {
	if m == nil {
		return nil
	}
	keys := make(map[int32]struct{})
	for k := range m.mem {
		keys[k] = struct{}{}
	}
	rdkeys := m.rdmem.keys()
	for k := range rdkeys {
		keys[k] = struct{}{}
	}
	return keys
}

func (m *Memory) readmem(idx int32) (smt.Int32, bool) {
	if m == nil {
		return smt.Int32{}, false
	}
	if v, ok := m.mem[idx]; ok {
		return v, true
	}
	return m.rdmem.readmem(idx)
}

// returns the value at idx, or false if it does not exist
func (m *Memory) read(idx smt.Int32, s *smt.Solver) (smt.Int32, bool) {
	for _, a := range m.arrs {
		if a.InBounds(idx, s) {
			if idx.Concrete() {
				if _, ok := m.valid[idx.C]; ok {
					return m.readmem(idx.C)
				}
			}

			return a.Read(idx, s), true
		}
	}
	if !idx.Concrete() {
		return smt.Int32{}, false
	}
	return m.readmem(idx.C)
}

// identical to read but returns the zero value of smt.Int32 if not found
func (m *Memory) readz(idx smt.Int32, s *smt.Solver) smt.Int32 {
	v, _ := m.read(idx, s)
	return v
}

// attempt to write val at idx; returns false if the access is out of bounds
// (only possible when using a symbolic address).
func (m *Memory) write(idx, val smt.Int32, s *smt.Solver) bool {

	for i := range m.arrs {
		if m.arrs[i].InBounds(idx, s) {
			if idx.Concrete() {
				m.mem[idx.C] = val
				m.valid[idx.C] = struct{}{}
				// even if the address is concrete we still need to perform a
				// symbolic write beceause in the future there may be a read
				// with a symbolic address
			} else {
				for k := range m.valid {
					if m.arrs[i].InBounds(smt.Int32{C: k}, s) {
						delete(m.valid, k)
					}
				}
			}

			m.arrs[i].Write(idx, val, s)
			return true
		}
	}
	if !idx.Concrete() {
		return false
	}
	m.mem[idx.C] = val
	return true
}

// Write32 writes a 32 bit value at addr. Returns false if the access is out of
// bounds.
func (m *Memory) Write32(addr, val smt.Int32, s *smt.Solver) bool {
	return m.write(addr.Srl(smt.Int32{C: 2}, s), val, s)
}

// Write16 writes a 16-bit value (truncated) at addr. Returns false if the
// access is out of bounds.
func (m *Memory) Write16(addr, val smt.Int32, s *smt.Solver) bool {
	// (addr & 0b11) << 3: this is the index of the first bit in the target
	// 32-bit word that val will be written to
	wrb := addr.And(smt.Int32{C: 0b011}, s).Sll(smt.Int32{C: 3}, s)
	// clear out the top 16 bits just in case
	val = val.ToInt16(s).ToInt32z(s)
	// val << wrb: shift 16-bit val over to the target
	// position.
	wrword := val.Sll(wrb, s)
	// ^(0xffff << wrb): shift the mask into position, masking only the correct
	// 16 bits. This mask will allow us to clear the right 16-bits before ORing
	// in the new ones.
	wrmask := smt.Int32{C: 0x0000ffff}.Sll(wrb, s).Not(s)
	idx := addr.Srl(smt.Int32{C: 2}, s)
	return m.write(idx, m.readz(idx, s).And(wrmask, s).Or(wrword, s), s)
}

// Write8 writes an 8-bit value (truncated) at addr.
func (m *Memory) Write8(addr, val smt.Int32, s *smt.Solver) bool {
	// same as Write16 but for 8 bits
	wrb := addr.And(smt.Int32{C: 0b011}, s).Sll(smt.Int32{C: 3}, s)
	val = val.ToInt8(s).ToInt32z(s)
	wrword := val.Sll(wrb, s)
	wrmask := smt.Int32{C: 0x000000ff}.Sll(wrb, s).Not(s)
	idx := addr.Srl(smt.Int32{C: 2}, s)
	return m.write(idx, m.readz(idx, s).And(wrmask, s).Or(wrword, s), s)
}

// Read32 reads the 32-bit value at addr. Returns false if not found.
func (m *Memory) Read32(addr smt.Int32, s *smt.Solver) (smt.Int32, bool) {
	idx := addr.Srl(smt.Int32{C: 2}, s)
	return m.read(idx, s)
}

// Read16 reads the 16-bit value at addr. Returns false if not found.
func (m *Memory) Read16(addr smt.Int32, s *smt.Solver) (smt.Int32, bool) {
	// (addr & 0b11) << 3
	rdb := addr.And(smt.Int32{C: 0b011}, s).Sll(smt.Int32{C: 3}, s)
	idx := addr.Srl(smt.Int32{C: 2}, s)
	if v, ok := m.read(idx, s); ok {
		return v.Srl(rdb, s).ToInt16(s).ToInt32s(s), true
	}
	return smt.Int32{}, false
}

// Read8 reads the 8-bit value at addr. Returns false if not found.
func (m *Memory) Read8(addr smt.Int32, s *smt.Solver) (smt.Int32, bool) {
	rdb := addr.And(smt.Int32{C: 0b011}, s).Sll(smt.Int32{C: 3}, s)
	idx := addr.Srl(smt.Int32{C: 2}, s)
	if v, ok := m.read(idx, s); ok {
		return v.Srl(rdb, s).ToInt8(s).ToInt32s(s), true
	}
	return smt.Int32{}, false
}

// Read16u reads the 16-bit unsigned value at addr. Returns false if not found.
func (m *Memory) Read16u(addr smt.Int32, s *smt.Solver) (smt.Int32, bool) {
	rdb := addr.And(smt.Int32{C: 0b011}, s).Sll(smt.Int32{C: 3}, s)
	idx := addr.Srl(smt.Int32{C: 2}, s)
	if v, ok := m.read(idx, s); ok {
		return v.Srl(rdb, s).ToInt16(s).ToInt32z(s), true
	}
	return smt.Int32{}, false
}

// Read8u reads the 8-bit unsigned value at addr. Returns false if not found.
func (m *Memory) Read8u(addr smt.Int32, s *smt.Solver) (smt.Int32, bool) {
	rdb := addr.And(smt.Int32{C: 0b011}, s).Sll(smt.Int32{C: 3}, s)
	idx := addr.Srl(smt.Int32{C: 2}, s)
	if v, ok := m.read(idx, s); ok {
		return v.Srl(rdb, s).ToInt8(s).ToInt32z(s), true
	}
	return smt.Int32{}, false
}
