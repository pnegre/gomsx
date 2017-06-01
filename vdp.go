package main

import "log"
import "github.com/pnegre/gogame"

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

var vdp_registers [8]byte
var vdp_writeToVRAM bool
var vdp_VRAM [0x10000]byte
var vdp_pointerVRAM uint16

func vdp_updateRegisters() {
	vdp_screenEnabled = vdp_registers[1]&0x40 != 0
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
	}

	log.Fatalf("Not implemented: VDP: In(%02x)", ad)
	return 0
}

func vdp_renderScreen() {
	if !vdp_screenEnabled {
		return
	}

	switch {
	case vdp_screenMode == SCREEN0:
		// Render SCREEN0
		// Pattern table: 0x0800 - 0x0FFF
		// Name table: 0x0000 - 0x03BF
		patTable := vdp_VRAM[0x800 : 0xFFF+1]
		doPattern := func(x, y int, pt int) {
			for i := 0; i < 8; i++ {
				b := patTable[i+pt]
				xx := 0
				for j := 0x80; j > 0; j >>= 1 {
					if byte(j)&b != 0 {
						gogame.DrawPixel(x+xx, y+i, gogame.COLOR_WHITE)
					}
					xx++
				}
			}
		}

		nameTable := vdp_VRAM[0x000 : 0x03BF+1]
		for y := 0; y < 5; y++ {
			for x := 0; x < 40; x++ {
				doPattern(x*8, y*8, int(nameTable[x+y*40])*8)
			}
		}
		//doPattern(100,100)
		return

	case vdp_screenMode == SCREEN1:
		// Render SCREEN1
		return

	case vdp_screenMode == SCREEN2:
		// Render SCREEN2
		return

	case vdp_screenMode == SCREEN3:
		// Render SCREEN3
		return

	default:
		panic("RenderScreen: impossible mode")

	}
}
