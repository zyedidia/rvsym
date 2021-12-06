package addr2line

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

type Converter struct {
	Elf string
}

type Position struct {
	File string
	Line int
}

func (c *Converter) Lookup(addr uint32) (Position, error) {
	out, err := exec.Command("addr2line", "-e", c.Elf, "-i", "-a", fmt.Sprintf("%x", addr)).Output()
	if err != nil {
		return Position{}, err
	}
	out = bytes.TrimSpace(out)
	lines := bytes.Split(out, []byte{'\n'})
	if len(lines) <= 1 {
		return Position{}, fmt.Errorf("no file position information found")
	}
	line := lines[len(lines)-1]
	parts := bytes.Split(line, []byte{':'})
	if len(parts) == 2 {
		if string(parts[0]) == "??" {
			return Position{}, fmt.Errorf("no file position for address 0x%x", addr)
		}
		n, err := strconv.Atoi(string(parts[1]))
		if err != nil {
			return Position{}, err
		}
		return Position{
			File: string(parts[0]),
			Line: n,
		}, nil
	} else {
		return Position{}, fmt.Errorf("invalid file position format: %s", string(line))
	}
}

func (p Position) String() string {
	var path string
	current, err := os.Getwd()
	if err != nil {
		path = p.File
	} else if rel, err := filepath.Rel(current, p.File); err != nil {
		path = p.File
	} else {
		path = rel
	}
	return fmt.Sprintf("%s:%d", path, p.Line)
}
