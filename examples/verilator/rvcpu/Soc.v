module Alu(
  input  [31:0] io_a,
  input  [31:0] io_b,
  input  [3:0]  io_op,
  output [31:0] io_out,
  output [31:0] io_sum
);
  wire [4:0] shamt = io_b[4:0]; // @[Alu.scala 22:19]
  wire [31:0] _b_T_1 = 32'h0 - io_b; // @[Alu.scala 29:10]
  wire [31:0] b = io_op == 4'h1 ? _b_T_1 : io_b; // @[Alu.scala 28:30 27:5 29:7]
  wire [31:0] _io_out_T = io_a & io_b; // @[Alu.scala 36:39]
  wire [31:0] _io_out_T_1 = io_a | io_b; // @[Alu.scala 37:39]
  wire [31:0] _io_out_T_2 = io_a ^ io_b; // @[Alu.scala 38:39]
  wire [31:0] _io_out_T_3 = io_a; // @[Alu.scala 39:39]
  wire [31:0] _io_out_T_4 = io_b; // @[Alu.scala 39:53]
  wire [62:0] _GEN_0 = {{31'd0}, io_a}; // @[Alu.scala 40:39]
  wire [62:0] _io_out_T_6 = _GEN_0 << shamt; // @[Alu.scala 40:39]
  wire [31:0] _io_out_T_8 = io_a >> shamt; // @[Alu.scala 42:39]
  wire [31:0] _io_out_T_11 = $signed(io_a) >>> shamt; // @[Alu.scala 43:57]
  wire [31:0] _GEN_1 = 4'hb == io_op ? io_b : io_a; // @[Alu.scala 24:10 33:18 44:31]
  wire [31:0] _GEN_2 = 4'h9 == io_op ? _io_out_T_11 : _GEN_1; // @[Alu.scala 33:18 43:31]
  wire [31:0] _GEN_3 = 4'h8 == io_op ? _io_out_T_8 : _GEN_2; // @[Alu.scala 33:18 42:31]
  wire [31:0] _GEN_4 = 4'h7 == io_op ? {{31'd0}, io_a < io_b} : _GEN_3; // @[Alu.scala 33:18 41:31]
  wire [62:0] _GEN_5 = 4'h6 == io_op ? _io_out_T_6 : {{31'd0}, _GEN_4}; // @[Alu.scala 33:18 40:31]
  wire [62:0] _GEN_6 = 4'h5 == io_op ? {{62'd0}, $signed(_io_out_T_3) < $signed(_io_out_T_4)} : _GEN_5; // @[Alu.scala 33:18 39:31]
  wire [62:0] _GEN_7 = 4'h4 == io_op ? {{31'd0}, _io_out_T_2} : _GEN_6; // @[Alu.scala 33:18 38:31]
  wire [62:0] _GEN_8 = 4'h3 == io_op ? {{31'd0}, _io_out_T_1} : _GEN_7; // @[Alu.scala 33:18 37:31]
  wire [62:0] _GEN_9 = 4'h2 == io_op ? {{31'd0}, _io_out_T} : _GEN_8; // @[Alu.scala 33:18 36:31]
  wire [62:0] _GEN_10 = 4'h1 == io_op ? {{31'd0}, io_sum} : _GEN_9; // @[Alu.scala 33:18 35:31]
  wire [62:0] _GEN_11 = 4'h0 == io_op ? {{31'd0}, io_sum} : _GEN_10; // @[Alu.scala 33:18 34:31]
  assign io_out = _GEN_11[31:0];
  assign io_sum = io_a + b; // @[Alu.scala 31:18]
