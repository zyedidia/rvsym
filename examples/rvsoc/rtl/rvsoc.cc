#include "rvsoc.cxx"
#include "rvsym.h"

int main() {
    cxxrtl_design::p_soc__top soc;

    // uint8_t instr[4];
    // 02a00093
    // instr[0] = 0x93;
    // instr[1] = 0x00;
    // instr[2] = 0xa0;
    // instr[3] = 0x02;
    // rvsym_mark_bytes(&instr[2], 2, "instr");
    // soc.memory_p_ram__unit_2e_mem[0].set<uint32_t>(*((uint32_t*) instr));

    uint32_t instr;
    rvsym_mark_bytes(&instr, 4, "instr");
    soc.memory_p_ram__unit_2e_mem[0].set<uint32_t>(instr);

    // rvsym_assume(((instr >> 7) & 0x1f) == 0);
    // rvsym_assume((instr & 0b1111111) == 0b0010011);
    // rvsym_assume(instr == 0x02a00093);

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

    rvsym_assert(x1 != 42);

    rvsym_quiet_exit();
    return 0;
}
