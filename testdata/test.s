addi x5, x0, 5
beq x5, x1, L1
beq x0, x0, L2
L1:
addi x5, x5, 5
L2:
beq x0, x0, L2
