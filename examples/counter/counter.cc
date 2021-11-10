#include <stdio.h>
#include <stdint.h>

#include <backends/cxxrtl/cxxrtl.h>

#include "rvsym.h"
#include "counter.hh"

int main() {
    cxxrtl_design::p_counter counter;

    uint32_t count;
    rvsym_mark_bytes(&count, 4);

    uint32_t result;
    rvsym_mark_bytes(&result, 4);

    rvsym_assume(count <= 10000);
    rvsym_assume(result > 10000);

    counter.p_q__reg.set<uint32_t>(count);

    counter.p_clk.set<bool>(false);
    counter.step();

    counter.p_clk.set<bool>(true);
    counter.step();

    counter.step();

    uint32_t next_count = counter.p_q.get<uint32_t>();

    if (next_count == result) {
        rvsym_fail();
    }

    return next_count;
}
