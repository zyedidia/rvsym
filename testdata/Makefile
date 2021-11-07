SRCS=$(wildcard *.s)
PROGS=$(SRCS:.s=.bin)

all: $(PROGS)

%.bin: %.s
	rvasm -base 0 -raw $< > $@

clean:
	rm -f *.bin

.PHONY: all clean
