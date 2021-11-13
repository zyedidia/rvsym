#include "rvsym.h"

int main() {
    int n;
    rvsym_mark_bytes(&n, sizeof(n), "n");

    rvsym_assume(n < 10);

    int x = 0;
    int y = 0;
    while (n > 0) {
        if (n % 2 == 0) {
            x = x + n;
            y = y + 1;
        }
        else {
            y = y + n;
            x = x + 1;
        }
        n = n - 1;
    }
    if (x > 0) {
        rvsym_exit();
    }
    rvsym_quiet_exit();
    return 0;
}
