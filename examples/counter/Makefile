include $(RVSYM_ROOT)/lib/rules.mk

YOSYS_INCLUDE=$(shell yosys-config --datdir)/include
LDLIBS=-lstdc++

all: counter.bin

counter.o: counter.cc counter.cxx
counter.bin: counter.o

%.cxx: %.sv
	yosys -q -p 'read_verilog -sv -noautowire $<; write_cxxrtl -g0 -O6 $@'

%.o: %.cc %.cxx
	$(CXX) $(CXXFLAGS) -I$(YOSYS_INCLUDE) -c $< -o $@

clean:
	rm -f *.o *.bin *.elf *.list *.hh
