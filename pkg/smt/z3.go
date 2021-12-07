//go:build z3

package smt

import (
	"github.com/zyedidia/rvsym/pkg/smt/z3"
)

type Solver struct {
	ctx    *z3.Context
	solver *z3.Solver

	sorts
}

type Model struct {
	model *z3.Model
}

type sorts struct {
	sortBool       z3.Sort
	sortInt8       z3.Sort
	sortInt16      z3.Sort
	sortInt32      z3.Sort
	sortInt64      z3.Sort
	sortArrayInt32 z3.Sort
}

func initSorts(ctx *z3.Context) sorts {
	var s sorts
	s.sortBool = ctx.BoolSort()
	s.sortInt8 = ctx.BVSort(8)
	s.sortInt16 = ctx.BVSort(16)
	s.sortInt32 = ctx.BVSort(32)
	s.sortInt64 = ctx.BVSort(64)
	s.sortArrayInt32 = ctx.ArraySort(s.sortInt32, s.sortInt32)
	return s
}

func NewSolver() *Solver {
	ctx := z3.NewContext(nil)
	return &Solver{
		ctx:    ctx,
		solver: z3.NewSolver(ctx),
		sorts:  initSorts(ctx),
	}
}

func (s *Solver) Push() {
	s.solver.Push()
}

func (s *Solver) Pop() {
	s.solver.Pop()
}

func (s *Solver) Assert(b Bool) {
	s.solver.Assert(b.Sym(s).Bool)
}

func (s *Solver) Check(model bool) CheckResult {
	sat, err := s.solver.Check()
	if sat && err == nil {
		return Sat
	} else if !sat && err == nil {
		return Unsat
	}
	return Unknown
}

func (s *Solver) Model() Model {
	return Model{s.solver.Model()}
}

func (m Model) Eval(a Int32) int32 {
	if a.Concrete() {
		return a.C
	}
	c := m.model.Eval(a.S, true).(z3.BV)
	val, ok, _ := c.AsInt64()
	if !ok {
		panic("model evaluation produced non-concrete value " + c.String())
	}
	return int32(val)
}

func (s *Solver) AnyInt32() Int32 {
	return Int32{S: SymInt32{s.ctx.FreshConst("", s.sortInt32).(z3.BV)}}
}

func (s *Solver) AnyArrayInt32(base, length uint32) ArrayInt32 {
	return ArrayInt32{
		base:   base,
		length: length,
		S:      SymArrayInt32{s.ctx.FreshConst("", s.sortArrayInt32).(z3.Array)},
	}
}

func (s *Solver) ToSymInt32(c int32) SymInt32 {
	return SymInt32{s.ctx.FromInt(int64(c), s.sortInt32).(z3.BV)}
}
func (s *Solver) ToSymInt64(c int64) SymInt64 {
	return SymInt64{s.ctx.FromInt(c, s.sortInt64).(z3.BV)}
}
func (s *Solver) ToSymBool(c bool) SymBool {
	return SymBool{s.ctx.FromBool(c)}
}

type SymInt8 struct {
	z3.BV
}

func (a SymInt8) Valid() bool {
	return a.Context() != nil
}

func (a SymInt8) ToInt32s(s *Solver) SymInt32 {
	return SymInt32{a.BV.SignExtend(24)}
}
func (a SymInt8) ToInt32z(s *Solver) SymInt32 {
	return SymInt32{a.BV.ZeroExtend(24)}
}

type SymInt16 struct {
	z3.BV
}

func (a SymInt16) Valid() bool {
	return a.Context() != nil
}

func (a SymInt16) ToInt32s(s *Solver) SymInt32 {
	return SymInt32{a.BV.SignExtend(16)}
}
func (a SymInt16) ToInt32z(s *Solver) SymInt32 {
	return SymInt32{a.BV.ZeroExtend(16)}
}

type SymInt64 struct {
	z3.BV
}

func (a SymInt64) Valid() bool {
	return a.Context() != nil
}

func (a SymInt64) Lower32(s *Solver) SymInt32 {
	return SymInt32{a.BV.Extract(31, 0)}
}
func (a SymInt64) Upper32(s *Solver) SymInt32 {
	return SymInt32{a.BV.Extract(63, 32)}
}
func (a SymInt64) Mul(b SymInt64, s *Solver) SymInt64 {
	return SymInt64{a.BV.Mul(b.BV)}
}

type SymBool struct {
	z3.Bool
}

func (a SymBool) Valid() bool {
	return a.Context() != nil
}

