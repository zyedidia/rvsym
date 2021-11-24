include $(RVSYM_ROOT)/lib/rules.mk

ASM=$(wildcard *.s)
C=$(wildcard *.c)
COBJ=$(C:.c=.o)
PROGS=$(ASM:.s=.bin) $(C:.c=.bin)

all: $(PROGS)

clean:
	rm -f *.bin *.o *.elf *.list

.PHONY: all clean
