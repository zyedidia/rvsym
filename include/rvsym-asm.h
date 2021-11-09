#pragma once

#include "cmds.h"

#define RVSYM_SYS(v) \
    li a0, v;        \
    ecall

#define rvsym_mark_regs_symbolic() RVSYM_SYS(RVSYM_SYMBOLIC_REGS)

#define rvsym_fail() RVSYM_SYS(RVSYM_FAIL)

#define rvsym_mark_reg_symbolic(r) \
    li a1, r;                      \
    RVSYM_SYS(RVSYM_SYMBOLIC_REG)

#define rvsym_exit() RVSYM_SYS(RVSYM_EXIT)

#define rvsym_quiet_exit() RVSYM_SYS(RVSYM_QUIET_EXIT)

#define rvsym_setup() RVSYM_SYS(RVSYM_SETUP)

#define rvsym_choose(x) \
    li a1, x;           \
    RVSYM_SYS(RVSYM_CHOOSE)
