package st

import (
	"fmt"
	"math/big"

	"github.com/zyedidia/rvsym/pkg/z3/z3"
)

type sorts struct {
	sortBool    z3.Sort
	sortInt     z3.Sort
	sortInt8    z3.Sort
	sortInt16   z3.Sort
	sortInt32   z3.Sort
	sortInt64   z3.Sort
	sortUint    z3.Sort
	sortUint8   z3.Sort
	sortUint16  z3.Sort
	sortUint32  z3.Sort
	sortUint64  z3.Sort
	sortUintptr z3.Sort
	sortInteger z3.Sort
	sortReal    z3.Sort
}

func initSorts(s *sorts, ctx *z3.Context) {
	s.sortBool = ctx.BoolSort()
	s.sortInt = ctx.BVSort(64)
	s.sortInt8 = ctx.BVSort(8)
	s.sortInt16 = ctx.BVSort(16)
	s.sortInt32 = ctx.BVSort(32)
	s.sortInt64 = ctx.BVSort(64)
	s.sortUint = ctx.BVSort(64)
	s.sortUint8 = ctx.BVSort(8)
	s.sortUint16 = ctx.BVSort(16)
	s.sortUint32 = ctx.BVSort(32)
	s.sortUint64 = ctx.BVSort(64)
	s.sortUintptr = ctx.BVSort(64)
	s.sortInteger = ctx.IntSort()
	s.sortReal = ctx.RealSort()
}

// Bool implements symbolic bool values.
type Bool struct {
	C bool
	S z3.Bool
}

// AnyBool returns an unconstrained symbolic Bool.
func AnyBool(ctx *z3.Context, name string) Bool {
	cache := getCache(ctx)
	sym := cache.z3.FreshConst(name, cache.sortBool).(z3.Bool)
	return Bool{S: sym}
}

// String returns x as a string.
func (x Bool) String() string {
	if x.IsConcrete() {
		return fmt.Sprint(x.C)
	}
	return x.S.String()
}

// IsConcrete returns true if x is concrete.
func (x Bool) IsConcrete() bool {
	return x.S.Context() == nil
}

func (x Bool) ToInt() Int {
	if x.IsConcrete() {
		if x.C {
			return Int{C: 1}
		}
		return Int{C: 0}
	}
	ctx := x.S.Context()
	cache := getCache(ctx)
	return Int{S: x.S.IfThenElse(ctx.FromInt(1, cache.sortInt), ctx.FromInt(0, cache.sortInt)).(z3.BV)}
}

func (x Bool) ToInt32() Int32 {
	if x.IsConcrete() {
		if x.C {
			return Int32{C: 1}
		}
		return Int32{C: 0}
	}
	ctx := x.S.Context()
	cache := getCache(ctx)
	return Int32{S: x.S.IfThenElse(ctx.FromInt(1, cache.sortInt32), ctx.FromInt(0, cache.sortInt32)).(z3.BV)}
}

func (x Bool) ToUint32() Uint32 {
	if x.IsConcrete() {
		if x.C {
			return Uint32{C: 1}
		}
		return Uint32{C: 0}
	}
	ctx := x.S.Context()
	cache := getCache(ctx)
	return Uint32{S: x.S.IfThenElse(ctx.FromInt(1, cache.sortUint32), ctx.FromInt(0, cache.sortUint32)).(z3.BV)}
}

// Eval returns x's concrete value in model m.
// This also evaluates x with model completion.
func (x Bool) Eval(m *z3.Model) bool {
	if x.IsConcrete() {
		return x.C
	}
	c := m.Eval(x.S, true).(z3.Bool)
	val, ok := c.AsBool()
	if !ok {
		panic("model evaluation produced non-concrete value " + c.String())
	}
	return (bool)(val)
}

// sym returns x's symbolic value, creating it if necessary.
func (x Bool) sym(c *cache) z3.Bool {
	if !x.IsConcrete() {
		return x.S
	}
	return c.z3.FromBool(x.C)
}

func (x Bool) And(y Bool) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C && y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).And(y.sym(cache))}
}

func (x Bool) Or(y Bool) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C || y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).Or(y.sym(cache))}
}

func (x Bool) Eq(y Bool) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C == y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).Eq(y.sym(cache))}
}

func (x Bool) NE(y Bool) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C != y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).NE(y.sym(cache))}
}

func (x Bool) Not() Bool {
	if x.IsConcrete() {
		return Bool{C: !x.C}
	}
	return Bool{S: x.S.Not()}
}

// Int implements symbolic int values.
type Int struct {
	C int
	S z3.BV
}

// AnyInt returns an unconstrained symbolic Int.
func AnyInt(ctx *z3.Context, name string) Int {
	cache := getCache(ctx)
	sym := cache.z3.FreshConst(name, cache.sortInt).(z3.BV)
	return Int{S: sym}
}

// String returns x as a string.
func (x Int) String() string {
	if x.IsConcrete() {
		return fmt.Sprint(x.C)
	}
	return x.S.String()
}

// IsConcrete returns true if x is concrete.
func (x Int) IsConcrete() bool {
	return x.S.Context() == nil
}

// Eval returns x's concrete value in model m.
// This also evaluates x with model completion.
func (x Int) Eval(m *z3.Model) int {
	if x.IsConcrete() {
		return x.C
	}
	c := m.Eval(x.S, true).(z3.BV)
	val, ok, _ := c.AsInt64()
	if !ok {
		panic("model evaluation produced non-concrete value " + c.String())
	}
	return (int)(val)
}

// sym returns x's symbolic value, creating it if necessary.
func (x Int) sym(c *cache) z3.BV {
	if !x.IsConcrete() {
		return x.S
	}
	return c.z3.FromInt(int64(x.C), c.sortInt).(z3.BV)
}

func (x Int) Add(y Int) Int {
	if x.IsConcrete() && y.IsConcrete() {
		return Int{C: x.C + y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int{S: x.sym(cache).Add(y.sym(cache))}
}

func (x Int) Sub(y Int) Int {
	if x.IsConcrete() && y.IsConcrete() {
		return Int{C: x.C - y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int{S: x.sym(cache).Sub(y.sym(cache))}
}

func (x Int) Mul(y Int) Int {
	if x.IsConcrete() && y.IsConcrete() {
		return Int{C: x.C * y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int{S: x.sym(cache).Mul(y.sym(cache))}
}

func (x Int) Quo(y Int) Int {
	if x.IsConcrete() && y.IsConcrete() {
		return Int{C: x.C / y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int{S: x.sym(cache).SDiv(y.sym(cache))}
}

func (x Int) Rem(y Int) Int {
	if x.IsConcrete() && y.IsConcrete() {
		return Int{C: x.C % y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int{S: x.sym(cache).SRem(y.sym(cache))}
}

func (x Int) And(y Int) Int {
	if x.IsConcrete() && y.IsConcrete() {
		return Int{C: x.C & y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int{S: x.sym(cache).And(y.sym(cache))}
}

func (x Int) Or(y Int) Int {
	if x.IsConcrete() && y.IsConcrete() {
		return Int{C: x.C | y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int{S: x.sym(cache).Or(y.sym(cache))}
}

func (x Int) Xor(y Int) Int {
	if x.IsConcrete() && y.IsConcrete() {
		return Int{C: x.C ^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int{S: x.sym(cache).Xor(y.sym(cache))}
}

func (x Int) Lsh(y Uint64) Int {
	if x.IsConcrete() && y.IsConcrete() {
		return Int{C: x.C << y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 64 {
			return Int{C: 0}
		}
	}
	rs = y.sym(cache)
	return Int{S: x.sym(cache).Lsh(rs)}
}

func (x Int) Rsh(y Uint64) Int {
	if x.IsConcrete() && y.IsConcrete() {
		return Int{C: x.C >> y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 64 {
			return Int{C: 0}
		}
	}
	rs = y.sym(cache)
	return Int{S: x.sym(cache).SRsh(rs)}
}

func (x Int) AndNot(y Int) Int {
	if x.IsConcrete() && y.IsConcrete() {
		return Int{C: x.C &^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int{S: x.sym(cache).And(y.Not().sym(cache))}
}

func (x Int) Eq(y Int) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C == y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).Eq(y.sym(cache))}
}

func (x Int) NE(y Int) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C != y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).NE(y.sym(cache))}
}

func (x Int) LT(y Int) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C < y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SLT(y.sym(cache))}
}

func (x Int) LE(y Int) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C <= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SLE(y.sym(cache))}
}

func (x Int) GT(y Int) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C > y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SGT(y.sym(cache))}
}

func (x Int) GE(y Int) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C >= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SGE(y.sym(cache))}
}

func (x Int) Neg() Int {
	if x.IsConcrete() {
		return Int{C: -x.C}
	}
	return Int{S: x.S.Neg()}
}

func (x Int) Not() Int {
	if x.IsConcrete() {
		return Int{C: ^x.C}
	}
	return Int{S: x.S.Not()}
}

func (x Int) ToInt() Int {
	if x.IsConcrete() {
		return Int{C: int(x.C)}
	}
	return Int{S: x.S.SignExtend(0)}
}

func (x Int) ToInt8() Int8 {
	if x.IsConcrete() {
		return Int8{C: int8(x.C)}
	}
	return Int8{S: x.S.Extract(7, 0)}
}

func (x Int) ToInt16() Int16 {
	if x.IsConcrete() {
		return Int16{C: int16(x.C)}
	}
	return Int16{S: x.S.Extract(15, 0)}
}

func (x Int) ToInt32() Int32 {
	if x.IsConcrete() {
		return Int32{C: int32(x.C)}
	}
	return Int32{S: x.S.Extract(31, 0)}
}

func (x Int) ToInt64() Int64 {
	if x.IsConcrete() {
		return Int64{C: int64(x.C)}
	}
	return Int64{S: x.S.SignExtend(0)}
}

func (x Int) ToUint() Uint {
	if x.IsConcrete() {
		return Uint{C: uint(x.C)}
	}
	return Uint{S: x.S.SignExtend(0)}
}

func (x Int) ToUint8() Uint8 {
	if x.IsConcrete() {
		return Uint8{C: uint8(x.C)}
	}
	return Uint8{S: x.S.Extract(7, 0)}
}

func (x Int) ToUint16() Uint16 {
	if x.IsConcrete() {
		return Uint16{C: uint16(x.C)}
	}
	return Uint16{S: x.S.Extract(15, 0)}
}

func (x Int) ToUint32() Uint32 {
	if x.IsConcrete() {
		return Uint32{C: uint32(x.C)}
	}
	return Uint32{S: x.S.Extract(31, 0)}
}

func (x Int) ToUint64() Uint64 {
	if x.IsConcrete() {
		return Uint64{C: uint64(x.C)}
	}
	return Uint64{S: x.S.SignExtend(0)}
}

func (x Int) ToUintptr() Uintptr {
	if x.IsConcrete() {
		return Uintptr{C: uintptr(x.C)}
	}
	return Uintptr{S: x.S.SignExtend(0)}
}

// Int8 implements symbolic int8 values.
type Int8 struct {
	C int8
	S z3.BV
}

// AnyInt8 returns an unconstrained symbolic Int8.
func AnyInt8(ctx *z3.Context, name string) Int8 {
	cache := getCache(ctx)
	sym := cache.z3.FreshConst(name, cache.sortInt8).(z3.BV)
	return Int8{S: sym}
}

// String returns x as a string.
func (x Int8) String() string {
	if x.IsConcrete() {
		return fmt.Sprint(x.C)
	}
	return x.S.String()
}

// IsConcrete returns true if x is concrete.
func (x Int8) IsConcrete() bool {
	return x.S.Context() == nil
}

// Eval returns x's concrete value in model m.
// This also evaluates x with model completion.
func (x Int8) Eval(m *z3.Model) int8 {
	if x.IsConcrete() {
		return x.C
	}
	c := m.Eval(x.S, true).(z3.BV)
	val, ok, _ := c.AsInt64()
	if !ok {
		panic("model evaluation produced non-concrete value " + c.String())
	}
	return (int8)(val)
}

// sym returns x's symbolic value, creating it if necessary.
func (x Int8) sym(c *cache) z3.BV {
	if !x.IsConcrete() {
		return x.S
	}
	return c.z3.FromInt(int64(x.C), c.sortInt8).(z3.BV)
}

func (x Int8) Add(y Int8) Int8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int8{C: x.C + y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int8{S: x.sym(cache).Add(y.sym(cache))}
}

func (x Int8) Sub(y Int8) Int8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int8{C: x.C - y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int8{S: x.sym(cache).Sub(y.sym(cache))}
}

func (x Int8) Mul(y Int8) Int8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int8{C: x.C * y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int8{S: x.sym(cache).Mul(y.sym(cache))}
}

func (x Int8) Quo(y Int8) Int8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int8{C: x.C / y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int8{S: x.sym(cache).SDiv(y.sym(cache))}
}

func (x Int8) Rem(y Int8) Int8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int8{C: x.C % y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int8{S: x.sym(cache).SRem(y.sym(cache))}
}

func (x Int8) And(y Int8) Int8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int8{C: x.C & y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int8{S: x.sym(cache).And(y.sym(cache))}
}

func (x Int8) Or(y Int8) Int8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int8{C: x.C | y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int8{S: x.sym(cache).Or(y.sym(cache))}
}

func (x Int8) Xor(y Int8) Int8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int8{C: x.C ^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int8{S: x.sym(cache).Xor(y.sym(cache))}
}

func (x Int8) Lsh(y Uint64) Int8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int8{C: x.C << y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 8 {
			return Int8{C: 0}
		}
		rs = Uint8{C: uint8(y.C)}.sym(cache)
	} else {
		rs = y.sym(cache)
		rs = rs.Extract(63, 7).AnyBits().Concat(rs.Extract(6, 0))
	}
	return Int8{S: x.sym(cache).Lsh(rs)}
}

func (x Int8) Rsh(y Uint64) Int8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int8{C: x.C >> y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 8 {
			return Int8{C: 0}
		}
		rs = Uint8{C: uint8(y.C)}.sym(cache)
	} else {
		rs = y.sym(cache)
		rs = rs.Extract(63, 7).AnyBits().Concat(rs.Extract(6, 0))
	}
	return Int8{S: x.sym(cache).SRsh(rs)}
}

func (x Int8) AndNot(y Int8) Int8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int8{C: x.C &^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int8{S: x.sym(cache).And(y.Not().sym(cache))}
}

func (x Int8) Eq(y Int8) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C == y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).Eq(y.sym(cache))}
}

func (x Int8) NE(y Int8) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C != y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).NE(y.sym(cache))}
}

