cons ?= lpf/orangecrab.lpf

synth=$(shell find . -type f -name '*.sv')
report=util.json

sim_top ?= $(top)
sim=$(sim_top)_sim.cc
tb=$(wildcard *tb.cc)

LINT_FLAGS=-Wno-PINCONNECTEMPTY -Wno-UNUSED
CXX_FLAGS=-g -O2 -std=c++14

MEM ?= mem/blinkrgb.mem
INC ?= -I.
DEFINE=-DRAMFILE="$(MEM)"

all: synth tb

synth: $(top).dfu

tb: $(sim_top)_tb

$(top).json: $(synth)
	yosys -q -p 'read_verilog $(INC) $(DEFINE) -noautowire -sv $(synth); synth_ecp5 -top $(top) -json $@'

$(top)_out.config: $(cons) $(top).json
	nextpnr-ecp5 -q --lpf-allow-unconstrained --report $(report) --25k --freq 48 --lpf $< --package CSFBGA285 --textcfg $@ --json $(top).json

$(top).bit: $(top)_out.config
	ecppack --compress --freq 38.8 --input $< --bit $@

$(top).dfu: $(top).bit
	cp $< $@
	dfu-suffix -v 1209 -p 5af0 -a $@

$(sim): $(synth)
	yosys -q -p 'read_verilog $(INC) -DSIM $(DEFINE) -noautowire -sv $(synth); hierarchy -top $(sim_top); write_cxxrtl -nohierarchy -O4 $(sim)'

$(sim_top)_tb: $(sim) $(tb)
	$(CXX) $(CXX_FLAGS) -I $(shell yosys-config --datdir)/include $(tb) -o $@

test: $(sim_top)_tb
	@./$<

waveform: test
	@gtkwave waveform.vcd >/dev/null 2>&1 &

lint:
	verilator $(INC) '$(DEFINE)' --lint-only -Wall -sv $(LINT_FLAGS) $(synth)

report: $(top).asc $(report)
	@cat $(report) | jq '.utilization'

prog: $(top).dfu
	sudo dfu-util -D $<

clean:
	rm -f $(top).bit $(top).dfu $(top)_out.config $(top).json $(top)_tb util.json *_sim.cc *.vcd

.PHONY: clean lint prog test report waveform all synth tb

