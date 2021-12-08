package smt

type Int32 struct {
	C int32
	S SymInt32
}

func (a Int32) Concrete() bool {
	return !a.S.Valid()
}

func (a Int32) Sym(s *Solver) SymInt32 {
	if !a.Concrete() {
		return a.S
	}
	return s.ToSymInt32(a.C)
}

func (a Int32) Add(b Int32, s *Solver) Int32 {
	if a.Concrete() && b.Concrete() {
		return Int32{C: a.C + b.C}
	}
	if a.Concrete() && a.C == 0 {
		return b
	}
	if b.Concrete() && b.C == 0 {
		return a
	}
	return Int32{S: a.Sym(s).Add(b.Sym(s), s)}
}
func (a Int32) Sub(b Int32, s *Solver) Int32 {
	if a.Concrete() && b.Concrete() {
		return Int32{C: a.C - b.C}
	}
	if a.Concrete() && a.C == 0 {
		return b
	}
	if b.Concrete() && b.C == 0 {
		return a
	}
	return Int32{S: a.Sym(s).Sub(b.Sym(s), s)}
}
func (a Int32) And(b Int32, s *Solver) Int32 {
	if a.Concrete() && b.Concrete() {
		return Int32{C: a.C & b.C}
	}
	if a.Concrete() && uint32(a.C) == 0xffffffff {
		return b
	}
	if b.Concrete() && uint32(b.C) == 0xffffffff {
		return a
	}
	if a.Concrete() && uint32(a.C) == 0 {
		return a
	}
	if b.Concrete() && uint32(b.C) == 0 {
		return b
	}
	return Int32{S: a.Sym(s).And(b.Sym(s), s)}
}
func (a Int32) Or(b Int32, s *Solver) Int32 {
	if a.Concrete() && b.Concrete() {
		return Int32{C: a.C | b.C}
	}
	if a.Concrete() && a.C == 0 {
		return b
	}
	if b.Concrete() && b.C == 0 {
		return a
	}
	if a.Concrete() && uint32(a.C) == 0xffffffff {
		return a
	}
	if b.Concrete() && uint32(b.C) == 0xffffffff {
		return b
	}
	return Int32{S: a.Sym(s).Or(b.Sym(s), s)}
}
func (a Int32) Xor(b Int32, s *Solver) Int32 {
	if a.Concrete() && b.Concrete() {
		return Int32{C: a.C ^ b.C}
	}
	if a.Concrete() && a.C == 0 {
		return b
	}
	if b.Concrete() && b.C == 0 {
		return a
	}
	return Int32{S: a.Sym(s).Xor(b.Sym(s), s)}
}
func (a Int32) Sll(b Int32, s *Solver) Int32 {
	if a.Concrete() && b.Concrete() {
		return Int32{C: a.C << b.C}
	}
	if a.Concrete() && a.C == 0 {
		return Int32{C: 0}
	}
	if b.Concrete() && b.C == 0 {
		return a
	}
	return Int32{S: a.Sym(s).Sll(b.Sym(s), s)}
}
func (a Int32) Srl(b Int32, s *Solver) Int32 {
	if a.Concrete() && b.Concrete() {
		return Int32{C: int32(uint32(a.C) >> uint32(b.C))}
	}
	if a.Concrete() && a.C == 0 {
		return Int32{C: 0}
	}
	if b.Concrete() && b.C == 0 {
		return a
	}
	return Int32{S: a.Sym(s).Srl(b.Sym(s), s)}
}
func (a Int32) Sra(b Int32, s *Solver) Int32 {
	if a.Concrete() && b.Concrete() {
		return Int32{C: a.C >> b.C}
	}
	if a.Concrete() && a.C == 0 {
		return Int32{C: 0}
	}
	if b.Concrete() && b.C == 0 {
		return a
	}
	return Int32{S: a.Sym(s).Sra(b.Sym(s), s)}
}
func (a Int32) Slt(b Int32, s *Solver) Int32 {
	if a.Concrete() && b.Concrete() {
		var c int32
		if a.C < b.C {
			c = 1
		}
		return Int32{C: c}
	}
	return Int32{S: a.Sym(s).Slt(b.Sym(s), s)}
}
func (a Int32) Ult(b Int32, s *Solver) Int32 {
	if a.Concrete() && b.Concrete() {
		var c int32
		if uint32(a.C) < uint32(b.C) {
			c = 1
		}
		return Int32{C: c}
	}
	return Int32{S: a.Sym(s).Ult(b.Sym(s), s)}
}
func (a Int32) Mul(b Int32, s *Solver) Int32 {
	if a.Concrete() && b.Concrete() {
		return Int32{C: a.C * b.C}
	}
	return Int32{S: a.Sym(s).Mul(b.Sym(s), s)}
}
func (a Int32) Div(b Int32, s *Solver) Int32 {
	if a.Concrete() && b.Concrete() {
		return Int32{C: a.C / b.C}
	}
	return Int32{S: a.Sym(s).Div(b.Sym(s), s)}
}
func (a Int32) Divu(b Int32, s *Solver) Int32 {
	if a.Concrete() && b.Concrete() {
		return Int32{C: int32(uint32(a.C) / uint32(b.C))}
	}
	return Int32{S: a.Sym(s).Divu(b.Sym(s), s)}
}
func (a Int32) Rem(b Int32, s *Solver) Int32 {
	if a.Concrete() && b.Concrete() {
		return Int32{C: a.C % b.C}
	}
	return Int32{S: a.Sym(s).Rem(b.Sym(s), s)}
}
func (a Int32) Remu(b Int32, s *Solver) Int32 {
	if a.Concrete() && b.Concrete() {
		return Int32{C: int32(uint32(a.C) % uint32(b.C))}
	}
	return Int32{S: a.Sym(s).Remu(b.Sym(s), s)}
}
func (a Int32) Not(s *Solver) Int32 {
	if a.Concrete() {
		return Int32{C: ^a.C}
	}
	return Int32{S: a.S.Not(s)}
}
func (a Int32) Eqz(s *Solver) Bool {
	if a.Concrete() {
		return Bool{C: a.C == 0}
	}
	return Bool{S: a.S.Eqz(s)}
}
func (a Int32) NEqz(s *Solver) Bool {
	if a.Concrete() {
		return Bool{C: a.C != 0}
	}
	return Bool{S: a.S.NEqz(s)}
}
func (a Int32) Eqb(b Int32, s *Solver) Bool {
	if a.Concrete() && b.Concrete() {
		return Bool{C: a.C == b.C}
	}
	// if reflect.DeepEqual(a.S, b.S) {
	// 	return Bool{C: true}
	// }
	return Bool{S: a.Sym(s).Eqb(b.Sym(s), s)}
}
func (a Int32) NEqb(b Int32, s *Solver) Bool {
	if a.Concrete() && b.Concrete() {
		return Bool{C: a.C != b.C}
	}
	return Bool{S: a.Sym(s).NEqb(b.Sym(s), s)}
}
func (a Int32) Sltb(b Int32, s *Solver) Bool {
	if a.Concrete() && b.Concrete() {
		return Bool{C: a.C < b.C}
	}
	return Bool{S: a.Sym(s).Sltb(b.Sym(s), s)}
}
func (a Int32) Ultb(b Int32, s *Solver) Bool {
	if a.Concrete() && b.Concrete() {
		return Bool{C: uint32(a.C) < uint32(b.C)}
	}
	return Bool{S: a.Sym(s).Ultb(b.Sym(s), s)}
}
func (a Int32) Sgeb(b Int32, s *Solver) Bool {
	if a.Concrete() && b.Concrete() {
		return Bool{C: a.C >= b.C}
	}
	return Bool{S: a.Sym(s).Sgeb(b.Sym(s), s)}
}
func (a Int32) Ugeb(b Int32, s *Solver) Bool {
	if a.Concrete() && b.Concrete() {
		return Bool{C: uint32(a.C) >= uint32(b.C)}
	}
	return Bool{S: a.Sym(s).Ugeb(b.Sym(s), s)}
}
func (a Int32) ToInt8(s *Solver) Int8 {
	if a.Concrete() {
		return Int8{C: int8(a.C)}
	}
	return Int8{S: a.S.ToInt8(s)}
}
func (a Int32) ToInt16(s *Solver) Int16 {
	if a.Concrete() {
		return Int16{C: int16(a.C)}
	}
	return Int16{S: a.S.ToInt16(s)}
}
func (a Int32) ToInt64s(s *Solver) Int64 {
	if a.Concrete() {
		return Int64{C: int64(a.C)}
	}
	return Int64{S: a.S.ToInt64s(s)}
}
func (a Int32) ToInt64z(s *Solver) Int64 {
	if a.Concrete() {
		return Int64{C: int64(uint32(a.C))}
	}
	return Int64{S: a.S.ToInt64z(s)}
}

