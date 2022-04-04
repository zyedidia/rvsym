package rvsym

import (
	"encoding/binary"

	"github.com/zyedidia/generic"
	"github.com/zyedidia/generic/hashmap"
	"github.com/zyedidia/generic/hashset"
	"github.com/zyedidia/rvsym/pkg/smt"
)

type Memory struct {
	mem *hashmap.Map[int32, smt.Int32]
	// regions of memory that may be symbolically addressed
	arrs []smt.ArrayInt32
	// valid concrete addresses within symbolic regions
	valid *hashset.Set[int32]
}

func NewMemory() *Memory {
	return &Memory{
		mem:   hashmap.New[int32, smt.Int32](1024, generic.Equals[int32], generic.HashInt32),
		arrs:  nil,
		valid: hashset.New[int32](64, generic.Equals[int32], generic.HashInt32),
	}
}

func (m *Memory) Copy() *Memory {
	arrs := make([]smt.ArrayInt32, len(m.arrs))
	copy(arrs, m.arrs)

	return &Memory{
		mem:   m.mem.Copy(),
		arrs:  arrs,
		valid: m.valid.Copy(),
	}
}

func (m *Memory) WriteWord(addr, val smt.Int32, s *smt.Solver) bool {
	return m.write(addr.Srl(smt.Int32{C: 2}, s), val, s)
}

func (m *Memory) WriteBytes(addr uint32, data []byte, s *smt.Solver) bool {
	for len(data) > 0 {
		if addr%4 == 0 && len(data) >= 4 {
			v := binary.LittleEndian.Uint32(data)
			ok := m.WriteWord(smt.Int32{C: int32(addr)}, smt.Int32{C: int32(v)}, s)
			if !ok {
				return false
			}
			data = data[4:]
			addr += 4
		} else {
			v := data[0]
			ok := m.Write8(smt.Int32{C: int32(addr)}, smt.Int32{C: int32(v)}, s)
			if !ok {
				return false
			}
			data = data[1:]
			addr++
		}
	}
	return true
}

func (m *Memory) ReadWord(addr smt.Int32, s *smt.Solver) (smt.Int32, bool) {
	return m.read(addr.Srl(smt.Int32{C: 2}, s), s)
}

func (m *Memory) ReadBytes(addr uint32, data []byte, s *smt.Solver) bool {
	for len(data) > 0 {
		// if addr%4 == 0 && len(data) >= 4 {
		// 	v, ok := m.ReadWord(smt.Int32{C: int32(addr)}, s)
		// 	if !ok || !v.Concrete() {
		// 		return false
		// 	}
		// 	binary.LittleEndian.PutUint32(data, uint32(v.C))
		// 	data = data[4:]
		// 	addr += 4
		// } else {
		v, ok := m.Read8u(smt.Int32{C: int32(addr)}, s)
		if !ok || !v.Concrete() {
			return false
		}
		data[0] = byte(v.C)
		addr++
		data = data[1:]
		// }
	}
	return true
}

type tmpval struct {
	idx uint32
	val smt.Int32
}

// MakeSymAddrRegion marks the given region as a region of memory that can be
// symbolically addressed.
func (m *Memory) MakeSymAddrRegion(base, length uint32, s *smt.Solver) {
	vals := make([]tmpval, 0, length)
	for i := base; i < base+length; i++ {
		v, ok := m.mem.Get(int32(i))
		m.valid.Put(int32(i))
		if ok {
			vals = append(vals, tmpval{
				idx: i,
				val: v,
			})
		}
	}
	m.arrs = append(m.arrs, s.AnyArrayInt32(base, length))
	for _, v := range vals {
		m.writeInit(smt.Int32{C: int32(v.idx)}, v.val, s)
	}
}

// returns the value at idx, or false if it does not exist
func (m *Memory) read(idx smt.Int32, s *smt.Solver) (smt.Int32, bool) {
	for _, a := range m.arrs {
		if a.InBounds(idx, s) {
			if idx.Concrete() && m.valid.Has(idx.C) {
				return m.mem.Get(idx.C)
			}
			return a.Read(idx, s), true
		}
	}

	if !idx.Concrete() {
		return smt.Int32{}, false
	}
	return m.mem.Get(idx.C)
}

// identical to read but returns the zero value of smt.Int32 if not found
func (m *Memory) readz(idx smt.Int32, s *smt.Solver) smt.Int32 {
	v, _ := m.read(idx, s)
	return v
}

func (m *Memory) write(idx, val smt.Int32, s *smt.Solver) bool {
	return m.writeVal(idx, val, s, false)
}

func (m *Memory) writeInit(idx, val smt.Int32, s *smt.Solver) bool {
	return m.writeVal(idx, val, s, true)
}

// attempt to write val at idx; returns false if the access is out of bounds
// (only possible when using a symbolic address).
// init indicates that this is an initial write to a symbolic memory.
func (m *Memory) writeVal(idx, val smt.Int32, s *smt.Solver, init bool) bool {
	for i := range m.arrs {
		if m.arrs[i].InBounds(idx, s) {
			if idx.Concrete() {
				if m.valid.Has(idx.C) {
					v, ok := m.read(idx, s)
					if !ok || init {
						// put the value at a concrete address in the concrete
						// memory and in the symbolic array memory.
						m.mem.Put(idx.C, val)
						m.arrs[i].WriteInitial(idx, val, s)
						return true
					}
					if val.Concrete() && v.Concrete() && val.C == v.C {
						return true
					}
				}
				// put the value at the concrete address and indicate that we
				// have a valid value at this concrete address
				m.mem.Put(idx.C, val)
				m.valid.Put(idx.C)
			} else {
				// the index is symbolic so we must invalidate all concrete
				// indexes within the associated symbolic array.
				m.valid.Each(func(k int32) {
					if m.arrs[i].InBounds(smt.Int32{C: k}, s) {
						m.valid.Remove(k)
					}
				})
			}

			// Note: even if the address is concrete we still need to perform a
			// symbolic write because in the future there may be a read
			// with a symbolic address

			m.arrs[i].Write(idx, val, s)
			return true
		}
	}
	if !idx.Concrete() {
		return false
	}
	m.mem.Put(idx.C, val)
	return true
}

// Write16 writes a 16-bit value (truncated) at addr. Returns false if the
// access is out of bounds.
func (m *Memory) Write16(addr, val smt.Int32, s *smt.Solver) bool {
	// (addr & 0b11) << 3: this is the index of the first bit in the target
	// 32-bit word that val will be written to
	wrb := addr.And(smt.Int32{C: 0b011}, s).Sll(smt.Int32{C: 3}, s)
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
	wrword := val.Sll(wrb, s)
	wrmask := smt.Int32{C: 0x000000ff}.Sll(wrb, s).Not(s)
	idx := addr.Srl(smt.Int32{C: 2}, s)
	return m.write(idx, m.readz(idx, s).And(wrmask, s).Or(wrword, s), s)
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
