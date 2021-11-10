#include "rvsym-asm.h"

.section ".text.boot"

.globl _start
_start:
	li sp, 0x100000
	jal main
	rvsym_quiet_exit()
