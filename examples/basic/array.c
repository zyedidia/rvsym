#include "rvsym.h"

int main() {
    int addr;
    rvsym_mark_bytes(&addr, sizeof(addr), "addr");
    rvsym_assume(addr >= 0);
    rvsym_assume(addr < 100);

    int array[100];

    rvsym_mark_array(&array, sizeof(array));
    for (int i = 0; i < 100; i++) {
        array[i] = i;
    }
    array[10] = 50;
    array[50] = 10;

    if (array[addr] == 50) {
        rvsym_exit();
    }
    rvsym_quiet_exit();
    return 0;
}
