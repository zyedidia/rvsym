package rvsym

const (
	Rzero = iota
	Rra
	Rsp
	Rgp
	Rtp
	Rt0
	Rt1
	Rt2
	Rs0
	Rs1
	Ra0
	Ra1
	Ra2
	Ra3
	Ra4
	Ra5
	Ra6
	Ra7
	Rs2
	Rs3
	Rs4
	Rs5
	Rs6
	Rs7
	Rs8
	Rs9
	Rs10
	Rs11
	Rt3
	Rt4
	Rt5
	Rt6
)

var RegNames = map[int]string{
	0:  "zero",
	1:  "ra",
	2:  "sp",
	3:  "gp",
	4:  "tp",
	5:  "t0",
	6:  "t1",
	7:  "t2",
	8:  "s0",
	9:  "s1",
	10: "a0",
	11: "a1",
	12: "a2",
	13: "a3",
	14: "a4",
	15: "a5",
	16: "a6",
	17: "a7",
	18: "s2",
	19: "s3",
	20: "s4",
	21: "s5",
	22: "s6",
	23: "s7",
	24: "s8",
	25: "s9",
	26: "s10",
	27: "s11",
	28: "t3",
	29: "t4",
	30: "t5",
	31: "t6",
}