endmodule
module RegFile(
  input         clock,
  input         io_wen,
  input  [4:0]  io_raddr1,
  input  [4:0]  io_raddr2,
  input  [4:0]  io_waddr,
  input  [31:0] io_wdata,
  output [31:0] io_rdata1,
  output [31:0] io_rdata2
);
`ifdef RANDOMIZE_MEM_INIT
  reg [31:0] _RAND_0;
`endif // RANDOMIZE_MEM_INIT
  reg [31:0] regs [0:31]; // @[RegFile.scala 21:17]
  wire  regs_io_rdata1_MPORT_en; // @[RegFile.scala 21:17]
  wire [4:0] regs_io_rdata1_MPORT_addr; // @[RegFile.scala 21:17]
  wire [31:0] regs_io_rdata1_MPORT_data; // @[RegFile.scala 21:17]
  wire  regs_io_rdata2_MPORT_en; // @[RegFile.scala 21:17]
  wire [4:0] regs_io_rdata2_MPORT_addr; // @[RegFile.scala 21:17]
  wire [31:0] regs_io_rdata2_MPORT_data; // @[RegFile.scala 21:17]
  wire [31:0] regs_MPORT_data; // @[RegFile.scala 21:17]
  wire [4:0] regs_MPORT_addr; // @[RegFile.scala 21:17]
  wire  regs_MPORT_mask; // @[RegFile.scala 21:17]
  wire  regs_MPORT_en; // @[RegFile.scala 21:17]
  wire [31:0] _GEN_5 = {{27'd0}, io_waddr}; // @[RegFile.scala 26:28]
  wire  _T = _GEN_5 != 32'h0; // @[RegFile.scala 26:28]
  assign regs_io_rdata1_MPORT_en = 1'h1;
  assign regs_io_rdata1_MPORT_addr = io_raddr1;
  assign regs_io_rdata1_MPORT_data = regs[regs_io_rdata1_MPORT_addr]; // @[RegFile.scala 21:17]
  assign regs_io_rdata2_MPORT_en = 1'h1;
  assign regs_io_rdata2_MPORT_addr = io_raddr2;
  assign regs_io_rdata2_MPORT_data = regs[regs_io_rdata2_MPORT_addr]; // @[RegFile.scala 21:17]
  assign regs_MPORT_data = io_wdata;
  assign regs_MPORT_addr = io_waddr;
  assign regs_MPORT_mask = 1'h1;
  assign regs_MPORT_en = io_wen & _T;
  assign io_rdata1 = regs_io_rdata1_MPORT_data; // @[RegFile.scala 23:13]
  assign io_rdata2 = regs_io_rdata2_MPORT_data; // @[RegFile.scala 24:13]
  always @(posedge clock) begin
    if (regs_MPORT_en & regs_MPORT_mask) begin
      regs[regs_MPORT_addr] <= regs_MPORT_data; // @[RegFile.scala 21:17]
    end
  end
// Register and memory initialization
`ifdef RANDOMIZE_GARBAGE_ASSIGN
`define RANDOMIZE
`endif
`ifdef RANDOMIZE_INVALID_ASSIGN
`define RANDOMIZE
`endif
`ifdef RANDOMIZE_REG_INIT
`define RANDOMIZE
`endif
`ifdef RANDOMIZE_MEM_INIT
`define RANDOMIZE
`endif
`ifndef RANDOM
`define RANDOM $random
`endif
`ifdef RANDOMIZE_MEM_INIT
  integer initvar;
`endif
`ifndef SYNTHESIS
`ifdef FIRRTL_BEFORE_INITIAL
`FIRRTL_BEFORE_INITIAL
`endif
initial begin
  `ifdef RANDOMIZE
    `ifdef INIT_RANDOM
      `INIT_RANDOM
    `endif
    `ifndef VERILATOR
      `ifdef RANDOMIZE_DELAY
        #`RANDOMIZE_DELAY begin end
      `else
        #0.002 begin end
      `endif
    `endif
`ifdef RANDOMIZE_MEM_INIT
  _RAND_0 = {1{`RANDOM}};
  for (initvar = 0; initvar < 32; initvar = initvar+1)
    regs[initvar] = _RAND_0[31:0];
`endif // RANDOMIZE_MEM_INIT
  `endif // RANDOMIZE
end // initial
`ifdef FIRRTL_AFTER_INITIAL
`FIRRTL_AFTER_INITIAL
`endif
`endif // SYNTHESIS
endmodule
module BrCond(
  input  [31:0] io_rs1,
  input  [31:0] io_rs2,
  input  [2:0]  io_br_type,
  output        io_taken
);
  wire  eq = io_rs1 == io_rs2; // @[BrCond.scala 16:19]
  wire  lt = $signed(io_rs1) < $signed(io_rs2); // @[BrCond.scala 17:26]
  wire  ltu = io_rs1 < io_rs2; // @[BrCond.scala 18:20]
  wire  _GEN_0 = 3'h4 == io_br_type & ~ltu; // @[BrCond.scala 20:12 22:23 28:32]
  wire  _GEN_1 = 3'h1 == io_br_type ? ltu : _GEN_0; // @[BrCond.scala 22:23 27:32]
  wire  _GEN_2 = 3'h5 == io_br_type ? ~lt : _GEN_1; // @[BrCond.scala 22:23 26:32]
  wire  _GEN_3 = 3'h2 == io_br_type ? lt : _GEN_2; // @[BrCond.scala 22:23 25:32]
  wire  _GEN_4 = 3'h6 == io_br_type ? ~eq : _GEN_3; // @[BrCond.scala 22:23 24:32]
  assign io_taken = 3'h3 == io_br_type ? eq : _GEN_4; // @[BrCond.scala 22:23 23:32]
endmodule
module ImmExtract(
  input  [31:0] io_inst,
  input  [2:0]  io_sel,
  output [31:0] io_out
);
  wire [11:0] _sint_T_1 = io_inst[31:20]; // @[ImmExtract.scala 20:45]
  wire [11:0] _sint_T_5 = {io_inst[31:25],io_inst[11:7]}; // @[ImmExtract.scala 21:66]
  wire [12:0] _sint_T_11 = {io_inst[31],io_inst[7],io_inst[30:25],io_inst[11:8],1'h0}; // @[ImmExtract.scala 22:101]
  wire [31:0] _sint_T_14 = {io_inst[31:12],12'h0}; // @[ImmExtract.scala 23:61]
  wire [20:0] _sint_T_21 = {io_inst[31],io_inst[19:12],io_inst[20],io_inst[30:25],io_inst[24:21],1'h0}; // @[ImmExtract.scala 24:120]
  wire [5:0] _sint_T_23 = {1'b0,$signed(io_inst[19:15])}; // @[ImmExtract.scala 25:45]
  wire [5:0] _GEN_0 = 3'h6 == io_sel ? $signed(_sint_T_23) : $signed(6'sh0); // @[ImmExtract.scala 19:19 25:26 18:8]
  wire [20:0] _GEN_1 = 3'h4 == io_sel ? $signed(_sint_T_21) : $signed({{15{_GEN_0[5]}},_GEN_0}); // @[ImmExtract.scala 19:19 24:26]
  wire [31:0] _GEN_2 = 3'h3 == io_sel ? $signed(_sint_T_14) : $signed({{11{_GEN_1[20]}},_GEN_1}); // @[ImmExtract.scala 19:19 23:26]
  wire [31:0] _GEN_3 = 3'h5 == io_sel ? $signed({{19{_sint_T_11[12]}},_sint_T_11}) : $signed(_GEN_2); // @[ImmExtract.scala 19:19 22:26]
  wire [31:0] _GEN_4 = 3'h2 == io_sel ? $signed({{20{_sint_T_5[11]}},_sint_T_5}) : $signed(_GEN_3); // @[ImmExtract.scala 19:19 21:26]
  assign io_out = 3'h1 == io_sel ? $signed({{20{_sint_T_1[11]}},_sint_T_1}) : $signed(_GEN_4); // @[ImmExtract.scala 16:18]
endmodule
module Datapath(
  input         clock,
  input         reset,
  output [31:0] io_imem_addr,
  input  [31:0] io_imem_rdata,
  output        io_dmem_req,
  output [31:0] io_dmem_addr,
  output        io_dmem_we,
  output [3:0]  io_dmem_be,
  output [31:0] io_dmem_wdata,
  input         io_dmem_gnt,
  input         io_dmem_rvalid,
  input  [31:0] io_dmem_rdata,
  output [31:0] io_ctrl_inst,
  input  [1:0]  io_ctrl_pc_sel,
  input         io_ctrl_inst_kill,
  input  [1:0]  io_ctrl_a_sel,
  input  [1:0]  io_ctrl_b_sel,
  input  [2:0]  io_ctrl_imm_sel,
  input  [3:0]  io_ctrl_alu_op,
  input  [2:0]  io_ctrl_br_type,
  input  [1:0]  io_ctrl_st_type,
  input  [2:0]  io_ctrl_ld_type,
  input  [1:0]  io_ctrl_wb_sel,
  input         io_ctrl_wb_en
);
`ifdef RANDOMIZE_REG_INIT
  reg [31:0] _RAND_0;
  reg [31:0] _RAND_1;
  reg [31:0] _RAND_2;
  reg [31:0] _RAND_3;
  reg [31:0] _RAND_4;
  reg [31:0] _RAND_5;
  reg [31:0] _RAND_6;
  reg [31:0] _RAND_7;
  reg [31:0] _RAND_8;
  reg [31:0] _RAND_9;
  reg [31:0] _RAND_10;
  reg [31:0] _RAND_11;
`endif // RANDOMIZE_REG_INIT
  wire [31:0] alu_io_a; // @[Datapath.scala 15:19]
  wire [31:0] alu_io_b; // @[Datapath.scala 15:19]
  wire [3:0] alu_io_op; // @[Datapath.scala 15:19]
  wire [31:0] alu_io_out; // @[Datapath.scala 15:19]
  wire [31:0] alu_io_sum; // @[Datapath.scala 15:19]
  wire  rf_clock; // @[Datapath.scala 16:18]
  wire  rf_io_wen; // @[Datapath.scala 16:18]
  wire [4:0] rf_io_raddr1; // @[Datapath.scala 16:18]
  wire [4:0] rf_io_raddr2; // @[Datapath.scala 16:18]
  wire [4:0] rf_io_waddr; // @[Datapath.scala 16:18]
  wire [31:0] rf_io_wdata; // @[Datapath.scala 16:18]
  wire [31:0] rf_io_rdata1; // @[Datapath.scala 16:18]
  wire [31:0] rf_io_rdata2; // @[Datapath.scala 16:18]
  wire [31:0] brCond_io_rs1; // @[Datapath.scala 17:22]
  wire [31:0] brCond_io_rs2; // @[Datapath.scala 17:22]
  wire [2:0] brCond_io_br_type; // @[Datapath.scala 17:22]
  wire  brCond_io_taken; // @[Datapath.scala 17:22]
  wire [31:0] immExt_io_inst; // @[Datapath.scala 18:22]
  wire [2:0] immExt_io_sel; // @[Datapath.scala 18:22]
  wire [31:0] immExt_io_out; // @[Datapath.scala 18:22]
  reg [31:0] fe_inst; // @[Datapath.scala 23:24]
  reg [31:0] fe_pc; // @[Datapath.scala 24:18]
  reg [31:0] ew_inst; // @[Datapath.scala 27:24]
  reg [31:0] ew_pc; // @[Datapath.scala 28:18]
  reg [31:0] ew_alu; // @[Datapath.scala 29:19]
  reg [1:0] st_type; // @[Datapath.scala 32:20]
  reg [2:0] ld_type; // @[Datapath.scala 33:20]
  reg [1:0] wb_sel; // @[Datapath.scala 34:19]
  reg  wb_en; // @[Datapath.scala 35:18]
  reg [31:0] wb_ld; // @[Datapath.scala 38:18]
  reg  started; // @[Datapath.scala 44:24]
  wire  dmem_rd_req = io_ctrl_ld_type != 3'h0; // @[Datapath.scala 115:34]
  wire  dmem_wr_req = io_ctrl_st_type != 2'h0; // @[Datapath.scala 116:34]
  wire  _stall_T_6 = dmem_wr_req & ~io_dmem_gnt; // @[Datapath.scala 51:25]
  wire  stall = dmem_rd_req & ~io_dmem_rvalid | _stall_T_6; // @[Datapath.scala 50:45]
  wire [31:0] _pc_T_1 = 32'h100000 - 32'h4; // @[Datapath.scala 53:34]
  reg [31:0] pc; // @[Datapath.scala 53:19]
  wire [31:0] _next_pc_T_1 = alu_io_sum & 32'hfffffffe; // @[Datapath.scala 60:27]
  wire [31:0] _next_pc_T_3 = pc + 32'h4; // @[Datapath.scala 64:19]
  wire [31:0] _GEN_0 = io_ctrl_pc_sel == 2'h2 ? pc : _next_pc_T_3; // @[Datapath.scala 61:48 62:13 64:13]
  wire [31:0] _GEN_1 = io_ctrl_pc_sel == 2'h1 | brCond_io_taken ? _next_pc_T_1 : _GEN_0; // @[Datapath.scala 59:65 60:13]
  wire  _inst_T_1 = started | io_ctrl_inst_kill | brCond_io_taken; // @[Datapath.scala 68:34]
  wire  _T_3 = ~stall; // @[Datapath.scala 76:9]
  wire [4:0] rs1_addr = fe_inst[19:15]; // @[Datapath.scala 86:25]
  wire [4:0] rs2_addr = fe_inst[24:20]; // @[Datapath.scala 87:25]
  wire [4:0] wb_rd_addr = ew_inst[11:7]; // @[Datapath.scala 96:27]
  wire  rs1hzd = wb_en & rs1_addr != 5'h0 & rs1_addr == wb_rd_addr; // @[Datapath.scala 98:42]
  wire  rs2hzd = wb_en & rs2_addr != 5'h0 & rs2_addr == wb_rd_addr; // @[Datapath.scala 99:42]
  wire  _rs1_T = wb_sel == 2'h0; // @[Datapath.scala 101:24]
  wire [31:0] rs1 = wb_sel == 2'h0 & rs1hzd ? ew_alu : rf_io_rdata1; // @[Datapath.scala 101:16]
  wire [31:0] rs2 = _rs1_T & rs2hzd ? ew_alu : rf_io_rdata2; // @[Datapath.scala 102:16]
  wire [4:0] _GEN_23 = {alu_io_sum[1], 4'h0}; // @[Datapath.scala 113:32]
  wire [7:0] _woffset_T_1 = {{3'd0}, _GEN_23}; // @[Datapath.scala 113:32]
  wire [3:0] _woffset_T_3 = {alu_io_sum[0], 3'h0}; // @[Datapath.scala 113:64]
  wire [7:0] _GEN_24 = {{4'd0}, _woffset_T_3}; // @[Datapath.scala 113:47]
  wire [7:0] woffset = _woffset_T_1 | _GEN_24; // @[Datapath.scala 113:47]
  wire [286:0] _GEN_2 = {{255'd0}, rs2}; // @[Datapath.scala 121:24]
  wire [286:0] _io_dmem_wdata_T = _GEN_2 << woffset; // @[Datapath.scala 121:24]
  wire [1:0] _T_4 = stall ? st_type : io_ctrl_st_type; // @[Datapath.scala 124:14]
  wire [4:0] _io_dmem_be_T_1 = 5'h3 << alu_io_sum[1:0]; // @[Datapath.scala 125:44]
  wire [3:0] _io_dmem_be_T_3 = 4'h1 << alu_io_sum[1:0]; // @[Datapath.scala 126:43]
  wire [3:0] _GEN_5 = 2'h3 == _T_4 ? _io_dmem_be_T_3 : 4'hf; // @[Datapath.scala 123:14 124:49 126:33]
  wire [4:0] _GEN_6 = 2'h2 == _T_4 ? _io_dmem_be_T_1 : {{1'd0}, _GEN_5}; // @[Datapath.scala 124:49 125:33]
  wire [4:0] _GEN_25 = {ew_alu[1], 4'h0}; // @[Datapath.scala 129:26]
  wire [7:0] _ldoff_T_1 = {{3'd0}, _GEN_25}; // @[Datapath.scala 129:26]
  wire [3:0] _ldoff_T_3 = {ew_alu[0], 3'h0}; // @[Datapath.scala 129:54]
  wire [7:0] _GEN_26 = {{4'd0}, _ldoff_T_3}; // @[Datapath.scala 129:41]
  wire [7:0] ldoff = _ldoff_T_1 | _GEN_26; // @[Datapath.scala 129:41]
  wire [31:0] ldshift = io_dmem_rdata >> ldoff; // @[Datapath.scala 130:31]
  wire [32:0] _ld_T = {1'b0,$signed(io_dmem_rdata)}; // @[Datapath.scala 133:23]
  wire [15:0] _ld_T_2 = ldshift[15:0]; // @[Datapath.scala 135:44]
  wire [7:0] _ld_T_4 = ldshift[7:0]; // @[Datapath.scala 136:43]
  wire [16:0] _ld_T_6 = {1'b0,$signed(ldshift[15:0])}; // @[Datapath.scala 137:44]
  wire [8:0] _ld_T_8 = {1'b0,$signed(ldshift[7:0])}; // @[Datapath.scala 138:43]
  wire [32:0] _GEN_7 = 3'h5 == ld_type ? $signed({{24{_ld_T_8[8]}},_ld_T_8}) : $signed(_ld_T); // @[Datapath.scala 134:20 138:26 133:6]
  wire [32:0] _GEN_8 = 3'h4 == ld_type ? $signed({{16{_ld_T_6[16]}},_ld_T_6}) : $signed(_GEN_7); // @[Datapath.scala 134:20 137:26]
  wire [32:0] _GEN_9 = 3'h3 == ld_type ? $signed({{25{_ld_T_4[7]}},_ld_T_4}) : $signed(_GEN_8); // @[Datapath.scala 134:20 136:26]
  wire [32:0] _GEN_10 = 3'h2 == ld_type ? $signed({{17{_ld_T_2[15]}},_ld_T_2}) : $signed(_GEN_9); // @[Datapath.scala 134:20 135:26]
  wire [31:0] _wb_ld_T = _GEN_10[31:0]; // @[Datapath.scala 145:17]
  wire [32:0] _regwr_T_1 = {1'b0,$signed(ew_alu)}; // @[Datapath.scala 157:24]
  wire [31:0] _regwr_T_3 = ew_pc + 32'h4; // @[Datapath.scala 160:37]
  wire [32:0] _GEN_21 = 2'h2 == wb_sel ? {{1'd0}, _regwr_T_3} : _regwr_T_1; // @[Datapath.scala 158:19 160:28 157:9]
  wire [32:0] _GEN_22 = 2'h1 == wb_sel ? {{1'd0}, wb_ld} : _GEN_21; // @[Datapath.scala 158:19 159:28]
  Alu alu ( // @[Datapath.scala 15:19]
    .io_a(alu_io_a),
    .io_b(alu_io_b),
    .io_op(alu_io_op),
    .io_out(alu_io_out),
    .io_sum(alu_io_sum)
  );
  RegFile rf ( // @[Datapath.scala 16:18]
    .clock(rf_clock),
    .io_wen(rf_io_wen),
    .io_raddr1(rf_io_raddr1),
    .io_raddr2(rf_io_raddr2),
    .io_waddr(rf_io_waddr),
    .io_wdata(rf_io_wdata),
    .io_rdata1(rf_io_rdata1),
    .io_rdata2(rf_io_rdata2)
  );
  BrCond brCond ( // @[Datapath.scala 17:22]
    .io_rs1(brCond_io_rs1),
    .io_rs2(brCond_io_rs2),
    .io_br_type(brCond_io_br_type),
    .io_taken(brCond_io_taken)
  );
  ImmExtract immExt ( // @[Datapath.scala 18:22]
    .io_inst(immExt_io_inst),
    .io_sel(immExt_io_sel),
    .io_out(immExt_io_out)
  );
  assign io_imem_addr = stall ? pc : _GEN_1; // @[Datapath.scala 57:16 58:13]
  assign io_dmem_req = dmem_rd_req | dmem_wr_req; // @[Datapath.scala 118:30]
  assign io_dmem_addr = alu_io_sum & 32'hfffffffc; // @[Datapath.scala 112:26]
  assign io_dmem_we = io_ctrl_st_type != 2'h0; // @[Datapath.scala 116:34]
  assign io_dmem_be = _GEN_6[3:0];
  assign io_dmem_wdata = _io_dmem_wdata_T[31:0]; // @[Datapath.scala 121:17]
  assign io_ctrl_inst = fe_inst; // @[Datapath.scala 83:16]
  assign alu_io_a = io_ctrl_a_sel == 2'h2 ? rs1 : fe_pc; // @[Datapath.scala 104:18]
  assign alu_io_b = io_ctrl_b_sel == 2'h2 ? rs2 : immExt_io_out; // @[Datapath.scala 105:18]
  assign alu_io_op = io_ctrl_alu_op; // @[Datapath.scala 106:13]
  assign rf_clock = clock;
  assign rf_io_wen = wb_en & _T_3; // @[Datapath.scala 163:22]
  assign rf_io_raddr1 = fe_inst[19:15]; // @[Datapath.scala 86:25]
  assign rf_io_raddr2 = fe_inst[24:20]; // @[Datapath.scala 87:25]
  assign rf_io_waddr = ew_inst[11:7]; // @[Datapath.scala 96:27]
  assign rf_io_wdata = _GEN_22[31:0]; // @[Datapath.scala 155:19]
  assign brCond_io_rs1 = wb_sel == 2'h0 & rs1hzd ? ew_alu : rf_io_rdata1; // @[Datapath.scala 101:16]
  assign brCond_io_rs2 = _rs1_T & rs2hzd ? ew_alu : rf_io_rdata2; // @[Datapath.scala 102:16]
  assign brCond_io_br_type = io_ctrl_br_type; // @[Datapath.scala 110:21]
  assign immExt_io_inst = fe_inst; // @[Datapath.scala 92:18]
  assign immExt_io_sel = io_ctrl_imm_sel; // @[Datapath.scala 93:17]
  always @(posedge clock) begin
    if (reset) begin // @[Datapath.scala 23:24]
      fe_inst <= 32'h13; // @[Datapath.scala 23:24]
    end else if (~stall) begin // @[Datapath.scala 76:17]
      if (_inst_T_1) begin // @[Datapath.scala 67:17]
        fe_inst <= 32'h13;
      end else begin
        fe_inst <= io_imem_rdata;
      end
    end
    if (~stall) begin // @[Datapath.scala 76:17]
      fe_pc <= pc; // @[Datapath.scala 77:11]
    end
    if (reset) begin // @[Datapath.scala 27:24]
      ew_inst <= 32'h13; // @[Datapath.scala 27:24]
    end else if (_T_3) begin // @[Datapath.scala 141:17]
      ew_inst <= fe_inst; // @[Datapath.scala 143:13]
    end
    if (_T_3) begin // @[Datapath.scala 141:17]
      ew_pc <= fe_pc; // @[Datapath.scala 142:11]
    end
    if (_T_3) begin // @[Datapath.scala 141:17]
      ew_alu <= alu_io_out; // @[Datapath.scala 144:12]
    end
    if (_T_3) begin // @[Datapath.scala 141:17]
      st_type <= io_ctrl_st_type; // @[Datapath.scala 146:13]
    end
    if (_T_3) begin // @[Datapath.scala 141:17]
      ld_type <= io_ctrl_ld_type; // @[Datapath.scala 147:13]
    end
    if (_T_3) begin // @[Datapath.scala 141:17]
      wb_sel <= io_ctrl_wb_sel; // @[Datapath.scala 148:12]
    end
    if (_T_3) begin // @[Datapath.scala 141:17]
      wb_en <= io_ctrl_wb_en; // @[Datapath.scala 149:11]
    end
    if (_T_3) begin // @[Datapath.scala 141:17]
      wb_ld <= _wb_ld_T; // @[Datapath.scala 145:11]
    end
    started <= reset; // @[Datapath.scala 44:31]
    if (reset) begin // @[Datapath.scala 53:19]
      pc <= _pc_T_1; // @[Datapath.scala 53:19]
    end else if (!(stall)) begin // @[Datapath.scala 57:16]
      if (io_ctrl_pc_sel == 2'h1 | brCond_io_taken) begin // @[Datapath.scala 59:65]
        pc <= _next_pc_T_1; // @[Datapath.scala 60:13]
      end else if (!(io_ctrl_pc_sel == 2'h2)) begin // @[Datapath.scala 61:48]
        pc <= _next_pc_T_3; // @[Datapath.scala 64:13]
      end
    end
  end
// Register and memory initialization
`ifdef RANDOMIZE_GARBAGE_ASSIGN
`define RANDOMIZE
`endif
`ifdef RANDOMIZE_INVALID_ASSIGN
`define RANDOMIZE
`endif
`ifdef RANDOMIZE_REG_INIT
`define RANDOMIZE
`endif
`ifdef RANDOMIZE_MEM_INIT
`define RANDOMIZE
`endif
`ifndef RANDOM
`define RANDOM $random
`endif
`ifdef RANDOMIZE_MEM_INIT
  integer initvar;
`endif
`ifndef SYNTHESIS
`ifdef FIRRTL_BEFORE_INITIAL
`FIRRTL_BEFORE_INITIAL
`endif
initial begin
  `ifdef RANDOMIZE
    `ifdef INIT_RANDOM
      `INIT_RANDOM
    `endif
    `ifndef VERILATOR
      `ifdef RANDOMIZE_DELAY
        #`RANDOMIZE_DELAY begin end
      `else
        #0.002 begin end
      `endif
    `endif
`ifdef RANDOMIZE_REG_INIT
  _RAND_0 = {1{`RANDOM}};
  fe_inst = _RAND_0[31:0];
  _RAND_1 = {1{`RANDOM}};
  fe_pc = _RAND_1[31:0];
  _RAND_2 = {1{`RANDOM}};
  ew_inst = _RAND_2[31:0];
  _RAND_3 = {1{`RANDOM}};
  ew_pc = _RAND_3[31:0];
  _RAND_4 = {1{`RANDOM}};
  ew_alu = _RAND_4[31:0];
  _RAND_5 = {1{`RANDOM}};
  st_type = _RAND_5[1:0];
  _RAND_6 = {1{`RANDOM}};
  ld_type = _RAND_6[2:0];
  _RAND_7 = {1{`RANDOM}};
  wb_sel = _RAND_7[1:0];
  _RAND_8 = {1{`RANDOM}};
  wb_en = _RAND_8[0:0];
  _RAND_9 = {1{`RANDOM}};
  wb_ld = _RAND_9[31:0];
  _RAND_10 = {1{`RANDOM}};
  started = _RAND_10[0:0];
  _RAND_11 = {1{`RANDOM}};
  pc = _RAND_11[31:0];
`endif // RANDOMIZE_REG_INIT
  `endif // RANDOMIZE
end // initial
`ifdef FIRRTL_AFTER_INITIAL
`FIRRTL_AFTER_INITIAL
`endif
`endif // SYNTHESIS
endmodule
module Control(
  input  [31:0] io_inst,
  output [1:0]  io_pc_sel,
  output        io_inst_kill,
  output [1:0]  io_a_sel,
  output [1:0]  io_b_sel,
  output [2:0]  io_imm_sel,
  output [3:0]  io_alu_op,
  output [2:0]  io_br_type,
  output [1:0]  io_st_type,
  output [2:0]  io_ld_type,
  output [1:0]  io_wb_sel,
  output        io_wb_en
);
  wire [31:0] _signals_T = io_inst & 32'h7f; // @[Lookup.scala 31:38]
  wire  _signals_T_1 = 32'h37 == _signals_T; // @[Lookup.scala 31:38]
  wire  _signals_T_3 = 32'h17 == _signals_T; // @[Lookup.scala 31:38]
  wire  _signals_T_5 = 32'h6f == _signals_T; // @[Lookup.scala 31:38]
  wire [31:0] _signals_T_6 = io_inst & 32'h707f; // @[Lookup.scala 31:38]
  wire  _signals_T_7 = 32'h67 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_9 = 32'h63 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_11 = 32'h1063 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_13 = 32'h4063 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_15 = 32'h5063 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_17 = 32'h6063 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_19 = 32'h7063 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_21 = 32'h3 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_23 = 32'h1003 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_25 = 32'h2003 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_27 = 32'h4003 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_29 = 32'h5003 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_31 = 32'h23 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_33 = 32'h1023 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_35 = 32'h2023 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_37 = 32'h13 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_39 = 32'h2013 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_41 = 32'h3013 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_43 = 32'h4013 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_45 = 32'h6013 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_47 = 32'h7013 == _signals_T_6; // @[Lookup.scala 31:38]
  wire [31:0] _signals_T_48 = io_inst & 32'hfe00707f; // @[Lookup.scala 31:38]
  wire  _signals_T_49 = 32'h1013 == _signals_T_48; // @[Lookup.scala 31:38]
  wire  _signals_T_51 = 32'h5013 == _signals_T_48; // @[Lookup.scala 31:38]
  wire  _signals_T_53 = 32'h40005013 == _signals_T_48; // @[Lookup.scala 31:38]
  wire  _signals_T_55 = 32'h33 == _signals_T_48; // @[Lookup.scala 31:38]
  wire  _signals_T_57 = 32'h40000033 == _signals_T_48; // @[Lookup.scala 31:38]
  wire  _signals_T_59 = 32'h1033 == _signals_T_48; // @[Lookup.scala 31:38]
  wire  _signals_T_61 = 32'h2033 == _signals_T_48; // @[Lookup.scala 31:38]
  wire  _signals_T_63 = 32'h3033 == _signals_T_48; // @[Lookup.scala 31:38]
  wire  _signals_T_65 = 32'h4033 == _signals_T_48; // @[Lookup.scala 31:38]
  wire  _signals_T_67 = 32'h5033 == _signals_T_48; // @[Lookup.scala 31:38]
  wire  _signals_T_69 = 32'h40005033 == _signals_T_48; // @[Lookup.scala 31:38]
  wire  _signals_T_71 = 32'h6033 == _signals_T_48; // @[Lookup.scala 31:38]
  wire  _signals_T_73 = 32'h7033 == _signals_T_48; // @[Lookup.scala 31:38]
  wire [31:0] _signals_T_74 = io_inst & 32'hf00fffff; // @[Lookup.scala 31:38]
  wire  _signals_T_75 = 32'hf == _signals_T_74; // @[Lookup.scala 31:38]
  wire  _signals_T_77 = 32'h100f == io_inst; // @[Lookup.scala 31:38]
  wire  _signals_T_79 = 32'h1073 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_81 = 32'h2073 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_83 = 32'h3073 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_85 = 32'h5073 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_87 = 32'h6073 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_89 = 32'h7073 == _signals_T_6; // @[Lookup.scala 31:38]
  wire  _signals_T_91 = 32'h73 == io_inst; // @[Lookup.scala 31:38]
  wire  _signals_T_93 = 32'h100073 == io_inst; // @[Lookup.scala 31:38]
  wire  _signals_T_95 = 32'h10000073 == io_inst; // @[Lookup.scala 31:38]
  wire  _signals_T_97 = 32'h10200073 == io_inst; // @[Lookup.scala 31:38]
  wire [1:0] _signals_T_98 = _signals_T_97 ? 2'h2 : 2'h0; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_99 = _signals_T_95 ? 2'h3 : _signals_T_98; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_100 = _signals_T_93 ? 2'h2 : _signals_T_99; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_101 = _signals_T_91 ? 2'h2 : _signals_T_100; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_102 = _signals_T_89 ? 2'h2 : _signals_T_101; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_103 = _signals_T_87 ? 2'h2 : _signals_T_102; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_104 = _signals_T_85 ? 2'h2 : _signals_T_103; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_105 = _signals_T_83 ? 2'h2 : _signals_T_104; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_106 = _signals_T_81 ? 2'h2 : _signals_T_105; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_107 = _signals_T_79 ? 2'h2 : _signals_T_106; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_108 = _signals_T_77 ? 2'h2 : _signals_T_107; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_109 = _signals_T_75 ? 2'h0 : _signals_T_108; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_110 = _signals_T_73 ? 2'h0 : _signals_T_109; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_111 = _signals_T_71 ? 2'h0 : _signals_T_110; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_112 = _signals_T_69 ? 2'h0 : _signals_T_111; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_113 = _signals_T_67 ? 2'h0 : _signals_T_112; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_114 = _signals_T_65 ? 2'h0 : _signals_T_113; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_115 = _signals_T_63 ? 2'h0 : _signals_T_114; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_116 = _signals_T_61 ? 2'h0 : _signals_T_115; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_117 = _signals_T_59 ? 2'h0 : _signals_T_116; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_118 = _signals_T_57 ? 2'h0 : _signals_T_117; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_119 = _signals_T_55 ? 2'h0 : _signals_T_118; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_120 = _signals_T_53 ? 2'h0 : _signals_T_119; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_121 = _signals_T_51 ? 2'h0 : _signals_T_120; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_122 = _signals_T_49 ? 2'h0 : _signals_T_121; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_123 = _signals_T_47 ? 2'h0 : _signals_T_122; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_124 = _signals_T_45 ? 2'h0 : _signals_T_123; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_125 = _signals_T_43 ? 2'h0 : _signals_T_124; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_126 = _signals_T_41 ? 2'h0 : _signals_T_125; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_127 = _signals_T_39 ? 2'h0 : _signals_T_126; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_128 = _signals_T_37 ? 2'h0 : _signals_T_127; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_129 = _signals_T_35 ? 2'h0 : _signals_T_128; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_130 = _signals_T_33 ? 2'h0 : _signals_T_129; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_131 = _signals_T_31 ? 2'h0 : _signals_T_130; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_132 = _signals_T_29 ? 2'h2 : _signals_T_131; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_133 = _signals_T_27 ? 2'h2 : _signals_T_132; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_134 = _signals_T_25 ? 2'h2 : _signals_T_133; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_135 = _signals_T_23 ? 2'h2 : _signals_T_134; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_136 = _signals_T_21 ? 2'h2 : _signals_T_135; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_137 = _signals_T_19 ? 2'h0 : _signals_T_136; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_138 = _signals_T_17 ? 2'h0 : _signals_T_137; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_139 = _signals_T_15 ? 2'h0 : _signals_T_138; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_140 = _signals_T_13 ? 2'h0 : _signals_T_139; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_141 = _signals_T_11 ? 2'h0 : _signals_T_140; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_142 = _signals_T_9 ? 2'h0 : _signals_T_141; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_143 = _signals_T_7 ? 2'h1 : _signals_T_142; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_144 = _signals_T_5 ? 2'h1 : _signals_T_143; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_145 = _signals_T_3 ? 2'h0 : _signals_T_144; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_153 = _signals_T_83 ? 2'h2 : 2'h0; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_154 = _signals_T_81 ? 2'h2 : _signals_T_153; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_155 = _signals_T_79 ? 2'h2 : _signals_T_154; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_156 = _signals_T_77 ? 2'h0 : _signals_T_155; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_157 = _signals_T_75 ? 2'h0 : _signals_T_156; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_158 = _signals_T_73 ? 2'h2 : _signals_T_157; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_159 = _signals_T_71 ? 2'h2 : _signals_T_158; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_160 = _signals_T_69 ? 2'h2 : _signals_T_159; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_161 = _signals_T_67 ? 2'h2 : _signals_T_160; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_162 = _signals_T_65 ? 2'h2 : _signals_T_161; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_163 = _signals_T_63 ? 2'h2 : _signals_T_162; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_164 = _signals_T_61 ? 2'h2 : _signals_T_163; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_165 = _signals_T_59 ? 2'h2 : _signals_T_164; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_166 = _signals_T_57 ? 2'h2 : _signals_T_165; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_167 = _signals_T_55 ? 2'h2 : _signals_T_166; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_168 = _signals_T_53 ? 2'h2 : _signals_T_167; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_169 = _signals_T_51 ? 2'h2 : _signals_T_168; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_170 = _signals_T_49 ? 2'h2 : _signals_T_169; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_171 = _signals_T_47 ? 2'h2 : _signals_T_170; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_172 = _signals_T_45 ? 2'h2 : _signals_T_171; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_173 = _signals_T_43 ? 2'h2 : _signals_T_172; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_174 = _signals_T_41 ? 2'h2 : _signals_T_173; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_175 = _signals_T_39 ? 2'h2 : _signals_T_174; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_176 = _signals_T_37 ? 2'h2 : _signals_T_175; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_177 = _signals_T_35 ? 2'h2 : _signals_T_176; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_178 = _signals_T_33 ? 2'h2 : _signals_T_177; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_179 = _signals_T_31 ? 2'h2 : _signals_T_178; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_180 = _signals_T_29 ? 2'h2 : _signals_T_179; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_181 = _signals_T_27 ? 2'h2 : _signals_T_180; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_182 = _signals_T_25 ? 2'h2 : _signals_T_181; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_183 = _signals_T_23 ? 2'h2 : _signals_T_182; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_184 = _signals_T_21 ? 2'h2 : _signals_T_183; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_185 = _signals_T_19 ? 2'h1 : _signals_T_184; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_186 = _signals_T_17 ? 2'h1 : _signals_T_185; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_187 = _signals_T_15 ? 2'h1 : _signals_T_186; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_188 = _signals_T_13 ? 2'h1 : _signals_T_187; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_189 = _signals_T_11 ? 2'h1 : _signals_T_188; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_190 = _signals_T_9 ? 2'h1 : _signals_T_189; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_191 = _signals_T_7 ? 2'h2 : _signals_T_190; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_192 = _signals_T_5 ? 2'h1 : _signals_T_191; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_193 = _signals_T_3 ? 2'h1 : _signals_T_192; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_206 = _signals_T_73 ? 2'h2 : 2'h0; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_207 = _signals_T_71 ? 2'h2 : _signals_T_206; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_208 = _signals_T_69 ? 2'h2 : _signals_T_207; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_209 = _signals_T_67 ? 2'h2 : _signals_T_208; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_210 = _signals_T_65 ? 2'h2 : _signals_T_209; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_211 = _signals_T_63 ? 2'h2 : _signals_T_210; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_212 = _signals_T_61 ? 2'h2 : _signals_T_211; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_213 = _signals_T_59 ? 2'h2 : _signals_T_212; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_214 = _signals_T_57 ? 2'h2 : _signals_T_213; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_215 = _signals_T_55 ? 2'h2 : _signals_T_214; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_216 = _signals_T_53 ? 2'h1 : _signals_T_215; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_217 = _signals_T_51 ? 2'h1 : _signals_T_216; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_218 = _signals_T_49 ? 2'h1 : _signals_T_217; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_219 = _signals_T_47 ? 2'h1 : _signals_T_218; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_220 = _signals_T_45 ? 2'h1 : _signals_T_219; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_221 = _signals_T_43 ? 2'h1 : _signals_T_220; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_222 = _signals_T_41 ? 2'h1 : _signals_T_221; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_223 = _signals_T_39 ? 2'h1 : _signals_T_222; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_224 = _signals_T_37 ? 2'h1 : _signals_T_223; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_225 = _signals_T_35 ? 2'h1 : _signals_T_224; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_226 = _signals_T_33 ? 2'h1 : _signals_T_225; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_227 = _signals_T_31 ? 2'h1 : _signals_T_226; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_228 = _signals_T_29 ? 2'h1 : _signals_T_227; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_229 = _signals_T_27 ? 2'h1 : _signals_T_228; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_230 = _signals_T_25 ? 2'h1 : _signals_T_229; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_231 = _signals_T_23 ? 2'h1 : _signals_T_230; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_232 = _signals_T_21 ? 2'h1 : _signals_T_231; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_233 = _signals_T_19 ? 2'h1 : _signals_T_232; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_234 = _signals_T_17 ? 2'h1 : _signals_T_233; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_235 = _signals_T_15 ? 2'h1 : _signals_T_234; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_236 = _signals_T_13 ? 2'h1 : _signals_T_235; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_237 = _signals_T_11 ? 2'h1 : _signals_T_236; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_238 = _signals_T_9 ? 2'h1 : _signals_T_237; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_239 = _signals_T_7 ? 2'h1 : _signals_T_238; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_240 = _signals_T_5 ? 2'h1 : _signals_T_239; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_241 = _signals_T_3 ? 2'h1 : _signals_T_240; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_246 = _signals_T_89 ? 3'h6 : 3'h0; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_247 = _signals_T_87 ? 3'h6 : _signals_T_246; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_248 = _signals_T_85 ? 3'h6 : _signals_T_247; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_249 = _signals_T_83 ? 3'h0 : _signals_T_248; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_250 = _signals_T_81 ? 3'h0 : _signals_T_249; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_251 = _signals_T_79 ? 3'h0 : _signals_T_250; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_252 = _signals_T_77 ? 3'h0 : _signals_T_251; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_253 = _signals_T_75 ? 3'h0 : _signals_T_252; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_254 = _signals_T_73 ? 3'h0 : _signals_T_253; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_255 = _signals_T_71 ? 3'h0 : _signals_T_254; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_256 = _signals_T_69 ? 3'h0 : _signals_T_255; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_257 = _signals_T_67 ? 3'h0 : _signals_T_256; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_258 = _signals_T_65 ? 3'h0 : _signals_T_257; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_259 = _signals_T_63 ? 3'h0 : _signals_T_258; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_260 = _signals_T_61 ? 3'h0 : _signals_T_259; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_261 = _signals_T_59 ? 3'h0 : _signals_T_260; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_262 = _signals_T_57 ? 3'h0 : _signals_T_261; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_263 = _signals_T_55 ? 3'h0 : _signals_T_262; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_264 = _signals_T_53 ? 3'h1 : _signals_T_263; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_265 = _signals_T_51 ? 3'h1 : _signals_T_264; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_266 = _signals_T_49 ? 3'h1 : _signals_T_265; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_267 = _signals_T_47 ? 3'h1 : _signals_T_266; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_268 = _signals_T_45 ? 3'h1 : _signals_T_267; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_269 = _signals_T_43 ? 3'h1 : _signals_T_268; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_270 = _signals_T_41 ? 3'h1 : _signals_T_269; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_271 = _signals_T_39 ? 3'h1 : _signals_T_270; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_272 = _signals_T_37 ? 3'h1 : _signals_T_271; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_273 = _signals_T_35 ? 3'h2 : _signals_T_272; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_274 = _signals_T_33 ? 3'h2 : _signals_T_273; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_275 = _signals_T_31 ? 3'h2 : _signals_T_274; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_276 = _signals_T_29 ? 3'h1 : _signals_T_275; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_277 = _signals_T_27 ? 3'h1 : _signals_T_276; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_278 = _signals_T_25 ? 3'h1 : _signals_T_277; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_279 = _signals_T_23 ? 3'h1 : _signals_T_278; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_280 = _signals_T_21 ? 3'h1 : _signals_T_279; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_281 = _signals_T_19 ? 3'h5 : _signals_T_280; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_282 = _signals_T_17 ? 3'h5 : _signals_T_281; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_283 = _signals_T_15 ? 3'h5 : _signals_T_282; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_284 = _signals_T_13 ? 3'h5 : _signals_T_283; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_285 = _signals_T_11 ? 3'h5 : _signals_T_284; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_286 = _signals_T_9 ? 3'h5 : _signals_T_285; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_287 = _signals_T_7 ? 3'h1 : _signals_T_286; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_288 = _signals_T_5 ? 3'h4 : _signals_T_287; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_289 = _signals_T_3 ? 3'h3 : _signals_T_288; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_297 = _signals_T_83 ? 4'ha : 4'hc; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_298 = _signals_T_81 ? 4'ha : _signals_T_297; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_299 = _signals_T_79 ? 4'ha : _signals_T_298; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_300 = _signals_T_77 ? 4'hc : _signals_T_299; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_301 = _signals_T_75 ? 4'hc : _signals_T_300; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_302 = _signals_T_73 ? 4'h2 : _signals_T_301; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_303 = _signals_T_71 ? 4'h3 : _signals_T_302; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_304 = _signals_T_69 ? 4'h9 : _signals_T_303; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_305 = _signals_T_67 ? 4'h8 : _signals_T_304; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_306 = _signals_T_65 ? 4'h4 : _signals_T_305; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_307 = _signals_T_63 ? 4'h7 : _signals_T_306; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_308 = _signals_T_61 ? 4'h5 : _signals_T_307; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_309 = _signals_T_59 ? 4'h6 : _signals_T_308; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_310 = _signals_T_57 ? 4'h1 : _signals_T_309; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_311 = _signals_T_55 ? 4'h0 : _signals_T_310; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_312 = _signals_T_53 ? 4'h9 : _signals_T_311; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_313 = _signals_T_51 ? 4'h8 : _signals_T_312; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_314 = _signals_T_49 ? 4'h6 : _signals_T_313; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_315 = _signals_T_47 ? 4'h2 : _signals_T_314; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_316 = _signals_T_45 ? 4'h3 : _signals_T_315; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_317 = _signals_T_43 ? 4'h4 : _signals_T_316; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_318 = _signals_T_41 ? 4'h7 : _signals_T_317; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_319 = _signals_T_39 ? 4'h5 : _signals_T_318; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_320 = _signals_T_37 ? 4'h0 : _signals_T_319; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_321 = _signals_T_35 ? 4'h0 : _signals_T_320; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_322 = _signals_T_33 ? 4'h0 : _signals_T_321; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_323 = _signals_T_31 ? 4'h0 : _signals_T_322; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_324 = _signals_T_29 ? 4'h0 : _signals_T_323; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_325 = _signals_T_27 ? 4'h0 : _signals_T_324; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_326 = _signals_T_25 ? 4'h0 : _signals_T_325; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_327 = _signals_T_23 ? 4'h0 : _signals_T_326; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_328 = _signals_T_21 ? 4'h0 : _signals_T_327; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_329 = _signals_T_19 ? 4'h0 : _signals_T_328; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_330 = _signals_T_17 ? 4'h0 : _signals_T_329; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_331 = _signals_T_15 ? 4'h0 : _signals_T_330; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_332 = _signals_T_13 ? 4'h0 : _signals_T_331; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_333 = _signals_T_11 ? 4'h0 : _signals_T_332; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_334 = _signals_T_9 ? 4'h0 : _signals_T_333; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_335 = _signals_T_7 ? 4'h0 : _signals_T_334; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_336 = _signals_T_5 ? 4'h0 : _signals_T_335; // @[Lookup.scala 34:39]
  wire [3:0] _signals_T_337 = _signals_T_3 ? 4'h0 : _signals_T_336; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_377 = _signals_T_19 ? 3'h4 : 3'h0; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_378 = _signals_T_17 ? 3'h1 : _signals_T_377; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_379 = _signals_T_15 ? 3'h5 : _signals_T_378; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_380 = _signals_T_13 ? 3'h2 : _signals_T_379; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_381 = _signals_T_11 ? 3'h6 : _signals_T_380; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_382 = _signals_T_9 ? 3'h3 : _signals_T_381; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_383 = _signals_T_7 ? 3'h0 : _signals_T_382; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_384 = _signals_T_5 ? 3'h0 : _signals_T_383; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_385 = _signals_T_3 ? 3'h0 : _signals_T_384; // @[Lookup.scala 34:39]
  wire  _signals_T_388 = _signals_T_93 ? 1'h0 : _signals_T_95; // @[Lookup.scala 34:39]
  wire  _signals_T_389 = _signals_T_91 ? 1'h0 : _signals_T_388; // @[Lookup.scala 34:39]
  wire  _signals_T_397 = _signals_T_75 ? 1'h0 : _signals_T_77 | (_signals_T_79 | (_signals_T_81 | (_signals_T_83 | (
    _signals_T_85 | (_signals_T_87 | (_signals_T_89 | _signals_T_389)))))); // @[Lookup.scala 34:39]
  wire  _signals_T_398 = _signals_T_73 ? 1'h0 : _signals_T_397; // @[Lookup.scala 34:39]
  wire  _signals_T_399 = _signals_T_71 ? 1'h0 : _signals_T_398; // @[Lookup.scala 34:39]
  wire  _signals_T_400 = _signals_T_69 ? 1'h0 : _signals_T_399; // @[Lookup.scala 34:39]
  wire  _signals_T_401 = _signals_T_67 ? 1'h0 : _signals_T_400; // @[Lookup.scala 34:39]
  wire  _signals_T_402 = _signals_T_65 ? 1'h0 : _signals_T_401; // @[Lookup.scala 34:39]
  wire  _signals_T_403 = _signals_T_63 ? 1'h0 : _signals_T_402; // @[Lookup.scala 34:39]
  wire  _signals_T_404 = _signals_T_61 ? 1'h0 : _signals_T_403; // @[Lookup.scala 34:39]
  wire  _signals_T_405 = _signals_T_59 ? 1'h0 : _signals_T_404; // @[Lookup.scala 34:39]
  wire  _signals_T_406 = _signals_T_57 ? 1'h0 : _signals_T_405; // @[Lookup.scala 34:39]
  wire  _signals_T_407 = _signals_T_55 ? 1'h0 : _signals_T_406; // @[Lookup.scala 34:39]
  wire  _signals_T_408 = _signals_T_53 ? 1'h0 : _signals_T_407; // @[Lookup.scala 34:39]
  wire  _signals_T_409 = _signals_T_51 ? 1'h0 : _signals_T_408; // @[Lookup.scala 34:39]
  wire  _signals_T_410 = _signals_T_49 ? 1'h0 : _signals_T_409; // @[Lookup.scala 34:39]
  wire  _signals_T_411 = _signals_T_47 ? 1'h0 : _signals_T_410; // @[Lookup.scala 34:39]
  wire  _signals_T_412 = _signals_T_45 ? 1'h0 : _signals_T_411; // @[Lookup.scala 34:39]
  wire  _signals_T_413 = _signals_T_43 ? 1'h0 : _signals_T_412; // @[Lookup.scala 34:39]
  wire  _signals_T_414 = _signals_T_41 ? 1'h0 : _signals_T_413; // @[Lookup.scala 34:39]
  wire  _signals_T_415 = _signals_T_39 ? 1'h0 : _signals_T_414; // @[Lookup.scala 34:39]
  wire  _signals_T_416 = _signals_T_37 ? 1'h0 : _signals_T_415; // @[Lookup.scala 34:39]
  wire  _signals_T_417 = _signals_T_35 ? 1'h0 : _signals_T_416; // @[Lookup.scala 34:39]
  wire  _signals_T_418 = _signals_T_33 ? 1'h0 : _signals_T_417; // @[Lookup.scala 34:39]
  wire  _signals_T_419 = _signals_T_31 ? 1'h0 : _signals_T_418; // @[Lookup.scala 34:39]
  wire  _signals_T_425 = _signals_T_19 ? 1'h0 : _signals_T_21 | (_signals_T_23 | (_signals_T_25 | (_signals_T_27 | (
    _signals_T_29 | _signals_T_419)))); // @[Lookup.scala 34:39]
  wire  _signals_T_426 = _signals_T_17 ? 1'h0 : _signals_T_425; // @[Lookup.scala 34:39]
  wire  _signals_T_427 = _signals_T_15 ? 1'h0 : _signals_T_426; // @[Lookup.scala 34:39]
  wire  _signals_T_428 = _signals_T_13 ? 1'h0 : _signals_T_427; // @[Lookup.scala 34:39]
  wire  _signals_T_429 = _signals_T_11 ? 1'h0 : _signals_T_428; // @[Lookup.scala 34:39]
  wire  _signals_T_430 = _signals_T_9 ? 1'h0 : _signals_T_429; // @[Lookup.scala 34:39]
  wire  _signals_T_433 = _signals_T_3 ? 1'h0 : _signals_T_5 | (_signals_T_7 | _signals_T_430); // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_466 = _signals_T_33 ? 2'h2 : {{1'd0}, _signals_T_35}; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_467 = _signals_T_31 ? 2'h3 : _signals_T_466; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_468 = _signals_T_29 ? 2'h0 : _signals_T_467; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_469 = _signals_T_27 ? 2'h0 : _signals_T_468; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_470 = _signals_T_25 ? 2'h0 : _signals_T_469; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_471 = _signals_T_23 ? 2'h0 : _signals_T_470; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_472 = _signals_T_21 ? 2'h0 : _signals_T_471; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_473 = _signals_T_19 ? 2'h0 : _signals_T_472; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_474 = _signals_T_17 ? 2'h0 : _signals_T_473; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_475 = _signals_T_15 ? 2'h0 : _signals_T_474; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_476 = _signals_T_13 ? 2'h0 : _signals_T_475; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_477 = _signals_T_11 ? 2'h0 : _signals_T_476; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_478 = _signals_T_9 ? 2'h0 : _signals_T_477; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_479 = _signals_T_7 ? 2'h0 : _signals_T_478; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_480 = _signals_T_5 ? 2'h0 : _signals_T_479; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_481 = _signals_T_3 ? 2'h0 : _signals_T_480; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_516 = _signals_T_29 ? 3'h4 : 3'h0; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_517 = _signals_T_27 ? 3'h5 : _signals_T_516; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_518 = _signals_T_25 ? 3'h1 : _signals_T_517; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_519 = _signals_T_23 ? 3'h2 : _signals_T_518; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_520 = _signals_T_21 ? 3'h3 : _signals_T_519; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_521 = _signals_T_19 ? 3'h0 : _signals_T_520; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_522 = _signals_T_17 ? 3'h0 : _signals_T_521; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_523 = _signals_T_15 ? 3'h0 : _signals_T_522; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_524 = _signals_T_13 ? 3'h0 : _signals_T_523; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_525 = _signals_T_11 ? 3'h0 : _signals_T_524; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_526 = _signals_T_9 ? 3'h0 : _signals_T_525; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_527 = _signals_T_7 ? 3'h0 : _signals_T_526; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_528 = _signals_T_5 ? 3'h0 : _signals_T_527; // @[Lookup.scala 34:39]
  wire [2:0] _signals_T_529 = _signals_T_3 ? 3'h0 : _signals_T_528; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_531 = _signals_T_95 ? 2'h3 : 2'h0; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_532 = _signals_T_93 ? 2'h3 : _signals_T_531; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_533 = _signals_T_91 ? 2'h3 : _signals_T_532; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_534 = _signals_T_89 ? 2'h3 : _signals_T_533; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_535 = _signals_T_87 ? 2'h3 : _signals_T_534; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_536 = _signals_T_85 ? 2'h3 : _signals_T_535; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_537 = _signals_T_83 ? 2'h3 : _signals_T_536; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_538 = _signals_T_81 ? 2'h3 : _signals_T_537; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_539 = _signals_T_79 ? 2'h3 : _signals_T_538; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_540 = _signals_T_77 ? 2'h0 : _signals_T_539; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_541 = _signals_T_75 ? 2'h0 : _signals_T_540; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_542 = _signals_T_73 ? 2'h0 : _signals_T_541; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_543 = _signals_T_71 ? 2'h0 : _signals_T_542; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_544 = _signals_T_69 ? 2'h0 : _signals_T_543; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_545 = _signals_T_67 ? 2'h0 : _signals_T_544; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_546 = _signals_T_65 ? 2'h0 : _signals_T_545; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_547 = _signals_T_63 ? 2'h0 : _signals_T_546; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_548 = _signals_T_61 ? 2'h0 : _signals_T_547; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_549 = _signals_T_59 ? 2'h0 : _signals_T_548; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_550 = _signals_T_57 ? 2'h0 : _signals_T_549; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_551 = _signals_T_55 ? 2'h0 : _signals_T_550; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_552 = _signals_T_53 ? 2'h0 : _signals_T_551; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_553 = _signals_T_51 ? 2'h0 : _signals_T_552; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_554 = _signals_T_49 ? 2'h0 : _signals_T_553; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_555 = _signals_T_47 ? 2'h0 : _signals_T_554; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_556 = _signals_T_45 ? 2'h0 : _signals_T_555; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_557 = _signals_T_43 ? 2'h0 : _signals_T_556; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_558 = _signals_T_41 ? 2'h0 : _signals_T_557; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_559 = _signals_T_39 ? 2'h0 : _signals_T_558; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_560 = _signals_T_37 ? 2'h0 : _signals_T_559; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_561 = _signals_T_35 ? 2'h0 : _signals_T_560; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_562 = _signals_T_33 ? 2'h0 : _signals_T_561; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_563 = _signals_T_31 ? 2'h0 : _signals_T_562; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_564 = _signals_T_29 ? 2'h1 : _signals_T_563; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_565 = _signals_T_27 ? 2'h1 : _signals_T_564; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_566 = _signals_T_25 ? 2'h1 : _signals_T_565; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_567 = _signals_T_23 ? 2'h1 : _signals_T_566; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_568 = _signals_T_21 ? 2'h1 : _signals_T_567; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_569 = _signals_T_19 ? 2'h0 : _signals_T_568; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_570 = _signals_T_17 ? 2'h0 : _signals_T_569; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_571 = _signals_T_15 ? 2'h0 : _signals_T_570; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_572 = _signals_T_13 ? 2'h0 : _signals_T_571; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_573 = _signals_T_11 ? 2'h0 : _signals_T_572; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_574 = _signals_T_9 ? 2'h0 : _signals_T_573; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_575 = _signals_T_7 ? 2'h2 : _signals_T_574; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_576 = _signals_T_5 ? 2'h2 : _signals_T_575; // @[Lookup.scala 34:39]
  wire [1:0] _signals_T_577 = _signals_T_3 ? 2'h0 : _signals_T_576; // @[Lookup.scala 34:39]
  wire  _signals_T_588 = _signals_T_77 ? 1'h0 : _signals_T_79 | (_signals_T_81 | (_signals_T_83 | (_signals_T_85 | (
    _signals_T_87 | _signals_T_89)))); // @[Lookup.scala 34:39]
  wire  _signals_T_589 = _signals_T_75 ? 1'h0 : _signals_T_588; // @[Lookup.scala 34:39]
  wire  _signals_T_609 = _signals_T_35 ? 1'h0 : _signals_T_37 | (_signals_T_39 | (_signals_T_41 | (_signals_T_43 | (
    _signals_T_45 | (_signals_T_47 | (_signals_T_49 | (_signals_T_51 | (_signals_T_53 | (_signals_T_55 | (_signals_T_57
     | (_signals_T_59 | (_signals_T_61 | (_signals_T_63 | (_signals_T_65 | (_signals_T_67 | (_signals_T_69 | (
    _signals_T_71 | (_signals_T_73 | _signals_T_589)))))))))))))))))); // @[Lookup.scala 34:39]
  wire  _signals_T_610 = _signals_T_33 ? 1'h0 : _signals_T_609; // @[Lookup.scala 34:39]
  wire  _signals_T_611 = _signals_T_31 ? 1'h0 : _signals_T_610; // @[Lookup.scala 34:39]
  wire  _signals_T_617 = _signals_T_19 ? 1'h0 : _signals_T_21 | (_signals_T_23 | (_signals_T_25 | (_signals_T_27 | (
    _signals_T_29 | _signals_T_611)))); // @[Lookup.scala 34:39]
  wire  _signals_T_618 = _signals_T_17 ? 1'h0 : _signals_T_617; // @[Lookup.scala 34:39]
  wire  _signals_T_619 = _signals_T_15 ? 1'h0 : _signals_T_618; // @[Lookup.scala 34:39]
  wire  _signals_T_620 = _signals_T_13 ? 1'h0 : _signals_T_619; // @[Lookup.scala 34:39]
  wire  _signals_T_621 = _signals_T_11 ? 1'h0 : _signals_T_620; // @[Lookup.scala 34:39]
  wire  _signals_T_622 = _signals_T_9 ? 1'h0 : _signals_T_621; // @[Lookup.scala 34:39]
  assign io_pc_sel = _signals_T_1 ? 2'h0 : _signals_T_145; // @[Lookup.scala 34:39]
  assign io_inst_kill = _signals_T_1 ? 1'h0 : _signals_T_433; // @[Lookup.scala 34:39]
  assign io_a_sel = _signals_T_1 ? 2'h1 : _signals_T_193; // @[Lookup.scala 34:39]
  assign io_b_sel = _signals_T_1 ? 2'h1 : _signals_T_241; // @[Lookup.scala 34:39]
  assign io_imm_sel = _signals_T_1 ? 3'h3 : _signals_T_289; // @[Lookup.scala 34:39]
  assign io_alu_op = _signals_T_1 ? 4'hb : _signals_T_337; // @[Lookup.scala 34:39]
  assign io_br_type = _signals_T_1 ? 3'h0 : _signals_T_385; // @[Lookup.scala 34:39]
  assign io_st_type = _signals_T_1 ? 2'h0 : _signals_T_481; // @[Lookup.scala 34:39]
  assign io_ld_type = _signals_T_1 ? 3'h0 : _signals_T_529; // @[Lookup.scala 34:39]
  assign io_wb_sel = _signals_T_1 ? 2'h0 : _signals_T_577; // @[Lookup.scala 34:39]
  assign io_wb_en = _signals_T_1 | (_signals_T_3 | (_signals_T_5 | (_signals_T_7 | _signals_T_622))); // @[Lookup.scala 34:39]
