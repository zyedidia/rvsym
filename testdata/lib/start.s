.section ".text.boot"

.globl _start
_start:
	li sp, 0x1000
	j main
