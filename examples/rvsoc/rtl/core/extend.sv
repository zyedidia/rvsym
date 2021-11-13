// The extend module sign-extends or zero-extends the lower byte or half-word
// of the input word. The operation is controlled by the ctrl input.
module extend
    (
        input logic [31:0] in,
        input logic [2:0] ctrl,
        output logic [31:0] out
    );

    `include "signals.svh"

    always_comb begin
        case (ctrl)
            EXT_BYTE: out = {{24{in[7]}}, in[7:0]};
            EXT_HALF: out = {{16{in[15]}}, in[15:0]};
            EXT_WORD: out = in;
            EXT_BYTEU: out = {24'b0, in[7:0]};
            EXT_HALFU: out = {16'b0, in[15:0]};
            default: out = 32'b0;
        endcase
    end
endmodule
