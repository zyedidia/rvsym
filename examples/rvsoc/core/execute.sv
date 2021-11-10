module execute
    (
        input logic clk, rst,
        input logic a_src, b_src,
        input logic negate_b,
        input logic shift_arith,
        input logic [2:0] alu_op,
        input logic [2:0] ext_ctrl,

        input logic [31:0] pc_reg, imm,
        input logic [31:0] rd1, rd2,

        output logic [31:0] alu_result,
        output logic [31:0] store_data,

        input logic mem_write,
        input logic [1:0] wb_src,
        output logic [2:0] exe_next
    );

    `include "ctrl_states.svh"
    `include "signals.svh"

    logic [31:0] alu_a, alu_b;

    // some instructions perform ALU operations on PC, regs[rs1] (rd1),
    // regs[rs2] (rd2), or the immediate, so we select the appropriate ALU
    // operands here.
    assign alu_a = a_src == A_SRC_PC ? pc_reg : rd1;
    assign alu_b = b_src == B_SRC_IMM ? imm : (negate_b ? -rd2 : rd2);

    alu alu_unit (
        .op(alu_op),
        .a(alu_a), .b(alu_b),
        .shift_arith,
        .out(alu_result)
    );

    // also the rs2 register holds the store data for sw/sh/sb. In the case of
    // sh/sb, we need to use an extend unit to properly extract the value by
    // sign-extending the lower byte/half-word of the register value.
    extend reg_extend_unit (
        .in(rd2),
        .ctrl(ext_ctrl),
        .out(store_data)
    );

    // if we need to access memory, transition to mem state, otherwise we can
    // go straight to writeback.
    assign exe_next = (mem_write || wb_src == WB_SRC_MEM) ? CTRL_STATE_MEM : CTRL_STATE_WB;
endmodule
