package main

import "log"

var vdp_valueRead byte
var vdp_writeState = 0

// 8 registres, però van de 1 a 8, no de 0 a 7...
var vdp_registers [9]byte
var vdp_writeToVRAM bool
var vdp_VRAM [0x10000]byte
var vdp_pointerVRAM uint16

func vdp_writePort(ad byte, val byte) {
	//log.Printf("VDP: Out(%02x, %02x)\n", ad, val)
	switch {
	case ad == 0x99:
		if vdp_writeState == 0 {
			vdp_valueRead = val
			vdp_writeState = 1
			return
		} else {
			vdp_writeState = 0
			// Bit 7 must be 1 for write
			if val&0x80 != 0 {
				regn := val - 127
				vdp_registers[regn] = vdp_valueRead
				log.Printf("vdp[%d] = %02x\n", regn, vdp_valueRead)
				log.Printf("VDPS: %v\n", vdp_registers)
				return
			} else {
				vdp_writeToVRAM = (val&0x40 != 0)
				val &= 0xBF
				vdp_pointerVRAM = 0
				vdp_pointerVRAM |= uint16(vdp_valueRead)
				vdp_pointerVRAM |= uint16(val) << 8
				return
			}
		}

	case ad == 0x98:
		// Writing to VRAM
		//log.Printf("Writing to VRAM: %04x -> %02x", vdp_pointerVRAM, val)
		vdp_VRAM[vdp_pointerVRAM] = val
		vdp_pointerVRAM++
		return

	}
	log.Fatalf("Not implemented: VDP: Out(%02x, %02x)", ad, val)
}

func vdp_readPort(ad byte) byte {
	switch {
	case ad == 0x98:
		// Reading from VRAM
		//log.Printf("Reading from VRAM: %04x", vdp_pointerVRAM)
		r := vdp_VRAM[vdp_pointerVRAM]
		vdp_pointerVRAM++
		return r
	}

	log.Fatalf("Not implemented: VDP: In(%02x)", ad)
	return 0
}
