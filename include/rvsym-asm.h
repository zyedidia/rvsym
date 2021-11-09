#pragma once

#include "cmds.h"

#define SYMX_SYS(v) \
    li a0, v;       \
    ecall

#define rvsym_mark_regs_symbolic() SYMX_SYS(SYMX_SYMBOLIC_REGS)

#define rvsym_fail() SYMX_SYS(SYMX_FAIL)

#define rvsym_mark_reg_symbolic(r) \
    li a1, r;                      \
    SYMX_SYS(SYMX_SYMBOLIC_REG)

#define rvsym_exit() SYMX_SYS(SYMX_EXIT)

#define rvsym_quiet_exit(x) \
    li a1, x;               \
    SYMX_SYS(SYMX_QUIET_EXIT)

#define rvsym_setup() SYMX_SYS(SYMX_SETUP)

#define rvsym_choose(x) \
    li a1, x;           \
    SYMX_SYS(SYMX_CHOOSE)
