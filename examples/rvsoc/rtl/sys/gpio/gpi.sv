module gpi
    #(
        parameter N = 32
    )
    (
        input logic clk, rst,

        input logic [N-1:0] gpi_data,

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

    logic gpi_addr;
    assign gpi_addr = i_addr >= GPI_BASE && i_addr < GPI_BASE + GPI_SIZE;

    logic rd_valid;
    logic [31:0] rd_data;

    always_comb begin
        rd_valid = 1'b0;
        rd_data = 32'b0;
        if (i_addr == GPI_LEVEL && i_rd) begin
            rd_data = gpi_data;
            rd_valid = 1'b1;
        end
    end

    assign o_wr_valid = 1'b0;
    assign o_data = gpi_addr ? rd_data : 32'b0;
    assign o_rd_valid = gpi_addr ? rd_valid : 1'b0;
endmodule
