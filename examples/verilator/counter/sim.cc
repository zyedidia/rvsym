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

    rvsym_assume(initial <= 10000);

    dut->counter->q_reg = initial;
    dut->rst = false;
    dut->eval();
    dut->clk = false;
    dut->eval();
    dut->clk = true;
    dut->eval();

    if (dut->counter->q > 10000) {
        rvsym_fail();
    }

    return 0;
}

