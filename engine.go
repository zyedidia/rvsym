package rvsym

type Engine struct {
}

func NewEngine(segs []Segment, entry uint32, mode EmuMode) *Engine {
	return nil
}

func (e *Engine) Step() bool {
	return false
}

func (e *Engine) handleExit(m *Machine) bool {
	return false
}

func (e *Engine) hasExit(m *Machine) bool {
	return false
}

func (e *Engine) TestCases() []TestCase {
	return nil
}

func (e *Engine) NumTestCases() int {
	return 0
}

func (e *Engine) Summary() string {
	return ""
}
