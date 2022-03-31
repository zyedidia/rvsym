CXX=riscv64-unknown-elf-g++
CXXFLAGS=-march=rv32im -mabi=ilp32

TOP=Soc

all: sim.elf

.PHONY: all

obj_dir: Soc.v
	verilator -cc $<

obj_dir/verilated.o: obj_dir
	make -C obj_dir -f V$(TOP).mk CXX=$(CXX) CXXFLAGS='$(CXXFLAGS)' verilated.o

obj_dir/V$(TOP)__ALL.a: obj_dir
	make -C obj_dir -f V$(TOP).mk CXX=$(CXX) CXXFLAGS='$(CXXFLAGS)' V$(TOP)__ALL.a

sim.o: obj_dir sim.cc
	make -C obj_dir -f V$(TOP).mk CXX=$(CXX) CXXFLAGS='$(CXXFLAGS) -I$(RVSYM_ROOT)/include' ../sim.o

sim.elf: sim.o obj_dir/verilated.o obj_dir/V$(TOP)__ALL.a
	$(CXX) $(CXXFLAGS) $^ -o $@

clean:
	rm -rf obj_dir
	rm -f *.o *.d *.elf
