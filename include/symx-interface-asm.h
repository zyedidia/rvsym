#ifndef __SYMX_INTERFACE_ASM_H__
#define __SYMX_INTERFACE_ASM_H__

#include "symx-syscalls.h"

#define SYMX_SYS(v)         li a0, v; ecall

#define symx_mark_regs_symbolic()    SYMX_SYS(SYMX_SYMBOLIC_REGS)
#define symx_fail()                  SYMX_SYS(SYMX_FAIL)
#define symx_mark_reg_symbolic(r)    li a1, r; SYMX_SYS(SYMX_SYMBOLIC_REG)
#define symx_exit()                  SYMX_SYS(SYMX_EXIT)
#define symx_quiet_exit(x)                  li a1, x; SYMX_SYS(SYMX_QUIET_EXIT)

#define symx_setup()                  SYMX_SYS(SYMX_SETUP)

#define symx_choose(x)               li a1, x; SYMX_SYS(SYMX_CHOOSE)
#endif
