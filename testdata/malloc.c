#include <stdlib.h>
#include <string.h>

#include "rvsym.h"

int main () {
    char* p = (char*) malloc(1000);
    char* h = (char*) malloc(1000);

    memset(p, 0, 1000);
    h[0] = 42;
    rvsym_assert(p[0] == 0);
    rvsym_assert(h[0] == 42);
    return 0;
}