func (a SymBool) Not(s *Solver) SymBool {
	return SymBool{a.Bool.Not()}
}

type SymInt32 struct {
	z3.BV
}

func (a SymInt32) Valid() bool {
	return a.Context() != nil
}

func (a SymInt32) Add(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{a.BV.Add(b.BV)}
}
func (a SymInt32) Sub(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{a.BV.Sub(b.BV)}
}
func (a SymInt32) And(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{a.BV.And(b.BV)}
}
func (a SymInt32) Or(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{a.BV.Or(b.BV)}
}
func (a SymInt32) Xor(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{a.BV.Xor(b.BV)}
}
func (a SymInt32) Sll(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{a.BV.Lsh(b.BV)}
}
func (a SymInt32) Srl(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{a.BV.URsh(b.BV)}
}
func (a SymInt32) Sra(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{a.BV.SRsh(b.BV)}
}
func (a SymInt32) Slt(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{
		a.BV.SLT(b.BV).IfThenElse(
			s.ctx.FromInt(1, s.sortInt32),
			s.ctx.FromInt(0, s.sortInt32),
		).(z3.BV),
	}
}
func (a SymInt32) Ult(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{
		a.BV.ULT(b.BV).IfThenElse(
			s.ctx.FromInt(1, s.sortInt32),
			s.ctx.FromInt(0, s.sortInt32),
		).(z3.BV),
	}
}
func (a SymInt32) Mul(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{a.BV.Mul(b.BV)}
}
func (a SymInt32) Div(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{a.BV.SDiv(b.BV)}
}
func (a SymInt32) Divu(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{a.BV.UDiv(b.BV)}
}
func (a SymInt32) Rem(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{a.BV.SRem(b.BV)}
}
func (a SymInt32) Remu(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{a.BV.URem(b.BV)}
}
func (a SymInt32) Not(s *Solver) SymInt32 {
	return SymInt32{a.BV.Not()}
}
func (a SymInt32) Eqz(s *Solver) SymBool {
	return SymBool{a.BV.Eq(s.ctx.FromInt(0, s.sortInt32).(z3.BV))}
}
func (a SymInt32) NEqz(s *Solver) SymBool {
	return SymBool{a.BV.NE(s.ctx.FromInt(0, s.sortInt32).(z3.BV))}
}
func (a SymInt32) Eqb(b SymInt32, s *Solver) SymBool {
	return SymBool{a.BV.Eq(b.BV)}
}
func (a SymInt32) NEqb(b SymInt32, s *Solver) SymBool {
	return SymBool{a.BV.NE(b.BV)}
}
func (a SymInt32) Sltb(b SymInt32, s *Solver) SymBool {
	return SymBool{a.BV.SLT(b.BV)}
}
func (a SymInt32) Sgeb(b SymInt32, s *Solver) SymBool {
	return SymBool{a.BV.SGE(b.BV)}
}
func (a SymInt32) Ultb(b SymInt32, s *Solver) SymBool {
	return SymBool{a.BV.ULT(b.BV)}
}
func (a SymInt32) Ugeb(b SymInt32, s *Solver) SymBool {
	return SymBool{a.BV.UGE(b.BV)}
}
func (a SymInt32) ToInt8(s *Solver) SymInt8 {
	return SymInt8{a.BV.Extract(7, 0)}
}
func (a SymInt32) ToInt16(s *Solver) SymInt16 {
	return SymInt16{a.BV.Extract(15, 0)}
}
func (a SymInt32) ToInt64s(s *Solver) SymInt64 {
	return SymInt64{a.BV.SignExtend(32)}
}
func (a SymInt32) ToInt64z(s *Solver) SymInt64 {
	return SymInt64{a.BV.ZeroExtend(32)}
}

type SymArrayInt32 struct {
	array z3.Array
}

func (a SymArrayInt32) Select(idx SymInt32, s *Solver) SymInt32 {
	return SymInt32{a.array.Select(idx.BV).(z3.BV)}
}

func (a SymArrayInt32) Store(idx SymInt32, val SymInt32, s *Solver) SymArrayInt32 {
	// arr := s.ctx.FreshConst("", s.sortArrayInt32).(z3.Array)
	// s.Assert(Bool{S: SymBool{arr.Eq(a.array.Store(idx.BV, val.BV))}})
	// return SymArrayInt32{arr}
	return SymArrayInt32{a.array.Store(idx.BV, val.BV)}
	// s.Assert(Bool{S: SymBool{a.array.Select(idx.BV).(z3.BV).Eq(val.BV)}})
	return a
}
