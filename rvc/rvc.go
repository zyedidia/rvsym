// Package rvc decompresses rv32c instructions
package rvc

import "github.com/zyedidia/rvsym/bits"

func op(insn uint32) uint32 {
	return bits.Get(insn, 1, 0)
}

func funct3(insn uint32) uint32 {
	return bits.Get(insn, 15, 13)
}

func imm(insn uint32) uint32 {
	return bits.Get(insn, 12, 5)
}

func Decompress(insn uint32) (out uint32, compressed bool, illegal bool) {
	out = insn
	illegal = false
	compressed = true

	switch op(insn) {
	case 0b00:
		switch funct3(insn) {
		case 0b000:
			illegal = imm(insn) == 0
			// c.addi4spn -> addi rd', x2, imm
			// instr_o = {2'b0, instr_i[10:7], instr_i[12:11], instr_i[5],
			//            instr_i[6], 2'b00, 5'h02, 3'b000, 2'b01, instr_i[4:2], {OPCODE_OP_IMM}};
		case 0b010:
			// c.lw -> lw rd', imm(rs1')
			// instr_o = {5'b0, instr_i[5], instr_i[12:10], instr_i[6],
			//            2'b00, 2'b01, instr_i[9:7], 3'b010, 2'b01, instr_i[4:2], {OPCODE_LOAD}};
		case 0b110:
			// c.sw -> sw rs2', imm(rs1')
			// instr_o = {5'b0, instr_i[5], instr_i[12], 2'b01, instr_i[4:2],
			//            2'b01, instr_i[9:7], 3'b010, instr_i[11:10], instr_i[6],
			//            2'b00, {OPCODE_STORE}};
		default:
			illegal = true
		}
	case 0b01:
		switch funct3(insn) {
		case 0b000:
			// c.addi -> addi rd, rd, nzimm
			// c.nop
			// instr_o = {{6 {instr_i[12]}}, instr_i[12], instr_i[6:2],
			//            instr_i[11:7], 3'b0, instr_i[11:7], {OPCODE_OP_IMM}};
		case 0b001, 0b101:
			// 001: c.jal -> jal x1, imm
			// 101: c.j   -> jal x0, imm
			// instr_o = {instr_i[12], instr_i[8], instr_i[10:9], instr_i[6],
			//            instr_i[7], instr_i[2], instr_i[11], instr_i[5:3],
			//            {9 {instr_i[12]}}, 4'b0, ~instr_i[15], {OPCODE_JAL}};
		case 0b010:
			// c.li -> addi rd, x0, nzimm
			// (c.li hints are translated into an addi hint)
			// instr_o = {{6 {instr_i[12]}}, instr_i[12], instr_i[6:2], 5'b0,
			//            3'b0, instr_i[11:7], {OPCODE_OP_IMM}};
		case 0b011:
			// c.lui -> lui rd, imm
			// (c.lui hints are translated into a lui hint)
			// instr_o = {{15 {instr_i[12]}}, instr_i[6:2], instr_i[11:7], {OPCODE_LUI}};

			if bits.Get(insn, 11, 7) == 2 {
				// c.addi16sp -> addi x2, x2, nzimm
				// instr_o = {{3 {instr_i[12]}}, instr_i[4:3], instr_i[5], instr_i[2],
				//            instr_i[6], 4'b0, 5'h02, 3'b000, 5'h02, {OPCODE_OP_IMM}};
			}
		case 0b100:
			switch bits.Get(insn, 11, 10) {
			case 0b00, 0b01:
				// 00: c.srli -> srli rd, rd, shamt
				// 01: c.srai -> srai rd, rd, shamt
				// (c.srli/c.srai hints are translated into a srli/srai hint)
				// instr_o = {1'b0, instr_i[10], 5'b0, instr_i[6:2], 2'b01, instr_i[9:7],
				//            3'b101, 2'b01, instr_i[9:7], {OPCODE_OP_IMM}};
				illegal = bits.GetBit(insn, 12) == 1
			case 0b10:
				// c.andi -> andi rd, rd, imm
				// instr_o = {{6 {instr_i[12]}}, instr_i[12], instr_i[6:2], 2'b01, instr_i[9:7],
				//            3'b111, 2'b01, instr_i[9:7], {OPCODE_OP_IMM}};

			case 0b11:
				switch /* {instr_i[12], instr_i[6:5]} */ {
				case 0b000:
				case 0b001:
				case 0b010:
				case 0b011:
				case 0b100, 0b101, 0b110, 0b111:
					// 100: c.subw
					// 101: c.addw
					illegal = true
				default:
					illegal = true
				}
			}
		case 0b110, 0b111:
			// 0: c.beqz -> beq rs1', x0, imm
			// 1: c.bnez -> bne rs1', x0, imm
			// instr_o = {{4 {instr_i[12]}}, instr_i[6:5], instr_i[2], 5'b0, 2'b01,
			//            instr_i[9:7], 2'b00, instr_i[13], instr_i[11:10], instr_i[4:3],
			//            instr_i[12], {OPCODE_BRANCH}};
		}
	case 0b10:
		switch funct3(insn) {
		case 0b000:
			// c.slli -> slli rd, rd, shamt
			// (c.ssli hints are translated into a slli hint)
			// instr_o = {7'b0, instr_i[6:2], instr_i[11:7], 3'b001, instr_i[11:7], {OPCODE_OP_IMM}};
			illegal = bits.GetBit(insn, 12) == 1 // reserved for custom extensions
		case 0b010:
			// c.lwsp -> lw rd, imm(x2)
			// instr_o = {4'b0, instr_i[3:2], instr_i[12], instr_i[6:4], 2'b00, 5'h02,
			//            3'b010, instr_i[11:7], OPCODE_LOAD};
			illegal = bits.Get(insn, 11, 7) == 0
		case 0b100:
			if bits.GetBit(insn, 12) == 0 {
				if bits.Get(insn, 6, 2) != 0 {
					// c.mv -> add rd/rs1, x0, rs2
					// (c.mv hints are translated into an add hint)
					// instr_o = {7'b0, instr_i[6:2], 5'b0, 3'b0, instr_i[11:7], {OPCODE_OP}};
				} else {
					// c.jr -> jalr x0, rd/rs1, 0
					// instr_o = {12'b0, instr_i[11:7], 3'b0, 5'b0, {OPCODE_JALR}};
					illegal = bits.Get(insn, 11, 7) == 0
				}
			} else {
				if bits.Get(insn, 6, 2) != 0 {
					// c.add -> add rd, rd, rs2
					// (c.add hints are translated into an add hint)
					// instr_o = {7'b0, instr_i[6:2], instr_i[11:7], 3'b0, instr_i[11:7], {OPCODE_OP}};
				} else {
					if bits.Get(insn, 11, 7) == 0 {
						// c.ebreak -> ebreak
						// instr_o = {32'h00_10_00_73};
					} else {
						// c.jalr -> jalr x1, rs1, 0
						// instr_o = {12'b0, instr_i[11:7], 3'b000, 5'b00001, {OPCODE_JALR}};
					}
				}
			}
		case 0b110:
			// c.swsp -> sw rs2, imm(x2)
			// instr_o = {4'b0, instr_i[8:7], instr_i[12], instr_i[6:2], 5'h02, 3'b010,
			//            instr_i[11:9], 2'b00, {OPCODE_STORE}};
		default:
			illegal = true
		}
	case 0b11:
		// not compressed
		compressed = false
	}

	return out, compressed, illegal
}