func (x Int8) LT(y Int8) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C < y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SLT(y.sym(cache))}
}

func (x Int8) LE(y Int8) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C <= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SLE(y.sym(cache))}
}

func (x Int8) GT(y Int8) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C > y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SGT(y.sym(cache))}
}

func (x Int8) GE(y Int8) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C >= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SGE(y.sym(cache))}
}

func (x Int8) Neg() Int8 {
	if x.IsConcrete() {
		return Int8{C: -x.C}
	}
	return Int8{S: x.S.Neg()}
}

func (x Int8) Not() Int8 {
	if x.IsConcrete() {
		return Int8{C: ^x.C}
	}
	return Int8{S: x.S.Not()}
}

func (x Int8) ToInt() Int {
	if x.IsConcrete() {
		return Int{C: int(x.C)}
	}
	return Int{S: x.S.SignExtend(56)}
}

func (x Int8) ToInt8() Int8 {
	if x.IsConcrete() {
		return Int8{C: int8(x.C)}
	}
	return Int8{S: x.S.SignExtend(0)}
}

func (x Int8) ToInt16() Int16 {
	if x.IsConcrete() {
		return Int16{C: int16(x.C)}
	}
	return Int16{S: x.S.SignExtend(8)}
}

func (x Int8) ToInt32() Int32 {
	if x.IsConcrete() {
		return Int32{C: int32(x.C)}
	}
	return Int32{S: x.S.SignExtend(24)}
}

func (x Int8) ToInt64() Int64 {
	if x.IsConcrete() {
		return Int64{C: int64(x.C)}
	}
	return Int64{S: x.S.SignExtend(56)}
}

func (x Int8) ToUint() Uint {
	if x.IsConcrete() {
		return Uint{C: uint(x.C)}
	}
	return Uint{S: x.S.SignExtend(56)}
}

func (x Int8) ToUint8() Uint8 {
	if x.IsConcrete() {
		return Uint8{C: uint8(x.C)}
	}
	return Uint8{S: x.S.SignExtend(0)}
}

func (x Int8) ToUint16() Uint16 {
	if x.IsConcrete() {
		return Uint16{C: uint16(x.C)}
	}
	return Uint16{S: x.S.SignExtend(8)}
}

func (x Int8) ToUint32() Uint32 {
	if x.IsConcrete() {
		return Uint32{C: uint32(x.C)}
	}
	return Uint32{S: x.S.SignExtend(24)}
}

func (x Int8) ToUint64() Uint64 {
	if x.IsConcrete() {
		return Uint64{C: uint64(x.C)}
	}
	return Uint64{S: x.S.SignExtend(56)}
}

func (x Int8) ToUintptr() Uintptr {
	if x.IsConcrete() {
		return Uintptr{C: uintptr(x.C)}
	}
	return Uintptr{S: x.S.SignExtend(56)}
}

// Int16 implements symbolic int16 values.
type Int16 struct {
	C int16
	S z3.BV
}

// AnyInt16 returns an unconstrained symbolic Int16.
func AnyInt16(ctx *z3.Context, name string) Int16 {
	cache := getCache(ctx)
	sym := cache.z3.FreshConst(name, cache.sortInt16).(z3.BV)
	return Int16{S: sym}
}

// String returns x as a string.
func (x Int16) String() string {
	if x.IsConcrete() {
		return fmt.Sprint(x.C)
	}
	return x.S.String()
}

// IsConcrete returns true if x is concrete.
func (x Int16) IsConcrete() bool {
	return x.S.Context() == nil
}

// Eval returns x's concrete value in model m.
// This also evaluates x with model completion.
func (x Int16) Eval(m *z3.Model) int16 {
	if x.IsConcrete() {
		return x.C
	}
	c := m.Eval(x.S, true).(z3.BV)
	val, ok, _ := c.AsInt64()
	if !ok {
		panic("model evaluation produced non-concrete value " + c.String())
	}
	return (int16)(val)
}

// sym returns x's symbolic value, creating it if necessary.
func (x Int16) sym(c *cache) z3.BV {
	if !x.IsConcrete() {
		return x.S
	}
	return c.z3.FromInt(int64(x.C), c.sortInt16).(z3.BV)
}

func (x Int16) Add(y Int16) Int16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int16{C: x.C + y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int16{S: x.sym(cache).Add(y.sym(cache))}
}

func (x Int16) Sub(y Int16) Int16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int16{C: x.C - y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int16{S: x.sym(cache).Sub(y.sym(cache))}
}

func (x Int16) Mul(y Int16) Int16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int16{C: x.C * y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int16{S: x.sym(cache).Mul(y.sym(cache))}
}

func (x Int16) Quo(y Int16) Int16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int16{C: x.C / y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int16{S: x.sym(cache).SDiv(y.sym(cache))}
}

func (x Int16) Rem(y Int16) Int16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int16{C: x.C % y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int16{S: x.sym(cache).SRem(y.sym(cache))}
}

func (x Int16) And(y Int16) Int16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int16{C: x.C & y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int16{S: x.sym(cache).And(y.sym(cache))}
}

func (x Int16) Or(y Int16) Int16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int16{C: x.C | y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int16{S: x.sym(cache).Or(y.sym(cache))}
}

func (x Int16) Xor(y Int16) Int16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int16{C: x.C ^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int16{S: x.sym(cache).Xor(y.sym(cache))}
}

func (x Int16) Lsh(y Uint64) Int16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int16{C: x.C << y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 16 {
			return Int16{C: 0}
		}
		rs = Uint16{C: uint16(y.C)}.sym(cache)
	} else {
		rs = y.sym(cache)
		rs = rs.Extract(63, 15).AnyBits().Concat(rs.Extract(14, 0))
	}
	return Int16{S: x.sym(cache).Lsh(rs)}
}

func (x Int16) Rsh(y Uint64) Int16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int16{C: x.C >> y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 16 {
			return Int16{C: 0}
		}
		rs = Uint16{C: uint16(y.C)}.sym(cache)
	} else {
		rs = y.sym(cache)
		rs = rs.Extract(63, 15).AnyBits().Concat(rs.Extract(14, 0))
	}
	return Int16{S: x.sym(cache).SRsh(rs)}
}

func (x Int16) AndNot(y Int16) Int16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int16{C: x.C &^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int16{S: x.sym(cache).And(y.Not().sym(cache))}
}

func (x Int16) Eq(y Int16) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C == y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).Eq(y.sym(cache))}
}

func (x Int16) NE(y Int16) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C != y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).NE(y.sym(cache))}
}

func (x Int16) LT(y Int16) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C < y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SLT(y.sym(cache))}
}

func (x Int16) LE(y Int16) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C <= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SLE(y.sym(cache))}
}

func (x Int16) GT(y Int16) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C > y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SGT(y.sym(cache))}
}

func (x Int16) GE(y Int16) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C >= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SGE(y.sym(cache))}
}

func (x Int16) Neg() Int16 {
	if x.IsConcrete() {
		return Int16{C: -x.C}
	}
	return Int16{S: x.S.Neg()}
}

func (x Int16) Not() Int16 {
	if x.IsConcrete() {
		return Int16{C: ^x.C}
	}
	return Int16{S: x.S.Not()}
}

func (x Int16) ToInt() Int {
	if x.IsConcrete() {
		return Int{C: int(x.C)}
	}
	return Int{S: x.S.SignExtend(48)}
}

