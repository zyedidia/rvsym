package smt

type CheckResult byte

const (
	Sat CheckResult = iota
	Unsat
	Unknown
)
