#include "symx-interface-asm.h"

    symx_mark_reg_symbolic(1)
    bne x1, x0, L1
	symx_fail()
L1: symx_exit()
