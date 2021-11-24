#include "rvsym-asm.h"

.section ".text.boot"

.globl _start
_start:
	li sp, 0x1000000
	jal main
	rvsym_quiet_exit()
