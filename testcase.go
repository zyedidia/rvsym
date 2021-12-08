package rvsym

import (
	"bytes"
	"fmt"
)

type Assignment struct {
	Name string
	Val  int32
}

type TestCase struct {
	Assignments []Assignment
	Exit        ExitStatus
	Pc          int32
	Err         error
	Dump        string
}

func (tc TestCase) String(hex bool) string {
	buf := &bytes.Buffer{}
	if tc.Err != nil {
		buf.WriteString(tc.Err.Error())
		buf.WriteByte('\n')
	}
	for _, a := range tc.Assignments {
		if hex {
			buf.WriteString(fmt.Sprintf("%s -> 0x%x\n", a.Name, uint32(a.Val)))
		} else {
			buf.WriteString(fmt.Sprintf("%s -> %d\n", a.Name, a.Val))
		}
	}
	return buf.String()
}
