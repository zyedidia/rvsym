#pragma once

#include <stdint.h>

#define RVSYM_ELAPSE_NS 0
#define RVSYM_FAIL 1
#define RVSYM_PRINT 2
#define RVSYM_EXIT 3
#define RVSYM_QUIET_EXIT 4
#define RVSYM_MARK_NBYTES 5
#define RVSYM_MARK_OUTPUT 6
#define RVSYM_MARK_ARRAY 7

static inline void symcall_0(int symno) {
    register uintptr_t a0 asm("a0") = symno;
    asm volatile("ecall" : "+r"(a0) : : "memory");
}
static inline void symcall_1(int symno, uintptr_t arg0) {
    register uintptr_t a0 asm("a0") = symno;
    register uintptr_t a1 asm("a1") = arg0;
    asm volatile("ecall" : "+r"(a0), "+r"(a1) : : "memory");
}
static inline void symcall_2(int symno, uintptr_t arg0, uintptr_t arg1) {
    register uintptr_t a0 asm("a0") = symno;
    register uintptr_t a1 asm("a1") = arg0;
    register uintptr_t a2 asm("a2") = arg1;
    asm volatile("ecall" : "+r"(a0), "+r"(a1), "+r"(a2) : : "memory");
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

static inline void rvsym_mark_output(volatile void* p, uint32_t nbytes, const char* name) {
    symcall_3(RVSYM_MARK_OUTPUT, (uintptr_t) p, nbytes, (uintptr_t) name);
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

static inline void rvsym_elapse_us(uint32_t val) {
    symcall_1(RVSYM_ELAPSE_NS, (uintptr_t) val * 1000);
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

