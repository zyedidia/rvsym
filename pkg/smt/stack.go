package smt

type stack struct {
	entries []Bool
	pushes  []int
	count   int
}

func newStack() *stack {
	return &stack{
		entries: make([]Bool, 0),
		pushes:  make([]int, 0),
	}
}

func (s *stack) add(b Bool) {
	s.entries = append(s.entries, b)
	s.count++
}

func (s *stack) push() {
	s.pushes = append(s.pushes, s.count)
	s.count = 0
}

func (s *stack) pop() {
	s.entries = s.entries[:len(s.entries)-s.count]
	s.count = s.pushes[len(s.pushes)-1]
	s.pushes = s.pushes[:len(s.pushes)-1]
}
