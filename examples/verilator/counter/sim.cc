#include <stdlib.h>
#include <stdint.h>

#include "Vcounter_counter.h"
#include "Vcounter.h"
#include "verilated.h"

#include "rvsym.h"

int main(int argc, char **argv) {
    Vcounter* dut = new Vcounter;

    uint32_t initial;
    rvsym_mark_bytes(&initial, 4, "initial");

    rvsym_assume(initial > 10000);

    dut->counter->q_reg = initial;
    dut->rst = false;
    dut->eval();
    dut->clk = false;
    dut->eval();
    dut->clk = true;
    dut->eval();

    // assert that we cannot reach a value above 10000 from any valid state
    rvsym_assert(dut->q > 10000);

    return 0;
}
