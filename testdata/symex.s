#include "rvsym-asm.h"

    rvsym_mark_reg_symbolic(1)
    bne x1, x0, L1
	rvsym_fail()
L1: rvsym_exit()