func (x Int16) ToInt8() Int8 {
	if x.IsConcrete() {
		return Int8{C: int8(x.C)}
	}
	return Int8{S: x.S.Extract(7, 0)}
}

func (x Int16) ToInt16() Int16 {
	if x.IsConcrete() {
		return Int16{C: int16(x.C)}
	}
	return Int16{S: x.S.SignExtend(0)}
}

func (x Int16) ToInt32() Int32 {
	if x.IsConcrete() {
		return Int32{C: int32(x.C)}
	}
	return Int32{S: x.S.SignExtend(16)}
}

func (x Int16) ToInt64() Int64 {
	if x.IsConcrete() {
		return Int64{C: int64(x.C)}
	}
	return Int64{S: x.S.SignExtend(48)}
}

func (x Int16) ToUint() Uint {
	if x.IsConcrete() {
		return Uint{C: uint(x.C)}
	}
	return Uint{S: x.S.SignExtend(48)}
}

func (x Int16) ToUint8() Uint8 {
	if x.IsConcrete() {
		return Uint8{C: uint8(x.C)}
	}
	return Uint8{S: x.S.Extract(7, 0)}
}

func (x Int16) ToUint16() Uint16 {
	if x.IsConcrete() {
		return Uint16{C: uint16(x.C)}
	}
	return Uint16{S: x.S.SignExtend(0)}
}

func (x Int16) ToUint32() Uint32 {
	if x.IsConcrete() {
		return Uint32{C: uint32(x.C)}
	}
	return Uint32{S: x.S.SignExtend(16)}
}

func (x Int16) ToUint64() Uint64 {
	if x.IsConcrete() {
		return Uint64{C: uint64(x.C)}
	}
	return Uint64{S: x.S.SignExtend(48)}
}

func (x Int16) ToUintptr() Uintptr {
	if x.IsConcrete() {
		return Uintptr{C: uintptr(x.C)}
	}
	return Uintptr{S: x.S.SignExtend(48)}
}

// Int32 implements symbolic int32 values.
type Int32 struct {
	C int32
	S z3.BV
}

// AnyInt32 returns an unconstrained symbolic Int32.
func AnyInt32(ctx *z3.Context, name string) Int32 {
	cache := getCache(ctx)
	sym := cache.z3.FreshConst(name, cache.sortInt32).(z3.BV)
	return Int32{S: sym}
}

// String returns x as a string.
func (x Int32) String() string {
	if x.IsConcrete() {
		return fmt.Sprint(x.C)
	}
	return x.S.String()
}

// IsConcrete returns true if x is concrete.
func (x Int32) IsConcrete() bool {
	return x.S.Context() == nil
}

// Eval returns x's concrete value in model m.
// This also evaluates x with model completion.
func (x Int32) Eval(m *z3.Model) int32 {
	if x.IsConcrete() {
		return x.C
	}
	c := m.Eval(x.S, true).(z3.BV)
	val, ok, _ := c.AsInt64()
	if !ok {
		panic("model evaluation produced non-concrete value " + c.String())
	}
	return (int32)(val)
}

// sym returns x's symbolic value, creating it if necessary.
func (x Int32) sym(c *cache) z3.BV {
	if !x.IsConcrete() {
		return x.S
	}
	return c.z3.FromInt(int64(x.C), c.sortInt32).(z3.BV)
}

func (x Int32) Add(y Int32) Int32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int32{C: x.C + y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int32{S: x.sym(cache).Add(y.sym(cache))}
}

func (x Int32) Sub(y Int32) Int32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int32{C: x.C - y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int32{S: x.sym(cache).Sub(y.sym(cache))}
}

func (x Int32) Mul(y Int32) Int32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int32{C: x.C * y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int32{S: x.sym(cache).Mul(y.sym(cache))}
}

func (x Int32) Quo(y Int32) Int32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int32{C: x.C / y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int32{S: x.sym(cache).SDiv(y.sym(cache))}
}

func (x Int32) Rem(y Int32) Int32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int32{C: x.C % y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int32{S: x.sym(cache).SRem(y.sym(cache))}
}

func (x Int32) And(y Int32) Int32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int32{C: x.C & y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int32{S: x.sym(cache).And(y.sym(cache))}
}

func (x Int32) Or(y Int32) Int32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int32{C: x.C | y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int32{S: x.sym(cache).Or(y.sym(cache))}
}

func (x Int32) Xor(y Int32) Int32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int32{C: x.C ^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int32{S: x.sym(cache).Xor(y.sym(cache))}
}

func (x Int32) Lsh(y Uint64) Int32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int32{C: x.C << y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 32 {
			return Int32{C: 0}
		}
		rs = Uint32{C: uint32(y.C)}.sym(cache)
	} else {
		rs = y.sym(cache)
		rs = rs.Extract(63, 31).AnyBits().Concat(rs.Extract(30, 0))
	}
	return Int32{S: x.sym(cache).Lsh(rs)}
}

func (x Int32) Rsh(y Uint64) Int32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int32{C: x.C >> y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 32 {
			return Int32{C: 0}
		}
		rs = Uint32{C: uint32(y.C)}.sym(cache)
	} else {
		rs = y.sym(cache)
		rs = rs.Extract(63, 31).AnyBits().Concat(rs.Extract(30, 0))
	}
	return Int32{S: x.sym(cache).SRsh(rs)}
}

func (x Int32) AndNot(y Int32) Int32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int32{C: x.C &^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int32{S: x.sym(cache).And(y.Not().sym(cache))}
}

func (x Int32) Eq(y Int32) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C == y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).Eq(y.sym(cache))}
}

func (x Int32) NE(y Int32) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C != y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).NE(y.sym(cache))}
}

func (x Int32) LT(y Int32) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C < y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SLT(y.sym(cache))}
}

func (x Int32) LE(y Int32) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C <= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SLE(y.sym(cache))}
}

func (x Int32) GT(y Int32) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C > y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SGT(y.sym(cache))}
}

func (x Int32) GE(y Int32) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C >= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SGE(y.sym(cache))}
}

func (x Int32) Neg() Int32 {
	if x.IsConcrete() {
		return Int32{C: -x.C}
	}
	return Int32{S: x.S.Neg()}
}

func (x Int32) Not() Int32 {
	if x.IsConcrete() {
		return Int32{C: ^x.C}
	}
	return Int32{S: x.S.Not()}
}

func (x Int32) ToInt() Int {
	if x.IsConcrete() {
		return Int{C: int(x.C)}
	}
	return Int{S: x.S.SignExtend(32)}
}

func (x Int32) ToInt8() Int8 {
	if x.IsConcrete() {
		return Int8{C: int8(x.C)}
	}
	return Int8{S: x.S.Extract(7, 0)}
}

func (x Int32) ToInt16() Int16 {
	if x.IsConcrete() {
		return Int16{C: int16(x.C)}
	}
	return Int16{S: x.S.Extract(15, 0)}
}

func (x Int32) ToInt32() Int32 {
	if x.IsConcrete() {
		return Int32{C: int32(x.C)}
	}
	return Int32{S: x.S.SignExtend(0)}
}

func (x Int32) ToInt64() Int64 {
	if x.IsConcrete() {
		return Int64{C: int64(x.C)}
	}
	return Int64{S: x.S.SignExtend(32)}
}

func (x Int32) ToUint() Uint {
	if x.IsConcrete() {
		return Uint{C: uint(x.C)}
	}
	return Uint{S: x.S.SignExtend(32)}
}

func (x Int32) ToUint8() Uint8 {
	if x.IsConcrete() {
		return Uint8{C: uint8(x.C)}
	}
	return Uint8{S: x.S.Extract(7, 0)}
}

func (x Int32) ToUint16() Uint16 {
	if x.IsConcrete() {
		return Uint16{C: uint16(x.C)}
	}
	return Uint16{S: x.S.Extract(15, 0)}
}

func (x Int32) ToUint32() Uint32 {
	if x.IsConcrete() {
		return Uint32{C: uint32(x.C)}
	}
	return Uint32{S: x.S.SignExtend(0)}
}

func (x Int32) ToUint64() Uint64 {
	if x.IsConcrete() {
		return Uint64{C: uint64(x.C)}
	}
	return Uint64{S: x.S.SignExtend(32)}
}

func (x Int32) ToUintptr() Uintptr {
	if x.IsConcrete() {
		return Uintptr{C: uintptr(x.C)}
	}
	return Uintptr{S: x.S.SignExtend(32)}
}

// Int64 implements symbolic int64 values.
type Int64 struct {
	C int64
	S z3.BV
}

// AnyInt64 returns an unconstrained symbolic Int64.
func AnyInt64(ctx *z3.Context, name string) Int64 {
	cache := getCache(ctx)
	sym := cache.z3.FreshConst(name, cache.sortInt64).(z3.BV)
	return Int64{S: sym}
}

// String returns x as a string.
func (x Int64) String() string {
	if x.IsConcrete() {
		return fmt.Sprint(x.C)
	}
	return x.S.String()
}

// IsConcrete returns true if x is concrete.
func (x Int64) IsConcrete() bool {
	return x.S.Context() == nil
}

// Eval returns x's concrete value in model m.
// This also evaluates x with model completion.
func (x Int64) Eval(m *z3.Model) int64 {
	if x.IsConcrete() {
		return x.C
	}
	c := m.Eval(x.S, true).(z3.BV)
	val, ok, _ := c.AsInt64()
	if !ok {
		panic("model evaluation produced non-concrete value " + c.String())
	}
	return (int64)(val)
}

// sym returns x's symbolic value, creating it if necessary.
func (x Int64) sym(c *cache) z3.BV {
	if !x.IsConcrete() {
		return x.S
	}
	return c.z3.FromInt(int64(x.C), c.sortInt64).(z3.BV)
}

func (x Int64) Add(y Int64) Int64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int64{C: x.C + y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int64{S: x.sym(cache).Add(y.sym(cache))}
}

func (x Int64) Sub(y Int64) Int64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int64{C: x.C - y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int64{S: x.sym(cache).Sub(y.sym(cache))}
}

func (x Int64) Mul(y Int64) Int64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int64{C: x.C * y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int64{S: x.sym(cache).Mul(y.sym(cache))}
}

func (x Int64) Quo(y Int64) Int64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int64{C: x.C / y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int64{S: x.sym(cache).SDiv(y.sym(cache))}
}

func (x Int64) Rem(y Int64) Int64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int64{C: x.C % y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int64{S: x.sym(cache).SRem(y.sym(cache))}
}

func (x Int64) And(y Int64) Int64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int64{C: x.C & y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int64{S: x.sym(cache).And(y.sym(cache))}
}

func (x Int64) Or(y Int64) Int64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int64{C: x.C | y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int64{S: x.sym(cache).Or(y.sym(cache))}
}

func (x Int64) Xor(y Int64) Int64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int64{C: x.C ^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int64{S: x.sym(cache).Xor(y.sym(cache))}
}

