//go:build !noboolector

package smt

//#cgo CFLAGS: -I./deps/include
//#cgo LDFLAGS: -L./deps/lib -lboolector -llgl -lbtor2parser -lm
//#include <stdlib.h>
//#include "boolector.h"
import "C"
import (
	"strconv"
)

type Solver struct {
	btor *C.Btor

	sorts
}

type Model struct {
	btor *C.Btor
}

type sorts struct {
	sortBool       C.BoolectorSort
	sortInt8       C.BoolectorSort
	sortInt16      C.BoolectorSort
	sortInt32      C.BoolectorSort
	sortInt64      C.BoolectorSort
	sortArrayInt32 C.BoolectorSort
}

func initSorts(btor *C.Btor) sorts {
	var s sorts
	s.sortBool = C.boolector_bitvec_sort(btor, 1)
	s.sortInt8 = C.boolector_bitvec_sort(btor, 8)
	s.sortInt16 = C.boolector_bitvec_sort(btor, 16)
	s.sortInt32 = C.boolector_bitvec_sort(btor, 32)
	s.sortInt64 = C.boolector_bitvec_sort(btor, 64)
	s.sortArrayInt32 = C.boolector_array_sort(btor, s.sortInt32, s.sortInt32)
	return s
}

func NewSolver() *Solver {
	btor := C.boolector_new()
	C.boolector_set_opt(btor, C.BTOR_OPT_MODEL_GEN, 1)
	C.boolector_set_opt(btor, C.BTOR_OPT_OUTPUT_NUMBER_FORMAT, C.BTOR_OUTPUT_BASE_DEC)
	C.boolector_set_opt(btor, C.BTOR_OPT_INCREMENTAL, 1)

	return &Solver{
		btor:  btor,
		sorts: initSorts(btor),
	}
}

func (s *Solver) Push() {
	C.boolector_push(s.btor, 1)
}

func (s *Solver) Pop() {
	C.boolector_pop(s.btor, 1)
}

func (s *Solver) Assert(b Bool) {
	C.boolector_assert(s.btor, b.Sym(s).BV)
}

func (s *Solver) Check() CheckResult {
	result := C.boolector_sat(s.btor)
	switch result {
	case C.BOOLECTOR_SAT:
		return Sat
	case C.BOOLECTOR_UNSAT:
		return Unsat
	default:
		return Unknown
	}
}

func (s *Solver) Model() Model {
	return Model{s.btor}
}

func (m Model) Eval(a Int32) int32 {
	if a.Concrete() {
		return a.C
	}
	i, err := strconv.Atoi(C.GoString(C.boolector_bv_assignment(m.btor, a.S.BV)))

	if err != nil {
		panic(err)
	}

	return int32(i)
}

func (s *Solver) AnyInt32() Int32 {
	return Int32{S: SymInt32{C.boolector_var(s.btor, s.sortInt32, nil)}}
}
func (s *Solver) AnyArrayInt32(base, length int) ArrayInt32 {
	return ArrayInt32{
		base:   int32(base),
		length: int32(length),
		S:      SymArrayInt32{C.boolector_array(s.btor, s.sortArrayInt32, nil)},
	}
}

func (s *Solver) ToSymInt32(c int32) SymInt32 {
	return SymInt32{C.boolector_const_uint64(s.btor, C.ulong(c), 32)}
}
func (s *Solver) ToSymInt64(c int64) SymInt64 {
	return SymInt64{C.boolector_const_uint64(s.btor, C.ulong(c), 64)}
}
func (s *Solver) ToSymBool(c bool) SymBool {
	var cu C.ulong
	if c {
		cu = 1
	}
	return SymBool{C.boolector_const_uint64(s.btor, cu, 1)}
}

type SymInt8 struct {
	BV *C.BoolectorNode
}

func (a SymInt8) Valid() bool {
	return a.BV != nil
}

func (a SymInt8) ToInt32s(s *Solver) SymInt32 {
	return SymInt32{C.boolector_sext(s.btor, a.BV, 24)}
}
func (a SymInt8) ToInt32z(s *Solver) SymInt32 {
	return SymInt32{C.boolector_uext(s.btor, a.BV, 24)}
}

type SymInt16 struct {
	BV *C.BoolectorNode
}

func (a SymInt16) Valid() bool {
	return a.BV != nil
}

func (a SymInt16) ToInt32s(s *Solver) SymInt32 {
	return SymInt32{C.boolector_sext(s.btor, a.BV, 16)}
}
func (a SymInt16) ToInt32z(s *Solver) SymInt32 {
	return SymInt32{C.boolector_uext(s.btor, a.BV, 16)}
}

type SymInt64 struct {
	BV *C.BoolectorNode
}

