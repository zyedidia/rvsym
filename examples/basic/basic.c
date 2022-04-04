#include "rvsym.h"
#include <stdio.h>

int main() {
    int x;
    rvsym_mark_bytes(&x, sizeof(x), "x");

    if (x == 42) {
        rvsym_fail();
    }

    printf("Hello world\n");

    return 0;
}
