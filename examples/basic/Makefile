RVSYM_ROOT ?= $(shell git rev-parse --show-toplevel)

include $(RVSYM_ROOT)/lib/rules.mk

C=$(wildcard *.c)
PROGS=$(C:.c=.hex)

all: $(PROGS)

clean:
	rm -f *.bin *.o *.elf *.list *.hex

.PHONY: all clean