type Int8 struct {
	C int8
	S SymInt8
}

func (a Int8) Concrete() bool {
	return !a.S.Valid()
}

func (a Int8) ToInt32s(s *Solver) Int32 {
	if a.Concrete() {
		return Int32{C: int32(a.C)}
	}
	return Int32{S: a.S.ToInt32s(s)}
}

func (a Int8) ToInt32z(s *Solver) Int32 {
	if a.Concrete() {
		return Int32{C: int32(uint8(a.C))}
	}
	return Int32{S: a.S.ToInt32z(s)}
}

type Int16 struct {
	C int16
	S SymInt16
}

func (a Int16) Concrete() bool {
	return !a.S.Valid()
}

func (a Int16) ToInt32s(s *Solver) Int32 {
	if a.Concrete() {
		return Int32{C: int32(a.C)}
	}
	return Int32{S: a.S.ToInt32s(s)}
}

func (a Int16) ToInt32z(s *Solver) Int32 {
	if a.Concrete() {
		return Int32{C: int32(uint16(a.C))}
	}
	return Int32{S: a.S.ToInt32z(s)}
}

type Int64 struct {
	C int64
	S SymInt64
}

func (a Int64) Sym(s *Solver) SymInt64 {
	if !a.Concrete() {
		return a.S
	}
	return s.ToSymInt64(a.C)
}

