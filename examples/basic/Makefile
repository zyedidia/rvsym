PREFIX ?= riscv32-unknown-elf
ifeq (, $(shell which riscv32-unknown-elf-gcc))
PREFIX=riscv64-unknown-elf
endif

CC=$(PREFIX)-gcc
CXX=$(PREFIX)-g++
AS=$(PREFIX)-as
LD=$(PREFIX)-ld
OBJCOPY=$(PREFIX)-objcopy
OBJDUMP=$(PREFIX)-objdump

ARCH=rv32im
O ?= 2
RVSYM_ROOT ?= $(shell git rev-parse --show-toplevel)
RVSYM_INCLUDE=$(RVSYM_ROOT)/include
INCLUDE=-I$(RVSYM_INCLUDE)

CXXFLAGS=-O$(O) $(INCLUDE) -g -Wall -march=$(ARCH) -mabi=ilp32 -std=c++14
CFLAGS=-O$(O) $(INCLUDE) -g -Wall -march=$(ARCH) -mabi=ilp32 -std=gnu99

C=$(wildcard *.c)
PROGS=$(C:.c=.elf)

all: $(PROGS)

%.elf: %.c
	$(CC) $(CFLAGS) $< -o $@

clean:
	rm -f *.bin *.o *.elf *.list *.hex *.out

.PHONY: all clean
