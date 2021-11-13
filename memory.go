package rvsym

import (
	"github.com/zyedidia/go-z3/st"
)

type Memory map[uint32]st.Int32

func (m Memory) Write32(addr uint32, val st.Int32) {
	m[addr>>2] = val
}

func (m Memory) Write16(addr uint32, val st.Int32) {
	wrb := addr & 0b11
	val = val.ToInt16().ToInt32()
	wrword := val.Lsh(st.Uint64{C: uint64(wrb * 8)})
	wrmask := st.Int32{C: int32(^uint32(0x0000ffff << (wrb * 8)))}
	m[addr>>2] = m[addr>>2].And(wrmask).Or(wrword)
}

func (m Memory) Write8(addr uint32, val st.Int32) {
	wrb := addr & 0b11
	val = val.ToInt8().ToInt32()
	wrword := val.Lsh(st.Uint64{C: uint64(wrb * 8)})
	wrmask := st.Int32{C: int32(^uint32(0x00ff << (wrb * 8)))}
	m[addr>>2] = m[addr>>2].And(wrmask).Or(wrword)
}

func (m Memory) Read32(addr uint32) (st.Int32, bool) {
	v, ok := m[addr>>2]
	return v, ok
}

func (m Memory) Read16(addr uint32) (st.Int32, bool) {
	rdb := st.Uint64{C: uint64(addr&0b11) * 8}
	if v, ok := m[addr>>2]; ok {
		return v.Rsh(rdb).ToInt16().ToInt32(), true
	}
	return st.Int32{}, false
}

func (m Memory) Read8(addr uint32) (st.Int32, bool) {
	rdb := st.Uint64{C: uint64(addr&0b11) * 8}
	if v, ok := m[addr>>2]; ok {
		return v.Rsh(rdb).ToInt8().ToInt32(), true
	}
	return st.Int32{}, false
}

func (m Memory) Read16u(addr uint32) (st.Int32, bool) {
	rdb := st.Uint64{C: uint64(addr&0b11) * 8}
	if v, ok := m[addr>>2]; ok {
		return v.Rsh(rdb).ToUint16().ToInt32(), true
	}
	return st.Int32{}, false
}

func (m Memory) Read8u(addr uint32) (st.Int32, bool) {
	rdb := st.Uint64{C: uint64(addr&0b11) * 8}
	if v, ok := m[addr>>2]; ok {
		return v.Rsh(rdb).ToUint8().ToInt32(), true
	}
	return st.Int32{}, false
}
