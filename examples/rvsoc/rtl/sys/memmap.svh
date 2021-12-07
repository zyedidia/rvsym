localparam [31:0] RAM_BASE = 32'h0000;
localparam [31:0] RAM_SIZE = 32'h2000;

localparam [31:0] GPO_BASE  = 32'h2000;
localparam [31:0] GPO_SIZE  = 8;
localparam [31:0] GPO_SET   = GPO_BASE + 0;
localparam [31:0] GPO_CLEAR = GPO_BASE + 4;

localparam [31:0] GPI_BASE  = 32'h2000 + GPO_SIZE;
localparam [31:0] GPI_SIZE  = 4;
localparam [31:0] GPI_LEVEL = GPI_BASE;

localparam [31:0] GPIO_BASE = 32'h3000;
localparam [31:0] GPIO_SIZE = 16;

localparam [31:0] GPIO_SET   = GPIO_BASE + 0;
localparam [31:0] GPIO_CLEAR = GPIO_BASE + 4;
localparam [31:0] GPIO_LEVEL = GPIO_BASE + 8;
localparam [31:0] GPIO_FSEL  = GPIO_BASE + 12;

localparam [31:0] TIMER_ADDR = 32'h4000;
