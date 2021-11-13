module soc_top
    (
        input logic clk, rst_n,
        
        input logic usr_btn,

        inout logic gpio_0,
        inout logic gpio_1,
        inout logic gpio_5,
        inout logic gpio_6,
        inout logic gpio_9,
        inout logic gpio_10,
        inout logic gpio_11,
        inout logic gpio_12,
        inout logic gpio_13,
        inout logic gpio_a0,
        inout logic gpio_a1,
        inout logic gpio_a2,
        inout logic gpio_a3,

        output logic rgb_led0_r,
        output logic rgb_led0_g,
        output logic rgb_led0_b
    );

    logic rst;
    assign rst = ~rst_n;

    `include "memmap.svh"

    logic bus_rd_valid;
    logic bus_wr_valid;
    logic [31:0] bus_mi_data;

    assign bus_rd_valid = ram_bus_rd_valid
        | timer_bus_rd_valid
        | gpo_bus_rd_valid
        | gpi_bus_rd_valid
        | gpio_bus_rd_valid;

    assign bus_wr_valid = ram_bus_wr_valid
        | timer_bus_wr_valid
        | gpo_bus_wr_valid
        | gpi_bus_wr_valid
        | gpio_bus_wr_valid;
    assign bus_mi_data = ram_bus_mi_data
        | timer_bus_mi_data
        | gpo_bus_mi_data
        | gpi_bus_mi_data
        | gpio_bus_mi_data;

    logic bus_rd;
    logic [31:0] bus_addr;
    logic bus_wr;
    logic [3:0] bus_wrmask;
    logic [31:0] bus_mo_data;

    cpu cpu_unit (
        .clk, .rst,

        .i_bus_rd_valid(bus_rd_valid),
        .i_bus_wr_valid(bus_wr_valid),
        .i_bus_data(bus_mi_data),

        .o_bus_rd(bus_rd),
        .o_bus_addr(bus_addr),
        .o_bus_wr(bus_wr),
        .o_bus_wrmask(bus_wrmask),
        .o_bus_data(bus_mo_data)
    );

    logic ram_bus_rd_valid;
    logic ram_bus_wr_valid;
    logic [31:0] ram_bus_mi_data;

    ram ram_unit (
        .clk,

        .i_rd(bus_rd),
        .i_addr(bus_addr),
        .i_wr(bus_wr),
        .i_wrmask(bus_wrmask),
        .i_data(bus_mo_data),

        .o_rd_valid(ram_bus_rd_valid),
        .o_wr_valid(ram_bus_wr_valid),
        .o_data(ram_bus_mi_data)
    );

    logic timer_bus_rd_valid;
    logic timer_bus_wr_valid;
    logic [31:0] timer_bus_mi_data;

    timer timer_unit (
        .clk, .rst,

        .i_rd(bus_rd),
        .i_addr(bus_addr),
        .i_wr(bus_wr),
        .i_wrmask(bus_wrmask),
        .i_data(bus_mo_data),

        .o_rd_valid(timer_bus_rd_valid),
        .o_wr_valid(timer_bus_wr_valid),
        .o_data(timer_bus_mi_data)
    );

    logic gpi_bus_rd_valid;
    logic gpi_bus_wr_valid;
    logic [31:0] gpi_bus_mi_data;

    logic [0:0] gpi_data;

    assign gpi_data[0] = ~usr_btn;

    gpi #(.N(1)) gpi_unit (
        .clk, .rst,

        .i_rd(bus_rd),
        .i_addr(bus_addr),
        .i_wr(bus_wr),
        .i_wrmask(bus_wrmask),
        .i_data(bus_mo_data),

        .o_rd_valid(gpi_bus_rd_valid),
        .o_wr_valid(gpi_bus_wr_valid),
        .o_data(gpi_bus_mi_data),

        .gpi_data
    );

    logic gpo_bus_rd_valid;
    logic gpo_bus_wr_valid;
    logic [31:0] gpo_bus_mi_data;

    logic [31:0] gpo_data;

    gpo gpo_unit (
        .clk, .rst,

        .i_rd(bus_rd),
        .i_addr(bus_addr),
        .i_wr(bus_wr),
        .i_wrmask(bus_wrmask),
        .i_data(bus_mo_data),

        .o_rd_valid(gpo_bus_rd_valid),
        .o_wr_valid(gpo_bus_wr_valid),
        .o_data(gpo_bus_mi_data),

        .gpo_data
    );

    assign rgb_led0_r = ~gpo_data[0];
    assign rgb_led0_g = ~gpo_data[1];
    assign rgb_led0_b = ~gpo_data[2];

    logic gpio_bus_rd_valid;
    logic gpio_bus_wr_valid;
    logic [31:0] gpio_bus_mi_data;

    tri [31:0] gpio_inout;

    gpio gpio_unit (
        .clk, .rst,

        .i_rd(bus_rd),
        .i_addr(bus_addr),
        .i_wr(bus_wr),
        .i_wrmask(bus_wrmask),
        .i_data(bus_mo_data),

        .o_rd_valid(gpio_bus_rd_valid),
        .o_wr_valid(gpio_bus_wr_valid),
        .o_data(gpio_bus_mi_data),

        .gpio_inout
    );

    assign gpio_0 = gpio_inout[0];
    assign gpio_1 = gpio_inout[1];
    assign gpio_5 = gpio_inout[2];
    assign gpio_6 = gpio_inout[3];
    assign gpio_9 = gpio_inout[4];
    assign gpio_10 = gpio_inout[5];
    assign gpio_11 = gpio_inout[6];
    assign gpio_12 = gpio_inout[7];
    assign gpio_13 = gpio_inout[8];
    assign gpio_a0 = gpio_inout[9];
    assign gpio_a1 = gpio_inout[10];
    assign gpio_a2 = gpio_inout[11];
    assign gpio_a3 = gpio_inout[12];
endmodule