func (x Int64) Lsh(y Uint64) Int64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int64{C: x.C << y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 64 {
			return Int64{C: 0}
		}
	}
	rs = y.sym(cache)
	return Int64{S: x.sym(cache).Lsh(rs)}
}

func (x Int64) Rsh(y Uint64) Int64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int64{C: x.C >> y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 64 {
			return Int64{C: 0}
		}
	}
	rs = y.sym(cache)
	return Int64{S: x.sym(cache).SRsh(rs)}
}

func (x Int64) AndNot(y Int64) Int64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Int64{C: x.C &^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Int64{S: x.sym(cache).And(y.Not().sym(cache))}
}

func (x Int64) Eq(y Int64) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C == y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).Eq(y.sym(cache))}
}

func (x Int64) NE(y Int64) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C != y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).NE(y.sym(cache))}
}

func (x Int64) LT(y Int64) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C < y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SLT(y.sym(cache))}
}

func (x Int64) LE(y Int64) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C <= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SLE(y.sym(cache))}
}

func (x Int64) GT(y Int64) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C > y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SGT(y.sym(cache))}
}

func (x Int64) GE(y Int64) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C >= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).SGE(y.sym(cache))}
}

func (x Int64) Neg() Int64 {
	if x.IsConcrete() {
		return Int64{C: -x.C}
	}
	return Int64{S: x.S.Neg()}
}

func (x Int64) Not() Int64 {
	if x.IsConcrete() {
		return Int64{C: ^x.C}
	}
	return Int64{S: x.S.Not()}
}

func (x Int64) Upper32() Int32 {
	if x.IsConcrete() {
		return Int32{C: int32(x.C >> 32)}
	}
	return Int32{S: x.S.Extract(63, 32)}
}

func (x Int64) ToInt() Int {
	if x.IsConcrete() {
		return Int{C: int(x.C)}
	}
	return Int{S: x.S.SignExtend(0)}
}

func (x Int64) ToInt8() Int8 {
	if x.IsConcrete() {
		return Int8{C: int8(x.C)}
	}
	return Int8{S: x.S.Extract(7, 0)}
}

func (x Int64) ToInt16() Int16 {
	if x.IsConcrete() {
		return Int16{C: int16(x.C)}
	}
	return Int16{S: x.S.Extract(15, 0)}
}

func (x Int64) ToInt32() Int32 {
	if x.IsConcrete() {
		return Int32{C: int32(x.C)}
	}
	return Int32{S: x.S.Extract(31, 0)}
}

func (x Int64) ToInt64() Int64 {
	if x.IsConcrete() {
		return Int64{C: int64(x.C)}
	}
	return Int64{S: x.S.SignExtend(0)}
}

func (x Int64) ToUint() Uint {
	if x.IsConcrete() {
		return Uint{C: uint(x.C)}
	}
	return Uint{S: x.S.SignExtend(0)}
}

func (x Int64) ToUint8() Uint8 {
	if x.IsConcrete() {
		return Uint8{C: uint8(x.C)}
	}
	return Uint8{S: x.S.Extract(7, 0)}
}

func (x Int64) ToUint16() Uint16 {
	if x.IsConcrete() {
		return Uint16{C: uint16(x.C)}
	}
	return Uint16{S: x.S.Extract(15, 0)}
}

func (x Int64) ToUint32() Uint32 {
	if x.IsConcrete() {
		return Uint32{C: uint32(x.C)}
	}
	return Uint32{S: x.S.Extract(31, 0)}
}

func (x Int64) ToUint64() Uint64 {
	if x.IsConcrete() {
		return Uint64{C: uint64(x.C)}
	}
	return Uint64{S: x.S.SignExtend(0)}
}

func (x Int64) ToUintptr() Uintptr {
	if x.IsConcrete() {
		return Uintptr{C: uintptr(x.C)}
	}
	return Uintptr{S: x.S.SignExtend(0)}
}

// Uint implements symbolic uint values.
type Uint struct {
	C uint
	S z3.BV
}

// AnyUint returns an unconstrained symbolic Uint.
func AnyUint(ctx *z3.Context, name string) Uint {
	cache := getCache(ctx)
	sym := cache.z3.FreshConst(name, cache.sortUint).(z3.BV)
	return Uint{S: sym}
}

// String returns x as a string.
func (x Uint) String() string {
	if x.IsConcrete() {
		return fmt.Sprint(x.C)
	}
	return x.S.String()
}

// IsConcrete returns true if x is concrete.
func (x Uint) IsConcrete() bool {
	return x.S.Context() == nil
}

// Eval returns x's concrete value in model m.
// This also evaluates x with model completion.
func (x Uint) Eval(m *z3.Model) uint {
	if x.IsConcrete() {
		return x.C
	}
	c := m.Eval(x.S, true).(z3.BV)
	val, ok, _ := c.AsUint64()
	if !ok {
		panic("model evaluation produced non-concrete value " + c.String())
	}
	return (uint)(val)
}

// sym returns x's symbolic value, creating it if necessary.
func (x Uint) sym(c *cache) z3.BV {
	if !x.IsConcrete() {
		return x.S
	}
	return c.z3.FromInt(int64(x.C), c.sortUint).(z3.BV)
}

func (x Uint) Add(y Uint) Uint {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint{C: x.C + y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint{S: x.sym(cache).Add(y.sym(cache))}
}

func (x Uint) Sub(y Uint) Uint {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint{C: x.C - y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint{S: x.sym(cache).Sub(y.sym(cache))}
}

func (x Uint) Mul(y Uint) Uint {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint{C: x.C * y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint{S: x.sym(cache).Mul(y.sym(cache))}
}

func (x Uint) Quo(y Uint) Uint {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint{C: x.C / y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint{S: x.sym(cache).UDiv(y.sym(cache))}
}

func (x Uint) Rem(y Uint) Uint {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint{C: x.C % y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint{S: x.sym(cache).URem(y.sym(cache))}
}

func (x Uint) And(y Uint) Uint {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint{C: x.C & y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint{S: x.sym(cache).And(y.sym(cache))}
}

func (x Uint) Or(y Uint) Uint {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint{C: x.C | y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint{S: x.sym(cache).Or(y.sym(cache))}
}

func (x Uint) Xor(y Uint) Uint {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint{C: x.C ^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint{S: x.sym(cache).Xor(y.sym(cache))}
}

func (x Uint) Lsh(y Uint64) Uint {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint{C: x.C << y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 64 {
			return Uint{C: 0}
		}
	}
	rs = y.sym(cache)
	return Uint{S: x.sym(cache).Lsh(rs)}
}

func (x Uint) Rsh(y Uint64) Uint {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint{C: x.C >> y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 64 {
			return Uint{C: 0}
		}
	}
	rs = y.sym(cache)
	return Uint{S: x.sym(cache).URsh(rs)}
}

func (x Uint) AndNot(y Uint) Uint {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint{C: x.C &^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint{S: x.sym(cache).And(y.Not().sym(cache))}
}

func (x Uint) Eq(y Uint) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C == y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).Eq(y.sym(cache))}
}

func (x Uint) NE(y Uint) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C != y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).NE(y.sym(cache))}
}

func (x Uint) LT(y Uint) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C < y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).ULT(y.sym(cache))}
}

func (x Uint) LE(y Uint) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C <= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).ULE(y.sym(cache))}
}

func (x Uint) GT(y Uint) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C > y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).UGT(y.sym(cache))}
}

func (x Uint) GE(y Uint) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C >= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).UGE(y.sym(cache))}
}

func (x Uint) Neg() Uint {
	if x.IsConcrete() {
		return Uint{C: -x.C}
	}
	return Uint{S: x.S.Neg()}
}

func (x Uint) Not() Uint {
	if x.IsConcrete() {
		return Uint{C: ^x.C}
	}
	return Uint{S: x.S.Not()}
}

func (x Uint) ToInt() Int {
	if x.IsConcrete() {
		return Int{C: int(x.C)}
	}
	return Int{S: x.S.ZeroExtend(0)}
}

func (x Uint) ToInt8() Int8 {
	if x.IsConcrete() {
		return Int8{C: int8(x.C)}
	}
	return Int8{S: x.S.Extract(7, 0)}
}

func (x Uint) ToInt16() Int16 {
	if x.IsConcrete() {
		return Int16{C: int16(x.C)}
	}
	return Int16{S: x.S.Extract(15, 0)}
}

func (x Uint) ToInt32() Int32 {
	if x.IsConcrete() {
		return Int32{C: int32(x.C)}
	}
	return Int32{S: x.S.Extract(31, 0)}
}

func (x Uint) ToInt64() Int64 {
	if x.IsConcrete() {
		return Int64{C: int64(x.C)}
	}
	return Int64{S: x.S.ZeroExtend(0)}
}

func (x Uint) ToUint() Uint {
	if x.IsConcrete() {
		return Uint{C: uint(x.C)}
	}
	return Uint{S: x.S.ZeroExtend(0)}
}

func (x Uint) ToUint8() Uint8 {
	if x.IsConcrete() {
		return Uint8{C: uint8(x.C)}
	}
	return Uint8{S: x.S.Extract(7, 0)}
}

func (x Uint) ToUint16() Uint16 {
	if x.IsConcrete() {
		return Uint16{C: uint16(x.C)}
	}
	return Uint16{S: x.S.Extract(15, 0)}
}

func (x Uint) ToUint32() Uint32 {
	if x.IsConcrete() {
		return Uint32{C: uint32(x.C)}
	}
	return Uint32{S: x.S.Extract(31, 0)}
}

func (x Uint) ToUint64() Uint64 {
	if x.IsConcrete() {
		return Uint64{C: uint64(x.C)}
	}
	return Uint64{S: x.S.ZeroExtend(0)}
}

func (x Uint) ToUintptr() Uintptr {
	if x.IsConcrete() {
		return Uintptr{C: uintptr(x.C)}
	}
	return Uintptr{S: x.S.ZeroExtend(0)}
}

// Uint8 implements symbolic uint8 values.
type Uint8 struct {
	C uint8
	S z3.BV
}

// AnyUint8 returns an unconstrained symbolic Uint8.
func AnyUint8(ctx *z3.Context, name string) Uint8 {
	cache := getCache(ctx)
	sym := cache.z3.FreshConst(name, cache.sortUint8).(z3.BV)
	return Uint8{S: sym}
}

// String returns x as a string.
func (x Uint8) String() string {
	if x.IsConcrete() {
		return fmt.Sprint(x.C)
	}
	return x.S.String()
}

// IsConcrete returns true if x is concrete.
func (x Uint8) IsConcrete() bool {
	return x.S.Context() == nil
}

// Eval returns x's concrete value in model m.
// This also evaluates x with model completion.
func (x Uint8) Eval(m *z3.Model) uint8 {
	if x.IsConcrete() {
		return x.C
	}
	c := m.Eval(x.S, true).(z3.BV)
	val, ok, _ := c.AsUint64()
	if !ok {
		panic("model evaluation produced non-concrete value " + c.String())
	}
	return (uint8)(val)
}

