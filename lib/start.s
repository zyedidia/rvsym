.section ".text.boot"

.globl _start
_start:
	la sp, 0x80000000
	j _cstart

