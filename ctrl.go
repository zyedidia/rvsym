package rvsym

import "github.com/zyedidia/rvsym/bits"

type ExitStatus byte

const (
	ExitNone = iota
	ExitNormal
	ExitQuiet
	ExitFail
	ExitMem
	ExitUnsat
)

func (e ExitStatus) String() string {
	switch e {
	case ExitNone:
		return "no exit"
	case ExitNormal:
		return "exit"
	case ExitFail:
		return "failure"
	case ExitUnsat:
		return "unsat"
	case ExitMem:
		return "memory failure"
	}
	return "quiet"
}

// control code definitions

const (
	OpRarith = 0b0110011
	OpIarith = 0b0010011
	OpBranch = 0b1100011
	OpLui    = 0b0110111
	OpAuipc  = 0b0010111
	OpJal    = 0b1101111
	OpJalr   = 0b1100111
	OpLoad   = 0b0000011
	OpStore  = 0b0100011
	OpFence  = 0b0001111
	OpSys    = 0b1110011
)

const (
	ExtByte  = 0b000
	ExtHalf  = 0b001
	ExtWord  = 0b010
	ExtByteU = 0b100
	ExtHalfU = 0b101
)

const (
	SymSymbolicRegs = iota
	SymFail
	SymPrint
	SymExit
	SymQuietExit
	SymMarkBytes
	SymDump
	SymMarkArray
)

const (
	InsnEcall  = 0x00000073
	InsnEbreak = 0x00100073
	InsnNop    = 0x00000013
)

func rd(insn uint32) uint32 {
	return bits.Get(insn, 11, 7)
}
func rs1(insn uint32) uint32 {
	return bits.Get(insn, 19, 15)
}
func rs2(insn uint32) uint32 {
	return bits.Get(insn, 24, 20)
}
func shamt(insn uint32) uint32 {
	return bits.Get(insn, 24, 20)
}
func funct3(insn uint32) uint32 {
	return bits.Get(insn, 14, 12)
}
func funct7(insn uint32) uint32 {
	return bits.Get(insn, 31, 25)
}
