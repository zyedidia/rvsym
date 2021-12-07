// 32x32 register file where register 32 is always 0. We use a synchronous
// memory so that the register can be mapped to BRAM.
module reg_file
    #(
        parameter DATA_WIDTH = 32, // bits per reg
                  ADDR_WIDTH = 5   // number of address bits
    )
    (
        input logic clk,
        input logic wr_en,
        input logic [ADDR_WIDTH-1:0] w_addr, r0_addr, r1_addr,
        input logic [DATA_WIDTH-1:0] w_data,
        output logic [DATA_WIDTH-1:0] r0_data, r1_data
    );

    logic [DATA_WIDTH-1:0] regs [0:2**ADDR_WIDTH-1];

    initial begin
        for (int i = 0; i < 32; i++) begin
            regs[i] = 32'b0;
        end
    end

    always_ff @(posedge clk) begin
        if (wr_en && w_addr != 0)
            regs[w_addr] <= w_data;

        r0_data <= regs[r0_addr];
        r1_data <= regs[r1_addr];
    end
endmodule