func (a SymInt64) Valid() bool {
	return a.BV != nil
}

func (a SymInt64) Lower32(s *Solver) SymInt32 {
	return SymInt32{C.boolector_slice(s.btor, a.BV, 31, 0)}
}
func (a SymInt64) Upper32(s *Solver) SymInt32 {
	return SymInt32{C.boolector_slice(s.btor, a.BV, 63, 32)}
}
func (a SymInt64) Mul(b SymInt64, s *Solver) SymInt64 {
	return SymInt64{C.boolector_mul(s.btor, a.BV, b.BV)}
}

type SymBool struct {
	BV *C.BoolectorNode
}

func (a SymBool) Valid() bool {
	return a.BV != nil
}

func (a SymBool) Not(s *Solver) SymBool {
	return SymBool{C.boolector_not(s.btor, a.BV)}
}

type SymInt32 struct {
	BV *C.BoolectorNode
}

func (a SymInt32) Valid() bool {
	return a.BV != nil
}

func (a SymInt32) Add(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{C.boolector_add(s.btor, a.BV, b.BV)}
}
func (a SymInt32) Sub(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{C.boolector_sub(s.btor, a.BV, b.BV)}
}
func (a SymInt32) And(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{C.boolector_and(s.btor, a.BV, b.BV)}
}
func (a SymInt32) Or(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{C.boolector_or(s.btor, a.BV, b.BV)}
}
func (a SymInt32) Xor(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{C.boolector_xor(s.btor, a.BV, b.BV)}
}
func (a SymInt32) Sll(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{C.boolector_sll(s.btor, a.BV, b.BV)}
}
func (a SymInt32) Srl(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{C.boolector_srl(s.btor, a.BV, b.BV)}
}
func (a SymInt32) Sra(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{C.boolector_sra(s.btor, a.BV, b.BV)}
}
func (a SymInt32) Slt(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{C.boolector_uext(s.btor, C.boolector_slt(s.btor, a.BV, b.BV), 31)}
}
func (a SymInt32) Ult(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{C.boolector_uext(s.btor, C.boolector_ult(s.btor, a.BV, b.BV), 31)}
}
func (a SymInt32) Mul(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{C.boolector_mul(s.btor, a.BV, b.BV)}
}
func (a SymInt32) Div(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{C.boolector_sdiv(s.btor, a.BV, b.BV)}
}
func (a SymInt32) Divu(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{C.boolector_udiv(s.btor, a.BV, b.BV)}
}
func (a SymInt32) Rem(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{C.boolector_srem(s.btor, a.BV, b.BV)}
}
func (a SymInt32) Remu(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{C.boolector_urem(s.btor, a.BV, b.BV)}
}
func (a SymInt32) Not(s *Solver) SymInt32 {
	return SymInt32{C.boolector_not(s.btor, a.BV)}
}
func (a SymInt32) Eqz(s *Solver) SymBool {
	return SymBool{C.boolector_eq(s.btor, a.BV, s.ToSymInt32(0).BV)}
}
func (a SymInt32) NEqz(s *Solver) SymBool {
	return SymBool{C.boolector_ne(s.btor, a.BV, s.ToSymInt32(0).BV)}
}
func (a SymInt32) Sltb(b SymInt32, s *Solver) SymBool {
	return SymBool{C.boolector_slt(s.btor, a.BV, b.BV)}
}
func (a SymInt32) Sgeb(b SymInt32, s *Solver) SymBool {
	return SymBool{C.boolector_sgte(s.btor, a.BV, b.BV)}
}
func (a SymInt32) ToInt8(s *Solver) SymInt8 {
	return SymInt8{C.boolector_slice(s.btor, a.BV, 7, 0)}
}
func (a SymInt32) ToInt16(s *Solver) SymInt16 {
	return SymInt16{C.boolector_slice(s.btor, a.BV, 15, 0)}
}
func (a SymInt32) ToInt64s(s *Solver) SymInt64 {
	return SymInt64{C.boolector_sext(s.btor, a.BV, 32)}
}
func (a SymInt32) ToInt64z(s *Solver) SymInt64 {
	return SymInt64{C.boolector_uext(s.btor, a.BV, 32)}
}

type SymArrayInt32 struct {
	array *C.BoolectorNode
}

func (a SymArrayInt32) Select(idx SymInt32, s *Solver) SymInt32 {
	return SymInt32{C.boolector_read(s.btor, a.array, idx.BV)}
}

func (a SymArrayInt32) Store(idx SymInt32, val SymInt32, s *Solver) SymArrayInt32 {
	return SymArrayInt32{C.boolector_write(s.btor, a.array, idx.BV, val.BV)}
}
