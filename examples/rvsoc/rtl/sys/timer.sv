module timer
    #(
        parameter CLKS_PER_TICK = 48
    )
    (
        input logic clk, rst,

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

    logic timer_addr;
    assign timer_addr = i_addr == TIMER_ADDR;

    logic [31:0] timer_reg, timer_next;
    logic [$clog2(CLKS_PER_TICK)-1:0] tick_reg, tick_next;

    logic [31:0] timer_data;
    logic rd_valid;

    always_ff @(posedge clk, posedge rst) begin
        if (rst) begin
            timer_reg <= 0;
            tick_reg <= 0;
        end else begin
            timer_reg <= timer_next;
            tick_reg <= tick_next;
        end
    end

    always_comb begin
        rd_valid = 1'b0;
        timer_data = 32'b0;
        if (timer_addr && i_rd) begin
            timer_data = timer_reg;
            rd_valid = 1'b1;
        end
    end

    assign timer_next = tick_reg == CLKS_PER_TICK ? timer_reg + 1 : timer_reg;
    assign tick_next = tick_reg == CLKS_PER_TICK ? 0 : tick_reg + 1;

    assign o_wr_valid = timer_addr ? 1'b0 : 1'b0;
    assign o_rd_valid = timer_addr ? rd_valid : 1'b0;
    assign o_data = timer_addr ? timer_data : 32'b0;
endmodule
