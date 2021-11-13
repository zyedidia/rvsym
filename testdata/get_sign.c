#include "rvsym.h"

int get_sign(int x) {
    if (x == 0)
        return 0;

    if (x < 0)
        return -1;
    else
        return 1;
}

int main() {
    int a;

    rvsym_mark_bytes(&a, sizeof(a), "a");

    return get_sign(a);
}
