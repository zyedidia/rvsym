include $(RVSYM_ROOT)/lib/rules.mk

SYNTH=$(shell find . -type f -name '*.sv')

LDFLAGS += -L$(RV_ROOT)/lib/gcc/$(PREFIX)/11.1.0/rv32i/ilp32
LDLIBS=-lstdc++ -lgcc

MEM = empty.mem
SV_INC=-Isys -Icore
DEFINE=-DRAMFILE="$(MEM)"

YOSIM_SO=$(YOSIM_ROOT)/cxx/cxx.so
YOSIM_INCLUDE=$(YOSIM_ROOT)/cxx
YOSYS_INCLUDE=$(shell yosys-config --datdir)/include

all: rvsoc.bin

rvsoc.o: rvsoc.cc rvsoc.cxx
rvsoc.bin: rvsoc.o

rvsoc.cxx: $(SYNTH)
	yosys -m $(YOSIM_SO) -q -p 'read_verilog $(SV_INC) $(DEFINE) -sv -noautowire $(SYNTH); write_yosimcxx -g0 -O6 $@'

%.o: %.cc %.cxx
	$(CXX) $(CXXFLAGS) -DRVSYM -DCXXRTL_NDEBUG -DNDEBUG -I$(YOSIM_INCLUDE) -I$(YOSYS_INCLUDE) -c $< -o $@

LINT_FLAGS=-Wno-PINCONNECTEMPTY -Wno-UNUSED

lint:
	verilator $(SV_INC) '$(DEFINE)' --lint-only -Wall -sv $(LINT_FLAGS) $(SYNTH)

clean:
	rm -f *.o *.bin *.elf *.list *.cxx
