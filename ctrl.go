package rvsym

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
	InsnEcall  = 0x00000073
	InsnEbreak = 0x00100073
	InsnNop    = 0x00000013
)

const (
	AluAdd  = 0b000
	AluShl  = 0b001
	AluSlt  = 0b010
	AluSltu = 0b011
	AluXor  = 0b100
	AluShr  = 0b101
	AluOr   = 0b110
	AluAnd  = 0b111

	MAluMul    = 0b000
	MAluMulH   = 0b001
	MAluMulHSU = 0b010
	MAluMulHU  = 0b011
	MAluDiv    = 0b100
	MAluDivU   = 0b101
	MAluRem    = 0b110
	MAluRemU   = 0b111
)

type ImmType byte

const (
	ImmI ImmType = iota
	ImmS
	ImmB
	ImmJ
	ImmU
)

const (
	SymSymbolicRegs = iota
	SymFail
	SymSymbolicReg
	SymExit
	SymQuietExit
	SymMarkNBytes
	SymDump
	SymMarkArray
)

type ExitStatus byte

const (
	ExitNone = iota
	ExitNormal
	ExitQuiet
	ExitFail
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
	}
	return "quiet"
}

const (
	ExtByte  = 0b000
	ExtHalf  = 0b001
	ExtWord  = 0b010
	ExtByteU = 0b100
	ExtHalfU = 0b101
)
