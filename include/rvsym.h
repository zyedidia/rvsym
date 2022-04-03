#pragma once

#include <stdint.h>

#define RVSYM_PRINT 0
#define RVSYM_FAIL 1
#define RVSYM_EXIT 2
#define RVSYM_QUIET_EXIT 3
#define RVSYM_MARK_ARRAY 4
#define RVSYM_MARK_NBYTES 5

static inline uintptr_t symcall_0(int symno) {
    register uintptr_t a0 asm("a0") = symno;
    asm volatile("ebreak" : "+r"(a0) : : "memory");
    return a0;
}
static inline uintptr_t symcall_1(int symno, uintptr_t arg0) {
    register uintptr_t a0 asm("a0") = symno;
    register uintptr_t a1 asm("a1") = arg0;
    asm volatile("ebreak" : "+r"(a0), "+r"(a1) : : "memory");
    return a0;
}
static inline uintptr_t symcall_2(int symno, uintptr_t arg0, uintptr_t arg1) {
    register uintptr_t a0 asm("a0") = symno;
    register uintptr_t a1 asm("a1") = arg0;
    register uintptr_t a2 asm("a2") = arg1;
    asm volatile("ebreak" : "+r"(a0), "+r"(a1), "+r"(a2) : : "memory");
    return a0;
}
static inline uintptr_t symcall_3(int symno, uintptr_t arg0, uintptr_t arg1, uintptr_t arg2) {
    register uintptr_t a0 asm("a0") = symno;
    register uintptr_t a1 asm("a1") = arg0;
    register uintptr_t a2 asm("a2") = arg1;
    register uintptr_t a3 asm("a3") = arg2;
    asm volatile("ebreak" : "+r"(a0), "+r"(a1), "+r"(a2), "+r"(a3) : : "memory");
    return a0;
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

static inline void rvsym_print(int val) {
    symcall_1(RVSYM_PRINT, val);
}

static inline void rvsym_fail() {
    symcall_0(RVSYM_FAIL);
}

static inline void rvsym_mark_array(volatile void* p, uint32_t nbytes) {
    symcall_2(RVSYM_MARK_ARRAY, (uintptr_t) p, nbytes);
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

