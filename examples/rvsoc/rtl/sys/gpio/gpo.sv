module gpo
    #(
        parameter N = 32
    )
    (
        input logic clk, rst,

        output logic [N-1:0] gpo_data,

        input logic        i_rd,
        input logic [31:0] i_addr,
        input logic        i_wr,
        input logic [3:0]  i_wrmask,
        input logic [31:0] i_data,

        output logic        o_rd_valid,
        output logic        o_wr_valid,
        output logic [31:0] o_data
    );

    `include "memmap.svh"

    logic gpo_addr;
    assign gpo_addr = i_addr >= GPO_BASE && i_addr < GPO_BASE + GPO_SIZE;

    logic [N-1:0] gpo_reg, gpo_next;
    logic wr_valid_reg, wr_valid_next;

    always_ff @(posedge clk, posedge rst) begin
        if (rst) begin
            gpo_reg <= '0;
            wr_valid_reg <= 1'b0;
        end else begin
            gpo_reg <= gpo_next;
            wr_valid_reg <= wr_valid_next;
        end
    end

    always_comb begin
        gpo_next = gpo_reg;
        if (i_addr == GPO_SET && i_wr) begin
            gpo_next = gpo_reg | i_data;
        end else if (i_addr == GPO_CLEAR && i_wr) begin
            gpo_next = gpo_reg & ~i_data;
        end

        wr_valid_next = i_wr;
        if (gpo_addr) begin
            o_rd_valid = 1'b0;
            o_data = 32'b0;
        end else begin
            o_rd_valid = 1'b0;
            o_data = 32'b0;
        end
    end

    assign o_wr_valid = gpo_addr ? wr_valid_reg : 1'b0;
    assign gpo_data = gpo_reg;
endmodule
