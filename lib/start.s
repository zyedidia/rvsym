.section ".text.boot"

.globl _start
_start:
	la sp, 0x1000000
	j _cstart

