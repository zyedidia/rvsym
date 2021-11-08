#ifndef __SYMX_INTERFACE_H___
#define __SYMX_INTERFACE_H___

#include "rpi.h"
#include "symx-syscalls.h"

uint32_t symx_call_sys(unsigned sysno, ...);

#define SYM_STRING(x) #x

// we really want to know the line number etc.  
// for the moment we don't worry about it.
static inline void symx_mark_nbytes(void *p, uint32_t nbytes) 
    { symx_call_sys(SYMX_MARK_NBYTES, p, nbytes); }

static inline void symx_quiet_exit(int exitcode)
    { symx_call_sys(SYMX_QUIET_EXIT, exitcode); }

#define symx_assume(x) do { if(!(x)) symx_quiet_exit(0); } while(0)

static inline uint32_t symx_rrot32(uint32_t x, uint32_t n) 
    { return symx_call_sys(SYMX_RROT, x,n); }


static inline uint32_t symx_prove(uint32_t x) 
    { return symx_call_sys(SYMX_PROVE, x); }

#define symx_print_expr(x) symx_call_sys(SYMX_PRINT_EXPR, (uint32_t)x)

#if 0
// macro hack to we pass the name of the variable.
static inline void symx_mark_uint32_raw(const char *name, uint32_t *x) 
    { symx_call_sys(SYMX_MARK_UINT32, name, x); }
#define symx_mark_uint32(_x) symx_mark_uint32_raw(SYM_STRING(_x), _x)


static inline uint32_t symx_solve_uint32(uint32_t *x) 
    { return symx_call_sys(SYMX_SOLVE_UINT32, x); }
#endif

#if 0

static inline void symx_mark_uint16(uint16_t *x) 
    { symx_call_sys(SYMX_MARK_UINT16, x); }
static inline void symx_mark_uint8(uint8_t *x) 
    { symx_call_sys(SYMX_MARK_UINT8, x); }

static inline void symx_follow_false(uint8_t *x) 
    { symx_call_sys(SYMX_MARK_UINT8, x); }
static inline void symx_follow_true(uint8_t *x) 
    { symx_call_sys(SYMX_FOLLOW_TRUE, x); }


#define SYMX_MARK_UINT32    1
#define SYMX_MARK_UINT16    2
#define SYMX_MARK_UINT8     3

#define SYMX_SOLVE_U32      4
#define SYMX_SOLVE_U16      5
#define SYMX_SOLVE_U8       6
#endif

#endif
