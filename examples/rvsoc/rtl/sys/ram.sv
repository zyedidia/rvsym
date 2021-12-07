// This module implements a single-port synchronous RAM with 32-bit words. It
// has valid signals for read and write requests and results so both the user
// and the module know when requests/results are ready. A write mask is used to
// specify that data smaller than a word should be read from memory. In order
// to read or write a full word, the address must be word-aligned. Otherwise
// less than a word will be read/written (up to the next word boundary).

module ram
    #(
        parameter SIZE = 1
    )
    (
        input logic clk,

        input logic        i_rd,
        input logic [31:0] i_addr,
        input logic        i_wr,
        input logic [3:0]  i_wrmask,
        input logic [31:0] i_data,

        output logic        o_rd_valid,
        output logic        o_wr_valid,
        output logic [31:0] o_data

    );

    logic ram_addr;
    assign ram_addr = i_addr >= RAM_BASE && i_addr < RAM_BASE + RAM_SIZE;

    `include "memmap.svh"

    logic [31:0] mem [0:SIZE-1];

    initial begin
        $readmemh(`RAMFILE, mem);
    end

    logic [31:0] read_word;
    logic [31:0] write_word;
    logic [3:0] write_mask;

    logic rd_valid, wr_valid;

    always_comb begin
        // index of the byte within this word that we are reading/writing
        logic [1:0] rbyte_n, wbyte_n;
        rbyte_n = i_addr[1:0];
        wbyte_n = i_addr[1:0];

        // the result data is the read word shifted to the exact byte that was
        // requested.
        o_data = ram_addr ? read_word >> (rbyte_n >> 3) : 32'b0;
        // same for the write word
        write_word = i_data << (wbyte_n >> 3);
        // the write mask also needs to be shifted. For example if the write
        // mask is 4'b0001, indicating that the lower byte of the
        // i_data should be written, the actual mask will have to be
        // shifted if we are writing to a particular byte within the word.
        write_mask = i_wrmask << wbyte_n;
    end

    always_ff @(posedge clk) begin
        if (i_wr && ram_addr) begin
            // write the appropriate bytes according to the write mask
            if (write_mask[3])
                mem[i_addr >> 2][31:24] <= write_word[31:24];
            if (write_mask[2])
                mem[i_addr >> 2][23:16] <= write_word[23:16];
            if (write_mask[1])
                mem[i_addr >> 2][15:8] <= write_word[15:8];
            if (write_mask[0])
                mem[i_addr >> 2][7:0] <= write_word[7:0];
        end

        if (i_rd && ram_addr) begin
            read_word <= mem[ram_addr ? i_addr >> 2 : 32'b0];
        end

        rd_valid <= i_rd;
        wr_valid <= i_wr;
    end

    assign o_rd_valid = ram_addr ? rd_valid : 1'b0;
    assign o_wr_valid = ram_addr ? wr_valid : 1'b0;
endmodule