// sym returns x's symbolic value, creating it if necessary.
func (x Uint8) sym(c *cache) z3.BV {
	if !x.IsConcrete() {
		return x.S
	}
	return c.z3.FromInt(int64(x.C), c.sortUint8).(z3.BV)
}

func (x Uint8) Add(y Uint8) Uint8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint8{C: x.C + y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint8{S: x.sym(cache).Add(y.sym(cache))}
}

func (x Uint8) Sub(y Uint8) Uint8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint8{C: x.C - y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint8{S: x.sym(cache).Sub(y.sym(cache))}
}

func (x Uint8) Mul(y Uint8) Uint8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint8{C: x.C * y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint8{S: x.sym(cache).Mul(y.sym(cache))}
}

func (x Uint8) Quo(y Uint8) Uint8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint8{C: x.C / y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint8{S: x.sym(cache).UDiv(y.sym(cache))}
}

func (x Uint8) Rem(y Uint8) Uint8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint8{C: x.C % y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint8{S: x.sym(cache).URem(y.sym(cache))}
}

func (x Uint8) And(y Uint8) Uint8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint8{C: x.C & y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint8{S: x.sym(cache).And(y.sym(cache))}
}

func (x Uint8) Or(y Uint8) Uint8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint8{C: x.C | y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint8{S: x.sym(cache).Or(y.sym(cache))}
}

func (x Uint8) Xor(y Uint8) Uint8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint8{C: x.C ^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint8{S: x.sym(cache).Xor(y.sym(cache))}
}

func (x Uint8) Lsh(y Uint64) Uint8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint8{C: x.C << y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 8 {
			return Uint8{C: 0}
		}
		rs = Uint8{C: uint8(y.C)}.sym(cache)
	} else {
		rs = y.sym(cache)
		rs = rs.Extract(63, 7).AnyBits().Concat(rs.Extract(6, 0))
	}
	return Uint8{S: x.sym(cache).Lsh(rs)}
}

func (x Uint8) Rsh(y Uint64) Uint8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint8{C: x.C >> y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 8 {
			return Uint8{C: 0}
		}
		rs = Uint8{C: uint8(y.C)}.sym(cache)
	} else {
		rs = y.sym(cache)
		rs = rs.Extract(63, 7).AnyBits().Concat(rs.Extract(6, 0))
	}
	return Uint8{S: x.sym(cache).URsh(rs)}
}

func (x Uint8) AndNot(y Uint8) Uint8 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint8{C: x.C &^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint8{S: x.sym(cache).And(y.Not().sym(cache))}
}

func (x Uint8) Eq(y Uint8) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C == y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).Eq(y.sym(cache))}
}

func (x Uint8) NE(y Uint8) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C != y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).NE(y.sym(cache))}
}

func (x Uint8) LT(y Uint8) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C < y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).ULT(y.sym(cache))}
}

func (x Uint8) LE(y Uint8) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C <= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).ULE(y.sym(cache))}
}

func (x Uint8) GT(y Uint8) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C > y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).UGT(y.sym(cache))}
}

func (x Uint8) GE(y Uint8) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C >= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).UGE(y.sym(cache))}
}

func (x Uint8) Neg() Uint8 {
	if x.IsConcrete() {
		return Uint8{C: -x.C}
	}
	return Uint8{S: x.S.Neg()}
}

func (x Uint8) Not() Uint8 {
	if x.IsConcrete() {
		return Uint8{C: ^x.C}
	}
	return Uint8{S: x.S.Not()}
}

func (x Uint8) ToInt() Int {
	if x.IsConcrete() {
		return Int{C: int(x.C)}
	}
	return Int{S: x.S.ZeroExtend(56)}
}

func (x Uint8) ToInt8() Int8 {
	if x.IsConcrete() {
		return Int8{C: int8(x.C)}
	}
	return Int8{S: x.S.ZeroExtend(0)}
}

func (x Uint8) ToInt16() Int16 {
	if x.IsConcrete() {
		return Int16{C: int16(x.C)}
	}
	return Int16{S: x.S.ZeroExtend(8)}
}

func (x Uint8) ToInt32() Int32 {
	if x.IsConcrete() {
		return Int32{C: int32(x.C)}
	}
	return Int32{S: x.S.ZeroExtend(24)}
}

func (x Uint8) ToInt64() Int64 {
	if x.IsConcrete() {
		return Int64{C: int64(x.C)}
	}
	return Int64{S: x.S.ZeroExtend(56)}
}

func (x Uint8) ToUint() Uint {
	if x.IsConcrete() {
		return Uint{C: uint(x.C)}
	}
	return Uint{S: x.S.ZeroExtend(56)}
}

func (x Uint8) ToUint8() Uint8 {
	if x.IsConcrete() {
		return Uint8{C: uint8(x.C)}
	}
	return Uint8{S: x.S.ZeroExtend(0)}
}

func (x Uint8) ToUint16() Uint16 {
	if x.IsConcrete() {
		return Uint16{C: uint16(x.C)}
	}
	return Uint16{S: x.S.ZeroExtend(8)}
}

func (x Uint8) ToUint32() Uint32 {
	if x.IsConcrete() {
		return Uint32{C: uint32(x.C)}
	}
	return Uint32{S: x.S.ZeroExtend(24)}
}

func (x Uint8) ToUint64() Uint64 {
	if x.IsConcrete() {
		return Uint64{C: uint64(x.C)}
	}
	return Uint64{S: x.S.ZeroExtend(56)}
}

func (x Uint8) ToUintptr() Uintptr {
	if x.IsConcrete() {
		return Uintptr{C: uintptr(x.C)}
	}
	return Uintptr{S: x.S.ZeroExtend(56)}
}

// Uint16 implements symbolic uint16 values.
type Uint16 struct {
	C uint16
	S z3.BV
}

// AnyUint16 returns an unconstrained symbolic Uint16.
func AnyUint16(ctx *z3.Context, name string) Uint16 {
	cache := getCache(ctx)
	sym := cache.z3.FreshConst(name, cache.sortUint16).(z3.BV)
	return Uint16{S: sym}
}

// String returns x as a string.
func (x Uint16) String() string {
	if x.IsConcrete() {
		return fmt.Sprint(x.C)
	}
	return x.S.String()
}

// IsConcrete returns true if x is concrete.
func (x Uint16) IsConcrete() bool {
	return x.S.Context() == nil
}

// Eval returns x's concrete value in model m.
// This also evaluates x with model completion.
func (x Uint16) Eval(m *z3.Model) uint16 {
	if x.IsConcrete() {
		return x.C
	}
	c := m.Eval(x.S, true).(z3.BV)
	val, ok, _ := c.AsUint64()
	if !ok {
		panic("model evaluation produced non-concrete value " + c.String())
	}
	return (uint16)(val)
}

// sym returns x's symbolic value, creating it if necessary.
func (x Uint16) sym(c *cache) z3.BV {
	if !x.IsConcrete() {
		return x.S
	}
	return c.z3.FromInt(int64(x.C), c.sortUint16).(z3.BV)
}

func (x Uint16) Add(y Uint16) Uint16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint16{C: x.C + y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint16{S: x.sym(cache).Add(y.sym(cache))}
}

func (x Uint16) Sub(y Uint16) Uint16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint16{C: x.C - y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint16{S: x.sym(cache).Sub(y.sym(cache))}
}

func (x Uint16) Mul(y Uint16) Uint16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint16{C: x.C * y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint16{S: x.sym(cache).Mul(y.sym(cache))}
}

func (x Uint16) Quo(y Uint16) Uint16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint16{C: x.C / y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint16{S: x.sym(cache).UDiv(y.sym(cache))}
}

func (x Uint16) Rem(y Uint16) Uint16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint16{C: x.C % y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint16{S: x.sym(cache).URem(y.sym(cache))}
}

func (x Uint16) And(y Uint16) Uint16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint16{C: x.C & y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint16{S: x.sym(cache).And(y.sym(cache))}
}

func (x Uint16) Or(y Uint16) Uint16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint16{C: x.C | y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint16{S: x.sym(cache).Or(y.sym(cache))}
}

func (x Uint16) Xor(y Uint16) Uint16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint16{C: x.C ^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint16{S: x.sym(cache).Xor(y.sym(cache))}
}

func (x Uint16) Lsh(y Uint64) Uint16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint16{C: x.C << y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 16 {
			return Uint16{C: 0}
		}
		rs = Uint16{C: uint16(y.C)}.sym(cache)
	} else {
		rs = y.sym(cache)
		rs = rs.Extract(63, 15).AnyBits().Concat(rs.Extract(14, 0))
	}
	return Uint16{S: x.sym(cache).Lsh(rs)}
}

func (x Uint16) Rsh(y Uint64) Uint16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint16{C: x.C >> y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 16 {
			return Uint16{C: 0}
		}
		rs = Uint16{C: uint16(y.C)}.sym(cache)
	} else {
		rs = y.sym(cache)
		rs = rs.Extract(63, 15).AnyBits().Concat(rs.Extract(14, 0))
	}
	return Uint16{S: x.sym(cache).URsh(rs)}
}

func (x Uint16) AndNot(y Uint16) Uint16 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint16{C: x.C &^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint16{S: x.sym(cache).And(y.Not().sym(cache))}
}

func (x Uint16) Eq(y Uint16) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C == y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).Eq(y.sym(cache))}
}

func (x Uint16) NE(y Uint16) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C != y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).NE(y.sym(cache))}
}

func (x Uint16) LT(y Uint16) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C < y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).ULT(y.sym(cache))}
}

func (x Uint16) LE(y Uint16) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C <= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).ULE(y.sym(cache))}
}

func (x Uint16) GT(y Uint16) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C > y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).UGT(y.sym(cache))}
}

func (x Uint16) GE(y Uint16) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C >= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).UGE(y.sym(cache))}
}

func (x Uint16) Neg() Uint16 {
	if x.IsConcrete() {
		return Uint16{C: -x.C}
	}
	return Uint16{S: x.S.Neg()}
}

func (x Uint16) Not() Uint16 {
	if x.IsConcrete() {
		return Uint16{C: ^x.C}
	}
	return Uint16{S: x.S.Not()}
}

func (x Uint16) ToInt() Int {
	if x.IsConcrete() {
		return Int{C: int(x.C)}
	}
	return Int{S: x.S.ZeroExtend(48)}
}

func (x Uint16) ToInt8() Int8 {
	if x.IsConcrete() {
		return Int8{C: int8(x.C)}
	}
	return Int8{S: x.S.Extract(7, 0)}
}

func (x Uint16) ToInt16() Int16 {
	if x.IsConcrete() {
		return Int16{C: int16(x.C)}
	}
	return Int16{S: x.S.ZeroExtend(0)}
}

func (x Uint16) ToInt32() Int32 {
	if x.IsConcrete() {
		return Int32{C: int32(x.C)}
	}
	return Int32{S: x.S.ZeroExtend(16)}
}

func (x Uint16) ToInt64() Int64 {
	if x.IsConcrete() {
		return Int64{C: int64(x.C)}
	}
	return Int64{S: x.S.ZeroExtend(48)}
}

