package z80

// http://map.grauw.nl/resources/z80instr.php

func initMSXTimings() {
    for i:=0; i<1536; i++ {
        timingsMSX[i] = func(z80 *Z80) uint64 { return 0 }
    }
	
	/* NOP */
	timingsMSX[0x0] = func(z80 *Z80) uint64 { return 5 }
	
	/* LD (BC),A */
	timingsMSX[0x2] = func(z80 *Z80) uint64 { return 8 }
	
	/* INC BC */
	timingsMSX[0x3] = func(z80 *Z80) uint64 { return 7 }
	
	/* INC B */
	timingsMSX[0x4] = func(z80 *Z80) uint64 { return 5 }
	
	/* DEC B */
	timingsMSX[0x5] = func(z80 *Z80) uint64 { return 5 }
	
	/* RLCA */
	timingsMSX[0x7] = func(z80 *Z80) uint64 { return 5 }
	
	/* EX AF,AF' */
	timingsMSX[0x8] = func(z80 *Z80) uint64 { return 5 }
	
	/* ADD HL,BC */
	timingsMSX[0x9] = func(z80 *Z80) uint64 { return 12 }
	
	/* LD (DE),A */
	timingsMSX[0x12] = func(z80 *Z80) uint64 { return 8 }
	
	/* INC DE */
	timingsMSX[0x13] = func(z80 *Z80) uint64 { return 7 }
	
	/* INC D */
	timingsMSX[0x14] = func(z80 *Z80) uint64 { return 5 }
	
	/* DEC D */
	timingsMSX[0x15] = func(z80 *Z80) uint64 { return 5 }
	
	/* RLA */
	timingsMSX[0x17] = func(z80 *Z80) uint64 { return 5 }
	
	/* ADD HL,DE */
	timingsMSX[0x19] = func(z80 *Z80) uint64 { return 12 }
	
	/* INC HL */
	timingsMSX[0x23] = func(z80 *Z80) uint64 { return 7 }
	
	/* INC H */
	timingsMSX[0x24] = func(z80 *Z80) uint64 { return 5 }
	
	/* DEC H */
	timingsMSX[0x25] = func(z80 *Z80) uint64 { return 5 }
	
	/* DAA */
	timingsMSX[0x27] = func(z80 *Z80) uint64 { return 5 }
	
	/* ADD HL,HL */
	timingsMSX[0x29] = func(z80 *Z80) uint64 { return 12 }
	
	/* INC SP */
	timingsMSX[0x33] = func(z80 *Z80) uint64 { return 7 }
	
	/* INC (HL) */
	timingsMSX[0x34] = func(z80 *Z80) uint64 { return 12 }
	
	/* DEC (HL) */
	timingsMSX[0x35] = func(z80 *Z80) uint64 { return 12 }
	
	/* SCF */
	timingsMSX[0x37] = func(z80 *Z80) uint64 { return 5 }
	
	/* ADD HL,SP */
	timingsMSX[0x39] = func(z80 *Z80) uint64 { return 12 }
	
	/* LD B,(HL) */
	timingsMSX[0x46] = func(z80 *Z80) uint64 { return 8 }
	
	/* LD D,(HL) */
	timingsMSX[0x56] = func(z80 *Z80) uint64 { return 8 }
	
	/* LD H,(HL) */
	timingsMSX[0x66] = func(z80 *Z80) uint64 { return 8 }
	
	/* HALT */
	timingsMSX[0x76] = func(z80 *Z80) uint64 { return 5 }
	
	/* ADD A,(HL) */
	timingsMSX[0x86] = func(z80 *Z80) uint64 { return 8 }
	
	/* SUB (HL) */
	timingsMSX[0x96] = func(z80 *Z80) uint64 { return 8 }
	
	/* LD BC,nn */
	timingsMSX[0x1] = func(z80 *Z80) uint64 { return 11 }
	
	/* LD B,n */
	timingsMSX[0x6] = func(z80 *Z80) uint64 { return 8 }
	
	/* LD A,(BC) */
	timingsMSX[0x0a] = func(z80 *Z80) uint64 { return 8 }
	
	/* DEC BC */
	timingsMSX[0x0b] = func(z80 *Z80) uint64 { return 7 }
	
	/* INC C */
	timingsMSX[0x0c] = func(z80 *Z80) uint64 { return 5 }
	
	/* DEC C */
	timingsMSX[0x0d] = func(z80 *Z80) uint64 { return 5 }
	
	/* LD C,n */
	timingsMSX[0x0e] = func(z80 *Z80) uint64 { return 8 }
	
	/* RRCA */
	timingsMSX[0x0f] = func(z80 *Z80) uint64 { return 5 }
	
	/* DJNZ o */
	timingsMSX[0x10] = func(z80 *Z80) uint64 { return 0 }
	
	/* LD DE,nn */
	timingsMSX[0x11] = func(z80 *Z80) uint64 { return 11 }
	
	/* LD D,n */
	timingsMSX[0x16] = func(z80 *Z80) uint64 { return 8 }
	
	/* JR o */
	timingsMSX[0x18] = func(z80 *Z80) uint64 { return 13 }
	
	/* LD A,(DE) */
	timingsMSX[0x1a] = func(z80 *Z80) uint64 { return 8 }
	
	/* DEC DE */
	timingsMSX[0x1b] = func(z80 *Z80) uint64 { return 7 }
	
	/* INC E */
	timingsMSX[0x1c] = func(z80 *Z80) uint64 { return 5 }
	
	/* DEC E */
	timingsMSX[0x1d] = func(z80 *Z80) uint64 { return 5 }
	
	/* LD E,n */
	timingsMSX[0x1e] = func(z80 *Z80) uint64 { return 8 }
	
	/* RRA */
	timingsMSX[0x1f] = func(z80 *Z80) uint64 { return 5 }
	
	/* JR NZ,o */
	timingsMSX[0x20] = func(z80 *Z80) uint64 { return 0 }
	
	/* LD HL,nn */
	timingsMSX[0x21] = func(z80 *Z80) uint64 { return 11 }
	
	/* LD (nn),HL */
	timingsMSX[0x22] = func(z80 *Z80) uint64 { return 17 }
	
	/* LD H,n */
	timingsMSX[0x26] = func(z80 *Z80) uint64 { return 8 }
	
	/* JR Z,o */
	timingsMSX[0x28] = func(z80 *Z80) uint64 { return 0 }
	
	/* LD HL,(nn) */
	timingsMSX[0x2a] = func(z80 *Z80) uint64 { return 17 }
	
	/* DEC HL */
	timingsMSX[0x2b] = func(z80 *Z80) uint64 { return 7 }
	
	/* INC L */
	timingsMSX[0x2c] = func(z80 *Z80) uint64 { return 5 }
	
	/* DEC L */
	timingsMSX[0x2d] = func(z80 *Z80) uint64 { return 5 }
	
	/* LD L,n */
	timingsMSX[0x2e] = func(z80 *Z80) uint64 { return 8 }
	
	/* CPL */
	timingsMSX[0x2f] = func(z80 *Z80) uint64 { return 5 }
	
	/* JR NC,o */
	timingsMSX[0x30] = func(z80 *Z80) uint64 { return 0 }
	
	/* LD SP,nn */
	timingsMSX[0x31] = func(z80 *Z80) uint64 { return 11 }
	
	/* LD (nn),A */
	timingsMSX[0x32] = func(z80 *Z80) uint64 { return 14 }
	
	/* LD (HL),n */
	timingsMSX[0x36] = func(z80 *Z80) uint64 { return 11 }
	
	/* JR C,o */
	timingsMSX[0x38] = func(z80 *Z80) uint64 { return 0 }
	
	/* LD A,(nn) */
	timingsMSX[0x3a] = func(z80 *Z80) uint64 { return 14 }
	
	/* DEC SP */
	timingsMSX[0x3b] = func(z80 *Z80) uint64 { return 7 }
	
	/* INC A */
	timingsMSX[0x3c] = func(z80 *Z80) uint64 { return 5 }
	
	/* DEC A */
	timingsMSX[0x3d] = func(z80 *Z80) uint64 { return 5 }
	
	/* LD A,n */
	timingsMSX[0x3e] = func(z80 *Z80) uint64 { return 8 }
	
	/* CCF */
	timingsMSX[0x3f] = func(z80 *Z80) uint64 { return 5 }
	
	/* LD B,r */
	timingsMSX[0x40] = func(z80 *Z80) uint64 { return 5 }
	
	/* LD C,r */
	timingsMSX[0x48] = func(z80 *Z80) uint64 { return 5 }
	
	/* LD C,(HL) */
	timingsMSX[0x4e] = func(z80 *Z80) uint64 { return 8 }
	
	/* LD D,r */
	timingsMSX[0x50] = func(z80 *Z80) uint64 { return 5 }
	
	/* LD E,r */
	timingsMSX[0x58] = func(z80 *Z80) uint64 { return 5 }
	
	/* LD E,(HL) */
	timingsMSX[0x5e] = func(z80 *Z80) uint64 { return 8 }
	
	/* LD H,r */
	timingsMSX[0x60] = func(z80 *Z80) uint64 { return 5 }
	
	/* LD L,r */
	timingsMSX[0x68] = func(z80 *Z80) uint64 { return 5 }
	
	/* LD L,(HL) */
	timingsMSX[0x6e] = func(z80 *Z80) uint64 { return 8 }
	
	/* LD (HL),r */
	timingsMSX[0x70] = func(z80 *Z80) uint64 { return 8 }
	
	/* LD A,r */
	timingsMSX[0x78] = func(z80 *Z80) uint64 { return 5 }
	
	/* LD A,(HL) */
	timingsMSX[0x7e] = func(z80 *Z80) uint64 { return 8 }
	
	/* ADD A,r */
	timingsMSX[0x80] = func(z80 *Z80) uint64 { return 5 }
	
	/* ADC A,r */
	timingsMSX[0x88] = func(z80 *Z80) uint64 { return 5 }
	
	/* ADC A,(HL) */
	timingsMSX[0x8e] = func(z80 *Z80) uint64 { return 8 }
	
	/* SUB r */
	timingsMSX[0x90] = func(z80 *Z80) uint64 { return 5 }
	
	/* SBC A,r */
	timingsMSX[0x98] = func(z80 *Z80) uint64 { return 5 }
	
	/* SBC A,(HL) */
	timingsMSX[0x9e] = func(z80 *Z80) uint64 { return 8 }
	
	/* AND r */
	timingsMSX[0xa0] = func(z80 *Z80) uint64 { return 5 }
	
	/* AND (HL) */
	timingsMSX[0xa6] = func(z80 *Z80) uint64 { return 8 }
	
	/* XOR r */
	timingsMSX[0xa8] = func(z80 *Z80) uint64 { return 5 }
	
	/* XOR (HL) */
	timingsMSX[0xae] = func(z80 *Z80) uint64 { return 8 }
	
	/* OR r */
	timingsMSX[0xb0] = func(z80 *Z80) uint64 { return 5 }
	
	/* OR (HL) */
	timingsMSX[0xb6] = func(z80 *Z80) uint64 { return 8 }
	
	/* CP r */
	timingsMSX[0xb8] = func(z80 *Z80) uint64 { return 5 }
	
	/* CP (HL) */
	timingsMSX[0xbe] = func(z80 *Z80) uint64 { return 8 }
	
	/* RET NZ */
	timingsMSX[0xc0] = func(z80 *Z80) uint64 { return 0 }
	
	/* POP BC */
	timingsMSX[0xc1] = func(z80 *Z80) uint64 { return 11 }
	
	/* JP NZ,nn */
	timingsMSX[0xc2] = func(z80 *Z80) uint64 { return 11 }
	
	/* JP nn */
	timingsMSX[0xc3] = func(z80 *Z80) uint64 { return 11 }
	
	/* CALL NZ,nn */
	timingsMSX[0xc4] = func(z80 *Z80) uint64 { return 0 }
	
	/* PUSH BC */
	timingsMSX[0xc5] = func(z80 *Z80) uint64 { return 12 }
	
	/* ADD A,n */
	timingsMSX[0xc6] = func(z80 *Z80) uint64 { return 8 }
	
	/* RST 0 */
	timingsMSX[0xc7] = func(z80 *Z80) uint64 { return 12 }
	
	/* RET Z */
	timingsMSX[0xc8] = func(z80 *Z80) uint64 { return 0 }
	
	/* RET */
	timingsMSX[0xc9] = func(z80 *Z80) uint64 { return 11 }
	
	/* JP Z,nn */
	timingsMSX[0xca] = func(z80 *Z80) uint64 { return 11 }
	
	/* RLC r */
	timingsMSX[SHIFT_0xCB+0x00] = func(z80 *Z80) uint64 { return 10 }
	
	/* RLC (HL) */
	timingsMSX[SHIFT_0xCB+0x06] = func(z80 *Z80) uint64 { return 17 }
	
	/* RRC r */
	timingsMSX[SHIFT_0xCB+0x08] = func(z80 *Z80) uint64 { return 10 }
	
	/* RRC (HL) */
	timingsMSX[SHIFT_0xCB+0x0e] = func(z80 *Z80) uint64 { return 17 }
	
	/* RL r */
	timingsMSX[SHIFT_0xCB+0x10] = func(z80 *Z80) uint64 { return 10 }
	
	/* RL (HL) */
	timingsMSX[SHIFT_0xCB+0x16] = func(z80 *Z80) uint64 { return 17 }
	
	/* RR r */
	timingsMSX[SHIFT_0xCB+0x18] = func(z80 *Z80) uint64 { return 10 }
	
	/* RR (HL) */
	timingsMSX[SHIFT_0xCB+0x1e] = func(z80 *Z80) uint64 { return 17 }
	
	/* SLA r */
	timingsMSX[SHIFT_0xCB+0x20] = func(z80 *Z80) uint64 { return 10 }
	
	/* SLA (HL) */
	timingsMSX[SHIFT_0xCB+0x26] = func(z80 *Z80) uint64 { return 17 }
	
	/* SRA r */
	timingsMSX[SHIFT_0xCB+0x28] = func(z80 *Z80) uint64 { return 10 }
	
	/* SRA (HL) */
	timingsMSX[SHIFT_0xCB+0x2e] = func(z80 *Z80) uint64 { return 17 }
	
	/* SRL r */
	timingsMSX[SHIFT_0xCB+0x38] = func(z80 *Z80) uint64 { return 10 }
	
	/* SRL (HL) */
	timingsMSX[SHIFT_0xCB+0x3e] = func(z80 *Z80) uint64 { return 17 }
	
	/* BIT b,r */
	timingsMSX[SHIFT_0xCB+0x40] = func(z80 *Z80) uint64 { return 10 }
	
	/* BIT b,(HL) */
	timingsMSX[SHIFT_0xCB+0x46] = func(z80 *Z80) uint64 { return 14 }
	
	/* RES b,r */
	timingsMSX[SHIFT_0xCB+0x80] = func(z80 *Z80) uint64 { return 10 }
	
	/* RES b,(HL) */
	timingsMSX[SHIFT_0xCB+0x86] = func(z80 *Z80) uint64 { return 17 }
	
	/* SET b,r */
	timingsMSX[SHIFT_0xCB+0xc0] = func(z80 *Z80) uint64 { return 10 }
	
	/* SET b,(HL) */
	timingsMSX[SHIFT_0xCB+0xc6] = func(z80 *Z80) uint64 { return 17 }
	
	/* CALL Z,nn */
	timingsMSX[0xcc] = func(z80 *Z80) uint64 { return 0 }
	
	/* CALL nn */
	timingsMSX[0xcd] = func(z80 *Z80) uint64 { return 18 }
	
	/* ADC A,n */
	timingsMSX[0xce] = func(z80 *Z80) uint64 { return 8 }
	
	/* RST 8H */
	timingsMSX[0xcf] = func(z80 *Z80) uint64 { return 12 }
	
	/* RET NC */
	timingsMSX[0xd0] = func(z80 *Z80) uint64 { return 0 }
	
	/* POP DE */
	timingsMSX[0xd1] = func(z80 *Z80) uint64 { return 11 }
	
	/* JP NC,nn */
	timingsMSX[0xd2] = func(z80 *Z80) uint64 { return 11 }
	
	/* OUT (n),A */
	timingsMSX[0xd3] = func(z80 *Z80) uint64 { return 12 }
	
	/* CALL NC,nn */
	timingsMSX[0xd4] = func(z80 *Z80) uint64 { return 0 }
	
	/* PUSH DE */
	timingsMSX[0xd5] = func(z80 *Z80) uint64 { return 12 }
	
	/* SUB n */
	timingsMSX[0xd6] = func(z80 *Z80) uint64 { return 8 }
	
	/* RST 10H */
	timingsMSX[0xd7] = func(z80 *Z80) uint64 { return 12 }
	
	/* RET C */
	timingsMSX[0xd8] = func(z80 *Z80) uint64 { return 0 }
	
	/* EXX */
	timingsMSX[0xd9] = func(z80 *Z80) uint64 { return 5 }
	
	/* JP C,nn */
	timingsMSX[0xda] = func(z80 *Z80) uint64 { return 11 }
	
	/* IN A,(n) */
	timingsMSX[0xdb] = func(z80 *Z80) uint64 { return 12 }
	
	/* CALL C,nn */
	timingsMSX[0xdc] = func(z80 *Z80) uint64 { return 0 }
	
	/* INC IXp */
	timingsMSX[SHIFT_0xDD+0x04] = func(z80 *Z80) uint64 { return 10 }
	
	/* DEC IXp */
	timingsMSX[SHIFT_0xDD+0x05] = func(z80 *Z80) uint64 { return 10 }
	
	/* ADD IX,BC */
	timingsMSX[SHIFT_0xDD+0x09] = func(z80 *Z80) uint64 { return 17 }
	
	/* ADD IX,DE */
	timingsMSX[SHIFT_0xDD+0x19] = func(z80 *Z80) uint64 { return 17 }
	
	/* LD IX,nn */
	timingsMSX[SHIFT_0xDD+0x21] = func(z80 *Z80) uint64 { return 16 }
	
	/* LD (nn),IX */
	timingsMSX[SHIFT_0xDD+0x22] = func(z80 *Z80) uint64 { return 22 }
	
	/* INC IX */
	timingsMSX[SHIFT_0xDD+0x23] = func(z80 *Z80) uint64 { return 12 }
	
	/* LD IXh,n */
	timingsMSX[SHIFT_0xDD+0x26] = func(z80 *Z80) uint64 { return 13 }
	
	/* ADD IX,IX */
	timingsMSX[SHIFT_0xDD+0x29] = func(z80 *Z80) uint64 { return 17 }
	
	/* LD IX,(nn) */
	timingsMSX[SHIFT_0xDD+0x2a] = func(z80 *Z80) uint64 { return 22 }
	
	/* DEC IX */
	timingsMSX[SHIFT_0xDD+0x2b] = func(z80 *Z80) uint64 { return 12 }
	
	/* LD IXl,n */
	timingsMSX[SHIFT_0xDD+0x2e] = func(z80 *Z80) uint64 { return 13 }
	
	/* INC (IX+o) */
	timingsMSX[SHIFT_0xDD+0x34] = func(z80 *Z80) uint64 { return 25 }
	
	/* DEC (IX+o) */
	timingsMSX[SHIFT_0xDD+0x35] = func(z80 *Z80) uint64 { return 25 }
	
	/* LD (IX+o),n */
	timingsMSX[SHIFT_0xDD+0x36] = func(z80 *Z80) uint64 { return 21 }
	
	/* ADD IX,SP */
	timingsMSX[SHIFT_0xDD+0x39] = func(z80 *Z80) uint64 { return 17 }
	
	/* LD B,IXp */
	timingsMSX[SHIFT_0xDD+0x40] = func(z80 *Z80) uint64 { return 10 }
	
	/* LD B,(IX+o) */
	timingsMSX[SHIFT_0xDD+0x46] = func(z80 *Z80) uint64 { return 21 }
	
	/* LD C,IXp */
	timingsMSX[SHIFT_0xDD+0x48] = func(z80 *Z80) uint64 { return 10 }
	
	/* LD C,(IX+o) */
	timingsMSX[SHIFT_0xDD+0x4e] = func(z80 *Z80) uint64 { return 21 }
	
	/* LD D,IXp */
	timingsMSX[SHIFT_0xDD+0x50] = func(z80 *Z80) uint64 { return 10 }
	
	/* LD D,(IX+o) */
	timingsMSX[SHIFT_0xDD+0x56] = func(z80 *Z80) uint64 { return 21 }
	
	/* LD E,IXp */
	timingsMSX[SHIFT_0xDD+0x58] = func(z80 *Z80) uint64 { return 10 }
	
	/* LD E,(IX+o) */
	timingsMSX[SHIFT_0xDD+0x5e] = func(z80 *Z80) uint64 { return 21 }
	
	/* LD IXh,p */
	timingsMSX[SHIFT_0xDD+0x60] = func(z80 *Z80) uint64 { return 10 }
	
	/* LD H,(IX+o) */
	timingsMSX[SHIFT_0xDD+0x66] = func(z80 *Z80) uint64 { return 21 }
	
	/* LD IXl,p */
	timingsMSX[SHIFT_0xDD+0x68] = func(z80 *Z80) uint64 { return 10 }
	
	/* LD L,(IX+o) */
	timingsMSX[SHIFT_0xDD+0x6e] = func(z80 *Z80) uint64 { return 21 }
	
	/* LD (IX+o),r */
	timingsMSX[SHIFT_0xDD+0x70] = func(z80 *Z80) uint64 { return 21 }
	
	/* LD A,IXp */
	timingsMSX[SHIFT_0xDD+0x78] = func(z80 *Z80) uint64 { return 10 }
	
	/* LD A,(IX+o) */
	timingsMSX[SHIFT_0xDD+0x7e] = func(z80 *Z80) uint64 { return 21 }
	
	/* ADD A,IXp */
	timingsMSX[SHIFT_0xDD+0x80] = func(z80 *Z80) uint64 { return 10 }
	
	/* ADD A,(IX+o) */
	timingsMSX[SHIFT_0xDD+0x86] = func(z80 *Z80) uint64 { return 21 }
	
	/* ADC A,IXp */
	timingsMSX[SHIFT_0xDD+0x88] = func(z80 *Z80) uint64 { return 10 }
	
	/* ADC A,(IX+o) */
	timingsMSX[SHIFT_0xDD+0x8e] = func(z80 *Z80) uint64 { return 21 }
	
	/* SUB IXp */
	timingsMSX[SHIFT_0xDD+0x90] = func(z80 *Z80) uint64 { return 10 }
	
	/* SUB (IX+o) */
	timingsMSX[SHIFT_0xDD+0x96] = func(z80 *Z80) uint64 { return 21 }
	
	/* SBC A,IXp */
	timingsMSX[SHIFT_0xDD+0x98] = func(z80 *Z80) uint64 { return 10 }
	
	/* SBC A,(IX+o) */
	timingsMSX[SHIFT_0xDD+0x9e] = func(z80 *Z80) uint64 { return 21 }
	
	/* AND IXp */
	timingsMSX[SHIFT_0xDD+0xa0] = func(z80 *Z80) uint64 { return 10 }
	
	/* AND (IX+o) */
	timingsMSX[SHIFT_0xDD+0xa6] = func(z80 *Z80) uint64 { return 21 }
	
	/* XOR IXp */
	timingsMSX[SHIFT_0xDD+0xa8] = func(z80 *Z80) uint64 { return 10 }
	
	/* XOR (IX+o) */
	timingsMSX[SHIFT_0xDD+0xae] = func(z80 *Z80) uint64 { return 21 }
	
	/* OR IXp */
	timingsMSX[SHIFT_0xDD+0xb0] = func(z80 *Z80) uint64 { return 10 }
	
	/* OR (IX+o) */
	timingsMSX[SHIFT_0xDD+0xb6] = func(z80 *Z80) uint64 { return 21 }
	
	/* CP IXp */
	timingsMSX[SHIFT_0xDD+0xb8] = func(z80 *Z80) uint64 { return 10 }
	
	/* CP (IX+o) */
	timingsMSX[SHIFT_0xDD+0xbe] = func(z80 *Z80) uint64 { return 21 }
	
	/* RLC (IX+o) */
	timingsMSX[SHIFT_0xDDCB+0x06] = func(z80 *Z80) uint64 { return 25 }
	
	/* RRC (IX+o) */
	timingsMSX[SHIFT_0xDDCB+0x0e] = func(z80 *Z80) uint64 { return 25 }
	
	/* RL (IX+o) */
	timingsMSX[SHIFT_0xDDCB+0x16] = func(z80 *Z80) uint64 { return 25 }
	
	/* RR (IX+o) */
	timingsMSX[SHIFT_0xDDCB+0x1e] = func(z80 *Z80) uint64 { return 25 }
	
	/* SLA (IX+o) */
	timingsMSX[SHIFT_0xDDCB+0x26] = func(z80 *Z80) uint64 { return 25 }
	
	/* SRA (IX+o) */
	timingsMSX[SHIFT_0xDDCB+0x2e] = func(z80 *Z80) uint64 { return 25 }
	
	/* SRL (IX+o) */
	timingsMSX[SHIFT_0xDDCB+0x3e] = func(z80 *Z80) uint64 { return 25 }
	
	/* BIT b,(IX+o) */
	timingsMSX[SHIFT_0xDDCB+0x46] = func(z80 *Z80) uint64 { return 22 }
	
	/* RES b,(IX+o) */
	timingsMSX[SHIFT_0xDDCB+0x86] = func(z80 *Z80) uint64 { return 25 }
	
	/* SET b,(IX+o) */
	timingsMSX[SHIFT_0xDDCB+0xc6] = func(z80 *Z80) uint64 { return 25 }
	
	/* POP IX */
	timingsMSX[SHIFT_0xDD+0xe1] = func(z80 *Z80) uint64 { return 16 }
	
	/* EX (SP),IX */
	timingsMSX[SHIFT_0xDD+0xe3] = func(z80 *Z80) uint64 { return 25 }
	
	/* PUSH IX */
	timingsMSX[SHIFT_0xDD+0xe5] = func(z80 *Z80) uint64 { return 17 }
	
	/* JP (IX) */
	timingsMSX[SHIFT_0xDD+0xe9] = func(z80 *Z80) uint64 { return 10 }
	
	/* LD SP,IX */
	timingsMSX[SHIFT_0xDD+0xf9] = func(z80 *Z80) uint64 { return 12 }
	
	/* SBC A,n */
	timingsMSX[0xde] = func(z80 *Z80) uint64 { return 8 }
	
	/* RST 18H */
	timingsMSX[0xdf] = func(z80 *Z80) uint64 { return 12 }
	
	/* RET PO */
	timingsMSX[0xe0] = func(z80 *Z80) uint64 { return 0 }
	
	/* POP HL */
	timingsMSX[0xe1] = func(z80 *Z80) uint64 { return 11 }
	
	/* JP PO,nn */
	timingsMSX[0xe2] = func(z80 *Z80) uint64 { return 11 }
	
	/* EX (SP),HL */
	timingsMSX[0xe3] = func(z80 *Z80) uint64 { return 20 }
	
	/* CALL PO,nn */
	timingsMSX[0xe4] = func(z80 *Z80) uint64 { return 0 }
	
	/* PUSH HL */
	timingsMSX[0xe5] = func(z80 *Z80) uint64 { return 12 }
	
	/* AND n */
	timingsMSX[0xe6] = func(z80 *Z80) uint64 { return 8 }
	
	/* RST 20H */
	timingsMSX[0xe7] = func(z80 *Z80) uint64 { return 12 }
	
	/* RET PE */
	timingsMSX[0xe8] = func(z80 *Z80) uint64 { return 0 }
	
	/* JP (HL) */
	timingsMSX[0xe9] = func(z80 *Z80) uint64 { return 5 }
	
	/* JP PE,nn */
	timingsMSX[0xea] = func(z80 *Z80) uint64 { return 11 }
	
	/* EX DE,HL */
	timingsMSX[0xeb] = func(z80 *Z80) uint64 { return 5 }
	
	/* CALL PE,nn */
	timingsMSX[0xec] = func(z80 *Z80) uint64 { return 0 }
	
	/* IN B,(C) */
	timingsMSX[SHIFT_0xED+0x40] = func(z80 *Z80) uint64 { return 14 }
	
	/* OUT (C),B */
	timingsMSX[SHIFT_0xED+0x41] = func(z80 *Z80) uint64 { return 14 }
	
	/* SBC HL,BC */
	timingsMSX[SHIFT_0xED+0x42] = func(z80 *Z80) uint64 { return 17 }
	
	/* LD (nn),BC */
	timingsMSX[SHIFT_0xED+0x43] = func(z80 *Z80) uint64 { return 22 }
	
	/* NEG */
	timingsMSX[SHIFT_0xED+0x44] = func(z80 *Z80) uint64 { return 10 }
	
	/* RETN */
	timingsMSX[SHIFT_0xED+0x45] = func(z80 *Z80) uint64 { return 16 }
	
	/* IM 0 */
	timingsMSX[SHIFT_0xED+0x46] = func(z80 *Z80) uint64 { return 10 }
	
	/* LD I,A */
	timingsMSX[SHIFT_0xED+0x47] = func(z80 *Z80) uint64 { return 11 }
	
	/* IN C,(C) */
	timingsMSX[SHIFT_0xED+0x48] = func(z80 *Z80) uint64 { return 14 }
	
	/* OUT (C),C */
	timingsMSX[SHIFT_0xED+0x49] = func(z80 *Z80) uint64 { return 14 }
	
	/* ADC HL,BC */
	timingsMSX[SHIFT_0xED+0x4a] = func(z80 *Z80) uint64 { return 17 }
	
	/* LD BC,(nn) */
	timingsMSX[SHIFT_0xED+0x4b] = func(z80 *Z80) uint64 { return 22 }
	
	/* RETI */
	timingsMSX[SHIFT_0xED+0x4d] = func(z80 *Z80) uint64 { return 16 }
	
	/* LD R,A */
	timingsMSX[SHIFT_0xED+0x4f] = func(z80 *Z80) uint64 { return 11 }
	
	/* IN D,(C) */
	timingsMSX[SHIFT_0xED+0x50] = func(z80 *Z80) uint64 { return 14 }
	
	/* OUT (C),D */
	timingsMSX[SHIFT_0xED+0x51] = func(z80 *Z80) uint64 { return 14 }
	
	/* SBC HL,DE */
	timingsMSX[SHIFT_0xED+0x52] = func(z80 *Z80) uint64 { return 17 }
	
	/* LD (nn),DE */
	timingsMSX[SHIFT_0xED+0x53] = func(z80 *Z80) uint64 { return 22 }
	
	/* IM 1 */
	timingsMSX[SHIFT_0xED+0x56] = func(z80 *Z80) uint64 { return 10 }
	
	/* LD A,I */
	timingsMSX[SHIFT_0xED+0x57] = func(z80 *Z80) uint64 { return 11 }
	
	/* IN E,(C) */
	timingsMSX[SHIFT_0xED+0x58] = func(z80 *Z80) uint64 { return 14 }
	
	/* OUT (C),E */
	timingsMSX[SHIFT_0xED+0x59] = func(z80 *Z80) uint64 { return 14 }
	
	/* ADC HL,DE */
	timingsMSX[SHIFT_0xED+0x5a] = func(z80 *Z80) uint64 { return 17 }
	
	/* LD DE,(nn) */
	timingsMSX[SHIFT_0xED+0x5b] = func(z80 *Z80) uint64 { return 22 }
	
	/* IM 2 */
	timingsMSX[SHIFT_0xED+0x5e] = func(z80 *Z80) uint64 { return 10 }
	
	/* LD A,R */
	timingsMSX[SHIFT_0xED+0x5f] = func(z80 *Z80) uint64 { return 11 }
	
	/* IN H,(C) */
	timingsMSX[SHIFT_0xED+0x60] = func(z80 *Z80) uint64 { return 14 }
	
	/* OUT (C),H */
	timingsMSX[SHIFT_0xED+0x61] = func(z80 *Z80) uint64 { return 14 }
	
	/* SBC HL,HL */
	timingsMSX[SHIFT_0xED+0x62] = func(z80 *Z80) uint64 { return 17 }
	
	/* RRD */
	timingsMSX[SHIFT_0xED+0x67] = func(z80 *Z80) uint64 { return 20 }
	
	/* IN L,(C) */
	timingsMSX[SHIFT_0xED+0x68] = func(z80 *Z80) uint64 { return 14 }
	
	/* OUT (C),L */
	timingsMSX[SHIFT_0xED+0x69] = func(z80 *Z80) uint64 { return 14 }
	
	/* ADC HL,HL */
	timingsMSX[SHIFT_0xED+0x6a] = func(z80 *Z80) uint64 { return 17 }
	
	/* RLD */
	timingsMSX[SHIFT_0xED+0x6f] = func(z80 *Z80) uint64 { return 20 }
	
	/* IN F,(C) */
	timingsMSX[SHIFT_0xED+0x70] = func(z80 *Z80) uint64 { return 14 }
	
	/* SBC HL,SP */
	timingsMSX[SHIFT_0xED+0x72] = func(z80 *Z80) uint64 { return 17 }
	
	/* LD (nn),SP */
	timingsMSX[SHIFT_0xED+0x73] = func(z80 *Z80) uint64 { return 22 }
	
	/* IN A,(C) */
	timingsMSX[SHIFT_0xED+0x78] = func(z80 *Z80) uint64 { return 14 }
	
	/* OUT (C),A */
	timingsMSX[SHIFT_0xED+0x79] = func(z80 *Z80) uint64 { return 14 }
	
	/* ADC HL,SP */
	timingsMSX[SHIFT_0xED+0x7a] = func(z80 *Z80) uint64 { return 17 }
	
	/* LD SP,(nn) */
	timingsMSX[SHIFT_0xED+0x7b] = func(z80 *Z80) uint64 { return 22 }
	
	/* LDI */
	timingsMSX[SHIFT_0xED+0xa0] = func(z80 *Z80) uint64 { return 18 }
	
	/* CPI */
	timingsMSX[SHIFT_0xED+0xa1] = func(z80 *Z80) uint64 { return 18 }
	
	/* INI */
	timingsMSX[SHIFT_0xED+0xa2] = func(z80 *Z80) uint64 { return 18 }
	
	/* OUTI */
	timingsMSX[SHIFT_0xED+0xa3] = func(z80 *Z80) uint64 { return 18 }
	
	/* LDD */
	timingsMSX[SHIFT_0xED+0xa8] = func(z80 *Z80) uint64 { return 18 }
	
	/* CPD */
	timingsMSX[SHIFT_0xED+0xa9] = func(z80 *Z80) uint64 { return 18 }
	
	/* IND */
	timingsMSX[SHIFT_0xED+0xaa] = func(z80 *Z80) uint64 { return 18 }
	
	/* OUTD */
	timingsMSX[SHIFT_0xED+0xab] = func(z80 *Z80) uint64 { return 18 }
	
	/* LDIR */
	timingsMSX[SHIFT_0xED+0xb0] = func(z80 *Z80) uint64 { return 0 }
	
	/* CPIR */
	timingsMSX[SHIFT_0xED+0xb1] = func(z80 *Z80) uint64 { return 0 }
	
	/* INIR */
	timingsMSX[SHIFT_0xED+0xb2] = func(z80 *Z80) uint64 { return 0 }
	
	/* OTIR */
	timingsMSX[SHIFT_0xED+0xb3] = func(z80 *Z80) uint64 { return 0 }
	
	/* LDDR */
	timingsMSX[SHIFT_0xED+0xb8] = func(z80 *Z80) uint64 { return 0 }
	
	/* CPDR */
	timingsMSX[SHIFT_0xED+0xb9] = func(z80 *Z80) uint64 { return 0 }
	
	/* INDR */
	timingsMSX[SHIFT_0xED+0xba] = func(z80 *Z80) uint64 { return 0 }
	
	/* OTDR */
	timingsMSX[SHIFT_0xED+0xbb] = func(z80 *Z80) uint64 { return 0 }
	
	/* MULUB A,r */
	timingsMSX[SHIFT_0xED+0xc1] = func(z80 *Z80) uint64 { return 0 }
	
	/* MULUW HL,BC */
	timingsMSX[SHIFT_0xED+0xc3] = func(z80 *Z80) uint64 { return 0 }
	
	/* MULUW HL,SP */
	timingsMSX[SHIFT_0xED+0xf3] = func(z80 *Z80) uint64 { return 0 }
	
	/* XOR n */
	timingsMSX[0xee] = func(z80 *Z80) uint64 { return 8 }
	
	/* RST 28H */
	timingsMSX[0xef] = func(z80 *Z80) uint64 { return 12 }
	
	/* RET P */
	timingsMSX[0xf0] = func(z80 *Z80) uint64 { return 0 }
	
	/* POP AF */
	timingsMSX[0xf1] = func(z80 *Z80) uint64 { return 11 }
	
	/* JP P,nn */
	timingsMSX[0xf2] = func(z80 *Z80) uint64 { return 11 }
	
	/* DI */
	timingsMSX[0xf3] = func(z80 *Z80) uint64 { return 5 }
	
	/* CALL P,nn */
	timingsMSX[0xf4] = func(z80 *Z80) uint64 { return 0 }
	
	/* PUSH AF */
	timingsMSX[0xf5] = func(z80 *Z80) uint64 { return 12 }
	
	/* OR n */
	timingsMSX[0xf6] = func(z80 *Z80) uint64 { return 8 }
	
	/* RST 30H */
	timingsMSX[0xf7] = func(z80 *Z80) uint64 { return 12 }
	
	/* RET M */
	timingsMSX[0xf8] = func(z80 *Z80) uint64 { return 0 }
	
	/* LD SP,HL */
	timingsMSX[0xf9] = func(z80 *Z80) uint64 { return 7 }
	
	/* JP M,nn */
	timingsMSX[0xfa] = func(z80 *Z80) uint64 { return 11 }
	
	/* EI */
	timingsMSX[0xfb] = func(z80 *Z80) uint64 { return 5 }
	
	/* CALL M,nn */
	timingsMSX[0xfc] = func(z80 *Z80) uint64 { return 0 }
	
	/* INC IYq */
	timingsMSX[SHIFT_0xFD+0x04] = func(z80 *Z80) uint64 { return 10 }
	
	/* DEC IYq */
	timingsMSX[SHIFT_0xFD+0x05] = func(z80 *Z80) uint64 { return 10 }
	
	/* ADD IY,BC */
	timingsMSX[SHIFT_0xFD+0x09] = func(z80 *Z80) uint64 { return 17 }
	
	/* ADD IY,DE */
	timingsMSX[SHIFT_0xFD+0x19] = func(z80 *Z80) uint64 { return 17 }
	
	/* LD IY,nn */
	timingsMSX[SHIFT_0xFD+0x21] = func(z80 *Z80) uint64 { return 16 }
	
	/* LD (nn),IY */
	timingsMSX[SHIFT_0xFD+0x22] = func(z80 *Z80) uint64 { return 22 }
	
	/* INC IY */
	timingsMSX[SHIFT_0xFD+0x23] = func(z80 *Z80) uint64 { return 12 }
	
	/* LD IYh,n */
	timingsMSX[SHIFT_0xFD+0x26] = func(z80 *Z80) uint64 { return 13 }
	
	/* ADD IY,IY */
	timingsMSX[SHIFT_0xFD+0x29] = func(z80 *Z80) uint64 { return 17 }
	
	/* LD IY,(nn) */
	timingsMSX[SHIFT_0xFD+0x2a] = func(z80 *Z80) uint64 { return 22 }
	
	/* DEC IY */
	timingsMSX[SHIFT_0xFD+0x2b] = func(z80 *Z80) uint64 { return 12 }
	
	/* LD IYl,n */
	timingsMSX[SHIFT_0xFD+0x2e] = func(z80 *Z80) uint64 { return 13 }
	
	/* INC (IY+o) */
	timingsMSX[SHIFT_0xFD+0x34] = func(z80 *Z80) uint64 { return 25 }
	
	/* DEC (IY+o) */
	timingsMSX[SHIFT_0xFD+0x35] = func(z80 *Z80) uint64 { return 25 }
	
	/* LD (IY+o),n */
	timingsMSX[SHIFT_0xFD+0x36] = func(z80 *Z80) uint64 { return 21 }
	
	/* ADD IY,SP */
	timingsMSX[SHIFT_0xFD+0x39] = func(z80 *Z80) uint64 { return 17 }
	
	/* LD B,IYq */
	timingsMSX[SHIFT_0xFD+0x40] = func(z80 *Z80) uint64 { return 10 }
	
	/* LD B,(IY+o) */
	timingsMSX[SHIFT_0xFD+0x46] = func(z80 *Z80) uint64 { return 21 }
	
	/* LD C,IYq */
	timingsMSX[SHIFT_0xFD+0x48] = func(z80 *Z80) uint64 { return 10 }
	
	/* LD C,(IY+o) */
	timingsMSX[SHIFT_0xFD+0x4e] = func(z80 *Z80) uint64 { return 21 }
	
	/* LD D,IYq */
	timingsMSX[SHIFT_0xFD+0x50] = func(z80 *Z80) uint64 { return 10 }
	
	/* LD D,(IY+o) */
	timingsMSX[SHIFT_0xFD+0x56] = func(z80 *Z80) uint64 { return 21 }
	
	/* LD E,IYq */
	timingsMSX[SHIFT_0xFD+0x58] = func(z80 *Z80) uint64 { return 10 }
	
	/* LD E,(IY+o) */
	timingsMSX[SHIFT_0xFD+0x5e] = func(z80 *Z80) uint64 { return 21 }
	
	/* LD IYh,q */
	timingsMSX[SHIFT_0xFD+0x60] = func(z80 *Z80) uint64 { return 10 }
	
	/* LD H,(IY+o) */
	timingsMSX[SHIFT_0xFD+0x66] = func(z80 *Z80) uint64 { return 21 }
	
	/* LD IYl,q */
	timingsMSX[SHIFT_0xFD+0x68] = func(z80 *Z80) uint64 { return 10 }
	
	/* LD L,(IY+o) */
	timingsMSX[SHIFT_0xFD+0x6e] = func(z80 *Z80) uint64 { return 21 }
	
	/* LD (IY+o),r */
	timingsMSX[SHIFT_0xFD+0x70] = func(z80 *Z80) uint64 { return 21 }
	
	/* LD A,IYq */
	timingsMSX[SHIFT_0xFD+0x78] = func(z80 *Z80) uint64 { return 10 }
	
	/* LD A,(IY+o) */
	timingsMSX[SHIFT_0xFD+0x7e] = func(z80 *Z80) uint64 { return 21 }
	
	/* ADD A,IYq */
	timingsMSX[SHIFT_0xFD+0x80] = func(z80 *Z80) uint64 { return 10 }
	
	/* ADD A,(IY+o) */
	timingsMSX[SHIFT_0xFD+0x86] = func(z80 *Z80) uint64 { return 21 }
	
	/* ADC A,IYq */
	timingsMSX[SHIFT_0xFD+0x88] = func(z80 *Z80) uint64 { return 10 }
	
	/* ADC A,(IY+o) */
	timingsMSX[SHIFT_0xFD+0x8e] = func(z80 *Z80) uint64 { return 21 }
	
	/* SUB IYq */
	timingsMSX[SHIFT_0xFD+0x90] = func(z80 *Z80) uint64 { return 10 }
	
	/* SUB (IY+o) */
	timingsMSX[SHIFT_0xFD+0x96] = func(z80 *Z80) uint64 { return 21 }
	
	/* SBC A,IYq */
	timingsMSX[SHIFT_0xFD+0x98] = func(z80 *Z80) uint64 { return 10 }
	
	/* SBC A,(IY+o) */
	timingsMSX[SHIFT_0xFD+0x9e] = func(z80 *Z80) uint64 { return 21 }
	
	/* AND IYq */
	timingsMSX[SHIFT_0xFD+0xa0] = func(z80 *Z80) uint64 { return 10 }
	
	/* AND (IY+o) */
	timingsMSX[SHIFT_0xFD+0xa6] = func(z80 *Z80) uint64 { return 21 }
	
	/* XOR IYq */
	timingsMSX[SHIFT_0xFD+0xa8] = func(z80 *Z80) uint64 { return 10 }
	
	/* XOR (IY+o) */
	timingsMSX[SHIFT_0xFD+0xae] = func(z80 *Z80) uint64 { return 21 }
	
	/* OR IYq */
	timingsMSX[SHIFT_0xFD+0xb0] = func(z80 *Z80) uint64 { return 10 }
	
	/* OR (IY+o) */
	timingsMSX[SHIFT_0xFD+0xb6] = func(z80 *Z80) uint64 { return 21 }
	
	/* CP IYq */
	timingsMSX[SHIFT_0xFD+0xb8] = func(z80 *Z80) uint64 { return 10 }
	
	/* CP (IY+o) */
	timingsMSX[SHIFT_0xFD+0xbe] = func(z80 *Z80) uint64 { return 21 }
	
	/* RLC (IY+o) */
	timingsMSX[SHIFT_0xDDCB+0x06] = func(z80 *Z80) uint64 { return 25 }
	
	/* RRC (IY+o) */
	timingsMSX[SHIFT_0xDDCB+0x0e] = func(z80 *Z80) uint64 { return 25 }
	
	/* RL (IY+o) */
	timingsMSX[SHIFT_0xDDCB+0x16] = func(z80 *Z80) uint64 { return 25 }
	
	/* RR (IY+o) */
	timingsMSX[SHIFT_0xDDCB+0x1e] = func(z80 *Z80) uint64 { return 25 }
	
	/* SLA (IY+o) */
	timingsMSX[SHIFT_0xDDCB+0x26] = func(z80 *Z80) uint64 { return 25 }
	
	/* SRA (IY+o) */
	timingsMSX[SHIFT_0xDDCB+0x2e] = func(z80 *Z80) uint64 { return 25 }
	
	/* SRL (IY+o) */
	timingsMSX[SHIFT_0xDDCB+0x3e] = func(z80 *Z80) uint64 { return 25 }
	
	/* BIT b,(IY+o) */
	timingsMSX[SHIFT_0xDDCB+0x46] = func(z80 *Z80) uint64 { return 22 }
	
	/* RES b,(IY+o) */
	timingsMSX[SHIFT_0xDDCB+0x86] = func(z80 *Z80) uint64 { return 25 }
	
	/* SET b,(IY+o) */
	timingsMSX[SHIFT_0xDDCB+0xc6] = func(z80 *Z80) uint64 { return 25 }
	
	/* POP IY */
	timingsMSX[SHIFT_0xFD+0xe1] = func(z80 *Z80) uint64 { return 16 }
	
	/* EX (SP),IY */
	timingsMSX[SHIFT_0xFD+0xe3] = func(z80 *Z80) uint64 { return 25 }
	
	/* PUSH IY */
	timingsMSX[SHIFT_0xFD+0xe5] = func(z80 *Z80) uint64 { return 17 }
	
	/* JP (IY) */
	timingsMSX[SHIFT_0xFD+0xe9] = func(z80 *Z80) uint64 { return 10 }
	
	/* LD SP,IY */
	timingsMSX[SHIFT_0xFD+0xf9] = func(z80 *Z80) uint64 { return 12 }
	
	/* CP n */
	timingsMSX[0xfe] = func(z80 *Z80) uint64 { return 8 }
	
	/* RST 38H */
	timingsMSX[0xff] = func(z80 *Z80) uint64 { return 12 }
	

    // CALL C, nn
    timingsMSX[0xdc] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_C) != 0 {
            return 18
        }
        return 11
    }

    // CALL M, nn
    timingsMSX[0xfc] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_S) != 0 {
            return 18
        }
        return 11
    }

    // CALL NC, nn
    timingsMSX[0xd4] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_C) == 0 {
            return 18
        }
        return 11
    }

    // CALL NZ, nn
    timingsMSX[0xc4] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_Z) == 0 {
            return 18
        }
        return 11
    }

    // CALL P, nn
    timingsMSX[0xf4] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_S) == 0 {
            return 18
        }
        return 11
    }

    // CALL PE, nn
    timingsMSX[0xec] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_P) != 0 {
            return 18
        }
        return 11
    }

    // CALL PO, nn
    timingsMSX[0xe4] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_P) == 0 {
            return 18
        }
        return 11
    }

    // CALL Z, nn
    timingsMSX[0xe4] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_Z) != 0 {
            return 18
        }
        return 11
    }

    // DJNZ o
    timingsMSX[0x10] = func(z80 *Z80) uint64 {
        if (z80.B != 0) {
            return 14
        }
        return 9
    }

    // JR C, nn
    timingsMSX[0x38] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_C) != 0 {
            return 13
        }
        return 8
    }

    // JR NC, nn
    timingsMSX[0x30] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_C) == 0 {
            return 13
        }
        return 8
    }

    // JR Z, nn
    timingsMSX[0x28] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_Z) != 0 {
            return 13
        }
        return 8
    }

    // JR NZ, nn
    timingsMSX[0x20] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_Z) == 0 {
            return 13
        }
        return 8
    }

    // RET C
    timingsMSX[0xd8] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_C) != 0 {
            return 12
        }
        return 6
    }

    // RET M
    timingsMSX[0xf8] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_S) != 0 {
            return 12
        }
        return 6
    }

    // RET NC
    timingsMSX[0xd0] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_C) == 0 {
            return 12
        }
        return 6
    }

    // RET NZ
    timingsMSX[0xc0] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_Z) == 0 {
            return 12
        }
        return 6
    }

    // RET P
    timingsMSX[0xf0] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_S) == 0 {
            return 12
        }
        return 6
    }

    // RET PE
    timingsMSX[0xe8] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_P) != 0 {
            return 12
        }
        return 6
    }

    // RET PO
    timingsMSX[0xe0] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_P) == 0 {
            return 12
        }
        return 6
    }

     // RET Z
    timingsMSX[0xc8] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_Z) != 0 {
            return 12
        }
        return 6
    }


}
