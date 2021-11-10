// This module extracts the immediate from a RISC-V instruction, depending on
// the instruction type.
module extract_imm
    (
        input logic [31:7] in,
        input logic [2:0] ctrl,
        output logic [31:0] out
    );

    `include "signals.svh"

    always_comb begin
        case (ctrl)
            IMM_ITYPE: out = {{20{in[31]}}, in[31:20]};
            IMM_STYPE: out = {{20{in[31]}}, in[31:25], in[11:7]};
            IMM_BTYPE: out = {{20{in[31]}}, in[7], in[30:25], in[11:8], 1'b0};
            IMM_JTYPE: out = {{12{in[31]}}, in[19:12], in[20], in[30:21], 1'b0};
            IMM_UTYPE: out = {in[31:12], 12'b0};
            default: out = 32'b0;
        endcase
    end
endmodule
