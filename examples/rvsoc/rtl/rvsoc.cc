#include "rvsoc.hh"
#include "rvsym.h"

int main() {
    cxxrtl_design::p_soc__top soc;

    uint8_t instr[4];
    // 02a00093
    // instr[0] = 0x93;
    // instr[1] = 0x00;
    instr[2] = 0xa0;
    instr[3] = 0x02;
    rvsym_mark_bytes(&instr[0], 2, "instr");
    soc.memory_p_ram__unit_2e_mem[0].set<uint32_t>(*((uint32_t*) instr));

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

    uint32_t x1 = soc.memory_p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_regs[1].get<uint32_t>();

    if (x1 != 0) {
        rvsym_fail();
    }
    rvsym_exit();
    return 0;
}
