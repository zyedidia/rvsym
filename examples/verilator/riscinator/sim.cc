#include <stdint.h>
#include <stdlib.h>
#include <vector>
#include "VCore.h"
#include "VCore___024root.h"
#include "verilated.h"

static uint32_t be2mask(uint8_t be) {
    uint32_t mask = 0;
    for (unsigned i = 0; i < 4; i++) {
        if (be & (1 << i)) {
            mask |= 0xff << i*8;
        }
    }
    return mask;
}

static void clock(VCore* core) {
    core->clock = 0;
    core->eval();
    core->clock = 1;
    core->eval();
}

static unsigned addr2idx(uint32_t addr, size_t mem_base) {
    return (addr-mem_base) / sizeof(uint32_t);
}

static bool failed = false;

static void simulate(VCore* core, uint32_t* mem, size_t len, size_t mem_base) {
    core->reset = 1;
    clock(core);
    core->reset = 0;

    uint32_t next_imem_rdata = 0;
    uint8_t next_imem_rvalid = 0;
    uint32_t next_dmem_rdata = 0;
    uint8_t next_dmem_rvalid = 0;

    unsigned ncyc = 50;
    for (unsigned i = 0; i < ncyc; i++) {
        core->eval();

        core->io_imem_rvalid = next_imem_rvalid;
        core->io_imem_rdata = next_imem_rdata;
        core->io_dmem_rvalid = next_dmem_rvalid;
        core->io_dmem_rdata = next_dmem_rdata;

        next_imem_rvalid = core->io_imem_req;
        next_dmem_rvalid = core->io_dmem_req;
        core->io_dmem_gnt = core->io_dmem_req;

        if (core->io_imem_req) {
            assert(core->io_imem_addr >= mem_base && core->io_imem_addr < mem_base + len);
            next_imem_rdata = mem[addr2idx(core->io_imem_addr, mem_base)];
        }

        if (core->io_dmem_req && core->io_dmem_we) {
            uint32_t write = core->io_dmem_wdata;
            uint32_t mask = be2mask(core->io_dmem_be);
            assert(core->io_dmem_addr >= mem_base && core->io_dmem_addr < mem_base + len);
            mem[addr2idx(core->io_dmem_addr, mem_base)] = write & mask;
        } else if (core->io_dmem_req) {
            assert(core->io_dmem_addr >= mem_base && core->io_dmem_addr < mem_base + len);
            next_dmem_rdata = mem[addr2idx(core->io_dmem_addr, mem_base)];
        }

        clock(core);
    }
}

int main(int argc, char **argv) {
    Verilated::commandArgs(argc, argv);

    VCore* core = new VCore;

    uint32_t mem[32];
    mem[0] = 0x0000006f;

    simulate(core, mem, sizeof(mem), 0x100000);

    if (failed) return 1;
    return 0;
}

