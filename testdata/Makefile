SRCS=$(wildcard *.s)
PROGS=$(SRCS:.s=.bin)

PREFIX ?= riscv32-unknown-elf-
ifeq (, $(shell which riscv32-unknown-elf-gcc))
PREFIX=riscv64-unknown-elf-
endif

CC=$(PREFIX)gcc
AS=$(PREFIX)as
LD=$(PREFIX)ld
OBJCOPY=$(PREFIX)objcopy
OBJDUMP=$(PREFIX)objdump

INC=-I../include
CFLAGS=$(INC) -Os -Wall -Werror -nostdlib -nostartfiles -ffreestanding -march=rv32i -mabi=ilp32 -std=gnu99
ASFLAGS=-march=rv32i -mabi=ilp32
LDFLAGS=-T memmap.ld -melf32lriscv -nostdlib

all: $(PROGS)

%.o: %.s
	cpp $(INC) $< $($@) | $(AS) $(ASFLAGS) -c -o $@

%.elf: %.o
	$(LD) $(LDFLAGS) $< -o $@

%.bin: %.elf
	$(OBJCOPY) $< -O binary $@

%.list: %.elf
	$(OBJDUMP) -D $< > $@

clean:
	rm -f *.bin *.o *.elf *.list *.p

.PHONY: all clean
