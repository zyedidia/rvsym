PREFIX ?= riscv32-unknown-elf-
ifeq (, $(shell which riscv32-unknown-elf-gcc))
PREFIX=riscv64-unknown-elf-
endif

CC=$(PREFIX)gcc
CXX=$(PREFIX)g++
AS=$(PREFIX)as
LD=$(PREFIX)ld
OBJCOPY=$(PREFIX)objcopy
OBJDUMP=$(PREFIX)objdump

INC=-I../../include
CXXFLAGS=$(INC) -Os -Wall -Werror -nostartfiles -ffreestanding -march=rv32i -mabi=ilp32 -std=c++14
CFLAGS=-Os -Wall -Werror -nostdlib -nostartfiles -ffreestanding -march=rv32i -mabi=ilp32 -std=gnu99
ASFLAGS=-march=rv32i -mabi=ilp32
LDFLAGS=-T memmap.ld -melf32lriscv -L/opt/riscv/riscv32-unknown-elf/lib

LIBOBJ=lib/start.o lib/libc.o

# YOSYS_INCLUDE=../../../yosys
YOSYS_INCLUDE=$(shell yosys-config --datdir)/include

all: counter.bin

counter.o: counter.cc counter.hh
counter.bin: counter.o

%.hh: %.sv
	yosys -q -p 'read_verilog -sv -noautowire $<; write_cxxrtl -g0 -O4 $@'

%.o: %.cc %.hh
	$(CXX) $(CXXFLAGS) -I$(YOSYS_INCLUDE) -c $< -o $@

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

%.o: %.s
	cpp $(INC) $< | $(AS) $(ASFLAGS) -c -o $@

%.elf: %.o $(LIBOBJ)
	$(LD) $(LDFLAGS) $(LIBOBJ) $< -lstdc++ -o $@

%.bin: %.elf
	$(OBJCOPY) $< -O binary $@

%.list: %.elf
	$(OBJDUMP) -D $< > $@

clean:
	rm -f *.o *.bin *.elf *.list *.hh
