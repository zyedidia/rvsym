#include <stdlib.h>
#include <stdint.h>
#include "VSoc_Ram.h"
#include "VSoc_Soc.h"
#include "VSoc.h"
#include "verilated.h"

#include "rvsym.h"

void cycle(VSoc* dut) {
    dut->clock = false;
    dut->eval();
    dut->clock = true;
    dut->eval();
}

int main(int argc, char **argv) {
    VSoc* dut = new VSoc;

    dut->reset = true;
    cycle(dut);
    dut->reset = false;

    for (int i = 0; i < 100; i++) {
        cycle(dut);
    }

    printf("%x\n", dut->Soc->bus_io_dev_0_ram->mem[25]);

    return 0;
}
