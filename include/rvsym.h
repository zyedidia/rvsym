#pragma once

#include <stdint.h>

#include "cmds.h"

static inline void make_syscall_0(int sysno) {
    register uintptr_t a0 asm("a0") = sysno;
    asm volatile ("ecall" : "+r" (a0) : );
}
static inline void make_syscall_1(int sysno, uintptr_t arg0) {
    register uintptr_t a0 asm("a0") = sysno;
    register uintptr_t a1 asm("a1") = arg0;
    asm volatile ("ecall" : "+r" (a0), "+r" (a1) :);

}
static inline void make_syscall_2(int sysno, uintptr_t arg0, uintptr_t arg1) {
    register uintptr_t a0 asm("a0") = sysno;
    register uintptr_t a1 asm("a1") = arg0;
    register uintptr_t a2 asm("a2") = arg1;
    asm volatile ("ecall" : "+r" (a0), "+r" (a1), "+r" (a2) :);
}

// we really want to know the line number etc.
// for the moment we don't worry about it.
static inline void rvsym_mark_nbytes(void* p, uint32_t nbytes) {
    make_syscall_2(SYMX_MARK_NBYTES, (uintptr_t) p, nbytes);
}

static inline void rvsym_quiet_exit(int exitcode) {
    make_syscall_1(SYMX_QUIET_EXIT, exitcode);
}

static inline void rvsym_fail() {
    make_syscall_0(SYMX_FAIL);
}

#define rvsym_assume(x)          \
    do {                        \
        if (!(x))               \
            rvsym_quiet_exit(0); \
    } while (0)
