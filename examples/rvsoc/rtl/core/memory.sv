module memory
    (
        input logic clk, rst,

        input logic [2:0] ext_ctrl,
        input logic [1:0] wb_src,
        input logic mem_write,

        input logic [31:0] alu_result, store_data,

        input logic mem_read_res_valid,
        input logic [31:0] mem_read_res_data,
        output logic mem_read_req_valid,
        output logic [31:0] mem_read_req_addr,

        input logic mem_write_res_valid,
        output logic mem_write_req_valid,
        output logic [31:0] mem_write_req_addr,
        output logic [31:0] mem_write_req_data,
        output logic [3:0] mem_write_req_mask,

        input logic [2:0] state_reg,
        output logic [2:0] mem_next,

        output logic [31:0] mem_rdata
    );

    `include "signals.svh"
    `include "ctrl_states.svh"

    // memory address is always going to be the result of an ALU computation.
    assign mem_read_req_addr = alu_result;
    assign mem_write_req_addr = alu_result;
    // write data is the data we want to store.
    assign mem_write_req_data = store_data;

    always_comb begin
        case (ext_ctrl)
            // set the memory write mask appropriately
            EXT_BYTE: mem_write_req_mask = MEM_MASK_BYTE;
            EXT_HALF: mem_write_req_mask = MEM_MASK_HALF;
            EXT_WORD: mem_write_req_mask = MEM_MASK_WORD;
            EXT_BYTEU: mem_write_req_mask = MEM_MASK_BYTEU; // shouldn't happen
            EXT_HALFU: mem_write_req_mask = MEM_MASK_HALFU; // shouldn't happen
            default: mem_write_req_mask = MEM_MASK_WORD; // shouldn't happen
        endcase
    end

    logic read_req_outstanding, write_req_outstanding;
    // we are making a read request if we are in the memory stage and need
    // a memory result for writeback.
    assign read_req_outstanding = (wb_src == WB_SRC_MEM) && state_reg == CTRL_STATE_MEM;
    // we are making a write request if this instruction writes to memory and
    // we are in the memory stage.
    assign write_req_outstanding = mem_write && state_reg == CTRL_STATE_MEM;

    assign mem_read_req_valid = read_req_outstanding;
    assign mem_write_req_valid = write_req_outstanding;

    always_comb begin
        if (read_req_outstanding && mem_read_res_valid)
            // got a result, can transition to writeback
            mem_next = CTRL_STATE_WB;
        else if (write_req_outstanding && mem_write_res_valid)
            // finished the write, we can fetch the next instruction
            mem_next = CTRL_STATE_FETCH;
        else
            // blocked waiting for memory result
            mem_next = CTRL_STATE_MEM;
    end

    logic [31:0] mem_rdata_reg, mem_rdata_next;

    // when we get data from reading memory, it may have to be sign/zero
    // extended if the instruction reads less than a full word (lh/lb/lhu/lbu).
    extend mem_read_extend_unit (
        .in(mem_read_res_data),
        .ctrl(ext_ctrl),
        .out(mem_rdata_next)
    );

    // register for the memory read result since we need it for the next stage.
    always_ff @(posedge clk, posedge rst) begin
        if (rst) begin
            mem_rdata_reg <= 32'b0;
        end else if (state_reg == CTRL_STATE_MEM && mem_read_res_valid) begin
            mem_rdata_reg <= mem_rdata_next;
        end
    end

    assign mem_rdata = mem_rdata_reg;
endmodule
