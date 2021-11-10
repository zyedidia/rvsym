// Machine Information Registers (always 0)
localparam CSR_MVENDORID = 12'hF11; // vendor ID
localparam CSR_MARCHID   = 12'hF12; // architecture ID
localparam CSR_MIMPID    = 12'hF13; // implementation ID
localparam CSR_MHARTID   = 12'hF14; // hardware thread ID

// Machine Trap Setup
// Status register for whether interrupts are enabled and more info that we
// don't support (privilege modes, virtualization, etc.).
localparam CSR_MSTATUS = 12'h300;
// ISA Register: reports the ISA implemented by this processor.
localparam CSR_MISA    = 12'h301;
// Interrupt enable bits.
localparam CSR_MIE     = 12'h304;
// Trap-Vector Base Address: determines where traps jump to.
localparam CSR_MTVEC   = 12'h305;

// Machine Trap Handling
// Dedicated for use by M-mode.
localparam CSR_MSCRATCH = 12'h340;
// Exception Program Counter. Stores the PC of the instruction that was
// interrupted or that encountered an exception.
localparam CSR_MEPC     = 12'h341;
// Cause Register. When a trap is taken into M-mode, mcause will contain a code
// that explains the event that caused the trap.
localparam CSR_MCAUSE   = 12'h342;
// Machine Trap Value. This register may contain exception-specific information
// when a trap is taken into M-mode.
localparam CSR_MTVAL    = 12'h343;
// Information about pending interrupts.
localparam CSP_MIP      = 12'h344;

// Performance counters
// M-mode versions
localparam CSR_NUM_MCYCLE        = 12'hB00;
localparam CSR_NUM_MINSTRET      = 12'hB02;
localparam CSR_NUM_MHPMCOUNTER3  = 12'hB03;
localparam CSR_NUM_MHPMCOUNTER31 = 12'hB1F;

localparam CSR_NUM_MCYCLEH        = 12'hB80;
localparam CSR_NUM_MINSTRETH      = 12'hB82;
localparam CSR_NUM_MHPMCOUNTER3H  = 12'hB83;
localparam CSR_NUM_MHPMCOUNTER31H = 12'hB9F;

localparam CSR_NUM_MHPMEVENT3  = 12'h323;
localparam CSR_NUM_MHPMEVENT31 = 12'h323;

// U-mode equivalents
localparam CSR_NUM_CYCLE        = 12'hC00;
localparam CSR_NUM_INSTRET      = 12'hC02;
localparam CSR_NUM_HPMCOUNTER3  = 12'hC03;
localparam CSR_NUM_HPMCOUNTER31 = 12'hC1F;

localparam CSR_NUM_CYCLEH        = 12'hC80;
localparam CSR_NUM_INSTRETH      = 12'hC82;
localparam CSR_NUM_HPMCOUNTER3H  = 12'hC83;
localparam CSR_NUM_HPMCOUNTER31H = 12'hC9F;

