module counter
    #(
        parameter MAX_VALUE = 10000
    )
    (
        input logic clk, rst,
        output logic [31:0] q
    );

    logic [31:0] q_reg = 0;
    logic [31:0] q_next;

    always_ff @(posedge clk, posedge rst) begin
        if (rst) begin
            q_reg <= 32'b0;
        end else begin
            q_reg <= q_next;
        end
    end

    // next-state logic
    assign q_next = q_reg == MAX_VALUE ? 32'b0 : q_reg + 1;

    // output logic
    assign q = q_reg;
endmodule
