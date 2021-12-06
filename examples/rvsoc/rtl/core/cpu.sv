module cpu
    (
        input logic clk, rst,

        input logic i_bus_rd_valid,
        input logic i_bus_wr_valid,
        input logic [31:0] i_bus_data,

        output logic o_bus_rd,
        output logic [31:0] o_bus_addr,
        output logic o_bus_wr,
        output logic [3:0] o_bus_wrmask,
        output logic [31:0] o_bus_data
    );

    // instruction reads
    logic instr_read_res_valid;
    logic [31:0] instr_read_res_data;
    logic instr_read_req_valid;
    logic [31:0] instr_read_req_addr;

    // memory reads
    logic mem_read_res_valid;
    logic [31:0] mem_read_res_data;
    logic mem_read_req_valid;
    logic [31:0] mem_read_req_addr;

    // memory writes
    logic mem_write_res_valid;
    logic mem_write_req_valid;
    logic [31:0] mem_write_req_addr;
    logic [31:0] mem_write_req_data;
    logic [3:0] mem_write_req_mask;

    // the bus unit multiplexes the memory/instruction
    // read/write signals and outputs a bus signal for
    // them.
    bus bus_unit (
        .instr_read_res_valid,
        .instr_read_res_data,
        .instr_read_req_valid,
        .instr_read_req_addr,

        .mem_read_res_valid,
        .mem_read_res_data,
        .mem_read_req_valid,
        .mem_read_req_addr,

        .mem_write_res_valid,
        .mem_write_req_valid,
        .mem_write_req_addr,
        .mem_write_req_data,
        .mem_write_req_mask,

        .i_bus_rd_valid,
        .i_bus_wr_valid,
        .i_bus_data,

        .o_bus_rd,
        .o_bus_addr,
        .o_bus_wr,
        .o_bus_wrmask,
        .o_bus_data
    );

    `include "signals.svh"
    `include "ctrl_states.svh"

    // control state signals
    logic [2:0] state_reg = CTRL_STATE_FETCH;
    logic [2:0] state_next;

    logic [2:0] fetch_next, decode_next, exe_next, mem_next, wb_next;
    
    // Each cycle we update the control state. The fetch_next, decode_next, etc.
    // wires help to assign transitions when in a particular state, for example
    // to block in a state while waiting for a memory result.
    always_ff @(posedge clk, posedge rst) begin
        if (rst) begin
            state_reg <= CTRL_STATE_FETCH;
        end else begin
            state_reg <= state_next;
        end
    end

    always_comb begin
        case (state_reg)
            CTRL_STATE_FETCH:  state_next = fetch_next;
            CTRL_STATE_DECODE: state_next = decode_next;
            CTRL_STATE_EXE:    state_next = exe_next;
            CTRL_STATE_MEM:    state_next = mem_next;
            CTRL_STATE_WB:     state_next = wb_next;
            default:           state_next = CTRL_STATE_ERR;
        endcase
    end

    // fetch stage

    logic [31:0] pc_reg, pc_next;
    logic [31:0] instr_reg;

    fetch fetch_unit (
        .clk, .rst,
        .state_reg, .state_next,
        .fetch_next,

        .instr_read_res_valid,
        .instr_read_res_data,
        .instr_read_req_valid,
        .instr_read_req_addr,

        .pc_next,
        .pc(pc_reg),
        .instr(instr_reg)
    );

    // decode stage

    // control unit inspects the instruction and gives us a bunch of useful
    // information.
    logic [4:0] rs1, rs2, rd;
    logic [31:0] imm;
    logic shift_arith;
    logic [2:0] alu_op;
    logic a_src, b_src, negate_b;
    logic reg_write, mem_write;
    logic [1:0] next_pc, wb_src;
    logic [2:0] ext_ctrl;
    logic nop;

    control control_unit (
        .instr(instr_reg),
        .rs1, .rs2, .rd,
        .imm,
        .shift_arith,
        .alu_op,
        .a_src, .b_src, .negate_b,
        .reg_write, .mem_write,
        .next_pc, .wb_src,
        .ext_ctrl,
        .nop
    );

    logic [31:0] rd1, rd2;
    logic [31:0] w_data;
    logic wr_en;

    decode decode_unit (
        .clk, .rst,

        .nop,

        .rs1, .rs2, .rd,
        .wr_en, .w_data,
        .rd1, .rd2,

        .state_reg,
        .decode_next
    );

    // execute stage

    logic [31:0] alu_result_reg, alu_result_next;
    logic [31:0] store_data_reg, store_data_next;

    execute execute_unit (
        .clk, .rst,
        .a_src, .b_src,
        .negate_b, .shift_arith,
        .alu_op, .ext_ctrl,

        .pc_reg, .imm, .rd1, .rd2,

        .alu_result(alu_result_next),
        .store_data(store_data_next),

        .mem_write,
        .wb_src,
        .exe_next
    );

    always_ff @(posedge clk, posedge rst) begin
        if (rst) begin
            alu_result_reg <= BOOT_ADDR;
            store_data_reg <= 32'b0;
        end else begin
            alu_result_reg <= alu_result_next;
            store_data_reg <= store_data_next;
        end
    end

    // memory stage
    logic [31:0] mem_rdata;

    memory memory_unit (
        .clk, .rst,

        .ext_ctrl, .wb_src, .mem_write,

        .alu_result(alu_result_reg), .store_data(store_data_reg),

        .mem_read_res_valid,
        .mem_read_res_data,
        .mem_read_req_valid,
        .mem_read_req_addr,

        .mem_write_res_valid,
        .mem_write_req_valid,
        .mem_write_req_addr,
        .mem_write_req_data,
        .mem_write_req_mask,

        .state_reg,
        .mem_next,

        .mem_rdata
    );

    // writeback stage

    logic [31:0] pc_plus4;
    assign pc_plus4 = pc_reg + 32'd4;

    // write enable for the register file. If we are in writeback and this
    // instruction needs to write to a register, we enable it.
    assign wr_en = state_reg == CTRL_STATE_WB && reg_write;
    always_comb begin
        case (wb_src)
            // the write data for the register file depends on the decoder's
            // wb_src control signal.
            WB_SRC_PC:  w_data = pc_plus4;
            WB_SRC_MEM: w_data = mem_rdata;
            WB_SRC_ALU: w_data = alu_result_reg;
            default:    w_data = 32'b0;
        endcase
    end

    // fetch the next instruction
    assign wb_next = CTRL_STATE_FETCH;

    // use result to determine input for fetch stage
    always_comb begin
        case (next_pc)
            NEXT_PC_ALU:      pc_next = {alu_result_reg[31:1], 1'b0};
            NEXT_PC_INC:      pc_next = pc_plus4;
            NEXT_PC_BR_IF_Z:  pc_next = (alu_result_reg == 32'b0) ? pc_reg + imm : pc_plus4;
            NEXT_PC_BR_IF_NZ: pc_next = (alu_result_reg == 32'b0) ? pc_plus4 : pc_reg + imm;
            default: pc_next = BOOT_ADDR;
        endcase
    end
endmodule
