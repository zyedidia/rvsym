module gpio
    (
        input logic clk, rst,

        inout logic [31:0] gpio_inout,

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

    logic gpio_addr;
    assign gpio_addr = i_addr >= GPIO_BASE && i_addr < GPIO_BASE + GPIO_SIZE;

    logic [31:0] gpo_reg, gpo_next;
    logic [31:0] gpio_fsel_reg, gpio_fsel_next;
    logic wr_valid, rd_valid;

    logic [31:0] rd_data;

    always_ff @(posedge clk, posedge rst) begin
        if (rst) begin
            gpo_reg <= 32'b0;
            gpio_fsel_reg <= 32'b0;
            wr_valid <= 1'b0;
        end else begin
            gpo_reg <= gpo_next;
            gpio_fsel_reg <= gpio_fsel_next;

            wr_valid <= i_wr && i_addr != GPIO_LEVEL; // level is read-only
        end
    end

    always_comb begin
        gpo_next = gpo_reg;
        gpio_fsel_next = gpio_fsel_reg;
        rd_valid = 1'b0;
        rd_data = 32'b0;
        if (i_addr == GPIO_FSEL) begin
            if (i_wr) begin
                gpio_fsel_next = i_data;
            end else if (i_rd) begin
                rd_data = gpio_fsel_reg;
                rd_valid = 1'b1;
            end
        end
        if (i_addr == GPIO_SET && i_wr) begin
            gpo_next = gpo_reg | i_data;
        end else if (i_addr == GPIO_CLEAR && i_wr) begin
            gpo_next = gpo_reg & ~i_data;
        end else if (i_addr == GPIO_LEVEL && i_rd) begin
            rd_data = gpio_inout;
            rd_valid = 1'b1;
        end
    end

    assign o_rd_valid = gpio_addr ? rd_valid : 1'b0;
    assign o_wr_valid = gpio_addr ? wr_valid : 1'b0;
    assign o_data = gpio_addr ? rd_data : 32'b0;

    generate
        genvar i;
        for (i = 0; i < 32; i++) begin
            assign gpio_inout[i] = gpio_fsel_reg[i] ? 1'b0 : gpo_reg[i];
        end
    endgenerate
endmodule
