include $(RVSYM_ROOT)/lib/rules.mk

C=$(wildcard *.c)
PROGS=$(C:.c=.bin)

all: $(PROGS)

clean:
	rm -f *.bin *.o *.elf *.list

.PHONY: all clean
