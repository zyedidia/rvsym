localparam BOOT_ADDR = 32'd0;

// control states
localparam CTRL_STATE_FETCH = 3'b000;
localparam CTRL_STATE_DECODE = 3'b001;
localparam CTRL_STATE_EXE = 3'b010;
localparam CTRL_STATE_MEM = 3'b011;
localparam CTRL_STATE_WB = 3'b100;
localparam CTRL_STATE_ERR = 3'b111;
