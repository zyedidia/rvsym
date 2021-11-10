localparam A_SRC_PC  = 1'b0;
localparam A_SRC_RS1 = 1'b1;
localparam B_SRC_IMM = 1'b0;
localparam B_SRC_RS2 = 1'b1;

localparam ALU_OP_ADD  = 3'b000;
localparam ALU_OP_SHL  = 3'b001;
localparam ALU_OP_SLT  = 3'b010;
localparam ALU_OP_SLTU = 3'b011;
localparam ALU_OP_XOR  = 3'b100;
localparam ALU_OP_SHR  = 3'b101;
localparam ALU_OP_OR   = 3'b110;
localparam ALU_OP_AND  = 3'b111;

localparam IMM_ITYPE = 3'b000;
localparam IMM_STYPE = 3'b001;
localparam IMM_BTYPE = 3'b010;
localparam IMM_JTYPE = 3'b011;
localparam IMM_UTYPE = 3'b100;
localparam IMM_ANY   = 3'b111;

localparam EXT_BYTE  = 3'b000;
localparam EXT_HALF  = 3'b001;
localparam EXT_WORD  = 3'b010;
localparam EXT_BYTEU = 3'b100;
localparam EXT_HALFU = 3'b101;

localparam MEM_MASK_BYTE  = 4'b0001;
localparam MEM_MASK_HALF  = 4'b0011;
localparam MEM_MASK_WORD  = 4'b1111;
localparam MEM_MASK_BYTEU = 4'b0001;
localparam MEM_MASK_HALFU = 4'b0011;

localparam NEXT_PC_ALU      = 2'd0;
localparam NEXT_PC_INC      = 2'd1;
localparam NEXT_PC_BR_IF_Z  = 2'd2;
localparam NEXT_PC_BR_IF_NZ = 2'd3;

localparam WB_SRC_PC  = 2'd0;
localparam WB_SRC_MEM = 2'd1;
localparam WB_SRC_ALU = 2'd2;
localparam WB_SRC_CSR = 2'd3;

localparam OP_LUI    = 7'b0110111;
localparam OP_AUIPC  = 7'b0010111;
localparam OP_JAL    = 7'b1101111;
localparam OP_JALR   = 7'b1100111;
localparam OP_LOAD   = 7'b0000011;
localparam OP_IARITH = 7'b0010011;
localparam OP_BRANCH = 7'b1100011;
localparam OP_STORE  = 7'b0100011;
localparam OP_RARITH = 7'b0110011;
localparam OP_FENCE  = 7'b0001111;
localparam OP_SYS    = 7'b1110011;

localparam INSTR_NOP = 32'h00000013;
