include $(RVSYM_ROOT)/lib/rules.mk

SYNTH=$(shell find . -type f -name '*.v')

LDFLAGS += -L$(RV_ROOT)/lib/gcc/$(PREFIX)/11.1.0/rv32i/ilp32
LDLIBS=-lstdc++ -lgcc

SV_INC=-Irtl

YOSIM_SO=$(YOSIM_ROOT)/cxx/cxx.so
YOSIM_INCLUDE=$(YOSIM_ROOT)/cxx
YOSYS_INCLUDE=$(shell yosys-config --datdir)/include

TARGET=or1200
TOP=or1200_cpu

all: $(TARGET).bin

$(TARGET).o: $(TARGET).cc $(TARGET).cxx
$(TARGET).bin: $(TARGET).o

$(TARGET).cxx: $(SYNTH)
	yosys -m $(YOSIM_SO) -q -p 'read_verilog $(SV_INC) -sv -noautowire $(SYNTH); hierarchy -top $(TOP); write_yosimcxx -nohierarchy -g0 -O6 $@'

%.o: %.cc %.cxx
	$(CXX) $(CXXFLAGS) -DRVSYM -DCXXRTL_NDEBUG -DNDEBUG -I$(YOSIM_INCLUDE) -I$(YOSYS_INCLUDE) -c $< -o $@

LINT_FLAGS=-Wno-PINCONNECTEMPTY -Wno-UNUSED

lint:
	verilator $(SV_INC) '$(DEFINE)' --lint-only -Wall -sv $(LINT_FLAGS) $(SYNTH)

clean:
	rm -f *.o *.bin *.elf *.list *.cxx

