RVSYM_ROOT ?= $(shell git rev-parse --show-toplevel)

include $(RVSYM_ROOT)/lib/rules.mk

YOSYS_INCLUDE=$(shell yosys-config --datdir)/include
LDLIBS=-lstdc++

all: counter.bin

counter.o: counter.cc counter.cxx
counter.bin: counter.o

%.cxx: %.sv
	yosys -q -p 'read_verilog -sv -noautowire $<; hierarchy -auto-top; write_cxxrtl -nohierarchy -g0 -O3 $@'

%.o: %.cc counter.cxx
	$(CXX) $(CXXFLAGS) -I$(YOSYS_INCLUDE) -c $< -o $@

clean:
	rm -f *.o *.bin *.elf *.list *.hh