endmodule
module Core(
  input         clock,
  input         reset,
  output [31:0] io_imem_addr,
  input  [31:0] io_imem_rdata,
  output        io_dmem_req,
  output [31:0] io_dmem_addr,
  output        io_dmem_we,
  output [3:0]  io_dmem_be,
  output [31:0] io_dmem_wdata,
  input         io_dmem_gnt,
  input         io_dmem_rvalid,
  input  [31:0] io_dmem_rdata
);
  wire  dpath_clock; // @[Core.scala 38:21]
  wire  dpath_reset; // @[Core.scala 38:21]
  wire [31:0] dpath_io_imem_addr; // @[Core.scala 38:21]
  wire [31:0] dpath_io_imem_rdata; // @[Core.scala 38:21]
  wire  dpath_io_dmem_req; // @[Core.scala 38:21]
  wire [31:0] dpath_io_dmem_addr; // @[Core.scala 38:21]
  wire  dpath_io_dmem_we; // @[Core.scala 38:21]
  wire [3:0] dpath_io_dmem_be; // @[Core.scala 38:21]
  wire [31:0] dpath_io_dmem_wdata; // @[Core.scala 38:21]
  wire  dpath_io_dmem_gnt; // @[Core.scala 38:21]
  wire  dpath_io_dmem_rvalid; // @[Core.scala 38:21]
  wire [31:0] dpath_io_dmem_rdata; // @[Core.scala 38:21]
  wire [31:0] dpath_io_ctrl_inst; // @[Core.scala 38:21]
  wire [1:0] dpath_io_ctrl_pc_sel; // @[Core.scala 38:21]
  wire  dpath_io_ctrl_inst_kill; // @[Core.scala 38:21]
  wire [1:0] dpath_io_ctrl_a_sel; // @[Core.scala 38:21]
  wire [1:0] dpath_io_ctrl_b_sel; // @[Core.scala 38:21]
  wire [2:0] dpath_io_ctrl_imm_sel; // @[Core.scala 38:21]
  wire [3:0] dpath_io_ctrl_alu_op; // @[Core.scala 38:21]
  wire [2:0] dpath_io_ctrl_br_type; // @[Core.scala 38:21]
  wire [1:0] dpath_io_ctrl_st_type; // @[Core.scala 38:21]
  wire [2:0] dpath_io_ctrl_ld_type; // @[Core.scala 38:21]
  wire [1:0] dpath_io_ctrl_wb_sel; // @[Core.scala 38:21]
  wire  dpath_io_ctrl_wb_en; // @[Core.scala 38:21]
  wire [31:0] ctrl_io_inst; // @[Core.scala 39:20]
  wire [1:0] ctrl_io_pc_sel; // @[Core.scala 39:20]
  wire  ctrl_io_inst_kill; // @[Core.scala 39:20]
  wire [1:0] ctrl_io_a_sel; // @[Core.scala 39:20]
  wire [1:0] ctrl_io_b_sel; // @[Core.scala 39:20]
  wire [2:0] ctrl_io_imm_sel; // @[Core.scala 39:20]
  wire [3:0] ctrl_io_alu_op; // @[Core.scala 39:20]
  wire [2:0] ctrl_io_br_type; // @[Core.scala 39:20]
  wire [1:0] ctrl_io_st_type; // @[Core.scala 39:20]
  wire [2:0] ctrl_io_ld_type; // @[Core.scala 39:20]
  wire [1:0] ctrl_io_wb_sel; // @[Core.scala 39:20]
  wire  ctrl_io_wb_en; // @[Core.scala 39:20]
  Datapath dpath ( // @[Core.scala 38:21]
    .clock(dpath_clock),
    .reset(dpath_reset),
    .io_imem_addr(dpath_io_imem_addr),
    .io_imem_rdata(dpath_io_imem_rdata),
    .io_dmem_req(dpath_io_dmem_req),
    .io_dmem_addr(dpath_io_dmem_addr),
    .io_dmem_we(dpath_io_dmem_we),
    .io_dmem_be(dpath_io_dmem_be),
    .io_dmem_wdata(dpath_io_dmem_wdata),
    .io_dmem_gnt(dpath_io_dmem_gnt),
    .io_dmem_rvalid(dpath_io_dmem_rvalid),
    .io_dmem_rdata(dpath_io_dmem_rdata),
    .io_ctrl_inst(dpath_io_ctrl_inst),
    .io_ctrl_pc_sel(dpath_io_ctrl_pc_sel),
    .io_ctrl_inst_kill(dpath_io_ctrl_inst_kill),
    .io_ctrl_a_sel(dpath_io_ctrl_a_sel),
    .io_ctrl_b_sel(dpath_io_ctrl_b_sel),
    .io_ctrl_imm_sel(dpath_io_ctrl_imm_sel),
    .io_ctrl_alu_op(dpath_io_ctrl_alu_op),
    .io_ctrl_br_type(dpath_io_ctrl_br_type),
    .io_ctrl_st_type(dpath_io_ctrl_st_type),
    .io_ctrl_ld_type(dpath_io_ctrl_ld_type),
    .io_ctrl_wb_sel(dpath_io_ctrl_wb_sel),
    .io_ctrl_wb_en(dpath_io_ctrl_wb_en)
  );
  Control ctrl ( // @[Core.scala 39:20]
    .io_inst(ctrl_io_inst),
    .io_pc_sel(ctrl_io_pc_sel),
    .io_inst_kill(ctrl_io_inst_kill),
    .io_a_sel(ctrl_io_a_sel),
    .io_b_sel(ctrl_io_b_sel),
    .io_imm_sel(ctrl_io_imm_sel),
    .io_alu_op(ctrl_io_alu_op),
    .io_br_type(ctrl_io_br_type),
    .io_st_type(ctrl_io_st_type),
    .io_ld_type(ctrl_io_ld_type),
    .io_wb_sel(ctrl_io_wb_sel),
    .io_wb_en(ctrl_io_wb_en)
  );
  assign io_imem_addr = dpath_io_imem_addr; // @[Core.scala 42:17]
  assign io_dmem_req = dpath_io_dmem_req; // @[Core.scala 41:17]
  assign io_dmem_addr = dpath_io_dmem_addr; // @[Core.scala 41:17]
  assign io_dmem_we = dpath_io_dmem_we; // @[Core.scala 41:17]
  assign io_dmem_be = dpath_io_dmem_be; // @[Core.scala 41:17]
  assign io_dmem_wdata = dpath_io_dmem_wdata; // @[Core.scala 41:17]
  assign dpath_clock = clock;
  assign dpath_reset = reset;
  assign dpath_io_imem_rdata = io_imem_rdata; // @[Core.scala 42:17]
  assign dpath_io_dmem_gnt = io_dmem_gnt; // @[Core.scala 41:17]
  assign dpath_io_dmem_rvalid = io_dmem_rvalid; // @[Core.scala 41:17]
  assign dpath_io_dmem_rdata = io_dmem_rdata; // @[Core.scala 41:17]
  assign dpath_io_ctrl_pc_sel = ctrl_io_pc_sel; // @[Core.scala 43:17]
  assign dpath_io_ctrl_inst_kill = ctrl_io_inst_kill; // @[Core.scala 43:17]
  assign dpath_io_ctrl_a_sel = ctrl_io_a_sel; // @[Core.scala 43:17]
  assign dpath_io_ctrl_b_sel = ctrl_io_b_sel; // @[Core.scala 43:17]
  assign dpath_io_ctrl_imm_sel = ctrl_io_imm_sel; // @[Core.scala 43:17]
  assign dpath_io_ctrl_alu_op = ctrl_io_alu_op; // @[Core.scala 43:17]
  assign dpath_io_ctrl_br_type = ctrl_io_br_type; // @[Core.scala 43:17]
  assign dpath_io_ctrl_st_type = ctrl_io_st_type; // @[Core.scala 43:17]
  assign dpath_io_ctrl_ld_type = ctrl_io_ld_type; // @[Core.scala 43:17]
  assign dpath_io_ctrl_wb_sel = ctrl_io_wb_sel; // @[Core.scala 43:17]
  assign dpath_io_ctrl_wb_en = ctrl_io_wb_en; // @[Core.scala 43:17]
  assign ctrl_io_inst = dpath_io_ctrl_inst; // @[Core.scala 43:17]
