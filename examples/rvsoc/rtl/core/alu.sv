module alu
    (
        input logic [31:0] a, b,
        input logic [2:0] op,
        input logic shift_arith,
        output logic [31:0] out
    );

    `include "signals.svh"

    // sra needs to be 'logic signed'
    logic signed [31:0] sra;
    assign sra = $signed(a) >>> $signed(b[4:0]);

    always_comb begin
        case (op)
            ALU_OP_ADD:  out = a + b;
            ALU_OP_SHL:  out = a << b[4:0];
            ALU_OP_SHR:  out = shift_arith ? sra : a >> b[4:0];
            ALU_OP_SLT:  out = $signed(a) < $signed(b) ? 32'b1 : 32'b0;
            ALU_OP_SLTU: out = a < b ? 32'b1 : 32'b0;
            ALU_OP_XOR:  out = a ^ b;
            ALU_OP_OR:   out = a | b;
            ALU_OP_AND:  out = a & b;
            default:     out = 32'b0;
        endcase
    end
endmodule