func (x Uint16) ToUint() Uint {
	if x.IsConcrete() {
		return Uint{C: uint(x.C)}
	}
	return Uint{S: x.S.ZeroExtend(48)}
}

func (x Uint16) ToUint8() Uint8 {
	if x.IsConcrete() {
		return Uint8{C: uint8(x.C)}
	}
	return Uint8{S: x.S.Extract(7, 0)}
}

func (x Uint16) ToUint16() Uint16 {
	if x.IsConcrete() {
		return Uint16{C: uint16(x.C)}
	}
	return Uint16{S: x.S.ZeroExtend(0)}
}

func (x Uint16) ToUint32() Uint32 {
	if x.IsConcrete() {
		return Uint32{C: uint32(x.C)}
	}
	return Uint32{S: x.S.ZeroExtend(16)}
}

func (x Uint16) ToUint64() Uint64 {
	if x.IsConcrete() {
		return Uint64{C: uint64(x.C)}
	}
	return Uint64{S: x.S.ZeroExtend(48)}
}

func (x Uint16) ToUintptr() Uintptr {
	if x.IsConcrete() {
		return Uintptr{C: uintptr(x.C)}
	}
	return Uintptr{S: x.S.ZeroExtend(48)}
}

// Uint32 implements symbolic uint32 values.
type Uint32 struct {
	C uint32
	S z3.BV
}

// AnyUint32 returns an unconstrained symbolic Uint32.
func AnyUint32(ctx *z3.Context, name string) Uint32 {
	cache := getCache(ctx)
	sym := cache.z3.FreshConst(name, cache.sortUint32).(z3.BV)
	return Uint32{S: sym}
}

// String returns x as a string.
func (x Uint32) String() string {
	if x.IsConcrete() {
		return fmt.Sprint(x.C)
	}
	return x.S.String()
}

// IsConcrete returns true if x is concrete.
func (x Uint32) IsConcrete() bool {
	return x.S.Context() == nil
}

// Eval returns x's concrete value in model m.
// This also evaluates x with model completion.
func (x Uint32) Eval(m *z3.Model) uint32 {
	if x.IsConcrete() {
		return x.C
	}
	c := m.Eval(x.S, true).(z3.BV)
	val, ok, _ := c.AsUint64()
	if !ok {
		panic("model evaluation produced non-concrete value " + c.String())
	}
	return (uint32)(val)
}

// sym returns x's symbolic value, creating it if necessary.
func (x Uint32) sym(c *cache) z3.BV {
	if !x.IsConcrete() {
		return x.S
	}
	return c.z3.FromInt(int64(x.C), c.sortUint32).(z3.BV)
}

func (x Uint32) Add(y Uint32) Uint32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint32{C: x.C + y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint32{S: x.sym(cache).Add(y.sym(cache))}
}

func (x Uint32) Sub(y Uint32) Uint32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint32{C: x.C - y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint32{S: x.sym(cache).Sub(y.sym(cache))}
}

func (x Uint32) Mul(y Uint32) Uint32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint32{C: x.C * y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint32{S: x.sym(cache).Mul(y.sym(cache))}
}

func (x Uint32) Quo(y Uint32) Uint32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint32{C: x.C / y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint32{S: x.sym(cache).UDiv(y.sym(cache))}
}

func (x Uint32) Rem(y Uint32) Uint32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint32{C: x.C % y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint32{S: x.sym(cache).URem(y.sym(cache))}
}

func (x Uint32) And(y Uint32) Uint32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint32{C: x.C & y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint32{S: x.sym(cache).And(y.sym(cache))}
}

func (x Uint32) Or(y Uint32) Uint32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint32{C: x.C | y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint32{S: x.sym(cache).Or(y.sym(cache))}
}

func (x Uint32) Xor(y Uint32) Uint32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint32{C: x.C ^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint32{S: x.sym(cache).Xor(y.sym(cache))}
}

func (x Uint32) Lsh(y Uint64) Uint32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint32{C: x.C << y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 32 {
			return Uint32{C: 0}
		}
		rs = Uint32{C: uint32(y.C)}.sym(cache)
	} else {
		rs = y.sym(cache)
		rs = rs.Extract(63, 31).AnyBits().Concat(rs.Extract(30, 0))
	}
	return Uint32{S: x.sym(cache).Lsh(rs)}
}

func (x Uint32) Rsh(y Uint64) Uint32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint32{C: x.C >> y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 32 {
			return Uint32{C: 0}
		}
		rs = Uint32{C: uint32(y.C)}.sym(cache)
	} else {
		rs = y.sym(cache)
		rs = rs.Extract(63, 31).AnyBits().Concat(rs.Extract(30, 0))
	}
	return Uint32{S: x.sym(cache).URsh(rs)}
}

func (x Uint32) AndNot(y Uint32) Uint32 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint32{C: x.C &^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint32{S: x.sym(cache).And(y.Not().sym(cache))}
}

func (x Uint32) Eq(y Uint32) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C == y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).Eq(y.sym(cache))}
}

func (x Uint32) NE(y Uint32) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C != y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).NE(y.sym(cache))}
}

func (x Uint32) LT(y Uint32) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C < y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).ULT(y.sym(cache))}
}

func (x Uint32) LE(y Uint32) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C <= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).ULE(y.sym(cache))}
}

func (x Uint32) GT(y Uint32) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C > y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).UGT(y.sym(cache))}
}

func (x Uint32) GE(y Uint32) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C >= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).UGE(y.sym(cache))}
}

func (x Uint32) Neg() Uint32 {
	if x.IsConcrete() {
		return Uint32{C: -x.C}
	}
	return Uint32{S: x.S.Neg()}
}

func (x Uint32) Not() Uint32 {
	if x.IsConcrete() {
		return Uint32{C: ^x.C}
	}
	return Uint32{S: x.S.Not()}
}

func (x Uint32) ToInt() Int {
	if x.IsConcrete() {
		return Int{C: int(x.C)}
	}
	return Int{S: x.S.ZeroExtend(32)}
}

func (x Uint32) ToInt8() Int8 {
	if x.IsConcrete() {
		return Int8{C: int8(x.C)}
	}
	return Int8{S: x.S.Extract(7, 0)}
}

func (x Uint32) ToInt16() Int16 {
	if x.IsConcrete() {
		return Int16{C: int16(x.C)}
	}
	return Int16{S: x.S.Extract(15, 0)}
}

func (x Uint32) ToInt32() Int32 {
	if x.IsConcrete() {
		return Int32{C: int32(x.C)}
	}
	return Int32{S: x.S.ZeroExtend(0)}
}

func (x Uint32) ToInt64() Int64 {
	if x.IsConcrete() {
		return Int64{C: int64(x.C)}
	}
	return Int64{S: x.S.ZeroExtend(32)}
}

func (x Uint32) ToUint() Uint {
	if x.IsConcrete() {
		return Uint{C: uint(x.C)}
	}
	return Uint{S: x.S.ZeroExtend(32)}
}

func (x Uint32) ToUint8() Uint8 {
	if x.IsConcrete() {
		return Uint8{C: uint8(x.C)}
	}
	return Uint8{S: x.S.Extract(7, 0)}
}

func (x Uint32) ToUint16() Uint16 {
	if x.IsConcrete() {
		return Uint16{C: uint16(x.C)}
	}
	return Uint16{S: x.S.Extract(15, 0)}
}

func (x Uint32) ToUint32() Uint32 {
	if x.IsConcrete() {
		return Uint32{C: uint32(x.C)}
	}
	return Uint32{S: x.S.ZeroExtend(0)}
}

func (x Uint32) ToUint64() Uint64 {
	if x.IsConcrete() {
		return Uint64{C: uint64(x.C)}
	}
	return Uint64{S: x.S.ZeroExtend(32)}
}

func (x Uint32) ToUintptr() Uintptr {
	if x.IsConcrete() {
		return Uintptr{C: uintptr(x.C)}
	}
	return Uintptr{S: x.S.ZeroExtend(32)}
}

// Uint64 implements symbolic uint64 values.
type Uint64 struct {
	C uint64
	S z3.BV
}

// AnyUint64 returns an unconstrained symbolic Uint64.
func AnyUint64(ctx *z3.Context, name string) Uint64 {
	cache := getCache(ctx)
	sym := cache.z3.FreshConst(name, cache.sortUint64).(z3.BV)
	return Uint64{S: sym}
}

// String returns x as a string.
func (x Uint64) String() string {
	if x.IsConcrete() {
		return fmt.Sprint(x.C)
	}
	return x.S.String()
}

// IsConcrete returns true if x is concrete.
func (x Uint64) IsConcrete() bool {
	return x.S.Context() == nil
}

// Eval returns x's concrete value in model m.
// This also evaluates x with model completion.
func (x Uint64) Eval(m *z3.Model) uint64 {
	if x.IsConcrete() {
		return x.C
	}
	c := m.Eval(x.S, true).(z3.BV)
	val, ok, _ := c.AsUint64()
	if !ok {
		panic("model evaluation produced non-concrete value " + c.String())
	}
	return (uint64)(val)
}

// sym returns x's symbolic value, creating it if necessary.
func (x Uint64) sym(c *cache) z3.BV {
	if !x.IsConcrete() {
		return x.S
	}
	return c.z3.FromInt(int64(x.C), c.sortUint64).(z3.BV)
}

func (x Uint64) Add(y Uint64) Uint64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint64{C: x.C + y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint64{S: x.sym(cache).Add(y.sym(cache))}
}

func (x Uint64) Sub(y Uint64) Uint64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint64{C: x.C - y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint64{S: x.sym(cache).Sub(y.sym(cache))}
}

func (x Uint64) Mul(y Uint64) Uint64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint64{C: x.C * y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint64{S: x.sym(cache).Mul(y.sym(cache))}
}

func (x Uint64) Quo(y Uint64) Uint64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint64{C: x.C / y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint64{S: x.sym(cache).UDiv(y.sym(cache))}
}

func (x Uint64) Rem(y Uint64) Uint64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint64{C: x.C % y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint64{S: x.sym(cache).URem(y.sym(cache))}
}

func (x Uint64) And(y Uint64) Uint64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint64{C: x.C & y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint64{S: x.sym(cache).And(y.sym(cache))}
}

func (x Uint64) Or(y Uint64) Uint64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint64{C: x.C | y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint64{S: x.sym(cache).Or(y.sym(cache))}
}

func (x Uint64) Xor(y Uint64) Uint64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint64{C: x.C ^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint64{S: x.sym(cache).Xor(y.sym(cache))}
}

func (x Uint64) Lsh(y Uint64) Uint64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint64{C: x.C << y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 64 {
			return Uint64{C: 0}
		}
	}
	rs = y.sym(cache)
	return Uint64{S: x.sym(cache).Lsh(rs)}
}

func (x Uint64) Rsh(y Uint64) Uint64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint64{C: x.C >> y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 64 {
			return Uint64{C: 0}
		}
	}
	rs = y.sym(cache)
	return Uint64{S: x.sym(cache).URsh(rs)}
}

