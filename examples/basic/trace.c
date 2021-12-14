#include "rvsym.h"

int a;

void func1() {
    rvsym_elapse_us(10);
    a = 42;
}

void func2() {
    rvsym_elapse_us(10);
    a = 0x2a;
}

int main() {
    rvsym_mark_output(&a, sizeof(a), "a");

    func1();
    int s0 = rvsym_snapshot();
    rvsym_trace_reset();

    func2();
    int s1 = rvsym_snapshot();

    int r = rvsym_snapshot_eq(s0, s1);
    rvsym_assert(r == 1);

    rvsym_quiet_exit();

    return 0;
}
