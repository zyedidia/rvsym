//go:build bitwuzla

package smt

//#cgo CFLAGS: -I./deps/include
//#cgo LDFLAGS: -L./deps/lib -lbitwuzla -lcadical -lm -lgmp -lbtor2parser -lstdc++
//#include <stdlib.h>
//#include "bitwuzla.h"
import "C"

type Solver struct {
	bzla *C.Bitwuzla

	sorts
}

type Model struct {
	bzla *C.Bitwuzla
}

type sorts struct {
	sortBool       *C.BitwuzlaSort
	sortInt8       *C.BitwuzlaSort
	sortInt16      *C.BitwuzlaSort
	sortInt32      *C.BitwuzlaSort
	sortInt64      *C.BitwuzlaSort
	sortArrayInt32 *C.BitwuzlaSort
}

func initSorts(bzla *C.Bitwuzla) sorts {
	var s sorts
	s.sortBool = C.bitwuzla_mk_bv_sort(bzla, 1)
	s.sortInt8 = C.bitwuzla_mk_bv_sort(bzla, 8)
	s.sortInt16 = C.bitwuzla_mk_bv_sort(bzla, 16)
	s.sortInt32 = C.bitwuzla_mk_bv_sort(bzla, 32)
	s.sortInt64 = C.bitwuzla_mk_bv_sort(bzla, 64)
	s.sortArrayInt32 = C.bitwuzla_mk_array_sort(bzla, s.sortInt32, s.sortInt32)
	return s
}

func NewSolver() *Solver {
	bzla := C.bitwuzla_new()
	C.bitwuzla_set_option(bzla, C.BITWUZLA_OPT_PRODUCE_MODELS, 1)
	C.bitwuzla_set_option(bzla, C.BITWUZLA_OPT_OUTPUT_NUMBER_FORMAT, C.BITWUZLA_BV_BASE_DEC)
	C.bitwuzla_set_option(bzla, C.BITWUZLA_OPT_INCREMENTAL, 1)

	return &Solver{
		bzla:  bzla,
		sorts: initSorts(bzla),
	}
}
