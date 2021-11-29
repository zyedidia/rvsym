#include <stdio.h>
#include <stdint.h>

#include <backends/cxxrtl/cxxrtl.h>

#include "rvsym.h"
#include "counter.cxx"

int main() {
    cxxrtl_design::p_counter counter;

    uint32_t state;
    rvsym_mark_bytes(&state, 4, "state");

    counter.p_q__reg.set<uint32_t>(state);
    counter.step();
    counter.step();

    rvsym_assume(counter.p_q.get<uint32_t>() <= 10000);

    counter.p_clk.set<bool>(false);
    counter.step();

    counter.p_clk.set<bool>(true);
    counter.step();

    counter.step();

    if (counter.p_q.get<uint32_t>() > 10000) {
        rvsym_fail();
    }

    return 0;
}