func (a Int64) Concrete() bool {
	return !a.S.Valid()
}

func (a Int64) Lower32(s *Solver) Int32 {
	if a.Concrete() {
		return Int32{C: int32(a.C)}
	}
	return Int32{S: a.S.Lower32(s)}
}

func (a Int64) Upper32(s *Solver) Int32 {
	if a.Concrete() {
		return Int32{C: int32(a.C >> 32)}
	}
	return Int32{S: a.S.Upper32(s)}
}

func (a Int64) Mul(b Int64, s *Solver) Int64 {
	if a.Concrete() && b.Concrete() {
		return Int64{C: a.C * b.C}
	}
	return Int64{S: a.Sym(s).Mul(b.Sym(s), s)}
}

type Bool struct {
	C bool
	S SymBool
}

func (a Bool) Concrete() bool {
	return !a.S.Valid()
}

func (a Bool) Sym(s *Solver) SymBool {
	if !a.Concrete() {
		return a.S
	}
	return s.ToSymBool(a.C)
}

func (a Bool) Not(s *Solver) Bool {
	if a.Concrete() {
		return Bool{C: !a.C}
	}
	return Bool{S: a.S.Not(s)}
}

type ArrayInt32 struct {
	base   uint32
	length uint32
	S      SymArrayInt32
}

func (a ArrayInt32) InBounds(idx Int32, s *Solver) bool {
	if idx.Concrete() {
		return uint32(idx.C) >= a.base && uint32(idx.C) < a.base+a.length
	}

	s.Push()
	s.Assert(Bool{S: idx.S.Ugeb(s.ToSymInt32(int32(a.base)), s)})
	s.Assert(Bool{S: idx.S.Ultb(s.ToSymInt32(int32(a.base+a.length)), s)})
	res := s.Check(false)
	s.Pop()
	return res == Sat
}

func (a ArrayInt32) Read(idx Int32, s *Solver) Int32 {
	return Int32{S: a.S.Select(idx.Sym(s), s)}
}

func (a *ArrayInt32) Write(idx, val Int32, s *Solver) {
	a.S = a.S.Store(idx.Sym(s), val.Sym(s), s)
}

func (a *ArrayInt32) WriteInitial(idx, val Int32, s *Solver) {
	a.S = a.S.StoreWithSelect(idx.Sym(s), val.Sym(s), s)
}
