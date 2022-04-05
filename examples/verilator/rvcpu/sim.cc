#include <stdlib.h>
#include <stdint.h>
#include "VSoc_Ram.h"
#include "VSoc_Soc.h"
#include "VSoc_Core.h"
#include "VSoc_Datapath.h"
#include "VSoc_RegFile.h"
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

    uint32_t insn;
    rvsym_mark_bytes(&insn, sizeof(insn), "insn");
    rvsym_assume((insn & 0b1111111) == 0b0010011);
    rvsym_assume(((insn >> 7) & 0x1f) == 1);
    // rvsym_assume(insn == 0x93);

    dut->Soc->bus_io_dev_0_ram->mem[0] = insn;
    rvsym_mark_array(&dut->Soc->core->dpath->rf->regs, sizeof(dut->Soc->core->dpath->rf->regs));
    rvsym_mark_array(&dut->Soc->bus_io_dev_0_ram->mem, sizeof(dut->Soc->bus_io_dev_0_ram->mem));

    for (int i = 0; i < 4; i++) {
        cycle(dut);
    }

    // printf("%x\n", dut->Soc->core->dpath->rf->regs[1]);
    if (dut->Soc->core->dpath->rf->regs[1] == 42) {
        rvsym_exit();
    }

    // printf("%x\n", dut->Soc->core->dpath->rf->regs[1]);
    // printf("%x\n", dut->Soc->bus_io_dev_0_ram->mem[25]);

    return 0;
}
