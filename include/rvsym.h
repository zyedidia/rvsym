#pragma once

#include <stdint.h>

#include "cmds.h"

static inline void symcall_0(int symno) {
    register uintptr_t a0 asm("a0") = symno;
    asm volatile("ecall" : "+r"(a0) : : "memory");
}
static inline void symcall_1(int symno, uintptr_t arg0) {
    register uintptr_t a0 asm("a0") = symno;
    register uintptr_t a1 asm("a1") = arg0;
    asm volatile("ecall" : "+r"(a0), "+r"(a1) : : "memory");
}
static inline void symcall_3(int symno, uintptr_t arg0, uintptr_t arg1, uintptr_t arg2) {
    register uintptr_t a0 asm("a0") = symno;
    register uintptr_t a1 asm("a1") = arg0;
    register uintptr_t a2 asm("a2") = arg1;
    register uintptr_t a3 asm("a3") = arg2;
    asm volatile("ecall" : "+r"(a0), "+r"(a1), "+r"(a2), "+r"(a3) : : "memory");
}

static inline void rvsym_mark_bytes(volatile void* p, uint32_t nbytes, const char* name) {
    symcall_3(RVSYM_MARK_NBYTES, (uintptr_t) p, nbytes, (uintptr_t) name);
}

static inline void rvsym_quiet_exit() {
    symcall_0(RVSYM_QUIET_EXIT);
}

static inline void rvsym_exit() {
    symcall_0(RVSYM_EXIT);
}

static inline void rvsym_mark_regs_symbolic() {
    symcall_0(RVSYM_SYMBOLIC_REGS);
}

static inline void rvsym_mark_reg_symbolic(int reg) {
    symcall_1(RVSYM_SYMBOLIC_REG, reg);
}

static inline void rvsym_fail() {
    symcall_0(RVSYM_FAIL);
}

static inline void rvsym_dump() {
    symcall_0(RVSYM_DUMP);
}

static inline void rvsym_mark_array(volatile void* p, uint32_t nbytes, const char* name) {
    symcall_3(RVSYM_MARK_ARRAY, (uintptr_t) p, nbytes, (uintptr_t) name);
}

#define rvsym_assert(x)         \
    do {                        \
        if (!(x))               \
            rvsym_fail();       \
    } while (0)

#define rvsym_assume(x)         \
    do {                        \
        if (!(x))               \
            rvsym_quiet_exit(); \
    } while (0)
