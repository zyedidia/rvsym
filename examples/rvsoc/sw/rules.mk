PREFIX ?= riscv32-unknown-elf-
ifeq (, $(shell which riscv32-unknown-elf-gcc))
PREFIX=riscv64-unknown-elf-
endif

CC=$(PREFIX)gcc
AS=$(PREFIX)as
LD=$(PREFIX)ld
OBJCOPY=$(PREFIX)objcopy
OBJDUMP=$(PREFIX)objdump

INC=-I../../librock
CFLAGS=$(INC) -Os -Wall -Werror -nostdlib -nostartfiles -ffreestanding -march=rv32i -mabi=ilp32 -std=gnu99
ASFLAGS=-march=rv32i -mabi=ilp32
LDFLAGS=-T ../../librock/memmap.ld -melf32lriscv -nostdlib

PROG ?= blink

OBJ=$(PROG).o
obj=$(OBJ) ../../librock/start.o

INSTALL_DIR=../../../rtl/mem/

all: $(PROG).mem

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

%.o: %.s
	$(AS) $(ASFLAGS) -c $< -o $@

%.elf: $(obj)
	$(LD) $(LDFLAGS) $(obj) -o $@

%.bin: %.elf
	$(OBJCOPY) $< -O binary $@

%.list: %.elf
	$(OBJDUMP) -D $< > $@

%.mem: %.bin
	bin2hex $< > $@

install: $(PROG).mem
	cp $< $(INSTALL_DIR)

clean:
	rm -f *.elf *.list *.bin *.mem *.o

.PHONY: all install clean