endmodule
module SimpleBus(
  input         clock,
  input         reset,
  input         io_host_0_req,
  input  [31:0] io_host_0_addr,
  input         io_host_0_we,
  input  [3:0]  io_host_0_be,
  input  [31:0] io_host_0_wdata,
  output        io_host_0_gnt,
  output        io_host_0_rvalid,
  output [31:0] io_host_0_rdata,
  output        io_dev_0_req,
  output [31:0] io_dev_0_addr,
  output        io_dev_0_we,
  output [3:0]  io_dev_0_be,
  output [31:0] io_dev_0_wdata,
  input         io_dev_0_rvalid,
  input  [31:0] io_dev_0_rdata,
  output        io_dev_1_req,
  output [31:0] io_dev_1_addr,
  input         io_dev_1_rvalid,
  input  [31:0] io_dev_1_rdata,
  output        io_dev_2_req,
  output [31:0] io_dev_2_addr,
  output        io_dev_2_we,
  output [31:0] io_dev_2_wdata,
  input         io_dev_2_rvalid,
  input  [31:0] io_dev_2_rdata
);
`ifdef RANDOMIZE_REG_INIT
  reg [31:0] _RAND_0;
`endif // RANDOMIZE_REG_INIT
  wire [31:0] _T_4 = io_host_0_addr & 32'hfffffc00; // @[Bus.scala 65:37]
  wire [1:0] devSelReq = _T_4 == 32'h20000 ? 2'h2 : {{1'd0}, _T_4 == 32'h10000}; // @[Bus.scala 65:76 66:17]
  reg [1:0] devSelResp; // @[Bus.scala 72:27]
  wire  _GEN_20 = 2'h1 == devSelResp ? io_dev_1_rvalid : io_dev_0_rvalid; // @[Bus.scala 93:{28,28}]
  wire [31:0] _GEN_26 = 2'h1 == devSelResp ? io_dev_1_rdata : io_dev_0_rdata; // @[Bus.scala 95:{27,27}]
  assign io_host_0_gnt = io_host_0_req; // @[Bus.scala 103:27]
  assign io_host_0_rvalid = 2'h2 == devSelResp ? io_dev_2_rvalid : _GEN_20; // @[Bus.scala 93:{28,28}]
  assign io_host_0_rdata = 2'h2 == devSelResp ? io_dev_2_rdata : _GEN_26; // @[Bus.scala 95:{27,27}]
  assign io_dev_0_req = 2'h0 == devSelReq & io_host_0_req; // @[Bus.scala 75:32 76:23 82:23]
  assign io_dev_0_addr = 2'h0 == devSelReq ? io_host_0_addr : 32'h0; // @[Bus.scala 75:32 78:24 84:24]
  assign io_dev_0_we = 2'h0 == devSelReq & io_host_0_we; // @[Bus.scala 75:32 77:22 83:22]
  assign io_dev_0_be = 2'h0 == devSelReq ? io_host_0_be : 4'h0; // @[Bus.scala 75:32 80:22 86:22]
  assign io_dev_0_wdata = 2'h0 == devSelReq ? io_host_0_wdata : 32'h0; // @[Bus.scala 75:32 79:25 85:25]
  assign io_dev_1_req = 2'h1 == devSelReq & io_host_0_req; // @[Bus.scala 75:32 76:23 82:23]
  assign io_dev_1_addr = 2'h1 == devSelReq ? io_host_0_addr : 32'h0; // @[Bus.scala 75:32 78:24 84:24]
  assign io_dev_2_req = 2'h2 == devSelReq & io_host_0_req; // @[Bus.scala 75:32 76:23 82:23]
  assign io_dev_2_addr = 2'h2 == devSelReq ? io_host_0_addr : 32'h0; // @[Bus.scala 75:32 78:24 84:24]
  assign io_dev_2_we = 2'h2 == devSelReq & io_host_0_we; // @[Bus.scala 75:32 77:22 83:22]
  assign io_dev_2_wdata = 2'h2 == devSelReq ? io_host_0_wdata : 32'h0; // @[Bus.scala 75:32 79:25 85:25]
  always @(posedge clock) begin
    if (reset) begin // @[Bus.scala 72:27]
      devSelResp <= 2'h0; // @[Bus.scala 72:27]
    end else if (_T_4 == 32'h20000) begin // @[Bus.scala 65:76]
      devSelResp <= 2'h2; // @[Bus.scala 66:17]
    end else begin
      devSelResp <= {{1'd0}, _T_4 == 32'h10000};
    end
  end
// Register and memory initialization
`ifdef RANDOMIZE_GARBAGE_ASSIGN
`define RANDOMIZE
`endif
`ifdef RANDOMIZE_INVALID_ASSIGN
`define RANDOMIZE
`endif
`ifdef RANDOMIZE_REG_INIT
`define RANDOMIZE
`endif
`ifdef RANDOMIZE_MEM_INIT
`define RANDOMIZE
`endif
`ifndef RANDOM
`define RANDOM $random
`endif
`ifdef RANDOMIZE_MEM_INIT
  integer initvar;
`endif
`ifndef SYNTHESIS
`ifdef FIRRTL_BEFORE_INITIAL
`FIRRTL_BEFORE_INITIAL
`endif
initial begin
  `ifdef RANDOMIZE
    `ifdef INIT_RANDOM
      `INIT_RANDOM
    `endif
    `ifndef VERILATOR
      `ifdef RANDOMIZE_DELAY
        #`RANDOMIZE_DELAY begin end
      `else
        #0.002 begin end
      `endif
    `endif
`ifdef RANDOMIZE_REG_INIT
  _RAND_0 = {1{`RANDOM}};
  devSelResp = _RAND_0[1:0];
`endif // RANDOMIZE_REG_INIT
  `endif // RANDOMIZE
end // initial
`ifdef FIRRTL_AFTER_INITIAL
`FIRRTL_AFTER_INITIAL
`endif
`endif // SYNTHESIS
endmodule
module Ram(
  input         clock,
  input  [31:0] io_imem_addr,
  output [31:0] io_imem_rdata,
  input         io_dmem_req,
  input  [31:0] io_dmem_addr,
  input         io_dmem_we,
  input  [3:0]  io_dmem_be,
  input  [31:0] io_dmem_wdata,
  output        io_dmem_rvalid,
  output [31:0] io_dmem_rdata
);
`ifdef RANDOMIZE_REG_INIT
  reg [31:0] _RAND_0;
  reg [31:0] _RAND_1;
  reg [31:0] _RAND_2;
  reg [31:0] _RAND_3;
  reg [31:0] _RAND_4;
  reg [31:0] _RAND_5;
  reg [31:0] _RAND_6;
  reg [31:0] _RAND_7;
  reg [31:0] _RAND_8;
  reg [31:0] _RAND_9;
  reg [31:0] _RAND_10;
  reg [31:0] _RAND_11;
  reg [31:0] _RAND_12;
`endif // RANDOMIZE_REG_INIT
  reg [31:0] mem [0:4095] /*verilator public*/; // @[Ram.scala 27:24]
  wire  mem_io_imem_rdata_MPORT_en; // @[Ram.scala 27:24]
  wire [11:0] mem_io_imem_rdata_MPORT_addr; // @[Ram.scala 27:24]
  wire [31:0] mem_io_imem_rdata_MPORT_data; // @[Ram.scala 27:24]
  wire  mem_io_dmem_rdata_MPORT_en; // @[Ram.scala 27:24]
  wire [11:0] mem_io_dmem_rdata_MPORT_addr; // @[Ram.scala 27:24]
  wire [31:0] mem_io_dmem_rdata_MPORT_data; // @[Ram.scala 27:24]
  wire  mem_write_MPORT_en; // @[Ram.scala 27:24]
  wire [11:0] mem_write_MPORT_addr; // @[Ram.scala 27:24]
  wire [31:0] mem_write_MPORT_data; // @[Ram.scala 27:24]
  wire  mem_write_MPORT_1_en; // @[Ram.scala 27:24]
  wire [11:0] mem_write_MPORT_1_addr; // @[Ram.scala 27:24]
  wire [31:0] mem_write_MPORT_1_data; // @[Ram.scala 27:24]
  wire  mem_write_MPORT_2_en; // @[Ram.scala 27:24]
  wire [11:0] mem_write_MPORT_2_addr; // @[Ram.scala 27:24]
  wire [31:0] mem_write_MPORT_2_data; // @[Ram.scala 27:24]
  wire  mem_write_MPORT_3_en; // @[Ram.scala 27:24]
  wire [11:0] mem_write_MPORT_3_addr; // @[Ram.scala 27:24]
  wire [31:0] mem_write_MPORT_3_data; // @[Ram.scala 27:24]
  wire [31:0] mem_MPORT_data; // @[Ram.scala 27:24]
  wire [11:0] mem_MPORT_addr; // @[Ram.scala 27:24]
  wire  mem_MPORT_mask; // @[Ram.scala 27:24]
  wire  mem_MPORT_en; // @[Ram.scala 27:24]
  reg  mem_io_imem_rdata_MPORT_en_pipe_0;
  reg [11:0] mem_io_imem_rdata_MPORT_addr_pipe_0;
  reg  mem_io_dmem_rdata_MPORT_en_pipe_0;
  reg [11:0] mem_io_dmem_rdata_MPORT_addr_pipe_0;
  reg  mem_write_MPORT_en_pipe_0;
  reg [11:0] mem_write_MPORT_addr_pipe_0;
  reg  mem_write_MPORT_1_en_pipe_0;
  reg [11:0] mem_write_MPORT_1_addr_pipe_0;
  reg  mem_write_MPORT_2_en_pipe_0;
  reg [11:0] mem_write_MPORT_2_addr_pipe_0;
  reg  mem_write_MPORT_3_en_pipe_0;
  reg [11:0] mem_write_MPORT_3_addr_pipe_0;
  reg  drvalid; // @[Ram.scala 35:24]
  wire [16:0] iaddr = io_imem_addr[18:2]; // @[Ram.scala 38:27]
  wire [16:0] daddr = io_dmem_addr[18:2]; // @[Ram.scala 39:27]
  wire  _write_T_1 = io_dmem_req & io_dmem_be[0]; // @[Ram.scala 46:19]
  wire [31:0] _write_T_3 = _write_T_1 ? io_dmem_wdata : mem_write_MPORT_data; // @[Ram.scala 45:9]
  wire [8:0] _write_T_5 = {{1'd0}, _write_T_3[7:0]}; // @[Ram.scala 49:31]
  wire [31:0] _write_T_6 = {{23'd0}, _write_T_5}; // @[Ram.scala 44:11]
  wire  _write_T_8 = io_dmem_req & io_dmem_be[1]; // @[Ram.scala 46:19]
  wire [31:0] _write_T_10 = _write_T_8 ? io_dmem_wdata : mem_write_MPORT_1_data; // @[Ram.scala 45:9]
  wire [15:0] _GEN_13 = {_write_T_10[15:8], 8'h0}; // @[Ram.scala 49:31]
  wire [22:0] _write_T_12 = {{7'd0}, _GEN_13}; // @[Ram.scala 49:31]
  wire [31:0] _GEN_14 = {{9'd0}, _write_T_12}; // @[Ram.scala 44:11]
  wire [31:0] _write_T_13 = _write_T_6 | _GEN_14; // @[Ram.scala 44:11]
  wire  _write_T_15 = io_dmem_req & io_dmem_be[2]; // @[Ram.scala 46:19]
  wire [31:0] _write_T_17 = _write_T_15 ? io_dmem_wdata : mem_write_MPORT_2_data; // @[Ram.scala 45:9]
  wire [23:0] _GEN_15 = {_write_T_17[23:16], 16'h0}; // @[Ram.scala 49:31]
  wire [38:0] _write_T_19 = {{15'd0}, _GEN_15}; // @[Ram.scala 49:31]
  wire [38:0] _GEN_16 = {{7'd0}, _write_T_13}; // @[Ram.scala 44:11]
  wire [38:0] _write_T_20 = _GEN_16 | _write_T_19; // @[Ram.scala 44:11]
  wire  _write_T_22 = io_dmem_req & io_dmem_be[3]; // @[Ram.scala 46:19]
  wire [31:0] _write_T_24 = _write_T_22 ? io_dmem_wdata : mem_write_MPORT_3_data; // @[Ram.scala 45:9]
  wire [31:0] _GEN_17 = {_write_T_24[31:24], 24'h0}; // @[Ram.scala 49:31]
  wire [38:0] _write_T_26 = {{7'd0}, _GEN_17}; // @[Ram.scala 49:31]
  wire [38:0] write = _write_T_20 | _write_T_26; // @[Ram.scala 44:11]
  assign mem_io_imem_rdata_MPORT_en = mem_io_imem_rdata_MPORT_en_pipe_0;
  assign mem_io_imem_rdata_MPORT_addr = mem_io_imem_rdata_MPORT_addr_pipe_0;
  assign mem_io_imem_rdata_MPORT_data = mem[mem_io_imem_rdata_MPORT_addr]; // @[Ram.scala 27:24]
  assign mem_io_dmem_rdata_MPORT_en = mem_io_dmem_rdata_MPORT_en_pipe_0;
  assign mem_io_dmem_rdata_MPORT_addr = mem_io_dmem_rdata_MPORT_addr_pipe_0;
  assign mem_io_dmem_rdata_MPORT_data = mem[mem_io_dmem_rdata_MPORT_addr]; // @[Ram.scala 27:24]
  assign mem_write_MPORT_en = mem_write_MPORT_en_pipe_0;
  assign mem_write_MPORT_addr = mem_write_MPORT_addr_pipe_0;
  assign mem_write_MPORT_data = mem[mem_write_MPORT_addr]; // @[Ram.scala 27:24]
  assign mem_write_MPORT_1_en = mem_write_MPORT_1_en_pipe_0;
  assign mem_write_MPORT_1_addr = mem_write_MPORT_1_addr_pipe_0;
  assign mem_write_MPORT_1_data = mem[mem_write_MPORT_1_addr]; // @[Ram.scala 27:24]
  assign mem_write_MPORT_2_en = mem_write_MPORT_2_en_pipe_0;
  assign mem_write_MPORT_2_addr = mem_write_MPORT_2_addr_pipe_0;
  assign mem_write_MPORT_2_data = mem[mem_write_MPORT_2_addr]; // @[Ram.scala 27:24]
  assign mem_write_MPORT_3_en = mem_write_MPORT_3_en_pipe_0;
  assign mem_write_MPORT_3_addr = mem_write_MPORT_3_addr_pipe_0;
  assign mem_write_MPORT_3_data = mem[mem_write_MPORT_3_addr]; // @[Ram.scala 27:24]
  assign mem_MPORT_data = write[31:0];
  assign mem_MPORT_addr = daddr[11:0];
  assign mem_MPORT_mask = 1'h1;
  assign mem_MPORT_en = io_dmem_req & io_dmem_we;
  assign io_imem_rdata = mem_io_imem_rdata_MPORT_data; // @[Ram.scala 40:17]
  assign io_dmem_rvalid = drvalid; // @[Ram.scala 36:18]
  assign io_dmem_rdata = mem_io_dmem_rdata_MPORT_data; // @[Ram.scala 41:17]
  always @(posedge clock) begin
    if (mem_MPORT_en & mem_MPORT_mask) begin
      mem[mem_MPORT_addr] <= mem_MPORT_data; // @[Ram.scala 27:24]
    end
    mem_io_imem_rdata_MPORT_en_pipe_0 <= 1'h1;
    if (1'h1) begin
      mem_io_imem_rdata_MPORT_addr_pipe_0 <= iaddr[11:0];
    end
    mem_io_dmem_rdata_MPORT_en_pipe_0 <= io_dmem_req;
    if (io_dmem_req) begin
      mem_io_dmem_rdata_MPORT_addr_pipe_0 <= daddr[11:0];
    end
    mem_write_MPORT_en_pipe_0 <= 1'h1;
    if (1'h1) begin
      mem_write_MPORT_addr_pipe_0 <= daddr[11:0];
    end
    mem_write_MPORT_1_en_pipe_0 <= 1'h1;
    if (1'h1) begin
      mem_write_MPORT_1_addr_pipe_0 <= daddr[11:0];
    end
    mem_write_MPORT_2_en_pipe_0 <= 1'h1;
    if (1'h1) begin
      mem_write_MPORT_2_addr_pipe_0 <= daddr[11:0];
    end
    mem_write_MPORT_3_en_pipe_0 <= 1'h1;
    if (1'h1) begin
      mem_write_MPORT_3_addr_pipe_0 <= daddr[11:0];
    end
    drvalid <= io_dmem_req; // @[Ram.scala 35:24]
  end
// Register and memory initialization
`ifdef RANDOMIZE_GARBAGE_ASSIGN
`define RANDOMIZE
`endif
`ifdef RANDOMIZE_INVALID_ASSIGN
`define RANDOMIZE
`endif
`ifdef RANDOMIZE_REG_INIT
`define RANDOMIZE
`endif
`ifdef RANDOMIZE_MEM_INIT
`define RANDOMIZE
`endif
`ifndef RANDOM
`define RANDOM $random
`endif
  integer initvar;
`ifndef SYNTHESIS
`ifdef FIRRTL_BEFORE_INITIAL
`FIRRTL_BEFORE_INITIAL
`endif
initial begin
  `ifdef RANDOMIZE
    `ifdef INIT_RANDOM
      `INIT_RANDOM
    `endif
    `ifndef VERILATOR
      `ifdef RANDOMIZE_DELAY
        #`RANDOMIZE_DELAY begin end
      `else
        #0.002 begin end
      `endif
    `endif
`ifdef RANDOMIZE_REG_INIT
  _RAND_0 = {1{`RANDOM}};
  mem_io_imem_rdata_MPORT_en_pipe_0 = _RAND_0[0:0];
  _RAND_1 = {1{`RANDOM}};
  mem_io_imem_rdata_MPORT_addr_pipe_0 = _RAND_1[11:0];
  _RAND_2 = {1{`RANDOM}};
  mem_io_dmem_rdata_MPORT_en_pipe_0 = _RAND_2[0:0];
  _RAND_3 = {1{`RANDOM}};
  mem_io_dmem_rdata_MPORT_addr_pipe_0 = _RAND_3[11:0];
  _RAND_4 = {1{`RANDOM}};
  mem_write_MPORT_en_pipe_0 = _RAND_4[0:0];
  _RAND_5 = {1{`RANDOM}};
  mem_write_MPORT_addr_pipe_0 = _RAND_5[11:0];
  _RAND_6 = {1{`RANDOM}};
  mem_write_MPORT_1_en_pipe_0 = _RAND_6[0:0];
  _RAND_7 = {1{`RANDOM}};
  mem_write_MPORT_1_addr_pipe_0 = _RAND_7[11:0];
  _RAND_8 = {1{`RANDOM}};
  mem_write_MPORT_2_en_pipe_0 = _RAND_8[0:0];
  _RAND_9 = {1{`RANDOM}};
  mem_write_MPORT_2_addr_pipe_0 = _RAND_9[11:0];
  _RAND_10 = {1{`RANDOM}};
  mem_write_MPORT_3_en_pipe_0 = _RAND_10[0:0];
  _RAND_11 = {1{`RANDOM}};
  mem_write_MPORT_3_addr_pipe_0 = _RAND_11[11:0];
  _RAND_12 = {1{`RANDOM}};
  drvalid = _RAND_12[0:0];
`endif // RANDOMIZE_REG_INIT
  `endif // RANDOMIZE
end // initial
`ifdef FIRRTL_AFTER_INITIAL
`FIRRTL_AFTER_INITIAL
`endif
`endif // SYNTHESIS
initial begin
  $readmemh("mem/riscvtest.hex", mem);
end
endmodule
module Timer(
  input         clock,
  input         reset,
  input         io_bus_req,
  input  [31:0] io_bus_addr,
  output        io_bus_rvalid,
  output [31:0] io_bus_rdata
);
`ifdef RANDOMIZE_REG_INIT
  reg [31:0] _RAND_0;
  reg [31:0] _RAND_1;
`endif // RANDOMIZE_REG_INIT
  reg [31:0] cyc_timer; // @[Timer.scala 23:26]
  wire [31:0] _cyc_timer_T_1 = cyc_timer + 32'h1; // @[Timer.scala 24:26]
  reg  io_bus_rvalid_REG; // @[Timer.scala 35:27]
  assign io_bus_rvalid = io_bus_rvalid_REG; // @[Timer.scala 35:17]
  assign io_bus_rdata = 16'h10 == io_bus_addr[15:0] ? cyc_timer : 32'h0; // @[Timer.scala 27:16 28:35 31:20]
  always @(posedge clock) begin
    if (reset) begin // @[Timer.scala 23:26]
      cyc_timer <= 32'h0; // @[Timer.scala 23:26]
    end else begin
      cyc_timer <= _cyc_timer_T_1; // @[Timer.scala 24:13]
    end
    io_bus_rvalid_REG <= io_bus_req; // @[Timer.scala 35:27]
  end
// Register and memory initialization
`ifdef RANDOMIZE_GARBAGE_ASSIGN
`define RANDOMIZE
`endif
`ifdef RANDOMIZE_INVALID_ASSIGN
`define RANDOMIZE
`endif
`ifdef RANDOMIZE_REG_INIT
`define RANDOMIZE
`endif
`ifdef RANDOMIZE_MEM_INIT
`define RANDOMIZE
`endif
`ifndef RANDOM
`define RANDOM $random
`endif
`ifdef RANDOMIZE_MEM_INIT
  integer initvar;
`endif
`ifndef SYNTHESIS
`ifdef FIRRTL_BEFORE_INITIAL
`FIRRTL_BEFORE_INITIAL
`endif
initial begin
  `ifdef RANDOMIZE
    `ifdef INIT_RANDOM
      `INIT_RANDOM
    `endif
    `ifndef VERILATOR
      `ifdef RANDOMIZE_DELAY
        #`RANDOMIZE_DELAY begin end
      `else
        #0.002 begin end
      `endif
    `endif
`ifdef RANDOMIZE_REG_INIT
  _RAND_0 = {1{`RANDOM}};
  cyc_timer = _RAND_0[31:0];
  _RAND_1 = {1{`RANDOM}};
  io_bus_rvalid_REG = _RAND_1[0:0];
`endif // RANDOMIZE_REG_INIT
  `endif // RANDOMIZE
end // initial
`ifdef FIRRTL_AFTER_INITIAL
`FIRRTL_AFTER_INITIAL
`endif
`endif // SYNTHESIS
endmodule
module Gpio(
  input         clock,
  input         reset,
  input         io_bus_req,
  input  [31:0] io_bus_addr,
  input         io_bus_we,
  input  [31:0] io_bus_wdata,
  output        io_bus_rvalid,
  output [31:0] io_bus_rdata,
  input         io_gpi_0,
  output        io_gpo_0,
  output        io_gpo_1,
  output        io_gpo_2
);
`ifdef RANDOMIZE_REG_INIT
  reg [31:0] _RAND_0;
  reg [31:0] _RAND_1;
  reg [31:0] _RAND_2;
`endif // RANDOMIZE_REG_INIT
  wire  we = io_bus_req & io_bus_we; // @[Gpio.scala 23:23]
  reg  inVal; // @[Gpio.scala 36:22]
  reg [2:0] outVal; // @[Gpio.scala 26:22]
  wire  outVal_rwe = we & io_bus_addr[16:0] == 17'h4; // @[Gpio.scala 27:18]
  wire [31:0] _GEN_0 = outVal_rwe ? io_bus_wdata : {{29'd0}, outVal}; // @[Gpio.scala 29:16 30:11 26:22]
  wire [2:0] _GEN_2 = 17'h4 == io_bus_addr[16:0] ? outVal : 3'h0; // @[Gpio.scala 40:16 41:35 48:20]
  wire [2:0] _GEN_4 = 17'h0 == io_bus_addr[16:0] ? {{2'd0}, inVal} : _GEN_2; // @[Gpio.scala 41:35 44:20]
  reg  io_bus_rvalid_REG; // @[Gpio.scala 54:27]
  wire [31:0] _GEN_5 = reset ? 32'h0 : _GEN_0; // @[Gpio.scala 26:{22,22}]
  assign io_bus_rvalid = io_bus_rvalid_REG; // @[Gpio.scala 54:17]
  assign io_bus_rdata = {{29'd0}, _GEN_4};
  assign io_gpo_0 = outVal[0]; // @[Gpio.scala 52:20]
  assign io_gpo_1 = outVal[1]; // @[Gpio.scala 52:20]
  assign io_gpo_2 = outVal[2]; // @[Gpio.scala 52:20]
  always @(posedge clock) begin
    if (reset) begin // @[Gpio.scala 36:22]
      inVal <= 1'h0; // @[Gpio.scala 36:22]
    end else begin
      inVal <= io_gpi_0; // @[Gpio.scala 36:22]
    end
    outVal <= _GEN_5[2:0]; // @[Gpio.scala 26:{22,22}]
    io_bus_rvalid_REG <= io_bus_req; // @[Gpio.scala 54:27]
  end
// Register and memory initialization
`ifdef RANDOMIZE_GARBAGE_ASSIGN
`define RANDOMIZE
`endif
`ifdef RANDOMIZE_INVALID_ASSIGN
`define RANDOMIZE
`endif
`ifdef RANDOMIZE_REG_INIT
`define RANDOMIZE
`endif
`ifdef RANDOMIZE_MEM_INIT
`define RANDOMIZE
`endif
`ifndef RANDOM
`define RANDOM $random
`endif
`ifdef RANDOMIZE_MEM_INIT
  integer initvar;
`endif
`ifndef SYNTHESIS
`ifdef FIRRTL_BEFORE_INITIAL
`FIRRTL_BEFORE_INITIAL
`endif
initial begin
  `ifdef RANDOMIZE
    `ifdef INIT_RANDOM
      `INIT_RANDOM
    `endif
    `ifndef VERILATOR
      `ifdef RANDOMIZE_DELAY
        #`RANDOMIZE_DELAY begin end
      `else
        #0.002 begin end
      `endif
    `endif
`ifdef RANDOMIZE_REG_INIT
  _RAND_0 = {1{`RANDOM}};
  inVal = _RAND_0[0:0];
  _RAND_1 = {1{`RANDOM}};
  outVal = _RAND_1[2:0];
  _RAND_2 = {1{`RANDOM}};
  io_bus_rvalid_REG = _RAND_2[0:0];
`endif // RANDOMIZE_REG_INIT
  `endif // RANDOMIZE
end // initial
`ifdef FIRRTL_AFTER_INITIAL
`FIRRTL_AFTER_INITIAL
`endif
`endif // SYNTHESIS
endmodule
module Soc(
  input   clock,
  input   reset,
  input   io_gpi_0,
  output  io_gpo_0,
  output  io_gpo_1,
  output  io_gpo_2
);
  wire  core_clock; // @[Soc.scala 38:20]
  wire  core_reset; // @[Soc.scala 38:20]
  wire [31:0] core_io_imem_addr; // @[Soc.scala 38:20]
  wire [31:0] core_io_imem_rdata; // @[Soc.scala 38:20]
  wire  core_io_dmem_req; // @[Soc.scala 38:20]
  wire [31:0] core_io_dmem_addr; // @[Soc.scala 38:20]
  wire  core_io_dmem_we; // @[Soc.scala 38:20]
  wire [3:0] core_io_dmem_be; // @[Soc.scala 38:20]
  wire [31:0] core_io_dmem_wdata; // @[Soc.scala 38:20]
  wire  core_io_dmem_gnt; // @[Soc.scala 38:20]
  wire  core_io_dmem_rvalid; // @[Soc.scala 38:20]
  wire [31:0] core_io_dmem_rdata; // @[Soc.scala 38:20]
  wire  bus_clock; // @[Soc.scala 65:19]
  wire  bus_reset; // @[Soc.scala 65:19]
  wire  bus_io_host_0_req; // @[Soc.scala 65:19]
  wire [31:0] bus_io_host_0_addr; // @[Soc.scala 65:19]
  wire  bus_io_host_0_we; // @[Soc.scala 65:19]
  wire [3:0] bus_io_host_0_be; // @[Soc.scala 65:19]
  wire [31:0] bus_io_host_0_wdata; // @[Soc.scala 65:19]
  wire  bus_io_host_0_gnt; // @[Soc.scala 65:19]
  wire  bus_io_host_0_rvalid; // @[Soc.scala 65:19]
  wire [31:0] bus_io_host_0_rdata; // @[Soc.scala 65:19]
  wire  bus_io_dev_0_req; // @[Soc.scala 65:19]
  wire [31:0] bus_io_dev_0_addr; // @[Soc.scala 65:19]
  wire  bus_io_dev_0_we; // @[Soc.scala 65:19]
  wire [3:0] bus_io_dev_0_be; // @[Soc.scala 65:19]
  wire [31:0] bus_io_dev_0_wdata; // @[Soc.scala 65:19]
  wire  bus_io_dev_0_rvalid; // @[Soc.scala 65:19]
  wire [31:0] bus_io_dev_0_rdata; // @[Soc.scala 65:19]
  wire  bus_io_dev_1_req; // @[Soc.scala 65:19]
  wire [31:0] bus_io_dev_1_addr; // @[Soc.scala 65:19]
  wire  bus_io_dev_1_rvalid; // @[Soc.scala 65:19]
  wire [31:0] bus_io_dev_1_rdata; // @[Soc.scala 65:19]
  wire  bus_io_dev_2_req; // @[Soc.scala 65:19]
  wire [31:0] bus_io_dev_2_addr; // @[Soc.scala 65:19]
  wire  bus_io_dev_2_we; // @[Soc.scala 65:19]
  wire [31:0] bus_io_dev_2_wdata; // @[Soc.scala 65:19]
  wire  bus_io_dev_2_rvalid; // @[Soc.scala 65:19]
  wire [31:0] bus_io_dev_2_rdata; // @[Soc.scala 65:19]
  wire  bus_io_dev_0_ram_clock; // @[Soc.scala 42:23]
  wire [31:0] bus_io_dev_0_ram_io_imem_addr; // @[Soc.scala 42:23]
  wire [31:0] bus_io_dev_0_ram_io_imem_rdata; // @[Soc.scala 42:23]
  wire  bus_io_dev_0_ram_io_dmem_req; // @[Soc.scala 42:23]
  wire [31:0] bus_io_dev_0_ram_io_dmem_addr; // @[Soc.scala 42:23]
  wire  bus_io_dev_0_ram_io_dmem_we; // @[Soc.scala 42:23]
  wire [3:0] bus_io_dev_0_ram_io_dmem_be; // @[Soc.scala 42:23]
  wire [31:0] bus_io_dev_0_ram_io_dmem_wdata; // @[Soc.scala 42:23]
  wire  bus_io_dev_0_ram_io_dmem_rvalid; // @[Soc.scala 42:23]
  wire [31:0] bus_io_dev_0_ram_io_dmem_rdata; // @[Soc.scala 42:23]
  wire  bus_io_dev_1_timer_clock; // @[Soc.scala 48:25]
  wire  bus_io_dev_1_timer_reset; // @[Soc.scala 48:25]
  wire  bus_io_dev_1_timer_io_bus_req; // @[Soc.scala 48:25]
  wire [31:0] bus_io_dev_1_timer_io_bus_addr; // @[Soc.scala 48:25]
  wire  bus_io_dev_1_timer_io_bus_rvalid; // @[Soc.scala 48:25]
  wire [31:0] bus_io_dev_1_timer_io_bus_rdata; // @[Soc.scala 48:25]
  wire  bus_io_dev_2_gpio_clock; // @[Soc.scala 52:24]
  wire  bus_io_dev_2_gpio_reset; // @[Soc.scala 52:24]
  wire  bus_io_dev_2_gpio_io_bus_req; // @[Soc.scala 52:24]
  wire [31:0] bus_io_dev_2_gpio_io_bus_addr; // @[Soc.scala 52:24]
  wire  bus_io_dev_2_gpio_io_bus_we; // @[Soc.scala 52:24]
  wire [31:0] bus_io_dev_2_gpio_io_bus_wdata; // @[Soc.scala 52:24]
  wire  bus_io_dev_2_gpio_io_bus_rvalid; // @[Soc.scala 52:24]
  wire [31:0] bus_io_dev_2_gpio_io_bus_rdata; // @[Soc.scala 52:24]
  wire  bus_io_dev_2_gpio_io_gpi_0; // @[Soc.scala 52:24]
  wire  bus_io_dev_2_gpio_io_gpo_0; // @[Soc.scala 52:24]
  wire  bus_io_dev_2_gpio_io_gpo_1; // @[Soc.scala 52:24]
  wire  bus_io_dev_2_gpio_io_gpo_2; // @[Soc.scala 52:24]
  Core core ( // @[Soc.scala 38:20]
    .clock(core_clock),
    .reset(core_reset),
    .io_imem_addr(core_io_imem_addr),
    .io_imem_rdata(core_io_imem_rdata),
    .io_dmem_req(core_io_dmem_req),
    .io_dmem_addr(core_io_dmem_addr),
    .io_dmem_we(core_io_dmem_we),
    .io_dmem_be(core_io_dmem_be),
    .io_dmem_wdata(core_io_dmem_wdata),
    .io_dmem_gnt(core_io_dmem_gnt),
    .io_dmem_rvalid(core_io_dmem_rvalid),
    .io_dmem_rdata(core_io_dmem_rdata)
  );
  SimpleBus bus ( // @[Soc.scala 65:19]
    .clock(bus_clock),
    .reset(bus_reset),
    .io_host_0_req(bus_io_host_0_req),
    .io_host_0_addr(bus_io_host_0_addr),
    .io_host_0_we(bus_io_host_0_we),
    .io_host_0_be(bus_io_host_0_be),
    .io_host_0_wdata(bus_io_host_0_wdata),
    .io_host_0_gnt(bus_io_host_0_gnt),
    .io_host_0_rvalid(bus_io_host_0_rvalid),
    .io_host_0_rdata(bus_io_host_0_rdata),
    .io_dev_0_req(bus_io_dev_0_req),
    .io_dev_0_addr(bus_io_dev_0_addr),
    .io_dev_0_we(bus_io_dev_0_we),
    .io_dev_0_be(bus_io_dev_0_be),
    .io_dev_0_wdata(bus_io_dev_0_wdata),
    .io_dev_0_rvalid(bus_io_dev_0_rvalid),
    .io_dev_0_rdata(bus_io_dev_0_rdata),
    .io_dev_1_req(bus_io_dev_1_req),
    .io_dev_1_addr(bus_io_dev_1_addr),
    .io_dev_1_rvalid(bus_io_dev_1_rvalid),
    .io_dev_1_rdata(bus_io_dev_1_rdata),
    .io_dev_2_req(bus_io_dev_2_req),
    .io_dev_2_addr(bus_io_dev_2_addr),
    .io_dev_2_we(bus_io_dev_2_we),
    .io_dev_2_wdata(bus_io_dev_2_wdata),
    .io_dev_2_rvalid(bus_io_dev_2_rvalid),
    .io_dev_2_rdata(bus_io_dev_2_rdata)
  );
  Ram bus_io_dev_0_ram ( // @[Soc.scala 42:23]
    .clock(bus_io_dev_0_ram_clock),
    .io_imem_addr(bus_io_dev_0_ram_io_imem_addr),
    .io_imem_rdata(bus_io_dev_0_ram_io_imem_rdata),
    .io_dmem_req(bus_io_dev_0_ram_io_dmem_req),
    .io_dmem_addr(bus_io_dev_0_ram_io_dmem_addr),
    .io_dmem_we(bus_io_dev_0_ram_io_dmem_we),
    .io_dmem_be(bus_io_dev_0_ram_io_dmem_be),
    .io_dmem_wdata(bus_io_dev_0_ram_io_dmem_wdata),
    .io_dmem_rvalid(bus_io_dev_0_ram_io_dmem_rvalid),
    .io_dmem_rdata(bus_io_dev_0_ram_io_dmem_rdata)
  );
  Timer bus_io_dev_1_timer ( // @[Soc.scala 48:25]
    .clock(bus_io_dev_1_timer_clock),
    .reset(bus_io_dev_1_timer_reset),
    .io_bus_req(bus_io_dev_1_timer_io_bus_req),
    .io_bus_addr(bus_io_dev_1_timer_io_bus_addr),
    .io_bus_rvalid(bus_io_dev_1_timer_io_bus_rvalid),
    .io_bus_rdata(bus_io_dev_1_timer_io_bus_rdata)
  );
  Gpio bus_io_dev_2_gpio ( // @[Soc.scala 52:24]
    .clock(bus_io_dev_2_gpio_clock),
    .reset(bus_io_dev_2_gpio_reset),
    .io_bus_req(bus_io_dev_2_gpio_io_bus_req),
    .io_bus_addr(bus_io_dev_2_gpio_io_bus_addr),
    .io_bus_we(bus_io_dev_2_gpio_io_bus_we),
    .io_bus_wdata(bus_io_dev_2_gpio_io_bus_wdata),
    .io_bus_rvalid(bus_io_dev_2_gpio_io_bus_rvalid),
    .io_bus_rdata(bus_io_dev_2_gpio_io_bus_rdata),
    .io_gpi_0(bus_io_dev_2_gpio_io_gpi_0),
    .io_gpo_0(bus_io_dev_2_gpio_io_gpo_0),
    .io_gpo_1(bus_io_dev_2_gpio_io_gpo_1),
    .io_gpo_2(bus_io_dev_2_gpio_io_gpo_2)
  );
  assign io_gpo_0 = bus_io_dev_2_gpio_io_gpo_0; // @[Soc.scala 53:19]
  assign io_gpo_1 = bus_io_dev_2_gpio_io_gpo_1; // @[Soc.scala 53:19]
  assign io_gpo_2 = bus_io_dev_2_gpio_io_gpo_2; // @[Soc.scala 53:19]
  assign core_clock = clock;
  assign core_reset = reset;
  assign core_io_imem_rdata = bus_io_dev_0_ram_io_imem_rdata; // @[Soc.scala 43:20]
  assign core_io_dmem_gnt = bus_io_host_0_gnt; // @[Soc.scala 66:16]
  assign core_io_dmem_rvalid = bus_io_host_0_rvalid; // @[Soc.scala 66:16]
  assign core_io_dmem_rdata = bus_io_host_0_rdata; // @[Soc.scala 66:16]
  assign bus_clock = clock;
  assign bus_reset = reset;
  assign bus_io_host_0_req = core_io_dmem_req; // @[Soc.scala 66:16]
  assign bus_io_host_0_addr = core_io_dmem_addr; // @[Soc.scala 66:16]
  assign bus_io_host_0_we = core_io_dmem_we; // @[Soc.scala 66:16]
  assign bus_io_host_0_be = core_io_dmem_be; // @[Soc.scala 66:16]
  assign bus_io_host_0_wdata = core_io_dmem_wdata; // @[Soc.scala 66:16]
  assign bus_io_dev_0_rvalid = bus_io_dev_0_ram_io_dmem_rvalid; // @[Soc.scala 70:19]
  assign bus_io_dev_0_rdata = bus_io_dev_0_ram_io_dmem_rdata; // @[Soc.scala 70:19]
  assign bus_io_dev_1_rvalid = bus_io_dev_1_timer_io_bus_rvalid; // @[Soc.scala 70:19]
  assign bus_io_dev_1_rdata = bus_io_dev_1_timer_io_bus_rdata; // @[Soc.scala 70:19]
  assign bus_io_dev_2_rvalid = bus_io_dev_2_gpio_io_bus_rvalid; // @[Soc.scala 70:19]
  assign bus_io_dev_2_rdata = bus_io_dev_2_gpio_io_bus_rdata; // @[Soc.scala 70:19]
  assign bus_io_dev_0_ram_clock = clock;
  assign bus_io_dev_0_ram_io_imem_addr = core_io_imem_addr; // @[Soc.scala 43:20]
  assign bus_io_dev_0_ram_io_dmem_req = bus_io_dev_0_req; // @[Soc.scala 70:19]
  assign bus_io_dev_0_ram_io_dmem_addr = bus_io_dev_0_addr; // @[Soc.scala 70:19]
  assign bus_io_dev_0_ram_io_dmem_we = bus_io_dev_0_we; // @[Soc.scala 70:19]
  assign bus_io_dev_0_ram_io_dmem_be = bus_io_dev_0_be; // @[Soc.scala 70:19]
  assign bus_io_dev_0_ram_io_dmem_wdata = bus_io_dev_0_wdata; // @[Soc.scala 70:19]
  assign bus_io_dev_1_timer_clock = clock;
  assign bus_io_dev_1_timer_reset = reset;
  assign bus_io_dev_1_timer_io_bus_req = bus_io_dev_1_req; // @[Soc.scala 70:19]
  assign bus_io_dev_1_timer_io_bus_addr = bus_io_dev_1_addr; // @[Soc.scala 70:19]
  assign bus_io_dev_2_gpio_clock = clock;
  assign bus_io_dev_2_gpio_reset = reset;
  assign bus_io_dev_2_gpio_io_bus_req = bus_io_dev_2_req; // @[Soc.scala 70:19]
  assign bus_io_dev_2_gpio_io_bus_addr = bus_io_dev_2_addr; // @[Soc.scala 70:19]
  assign bus_io_dev_2_gpio_io_bus_we = bus_io_dev_2_we; // @[Soc.scala 70:19]
  assign bus_io_dev_2_gpio_io_bus_wdata = bus_io_dev_2_wdata; // @[Soc.scala 70:19]
  assign bus_io_dev_2_gpio_io_gpi_0 = io_gpi_0; // @[Soc.scala 54:19]
endmodule
