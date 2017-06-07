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
	scm := vdp_screenMode
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
	}
	if scm != vdp_screenMode {
		log.Printf("Change screen mode: %d\n", vdp_screenMode)
		graphics_setLogicalResolution()
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
				// log.Printf("vdp[%d] = %02x\n", regn, vdp_valueRead)
				// log.Printf("VDPS: %v\n", vdp_registers)
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

func vdp_updateBuffer() {
	if !vdp_screenEnabled {
		return
	}
	nameTable := vdp_VRAM[(uint16(vdp_registers[2]) << 10):]
	patTable := vdp_VRAM[(uint16(vdp_registers[4]) << 11):]
	colorTable := vdp_VRAM[(uint16(vdp_registers[3]) << 6):]
	switch {
	case vdp_screenMode == SCREEN0:
		// Render SCREEN0 (40x24)
		color1 := int((vdp_registers[7] & 0xF0) >> 4)
		color2 := int((vdp_registers[7] & 0x0F))
		for y := 0; y < 24; y++ {
			for x := 0; x < 40; x++ {
				vdp_drawPatternsS0(x*8, y*8, int(nameTable[x+y*40])*8, patTable, color1, color2)
			}
		}
		break

	case vdp_screenMode == SCREEN1:
		// Render SCREEN1 (32x24)
		for y := 0; y < 24; y++ {
			for x := 0; x < 32; x++ {
				pat := int(nameTable[x+y*32])
				vdp_drawPatternsS1(x*8, y*8, pat*8, patTable, colorTable[pat/8])
			}
		}
		vdp_drawSprites()
		break

	case vdp_screenMode == SCREEN2:
		// Render SCREEN2
		// Pattern table: 0000H to 17FFH
		patTable := vdp_VRAM[(uint16(vdp_registers[4]&0x04) << 11):]
		colorTable := vdp_VRAM[(uint16(vdp_registers[3]&0x80) << 6):]
		for y := 0; y < 24; y++ {
			for x := 0; x < 32; x++ {
				pat := int(nameTable[x+y*32])
				vdp_drawPatternsS2(x*8, y*8, pat*8, patTable, colorTable)
			}
		}
		vdp_drawSprites()
		break

	case vdp_screenMode == SCREEN3:
		// Render SCREEN3
		log.Println("Drawing in screen3 not implemented yet")
		break

	default:
		panic("RenderScreen: impossible mode")

	}
}

func vdp_drawPatternsS0(x, y int, pt int, patTable []byte, color1, color2 int) {
	var mask byte
	for i := 0; i < 8; i++ {
		b := patTable[i+pt]
		xx := 0
		for mask = 0x80; mask > 0; mask >>= 1 {
			if mask&b != 0 {
				graphics_drawPixel(x+xx, y+i, color1)
			} else {
				graphics_drawPixel(x+xx, y+i, color2)
			}
			xx++
		}
	}
}
func vdp_drawPatternsS1(x, y int, pt int, patTable []byte, color byte) {
	color1 := int((color & 0xF0) >> 4)
	color2 := int(color & 0x0F)
	var mask byte
	for i := 0; i < 8; i++ {
		b := patTable[i+pt]
		xx := 0
		for mask = 0x80; mask > 0; mask >>= 1 {
			if mask&b != 0 {
				graphics_drawPixel(x+xx, y+i, color1)
			} else {
				graphics_drawPixel(x+xx, y+i, color2)
			}
			xx++
		}
	}
}

func vdp_drawPatternsS2(x, y int, pt int, patTable []byte, colorTable []byte) {
	var mask byte
	var b byte
	var color byte
	for i := 0; i < 8; i++ {
		if y < 64 {
			b = patTable[i+pt]
			color = colorTable[i+pt]
		} else if y < 128 {
			b = patTable[i+pt+2048]
			color = colorTable[i+pt+2048]
		} else {
			b = patTable[i+pt+2048*2]
			color = colorTable[i+pt+2048*2]
		}
		color1 := int((color & 0xF0) >> 4)
		color2 := int(color & 0x0F)
		xx := 0
		for mask = 0x80; mask > 0; mask >>= 1 {
			if mask&b != 0 {
				graphics_drawPixel(x+xx, y+i, color1)
			} else {
				graphics_drawPixel(x+xx, y+i, color2)
			}
			xx++
		}
	}
}

func vdp_drawSprites() {
	// Sprite name table: 1B00H to 1B7FH
	// Sprite pattern table: 3800H to 3FFFH
	sprTable := vdp_VRAM[(uint16(vdp_registers[5]) << 7):]
	sprPatTable := vdp_VRAM[(uint16(vdp_registers[6]) << 11):]
	magnif := (vdp_registers[1] & 0x01) != 0
	spr16x16 := (vdp_registers[1] & 0x02) != 0
	for i, j := 0, 0; i < 32; i, j = i+1, j+4 {
		ypos := int(sprTable[j])
		xpos := int(sprTable[j+1])
		patn := sprTable[j+2]
		ec := (sprTable[j+3] & 0x80) != 0
		color := int(sprTable[j+3] & 0x0F)
		if !spr16x16 {
			patt := sprPatTable[uint16(patn)*8:]
			drawSpr(magnif, xpos, ypos, patt, ec, color)
		} else {
			patt := sprPatTable[uint16((patn>>2))*8*4:]
			drawSpr(magnif, xpos, ypos, patt, ec, color)
			drawSpr(magnif, xpos, ypos+8, patt[8:], ec, color)
			drawSpr(magnif, xpos+8, ypos, patt[16:], ec, color)
			drawSpr(magnif, xpos+8, ypos+8, patt[24:], ec, color)
		}
	}
}

// TODO: sprite magnification not implemented
func drawSpr(magnif bool, xpos, ypos int, patt []byte, ec bool, color int) {
	if ypos > 191 {
		return
	}

	for y := 0; y < 8; y++ {
		b := patt[y]
		for x, mask := 0, byte(0x80); mask > 0; mask >>= 1 {
			if mask&b != 0 {
				graphics_drawPixel(xpos+x, ypos+y, color)
			}
			x++
		}
	}
}
