#include "rvsym.h"

int main() {
    int a;
    int b;

    const int sz = 16;
    int arr[sz];

    rvsym_mark_bytes(&a, sizeof(a), "a");
    rvsym_mark_bytes(&b, sizeof(b), "b");

    rvsym_assume(a > 8);
    rvsym_assume(b < 8);

    rvsym_assume(a < sz);
    rvsym_assume(b < sz);
    rvsym_assume(b >= 0);
    rvsym_assume(a >= 0);

    arr[a] = 0;
    arr[b] = 0;

    return arr[0];
}
