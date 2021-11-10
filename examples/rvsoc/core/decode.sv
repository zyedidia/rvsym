module decode
    (
        input logic clk, rst,
        input logic nop,

        // register file inputs
        input logic [4:0] rs1, rs2, rd,
        input logic wr_en,
        input logic [31:0] w_data,
        output logic [31:0] rd1, rd2,

        // state
        input logic [2:0] state_reg,
        output logic [2:0] decode_next
    );

    `include "ctrl_states.svh"
    `include "signals.svh"

    // register file. We the rs1 and rs2 fields from the instruction and their
    // outputs will be available at rd1 and rd2. We also write to rd, if
    // writing is enabled for this instruction.
    reg_file reg_file_unit (
        .clk,
        .r0_addr(rs1), .r1_addr(rs2), .w_addr(rd),
        .wr_en, .w_data,
        .r0_data(rd1), .r1_data(rd2)
    );

    // transition to execute from decode, except if we decode to a nop, we can
    // immediately fetch the next instruction. By default we always increment
    // the PC by 4, so the next PC for a nop is ready to go.
    assign decode_next = nop ? CTRL_STATE_FETCH : CTRL_STATE_EXE;
endmodule
