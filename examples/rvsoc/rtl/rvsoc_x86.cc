#include <stdio.h>
#include "rvsoc.hh"

int main() {
    cxxrtl_design::p_soc__top soc;

    uint32_t instr = 0x60e7;
    soc.memory_p_ram__unit_2e_mem[0].set<uint32_t>(instr);
    // soc.memory_p_ram__unit_2e_mem[1].set<uint32_t>(0x00000063);

    soc.p_clk.set<bool>(false);
    soc.step();
    soc.p_clk.set<bool>(true);
    soc.step();

    soc.p_clk.set<bool>(false);
    soc.step();
    soc.p_clk.set<bool>(true);
    soc.step();

    soc.p_clk.set<bool>(false);
    soc.step();
    soc.p_clk.set<bool>(true);
    soc.step();

    soc.p_clk.set<bool>(false);
    soc.step();
    soc.p_clk.set<bool>(true);
    soc.step();

    soc.p_clk.set<bool>(false);
    soc.step();
    soc.p_clk.set<bool>(true);
    soc.step();

    soc.step();

    // uint32_t x1 = soc.p_cpu__unit_2e_instr__reg.get<uint32_t>();
    uint32_t x1 = soc.memory_p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_regs[1].get<uint32_t>();
    printf("x1: %x\n", x1);


    return 0;
}

