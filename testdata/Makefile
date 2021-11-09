ASM=$(wildcard *.s)
C=$(wildcard *.c)
COBJ=$(C:.c=.o)
PROGS=$(ASM:.s=.bin) $(C:.c=.bin)

PREFIX ?= riscv32-unknown-elf-
ifeq (, $(shell which riscv32-unknown-elf-gcc))
PREFIX=riscv64-unknown-elf-
endif

CC=$(PREFIX)gcc
AS=$(PREFIX)as
LD=$(PREFIX)ld
OBJCOPY=$(PREFIX)objcopy
OBJDUMP=$(PREFIX)objdump

startobj = lib/start.o

INC=-I../include
CFLAGS=$(INC) -Os -Wall -Werror -nostdlib -nostartfiles -ffreestanding -march=rv32i -mabi=ilp32 -std=gnu99
ASFLAGS=-march=rv32i -mabi=ilp32
LDFLAGS=-T memmap.ld -melf32lriscv -nostdlib

all: $(PROGS)

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

%.o: %.s
	cpp $(INC) $< | $(AS) $(ASFLAGS) -c -o $@

%.elf: %.o $(startobj)
	$(LD) $(LDFLAGS) $(if $(filter $<,$(COBJ)), $(startobj)) $< -o $@

%.bin: %.elf
	$(OBJCOPY) $< -O binary $@

%.list: %.elf
	$(OBJDUMP) -D $< > $@

clean:
	rm -f *.bin *.o *.elf *.list

.PHONY: all clean
