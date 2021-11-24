PREFIX ?= riscv32-unknown-elf-
ifeq (, $(shell which riscv32-unknown-elf-gcc))
PREFIX=riscv64-unknown-elf-
endif

synth=picorv32.v

CC=$(PREFIX)gcc
CXX=$(PREFIX)g++
AS=$(PREFIX)as
LD=$(PREFIX)ld
OBJCOPY=$(PREFIX)objcopy
OBJDUMP=$(PREFIX)objdump

INC=-I../../include
CXXFLAGS=$(INC) -O2 -g -Wall -nostartfiles -ffreestanding -march=rv32i -mabi=ilp32 -std=c++14 -fpermissive
CFLAGS=-O2 -g -Wall -Werror -nostdlib -nostartfiles -ffreestanding -march=rv32i -mabi=ilp32 -std=gnu99
ASFLAGS=-g -march=rv32i -mabi=ilp32
LDFLAGS=-T memmap.ld -melf32lriscv -L/opt/riscv/riscv64-unknown-elf/lib/rv32i/ilp32 -L/opt/riscv/lib/gcc/riscv64-unknown-elf/11.1.0/rv32i/ilp32

LIBOBJ=lib/start.o lib/libc.o

YOSIM_SO=/home/zyedidia/programming/yosim/cxx/cxx.so
YOSIM_INCLUDE=/home/zyedidia/programming/yosim/cxx
YOSYS_INCLUDE=$(shell yosys-config --datdir)/include

all: pico.bin

pico.o: pico.cc pico.hh
pico.bin: pico.o

pico.hh: $(synth)
	yosys -m $(YOSIM_SO) -q -p 'read_verilog $(SV_INC) $(DEFINE) -sv -noautowire $(synth); hierarchy -top picorv32; write_yosimcxx -nohierarchy -g0 -O6 $@'

%.o: %.cc %.hh
	$(CXX) $(CXXFLAGS) -DRVSYM -DCXXRTL_NDEBUG -DNDEBUG -I$(YOSIM_INCLUDE) -I$(YOSYS_INCLUDE) -c $< -o $@

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

%.o: %.s
	cpp $(INC) $< | $(AS) $(ASFLAGS) -c -o $@

%.elf: %.o $(LIBOBJ)
	$(LD) $(LDFLAGS) $(LIBOBJ) $< -lstdc++ -lgcc -o $@

%.bin: %.elf
	$(OBJCOPY) $< -O binary $@

%.list: %.elf
	$(OBJDUMP) -D $< > $@

clean:
	rm -f *.o *.bin *.elf *.list *.hh
