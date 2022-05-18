package addr2line

import (
	"bytes"
	"debug/elf"
	"fmt"
	"log"
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

func (c *Converter) FuncToAddr(fn string) (uint32, error) {
	r, err := os.Open(c.Elf)
	if err != nil {
		return 0, err
	}
	defer r.Close()
	f, err := elf.NewFile(r)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	if f.Type != elf.ET_EXEC {
		return 0, fmt.Errorf("invalid elf file type: %v", f.Type)
	}

	symbols, err := f.Symbols()
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range symbols {
		if elf.ST_TYPE(s.Info) == elf.STT_FUNC && s.Name == fn {
			return uint32(s.Value), nil
		}
	}
	return 0, fmt.Errorf("function %s not found", fn)
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
