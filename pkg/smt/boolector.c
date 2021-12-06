#include "boolector.h"
#include <time.h>
#include <stdio.h>

void boolector_multi_assert(Btor* btor, BoolectorNode** bools, int nbools) {
    for (int i = 0; i < nbools; i++) {
        if (!bools[i]) {
            boolector_push(btor, 1);
        } else {
            boolector_assert(btor, bools[i]);
        }
    }
}
