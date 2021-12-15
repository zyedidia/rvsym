#include "rvsym.h"

#include <backends/cxxrtl/cxxrtl.h>
#include "counter.cxx"

void add_n(uint32_t initial, uint32_t* output, int n) {
    for (int i = 0; i < n; i++) {
        initial++;
        if (initial > 10000) {
            initial = 0;
        }
        rvsym_elapse_us(2);
    }
    *output = initial;
}

void add_n_counter(uint32_t initial, uint32_t* output, int n) {
    cxxrtl_design::p_counter counter;

    counter.p_q__reg.set<uint32_t>(initial);

    for (int i = 0; i < n; i++) {
        counter.p_clk.set<bool>(false);
        counter.step();

        counter.p_clk.set<bool>(true);
        counter.step();
        rvsym_elapse_us(2);
    }

    *output = counter.p_q.get<uint32_t>();
}

int main() {
    uint32_t initial;
    rvsym_mark_bytes(&initial, 4, "initial");
    rvsym_assume(initial <= 10000);

    uint32_t output;
    rvsym_mark_output(&output, 4, "output");

    add_n(initial, &output, 1);
    int s1 = rvsym_snapshot();

    rvsym_trace_reset();

    add_n_counter(initial, &output, 1);
    int s2 = rvsym_snapshot();

    rvsym_assert(rvsym_snapshot_eq(s1, s2));

    rvsym_quiet_exit();
    return 0;
}
