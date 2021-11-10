// This module creates control signals according to the decoded instruction.
module control
    (
        input logic [31:0] instr,
        output logic [4:0] rs1, rs2, rd,
        output logic [31:0] imm,
        output logic shift_arith,
        output logic [2:0] alu_op,
        output logic a_src, b_src, negate_b,
        output logic reg_write, mem_write,
        output logic [1:0] next_pc, wb_src,
        output logic [2:0] ext_ctrl,
        output logic nop
    );

    `include "signals.svh"

    logic [6:0] op;
    logic [2:0] funct3;
    logic [6:0] funct7;

    assign op = instr[6:0];
    assign funct3 = instr[14:12];
    assign funct7 = instr[31:25];

    logic is_lui, is_auipc, is_jal, is_jalr, is_branch, is_load, is_store,
          is_iarith, is_rarith, is_fence, is_sys, is_illegal;

    // determine instruction type
    always_comb begin
        is_lui = 1'b0;
        is_auipc = 1'b0;
        is_jal = 1'b0;
        is_jalr = 1'b0;
        is_branch = 1'b0;
        is_load = 1'b0;
        is_store = 1'b0;
        is_iarith = 1'b0;
        is_rarith = 1'b0;
        is_fence = 1'b0;
        is_sys = 1'b0;
        is_illegal = 1'b0;
        case (op)
            OP_LUI: is_lui = 1'b1;
            OP_AUIPC: is_auipc = 1'b1;
            OP_JAL: is_jal = 1'b1;
            OP_JALR: is_jalr = 1'b1;
            OP_LOAD: is_load = 1'b1;
            OP_STORE: is_store = 1'b1;
            OP_BRANCH: is_branch = 1'b1;
            OP_IARITH: is_iarith = 1'b1;
            OP_RARITH: is_rarith = 1'b1;
            OP_FENCE: is_fence = 1'b1;
            OP_SYS: is_sys = 1'b1;
            default: is_illegal = 1'b1;
        endcase
    end

    // extract immediate
    logic [2:0] imm_ctrl;
    always_comb begin
        if (is_lui || is_auipc)                   imm_ctrl = IMM_UTYPE;
        else if (is_jal)                          imm_ctrl = IMM_JTYPE;
        else if (is_branch)                       imm_ctrl = IMM_BTYPE;
        else if (is_iarith || is_jalr || is_load) imm_ctrl = IMM_ITYPE;
        else if (is_store)                        imm_ctrl = IMM_STYPE;
        else                                      imm_ctrl = IMM_ANY;
    end

    assign ext_ctrl = funct3;

    extract_imm extract_imm_unit (
        .in(instr[31:7]), .ctrl(imm_ctrl), .out(imm)
    );

    // register indices
    assign rs1 = (is_jal || is_lui || is_auipc) ? 5'b0 : instr[19:15];
    assign rs2 = instr[24:20];
    assign rd = instr[11:7];

    // whether this instruction is an arithmetic shift
    assign shift_arith = instr[30];

    // whether this instruction writes to memory or the register file
    assign mem_write = is_store;
    assign reg_write = !(is_store || is_branch);

    // where the ALU a and b src should come from for this instruction
    assign a_src = is_auipc ? A_SRC_PC : A_SRC_RS1;
    assign b_src = (is_branch || is_rarith) ? B_SRC_RS2 : B_SRC_IMM;
    assign negate_b = (is_rarith && funct3 == ALU_OP_ADD && instr[30] == 1'b1) ? 1'b1 : 1'b0;

    // alu operation required by this instruction
    always_comb begin
        if (is_rarith || is_iarith)             alu_op = funct3;
        else if (is_lui || is_auipc || is_jalr) alu_op = ALU_OP_ADD;
        else if (is_store || is_load)           alu_op = ALU_OP_ADD;
        else if (is_jal)                        alu_op = ALU_OP_AND;
        else if (is_branch) begin
            case (funct3)
                3'b000, 3'b001: alu_op = ALU_OP_XOR;
                3'b100, 3'b101: alu_op = ALU_OP_SLT;
                3'b110, 3'b111: alu_op = ALU_OP_SLTU;
                default: alu_op = ALU_OP_XOR; // don't care
            endcase
        end
        else alu_op = ALU_OP_AND; // don't care
    end

    // determine how the next PC should be calculated for this instruction.
    always_comb begin
        if (is_branch) begin
            case (funct3)
                3'b000, 3'b101, 3'b111: next_pc = NEXT_PC_BR_IF_Z;
                3'b001, 3'b100, 3'b110: next_pc = NEXT_PC_BR_IF_NZ;
                default: next_pc = 0;
            endcase
        end
        else if (is_jal)  next_pc = NEXT_PC_BR_IF_Z;
        else if (is_jalr) next_pc = NEXT_PC_ALU;
        else              next_pc = NEXT_PC_INC;
    end

    // determine where the writeback value comes from for this instruction.
    always_comb begin
        if (is_load)                wb_src = WB_SRC_MEM;
        else if (is_jal || is_jalr) wb_src = WB_SRC_PC;
        else                        wb_src = WB_SRC_ALU;
    end

    // fences are nops, and the "nop" instruction which is commonly encoded as
    // `addi x0, x0, 0`, so we have a special optimization for this.
    assign nop = is_fence || instr == INSTR_NOP;
endmodule
