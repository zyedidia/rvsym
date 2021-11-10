// Module for the fetch stage. We issue an instruction memory request with the
// current PC, wait until a result comes back, and latch it into a register to
// store the instruction data. We also store the current PC, and latch a new PC
// in if the CPU is about to transition to the fetch control state.
module fetch
    (
        input logic clk, rst,
        // current control state
        input logic [2:0] state_reg, state_next,
        output logic [2:0] fetch_next,

        // instruction mem signals needed for fetch stage
        input logic instr_read_res_valid,
        input logic [31:0] instr_read_res_data,
        output logic instr_read_req_valid,
        output logic [31:0] instr_read_req_addr,

        // need to know the next PC (to latch it right before fetch)
        // and we output the current PC and instruction.
        input logic [31:0] pc_next,
        output logic [31:0] pc, instr
    );

    `include "ctrl_states.svh"

    logic [31:0] pc_reg = BOOT_ADDR;
    assign pc = pc_reg;

    // PC register
    always_ff @(posedge clk, posedge rst) begin
        if (rst) begin
            pc_reg <= BOOT_ADDR;
        end else if (state_next == CTRL_STATE_FETCH && state_reg != CTRL_STATE_FETCH) begin
            // if we are about to transition to fetch, then we load in a new PC.
            pc_reg <= pc_next;
        end
    end

    // read the instruction at PC when in the fetch state.
    assign instr_read_req_addr = pc_reg;
    assign instr_read_req_valid = state_reg == CTRL_STATE_FETCH;

    // transition to decode state if we get a valid result from memory,
    // otherwise block in fetch while waiting.
    assign fetch_next = instr_read_res_valid ? CTRL_STATE_DECODE : CTRL_STATE_FETCH;

    always_ff @(posedge clk, posedge rst) begin
        // if we are about to transition from fetch to decode, then we know the
        // instruction is ready from memory and we can latch it into the
        // instruction register.
        if (rst) begin
            instr <= 32'b0;
        end else if (state_reg == CTRL_STATE_FETCH && fetch_next == CTRL_STATE_DECODE) begin
            instr <= instr_read_res_data;
        end
    end
endmodule
