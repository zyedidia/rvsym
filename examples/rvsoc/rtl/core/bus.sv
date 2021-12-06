module bus
    (
        output logic instr_read_res_valid,
        output logic [31:0] instr_read_res_data,
        input logic instr_read_req_valid,
        input logic [31:0] instr_read_req_addr,

        output logic mem_read_res_valid,
        output logic [31:0] mem_read_res_data,
        input logic mem_read_req_valid,
        input logic [31:0] mem_read_req_addr,

        output logic mem_write_res_valid,
        input logic mem_write_req_valid,
        input logic [31:0] mem_write_req_addr,
        input logic [31:0] mem_write_req_data,
        input logic [3:0] mem_write_req_mask,

        input logic i_bus_rd_valid,
        input logic i_bus_wr_valid,
        input logic [31:0] i_bus_data,

        output logic o_bus_rd,
        output logic [31:0] o_bus_addr,
        output logic o_bus_wr,
        output logic [3:0] o_bus_wrmask,
        output logic [31:0] o_bus_data
    );

    always_comb begin
        o_bus_rd = 1'b0;
        o_bus_addr = 32'b0;
        o_bus_wr = 1'b0;
        o_bus_wrmask = 4'b0;
        o_bus_data = 32'b0;

        instr_read_res_valid = 1'b0;
        instr_read_res_data = 32'b0;
        mem_read_res_valid = 1'b0;
        mem_read_res_data = 32'b0;
        mem_write_res_valid = 1'b0;
        if (instr_read_req_valid) begin
            o_bus_rd = 1'b1;
            o_bus_addr = instr_read_req_addr;
            instr_read_res_valid = i_bus_rd_valid;
            instr_read_res_data = i_bus_data;
        end else if (mem_read_req_valid) begin
            o_bus_rd = 1'b1;
            o_bus_addr = mem_read_req_addr;
            mem_read_res_valid = i_bus_rd_valid;
            mem_read_res_data = i_bus_data;
        end else if (mem_write_req_valid) begin
            o_bus_wr = 1'b1;
            o_bus_wrmask = mem_write_req_mask;
            o_bus_addr = mem_write_req_addr;
            o_bus_data = mem_write_req_data;
            mem_write_res_valid = i_bus_wr_valid;
        end
    end
endmodule
