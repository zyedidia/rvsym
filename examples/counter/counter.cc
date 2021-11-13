#include <stdio.h>
#include <stdint.h>

#include <backends/cxxrtl/cxxrtl.h>

#include "rvsym.h"
#include "counter.hh"

uint32_t eval_one_cycle(cxxrtl_design::p_counter& counter, uint32_t current) {
    counter.p_q__reg.set<uint32_t>(current);

    counter.p_clk.set<bool>(false);
    counter.step();

    counter.p_clk.set<bool>(true);
    counter.step();

    counter.step();

    return counter.p_q.get<uint32_t>();
}

int main() {
    cxxrtl_design::p_counter counter;

    uint32_t initial;
    rvsym_mark_bytes(&initial, 4, "initial");

    uint32_t result;
    rvsym_mark_bytes(&result, 4, "result");

    rvsym_assume(initial > 10000);
    rvsym_assume(result <= 10000);

    uint32_t next_count = eval_one_cycle(counter, initial);

    if (next_count == result) {
        rvsym_fail();
    }

    rvsym_quiet_exit();
    return 0;
}
