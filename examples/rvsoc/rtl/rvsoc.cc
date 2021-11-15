#include "rvsoc.hh"
#include "rvsym.h"

int main() {
    cxxrtl_design::p_soc__top soc;

    uint32_t instr;
    rvsym_mark_bytes(&instr, sizeof(instr), "instr");
    soc.memory_p_ram__unit_2e_mem[0].set<uint32_t>(instr);

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

    for (int i = 1; i < 32; i++) {
        uint32_t x1 = soc.memory_p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_regs[i].get<uint32_t>();

        if (x1 != 0) {
            rvsym_fail();
        }
    }
    return 0;
}