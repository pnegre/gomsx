package main

import "log"

const (
	SCREEN0 = 0
	SCREEN1 = 1
	SCREEN2 = 2
	SCREEN3 = 3
)

var vdp_screenEnabled bool = false
var vdp_screenMode int
var vdp_valueRead byte
var vdp_writeState = 0
var vdp_enabledInterrupts = false
var vdp_registers [8]byte
var vdp_writeToVRAM bool
var vdp_VRAM [0x10000]byte
var vdp_pointerVRAM uint16
var vdp_statusReg byte = 0

func vdp_setFrameFlag() {
	vdp_statusReg |= 0x80
}

func vdp_updateRegisters() {
	vdp_screenEnabled = vdp_registers[1]&0x40 != 0
	vdp_enabledInterrupts = vdp_registers[1]&0x20 != 0
	m1 := vdp_registers[1]&0x10 != 0
	m2 := vdp_registers[1]&0x08 != 0
	m3 := vdp_registers[0]&0x02 != 0
	m4 := vdp_registers[0]&0x04 != 0
	m5 := vdp_registers[0]&0x08 != 0
	switch {
	case m1 == false && m2 == false && m3 == false && m4 == false && m5 == false:
		vdp_screenMode = SCREEN1
		break

	case m1 == true && m2 == false && m3 == false && m4 == false && m5 == false:
		vdp_screenMode = SCREEN0
		break

	case m1 == false && m2 == true && m3 == false && m4 == false && m5 == false:
		vdp_screenMode = SCREEN3
		break

	case m1 == false && m2 == false && m3 == true && m4 == false && m5 == false:
		vdp_screenMode = SCREEN2
		break

	default:
		println(m1)
		println(m2)
		println(m3)
		println(m4)
		println(m5)
		panic("VDP: Screen mode not implemented")
	}
}

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
				regn := val - 128
				vdp_registers[regn] = vdp_valueRead
				log.Printf("vdp[%d] = %02x\n", regn, vdp_valueRead)
				log.Printf("VDPS: %v\n", vdp_registers)
				vdp_updateRegisters()
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

	case ad == 0x99:
		// Reading status register
		// TODO: mirar-ho bé....
		var r = vdp_statusReg
		vdp_statusReg &= 0x7F // Clear frame flag
		return r
	}

	log.Fatalf("Not implemented: VDP: In(%02x)", ad)
	return 0
}