func (x Uint64) AndNot(y Uint64) Uint64 {
	if x.IsConcrete() && y.IsConcrete() {
		return Uint64{C: x.C &^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uint64{S: x.sym(cache).And(y.Not().sym(cache))}
}

func (x Uint64) Eq(y Uint64) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C == y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).Eq(y.sym(cache))}
}

func (x Uint64) NE(y Uint64) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C != y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).NE(y.sym(cache))}
}

func (x Uint64) LT(y Uint64) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C < y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).ULT(y.sym(cache))}
}

func (x Uint64) LE(y Uint64) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C <= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).ULE(y.sym(cache))}
}

func (x Uint64) GT(y Uint64) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C > y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).UGT(y.sym(cache))}
}

func (x Uint64) GE(y Uint64) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C >= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).UGE(y.sym(cache))}
}

func (x Uint64) Neg() Uint64 {
	if x.IsConcrete() {
		return Uint64{C: -x.C}
	}
	return Uint64{S: x.S.Neg()}
}

func (x Uint64) Not() Uint64 {
	if x.IsConcrete() {
		return Uint64{C: ^x.C}
	}
	return Uint64{S: x.S.Not()}
}

func (x Uint64) Upper32() Int32 {
	if x.IsConcrete() {
		return Int32{C: int32(x.C >> 32)}
	}
	return Int32{S: x.S.Extract(63, 32)}
}

func (x Uint64) ToInt() Int {
	if x.IsConcrete() {
		return Int{C: int(x.C)}
	}
	return Int{S: x.S.ZeroExtend(0)}
}

func (x Uint64) ToInt8() Int8 {
	if x.IsConcrete() {
		return Int8{C: int8(x.C)}
	}
	return Int8{S: x.S.Extract(7, 0)}
}

func (x Uint64) ToInt16() Int16 {
	if x.IsConcrete() {
		return Int16{C: int16(x.C)}
	}
	return Int16{S: x.S.Extract(15, 0)}
}

func (x Uint64) ToInt32() Int32 {
	if x.IsConcrete() {
		return Int32{C: int32(x.C)}
	}
	return Int32{S: x.S.Extract(31, 0)}
}

func (x Uint64) ToInt64() Int64 {
	if x.IsConcrete() {
		return Int64{C: int64(x.C)}
	}
	return Int64{S: x.S.ZeroExtend(0)}
}

func (x Uint64) ToUint() Uint {
	if x.IsConcrete() {
		return Uint{C: uint(x.C)}
	}
	return Uint{S: x.S.ZeroExtend(0)}
}

func (x Uint64) ToUint8() Uint8 {
	if x.IsConcrete() {
		return Uint8{C: uint8(x.C)}
	}
	return Uint8{S: x.S.Extract(7, 0)}
}

func (x Uint64) ToUint16() Uint16 {
	if x.IsConcrete() {
		return Uint16{C: uint16(x.C)}
	}
	return Uint16{S: x.S.Extract(15, 0)}
}

func (x Uint64) ToUint32() Uint32 {
	if x.IsConcrete() {
		return Uint32{C: uint32(x.C)}
	}
	return Uint32{S: x.S.Extract(31, 0)}
}

func (x Uint64) ToUint64() Uint64 {
	if x.IsConcrete() {
		return Uint64{C: uint64(x.C)}
	}
	return Uint64{S: x.S.ZeroExtend(0)}
}

func (x Uint64) ToUintptr() Uintptr {
	if x.IsConcrete() {
		return Uintptr{C: uintptr(x.C)}
	}
	return Uintptr{S: x.S.ZeroExtend(0)}
}

// Uintptr implements symbolic uintptr values.
type Uintptr struct {
	C uintptr
	S z3.BV
}

// AnyUintptr returns an unconstrained symbolic Uintptr.
func AnyUintptr(ctx *z3.Context, name string) Uintptr {
	cache := getCache(ctx)
	sym := cache.z3.FreshConst(name, cache.sortUintptr).(z3.BV)
	return Uintptr{S: sym}
}

// String returns x as a string.
func (x Uintptr) String() string {
	if x.IsConcrete() {
		return fmt.Sprint(x.C)
	}
	return x.S.String()
}

// IsConcrete returns true if x is concrete.
func (x Uintptr) IsConcrete() bool {
	return x.S.Context() == nil
}

// Eval returns x's concrete value in model m.
// This also evaluates x with model completion.
func (x Uintptr) Eval(m *z3.Model) uintptr {
	if x.IsConcrete() {
		return x.C
	}
	c := m.Eval(x.S, true).(z3.BV)
	val, ok, _ := c.AsUint64()
	if !ok {
		panic("model evaluation produced non-concrete value " + c.String())
	}
	return (uintptr)(val)
}

// sym returns x's symbolic value, creating it if necessary.
func (x Uintptr) sym(c *cache) z3.BV {
	if !x.IsConcrete() {
		return x.S
	}
	return c.z3.FromInt(int64(x.C), c.sortUintptr).(z3.BV)
}

func (x Uintptr) Add(y Uintptr) Uintptr {
	if x.IsConcrete() && y.IsConcrete() {
		return Uintptr{C: x.C + y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uintptr{S: x.sym(cache).Add(y.sym(cache))}
}

func (x Uintptr) Sub(y Uintptr) Uintptr {
	if x.IsConcrete() && y.IsConcrete() {
		return Uintptr{C: x.C - y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uintptr{S: x.sym(cache).Sub(y.sym(cache))}
}

func (x Uintptr) Mul(y Uintptr) Uintptr {
	if x.IsConcrete() && y.IsConcrete() {
		return Uintptr{C: x.C * y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uintptr{S: x.sym(cache).Mul(y.sym(cache))}
}

func (x Uintptr) Quo(y Uintptr) Uintptr {
	if x.IsConcrete() && y.IsConcrete() {
		return Uintptr{C: x.C / y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uintptr{S: x.sym(cache).UDiv(y.sym(cache))}
}

func (x Uintptr) Rem(y Uintptr) Uintptr {
	if x.IsConcrete() && y.IsConcrete() {
		return Uintptr{C: x.C % y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uintptr{S: x.sym(cache).URem(y.sym(cache))}
}

func (x Uintptr) And(y Uintptr) Uintptr {
	if x.IsConcrete() && y.IsConcrete() {
		return Uintptr{C: x.C & y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uintptr{S: x.sym(cache).And(y.sym(cache))}
}

func (x Uintptr) Or(y Uintptr) Uintptr {
	if x.IsConcrete() && y.IsConcrete() {
		return Uintptr{C: x.C | y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uintptr{S: x.sym(cache).Or(y.sym(cache))}
}

func (x Uintptr) Xor(y Uintptr) Uintptr {
	if x.IsConcrete() && y.IsConcrete() {
		return Uintptr{C: x.C ^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uintptr{S: x.sym(cache).Xor(y.sym(cache))}
}

func (x Uintptr) Lsh(y Uint64) Uintptr {
	if x.IsConcrete() && y.IsConcrete() {
		return Uintptr{C: x.C << y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 64 {
			return Uintptr{C: 0}
		}
	}
	rs = y.sym(cache)
	return Uintptr{S: x.sym(cache).Lsh(rs)}
}

func (x Uintptr) Rsh(y Uint64) Uintptr {
	if x.IsConcrete() && y.IsConcrete() {
		return Uintptr{C: x.C >> y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	var rs z3.BV
	if y.IsConcrete() {
		if y.C >= 64 {
			return Uintptr{C: 0}
		}
	}
	rs = y.sym(cache)
	return Uintptr{S: x.sym(cache).URsh(rs)}
}

func (x Uintptr) AndNot(y Uintptr) Uintptr {
	if x.IsConcrete() && y.IsConcrete() {
		return Uintptr{C: x.C &^ y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Uintptr{S: x.sym(cache).And(y.Not().sym(cache))}
}

func (x Uintptr) Eq(y Uintptr) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C == y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).Eq(y.sym(cache))}
}

func (x Uintptr) NE(y Uintptr) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C != y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).NE(y.sym(cache))}
}

func (x Uintptr) LT(y Uintptr) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C < y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).ULT(y.sym(cache))}
}

func (x Uintptr) LE(y Uintptr) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C <= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).ULE(y.sym(cache))}
}

func (x Uintptr) GT(y Uintptr) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C > y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).UGT(y.sym(cache))}
}

func (x Uintptr) GE(y Uintptr) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C >= y.C}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).UGE(y.sym(cache))}
}

func (x Uintptr) Neg() Uintptr {
	if x.IsConcrete() {
		return Uintptr{C: -x.C}
	}
	return Uintptr{S: x.S.Neg()}
}

func (x Uintptr) Not() Uintptr {
	if x.IsConcrete() {
		return Uintptr{C: ^x.C}
	}
	return Uintptr{S: x.S.Not()}
}

func (x Uintptr) ToInt() Int {
	if x.IsConcrete() {
		return Int{C: int(x.C)}
	}
	return Int{S: x.S.ZeroExtend(0)}
}

func (x Uintptr) ToInt8() Int8 {
	if x.IsConcrete() {
		return Int8{C: int8(x.C)}
	}
	return Int8{S: x.S.Extract(7, 0)}
}

func (x Uintptr) ToInt16() Int16 {
	if x.IsConcrete() {
		return Int16{C: int16(x.C)}
	}
	return Int16{S: x.S.Extract(15, 0)}
}

func (x Uintptr) ToInt32() Int32 {
	if x.IsConcrete() {
		return Int32{C: int32(x.C)}
	}
	return Int32{S: x.S.Extract(31, 0)}
}

func (x Uintptr) ToInt64() Int64 {
	if x.IsConcrete() {
		return Int64{C: int64(x.C)}
	}
	return Int64{S: x.S.ZeroExtend(0)}
}

func (x Uintptr) ToUint() Uint {
	if x.IsConcrete() {
		return Uint{C: uint(x.C)}
	}
	return Uint{S: x.S.ZeroExtend(0)}
}

func (x Uintptr) ToUint8() Uint8 {
	if x.IsConcrete() {
		return Uint8{C: uint8(x.C)}
	}
	return Uint8{S: x.S.Extract(7, 0)}
}

func (x Uintptr) ToUint16() Uint16 {
	if x.IsConcrete() {
		return Uint16{C: uint16(x.C)}
	}
	return Uint16{S: x.S.Extract(15, 0)}
}

func (x Uintptr) ToUint32() Uint32 {
	if x.IsConcrete() {
		return Uint32{C: uint32(x.C)}
	}
	return Uint32{S: x.S.Extract(31, 0)}
}

func (x Uintptr) ToUint64() Uint64 {
	if x.IsConcrete() {
		return Uint64{C: uint64(x.C)}
	}
	return Uint64{S: x.S.ZeroExtend(0)}
}

func (x Uintptr) ToUintptr() Uintptr {
	if x.IsConcrete() {
		return Uintptr{C: uintptr(x.C)}
	}
	return Uintptr{S: x.S.ZeroExtend(0)}
}

// Integer implements symbolic *big.Int values.
type Integer struct {
	C *big.Int
	S z3.Int
}

// AnyInteger returns an unconstrained symbolic Integer.
func AnyInteger(ctx *z3.Context, name string) Integer {
	cache := getCache(ctx)
	sym := cache.z3.FreshConst(name, cache.sortInteger).(z3.Int)
	return Integer{S: sym}
}

// String returns x as a string.
func (x Integer) String() string {
	if x.IsConcrete() {
		return fmt.Sprint(x.C)
	}
	return x.S.String()
}

// IsConcrete returns true if x is concrete.
func (x Integer) IsConcrete() bool {
	return x.S.Context() == nil
}

// Eval returns x's concrete value in model m.
// This also evaluates x with model completion.
func (x Integer) Eval(m *z3.Model) *big.Int {
	if x.IsConcrete() {
		return x.C
	}
	c := m.Eval(x.S, true).(z3.Int)
	val, ok := c.AsBigInt()
	if !ok {
		panic("model evaluation produced non-concrete value " + c.String())
	}
	return (*big.Int)(val)
}

// sym returns x's symbolic value, creating it if necessary.
func (x Integer) sym(c *cache) z3.Int {
	if !x.IsConcrete() {
		return x.S
	}
	return c.z3.FromBigInt(x.C, c.sortInteger).(z3.Int)
}

func (x Integer) Add(y Integer) Integer {
	if x.IsConcrete() && y.IsConcrete() {
		z := Integer{C: new(big.Int)}
		z.C.Add(x.C, y.C)
		return z
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Integer{S: x.sym(cache).Add(y.sym(cache))}
}

func (x Integer) Sub(y Integer) Integer {
	if x.IsConcrete() && y.IsConcrete() {
		z := Integer{C: new(big.Int)}
		z.C.Sub(x.C, y.C)
		return z
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Integer{S: x.sym(cache).Sub(y.sym(cache))}
}

func (x Integer) Mul(y Integer) Integer {
	if x.IsConcrete() && y.IsConcrete() {
		z := Integer{C: new(big.Int)}
		z.C.Mul(x.C, y.C)
		return z
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Integer{S: x.sym(cache).Mul(y.sym(cache))}
}

func (x Integer) Quo(y Integer) Integer {
	if x.IsConcrete() && y.IsConcrete() {
		z := Integer{C: new(big.Int)}
		z.C.Quo(x.C, y.C)
		return z
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	xs, ys := x.sym(cache), y.sym(cache)
	zero := cache.z3.FromInt(0, cache.sortInteger).(z3.Int)
	one := cache.z3.FromInt(1, cache.sortInteger).(z3.Int)
	return Integer{S: xs.Div(ys).Add(xs.Mod(ys).Eq(zero).Or(xs.GE(zero)).IfThenElse(zero, ys.GE(zero).IfThenElse(one, one.Neg())).(z3.Int))}
}

func (x Integer) Rem(y Integer) Integer {
	if x.IsConcrete() && y.IsConcrete() {
		z := Integer{C: new(big.Int)}
		z.C.Rem(x.C, y.C)
		return z
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	xs, ys := x.sym(cache), y.sym(cache)
	zero := cache.z3.FromInt(0, cache.sortInteger).(z3.Int)
	one := cache.z3.FromInt(1, cache.sortInteger).(z3.Int)
	return Integer{S: xs.Sub(xs.Div(ys).Add(xs.Mod(ys).Eq(zero).Or(xs.GE(zero)).IfThenElse(zero, ys.GE(zero).IfThenElse(one, one.Neg())).(z3.Int)).Mul(ys))}
}

func (x Integer) Eq(y Integer) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C.Cmp(y.C) == 0}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).Eq(y.sym(cache))}
}

func (x Integer) NE(y Integer) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C.Cmp(y.C) != 0}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).NE(y.sym(cache))}
}

