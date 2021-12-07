PREFIX ?= riscv32-unknown-elf
ifeq (, $(shell which riscv32-unknown-elf-gcc))
PREFIX=riscv64-unknown-elf
endif

RV_ROOT ?= /opt/riscv

CC=$(PREFIX)-gcc
CXX=$(PREFIX)-g++
AS=$(PREFIX)-as
LD=$(PREFIX)-ld
OBJCOPY=$(PREFIX)-objcopy
OBJDUMP=$(PREFIX)-objdump

RVSYM_ROOT ?=
RVSYM_LIB=$(RVSYM_ROOT)/lib
RVSYM_INCLUDE=$(RVSYM_ROOT)/include
INCLUDE=-I$(RVSYM_INCLUDE)

O ?= 0

CXXFLAGS=-O$(O) $(INCLUDE) -g -Wall -nostdlib -nostartfiles -ffreestanding -march=rv32im -mabi=ilp32 -std=c++14
CFLAGS=-O$(O) $(INCLUDE) -g -Wall -nostdlib -nostartfiles -ffreestanding -march=rv32im -mabi=ilp32 -std=gnu99
ASFLAGS=-march=rv32im -mabi=ilp32
LDFLAGS=-T $(RVSYM_LIB)/memmap.ld -melf32lriscv -L$(RV_ROOT)/$(PREFIX)/lib/rv32im/ilp32

LIBOBJ += $(RVSYM_LIB)/start.o $(RVSYM_LIB)/libc.o

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

%.o: %.s
	cpp -I$(RVSYM_INCLUDE) $< | $(AS) $(ASFLAGS) -c -o $@

%.elf: %.o $(LIBOBJ)
	$(LD) $(LDFLAGS) $(LIBOBJ) $< $(LDLIBS) -o $@

%.bin: %.elf
	$(OBJCOPY) $< -O binary --set-section-flags .sbss=alloc,load,contents $@

%.list: %.elf
	$(OBJDUMP) -D $< > $@
