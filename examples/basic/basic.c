#include "rvsym.h"

int main() {
    int x;
    rvsym_mark_bytes(&x, sizeof(x), "x");

    if (x == 42) {
        rvsym_fail();
    }

    return 0;
}
