//go:build nosmt

package smt

type Solver struct{}
type Model struct{}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) Push()         {}
func (s *Solver) Pop()          {}
func (s *Solver) Assert(b Bool) {}
func (s *Solver) Check() CheckResult {
	return Unknown
}
func (s *Solver) Model() Model {
	return Model{}
}
func (m Model) Eval(a Int32) int32 {
	return 0
}
func (s *Solver) AnyInt32() Int32 {
	return Int32{}
}
func (s *Solver) AnyArrayInt32(base, length int) ArrayInt32 {
	return ArrayInt32{}
}
func (s *Solver) ToSymInt32(c int32) SymInt32 {
	return SymInt32{}
}
func (s *Solver) ToSymInt64(c int64) SymInt64 {
	return SymInt64{}
}
func (s *Solver) ToSymBool(c bool) SymBool {
	return SymBool{}
}

type SymInt8 struct{}

func (a SymInt8) Valid() bool {
	return false
}

func (a SymInt8) ToInt32s(s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt8) ToInt32z(s *Solver) SymInt32 {
	return SymInt32{}
}

type SymInt16 struct{}

func (a SymInt16) Valid() bool {
	return false
}

func (a SymInt16) ToInt32s(s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt16) ToInt32z(s *Solver) SymInt32 {
	return SymInt32{}
}

type SymInt64 struct{}

func (a SymInt64) Valid() bool {
	return false
}

func (a SymInt64) Lower32(s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt64) Upper32(s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt64) Mul(b SymInt64, s *Solver) SymInt64 {
	return SymInt64{}
}

type SymBool struct{}

func (a SymBool) Valid() bool {
	return false
}
func (a SymBool) Not(s *Solver) SymBool {
	return SymBool{}
}

type SymInt32 struct{}

func (a SymInt32) Valid() bool {
	return false
}

func (a SymInt32) Add(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt32) Sub(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt32) And(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt32) Or(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt32) Xor(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt32) Sll(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt32) Srl(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt32) Sra(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt32) Slt(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt32) Ult(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt32) Mul(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt32) Div(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt32) Divu(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt32) Rem(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt32) Remu(b SymInt32, s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt32) Not(s *Solver) SymInt32 {
	return SymInt32{}
}
func (a SymInt32) Eqz(s *Solver) SymBool {
	return SymBool{}
}
func (a SymInt32) NEqz(s *Solver) SymBool {
	return SymBool{}
}
func (a SymInt32) Eqb(b SymInt32, s *Solver) SymBool {
	return SymBool{}
}
func (a SymInt32) NEqb(b SymInt32, s *Solver) SymBool {
	return SymBool{}
}
func (a SymInt32) Sltb(b SymInt32, s *Solver) SymBool {
	return SymBool{}
}
func (a SymInt32) Sgeb(b SymInt32, s *Solver) SymBool {
	return SymBool{}
}
func (a SymInt32) Ultb(b SymInt32, s *Solver) SymBool {
	return SymBool{}
}
func (a SymInt32) Ugeb(b SymInt32, s *Solver) SymBool {
	return SymBool{}
}
func (a SymInt32) ToInt8(s *Solver) SymInt8 {
	return SymInt8{}
}
func (a SymInt32) ToInt16(s *Solver) SymInt16 {
	return SymInt16{}
}
func (a SymInt32) ToInt64s(s *Solver) SymInt64 {
	return SymInt64{}
}
func (a SymInt32) ToInt64z(s *Solver) SymInt64 {
	return SymInt64{}
}

type SymArrayInt32 struct{}

func (a SymArrayInt32) Select(idx SymInt32, s *Solver) SymInt32 {
	return SymInt32{}
}

func (a SymArrayInt32) Store(idx SymInt32, val SymInt32, s *Solver) SymArrayInt32 {
	return SymArrayInt32{}
}
