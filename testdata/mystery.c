#include "rvsym.h"

int mystery(int n[6]) {
    for (int i = 1; i < 6; i++) {
        if (n[i] != n[0] + n[i-1] + i) {
            return 0;
        }
    }
    return 1;
}

int main() {
    int n[6];
    rvsym_mark_bytes(&n, sizeof(n), "n");

    rvsym_assume(n[0] == 0);

    int r = mystery(n);
    if (r == 1) {
        rvsym_fail();
    }
    rvsym_quiet_exit();
    return 0;
}