func (x Integer) LT(y Integer) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C.Cmp(y.C) < 0}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).LT(y.sym(cache))}
}

func (x Integer) LE(y Integer) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C.Cmp(y.C) <= 0}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).LE(y.sym(cache))}
}

func (x Integer) GT(y Integer) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C.Cmp(y.C) > 0}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).GT(y.sym(cache))}
}

func (x Integer) GE(y Integer) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C.Cmp(y.C) >= 0}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).GE(y.sym(cache))}
}

func (x Integer) Neg() Integer {
	if x.IsConcrete() {
		z := Integer{C: new(big.Int)}
		z.C.Neg(x.C)
		return z
	}
	return Integer{S: x.S.Neg()}
}

// Real implements symbolic *big.Rat values.
type Real struct {
	C *big.Rat
	S z3.Real
}

// AnyReal returns an unconstrained symbolic Real.
func AnyReal(ctx *z3.Context, name string) Real {
	cache := getCache(ctx)
	sym := cache.z3.FreshConst(name, cache.sortReal).(z3.Real)
	return Real{S: sym}
}

// String returns x as a string.
func (x Real) String() string {
	if x.IsConcrete() {
		return fmt.Sprint(x.C)
	}
	return x.S.String()
}

// IsConcrete returns true if x is concrete.
func (x Real) IsConcrete() bool {
	return x.S.Context() == nil
}

// Eval returns x's concrete value in model m.
// This also evaluates x with model completion.
func (x Real) Eval(m *z3.Model) *big.Rat {
	if x.IsConcrete() {
		return x.C
	}
	c := m.Eval(x.S, true).(z3.Real)
	if c2, _, ok := c.Approx(RealApproxDigits); ok {
		c = c2
	}
	val, ok := c.AsBigRat()
	if !ok {
		panic("model evaluation produced non-concrete value " + c.String())
	}
	return (*big.Rat)(val)
}

// sym returns x's symbolic value, creating it if necessary.
func (x Real) sym(c *cache) z3.Real {
	if !x.IsConcrete() {
		return x.S
	}
	return c.z3.FromBigRat(x.C)
}

func (x Real) Add(y Real) Real {
	if x.IsConcrete() && y.IsConcrete() {
		z := Real{C: new(big.Rat)}
		z.C.Add(x.C, y.C)
		return z
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Real{S: x.sym(cache).Add(y.sym(cache))}
}

func (x Real) Sub(y Real) Real {
	if x.IsConcrete() && y.IsConcrete() {
		z := Real{C: new(big.Rat)}
		z.C.Sub(x.C, y.C)
		return z
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Real{S: x.sym(cache).Sub(y.sym(cache))}
}

func (x Real) Mul(y Real) Real {
	if x.IsConcrete() && y.IsConcrete() {
		z := Real{C: new(big.Rat)}
		z.C.Mul(x.C, y.C)
		return z
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Real{S: x.sym(cache).Mul(y.sym(cache))}
}

func (x Real) Quo(y Real) Real {
	if x.IsConcrete() && y.IsConcrete() {
		z := Real{C: new(big.Rat)}
		z.C.Quo(x.C, y.C)
		return z
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Real{S: x.sym(cache).Div(y.sym(cache))}
}

func (x Real) Eq(y Real) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C.Cmp(y.C) == 0}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).Eq(y.sym(cache))}
}

func (x Real) NE(y Real) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C.Cmp(y.C) != 0}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).NE(y.sym(cache))}
}

func (x Real) LT(y Real) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C.Cmp(y.C) < 0}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).LT(y.sym(cache))}
}

func (x Real) LE(y Real) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C.Cmp(y.C) <= 0}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).LE(y.sym(cache))}
}

func (x Real) GT(y Real) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C.Cmp(y.C) > 0}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).GT(y.sym(cache))}
}

func (x Real) GE(y Real) Bool {
	if x.IsConcrete() && y.IsConcrete() {
		return Bool{C: x.C.Cmp(y.C) >= 0}
	}
	ctx := x.S.Context()
	if ctx == nil {
		ctx = y.S.Context()
	}
	cache := getCache(ctx)
	return Bool{S: x.sym(cache).GE(y.sym(cache))}
}

func (x Real) Neg() Real {
	if x.IsConcrete() {
		z := Real{C: new(big.Rat)}
		z.C.Neg(x.C)
		return z
	}
	return Real{S: x.S.Neg()}
}

type ArrayInt32 struct {
	base   int64
	length int64
	ctx    *z3.Context
	mem    z3.Array
}

func AnyArrayInt32(ctx *z3.Context, name string, base, length int) ArrayInt32 {
	cache := getCache(ctx)
	arr32 := ctx.ArraySort(cache.sortUint32, cache.sortInt32)
	return ArrayInt32{
		ctx:    ctx,
		base:   int64(base),
		length: int64(length),
		mem:    ctx.FreshConst(name, arr32).(z3.Array),
	}
}

func (a *ArrayInt32) InBounds(idx Uint32, solv *z3.Solver) bool {
	if idx.IsConcrete() {
		return int64(idx.C) >= a.base && int64(idx.C) < a.base+a.length
	}

	cache := getCache(a.ctx)
	solv.Push()
	solv.Assert(idx.S.SGE(a.ctx.FromInt(a.base, cache.sortUint32).(z3.BV)))
	solv.Assert(idx.S.SLT(a.ctx.FromInt(a.base+a.length, cache.sortUint32).(z3.BV)))
	sat, err := solv.Check()
	if err != nil {
		panic(err)
	}
	solv.Pop()
	if !sat && err == nil {
		return false
	}
	return true
}

func (a *ArrayInt32) Read(idx Uint32, solv *z3.Solver) Int32 {
	cache := getCache(a.ctx)
	sidx := idx
	if idx.IsConcrete() {
		sidx = Uint32{S: cache.z3.FreshConst(fmt.Sprintf("addr(%x)", idx.C), cache.sortUint32).(z3.BV)}
		solv.Assert(sidx.S.Eq(a.ctx.FromInt(int64(idx.C), cache.sortUint32).(z3.BV)))
	}
	return Int32{S: a.mem.Select(sidx.S).(z3.BV)}
}

func (a *ArrayInt32) Write(idx Uint32, val Int32, solv *z3.Solver) {
	cache := getCache(a.ctx)
	sidx := idx
	if idx.IsConcrete() {
		sidx = Uint32{S: cache.z3.FreshConst(fmt.Sprintf("addr(%x)", idx.C), cache.sortUint32).(z3.BV)}
		solv.Assert(sidx.S.Eq(a.ctx.FromInt(int64(idx.C), cache.sortUint32).(z3.BV)))
	}
	sval := val
	if val.IsConcrete() {
		sval = Int32{S: cache.z3.FreshConst(fmt.Sprintf("val(%x)", val.C), cache.sortInt32).(z3.BV)}
		solv.Assert(sval.S.Eq(a.ctx.FromInt(int64(val.C), cache.sortInt32).(z3.BV)))
	}
	a.mem = a.mem.Store(sidx.S, sval.S)
}

func (a *ArrayInt32) String() string {
	return fmt.Sprintf("%v", a.mem)
}
