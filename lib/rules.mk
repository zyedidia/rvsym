PREFIX ?= riscv32-unknown-elf
ifeq (, $(shell which riscv32-unknown-elf-gcc))
PREFIX=riscv64-unknown-elf
endif

RISCV ?= /opt/riscv

CC=$(PREFIX)-gcc
CXX=$(PREFIX)-g++
AS=$(PREFIX)-as
LD=$(PREFIX)-ld
OBJCOPY=$(PREFIX)-objcopy
OBJDUMP=$(PREFIX)-objdump

RVSYM_ROOT ?= $(shell git rev-parse --show-toplevel)
RVSYM_LIB=$(RVSYM_ROOT)/lib
RVSYM_INCLUDE=$(RVSYM_ROOT)/include
INCLUDE=-I$(RVSYM_INCLUDE)

O ?= 2

ARCH=rv32im
CXXFLAGS=-O$(O) $(INCLUDE) -g -Wall -nostdlib -nostartfiles -ffreestanding -march=$(ARCH) -mabi=ilp32 -std=c++14
CFLAGS=-O$(O) $(INCLUDE) -g -Wall -nostdlib -nostartfiles -ffreestanding -march=$(ARCH) -mabi=ilp32 -std=gnu99
ASFLAGS=-march=$(ARCH) -mabi=ilp32
LDFLAGS=-T $(RVSYM_LIB)/memmap.ld -melf32lriscv -L$(RISCV)/$(PREFIX)/lib/$(ARCH)/ilp32

LIBOBJ += $(RVSYM_LIB)/start.o $(RVSYM_LIB)/libc.o $(RVSYM_LIB)/cstart.o

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

%.o: %.s
	cpp -I$(RVSYM_INCLUDE) $< | $(AS) $(ASFLAGS) -c -o $@

%.elf: %.o $(LIBOBJ)
	$(LD) $(LDFLAGS) $(LIBOBJ) $< $(LDLIBS) -o $@

%.bin: %.elf
	$(OBJCOPY) $< -S -O binary $@

%.hex: %.elf
	$(OBJCOPY) $< -S -O ihex $@

%.list: %.elf
	$(OBJDUMP) -D $< > $@
