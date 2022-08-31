#include <stdlib.h>
#include <stdint.h>

#include "Vcounter_counter.h"
#include "Vcounter.h"
#include "verilated.h"

int main(int argc, char **argv) {
    Vcounter* dut = new Vcounter;

    dut->counter->q_reg = 100;
    dut->rst = false;
    dut->eval();
    dut->clk = false;
    dut->eval();
    dut->clk = true;
    dut->eval();
    return 0;
}
