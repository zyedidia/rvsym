PREFIX ?= riscv32-unknown-elf
ifeq (, $(shell which riscv32-unknown-elf-gcc))
PREFIX=riscv64-unknown-elf
endif

CXX=$(PREFIX)-g++
CXXFLAGS=-g -march=rv32im -mabi=ilp32 -I$(RVSYM_ROOT)/include

TOP=Core

VERILATOR_DIR=.verilator

VERILATOR=$(VERILATOR_ROOT)/bin/verilator

all: sim.elf

.PHONY: all

Core.sv:
	knit -C riscinator generated/Core.sv
	mv riscinator/generated/Core.sv .

$(VERILATOR_DIR): Core.sv
	$(VERILATOR) -fno-table --public -O0 -cc -sv --Mdir $(VERILATOR_DIR) $^

$(VERILATOR_DIR)/verilated.o: $(VERILATOR_DIR)
	$(MAKE) -C $(VERILATOR_DIR) -f V$(TOP).mk CXX=$(CXX) CXXFLAGS='$(CXXFLAGS)' verilated.o

$(VERILATOR_DIR)/V$(TOP)__ALL.a: $(VERILATOR_DIR)
	$(MAKE) -C $(VERILATOR_DIR) -f V$(TOP).mk CXX=$(CXX) CXXFLAGS='$(CXXFLAGS)' V$(TOP)__ALL.a

sim.o: $(VERILATOR_DIR) sim.cc
	$(MAKE) -C $(VERILATOR_DIR) -f V$(TOP).mk CXX=$(CXX) CXXFLAGS='$(CXXFLAGS) -I$(RVSYM_ROOT)/include' ../sim.o

sim.elf: sim.o $(VERILATOR_DIR)/verilated.o $(VERILATOR_DIR)/V$(TOP)__ALL.a
	$(CXX) $(CXXFLAGS) $^ -o $@

sim.x86: $(SRC) sim.cc
	$(VERILATOR) -cc -sv --Mdir $(VERILATOR_DIR)-x86 --exe --build $^ -o ../$@

clean:
	rm -rf $(VERILATOR_DIR) $(VERILATOR_DIR)-x86
	rm -f *.o *.d *.elf
