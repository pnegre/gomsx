package main

import "log"

var vdp_valueRead byte
var vdp_writeState = 0
var vdp_registers [27]byte
var vdp_writeToVRAM bool
var vdp_VRAM [0x10000]byte
var vdp_pointerVRAM uint16

func vdp_writePort(ad byte, val byte) {
	log.Printf("VDP: Out(%02x, %02x)\n", ad, val)
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
		log.Printf("Writing to VRAM: %04x -> %02x", vdp_pointerVRAM, val)
		vdp_VRAM[vdp_pointerVRAM] = val
		vdp_pointerVRAM++
		return

	}
	log.Fatalf("Not implemented: VDP: Out(%02x, %02x)", ad, val)
}
