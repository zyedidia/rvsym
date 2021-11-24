#pragma once

#include "cmds.h"

#define RVSYM_SYS(v) \
    li a0, v;        \
    ecall

#define rvsym_mark_regs_symbolic() RVSYM_SYS(RVSYM_SYMBOLIC_REGS)

#define rvsym_fail() RVSYM_SYS(RVSYM_FAIL)

#define rvsym_exit() RVSYM_SYS(RVSYM_EXIT)

#define rvsym_dump() RVSYM_SYS(RVSYM_DUMP)

#define rvsym_quiet_exit() RVSYM_SYS(RVSYM_QUIET_EXIT)

#define rvsym_choose(x) \
    li a1, x;           \
    RVSYM_SYS(RVSYM_CHOOSE)
