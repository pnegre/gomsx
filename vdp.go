package main

import "log"

const (
	SCREEN0 = 0
	SCREEN1 = 1
	SCREEN2 = 2
	SCREEN3 = 3
)

type Vdp struct {
	screenEnabled     bool
	screenMode        int
	valueRead         byte
	writeState        int
	enabledInterrupts bool
	registers         [10]byte
	writeToVRAM       bool
	vram              [0x10000]byte
	pointerVRAM       uint16
	statusReg         byte
}

func NewVdp() *Vdp {
	return &Vdp{}
}

func (vdp *Vdp) setFrameFlag() {
	vdp.statusReg |= 0x80
}

func (vdp *Vdp) updateRegisters() {
	vdp.screenEnabled = vdp.registers[1]&0x40 != 0
	vdp.enabledInterrupts = vdp.registers[1]&0x20 != 0
	m1 := vdp.registers[1]&0x10 != 0
	m2 := vdp.registers[1]&0x08 != 0
	m3 := vdp.registers[0]&0x02 != 0
	m4 := vdp.registers[0]&0x04 != 0
	m5 := vdp.registers[0]&0x08 != 0
	scm := vdp.screenMode
	switch {
	case m1 == false && m2 == false && m3 == false && m4 == false && m5 == false:
		vdp.screenMode = SCREEN1
		break

	case m1 == true && m2 == false && m3 == false && m4 == false && m5 == false:
		vdp.screenMode = SCREEN0
		break

	case m1 == false && m2 == true && m3 == false && m4 == false && m5 == false:
		vdp.screenMode = SCREEN3
		break

	case m1 == false && m2 == false && m3 == true && m4 == false && m5 == false:
		vdp.screenMode = SCREEN2
		break
	}
	if scm != vdp.screenMode {
		log.Printf("Change screen mode: %d\n", vdp.screenMode)
		graphics_setLogicalResolution(vdp.screenMode)
	}
}

func (vdp *Vdp) writePort(ad byte, val byte) {
	//log.Printf("VDP: Out(%02x, %02x)\n", ad, val)
	switch {
	case ad == 0x99:
		if vdp.writeState == 0 {
			vdp.valueRead = val
			vdp.writeState = 1
			return
		} else {
			vdp.writeState = 0
			// Bit 7 must be 1 for write
			if val&0x80 != 0 {
				regn := val - 128
				// log.Printf("vdp[%d] = %02x\n", regn, vdp.valueRead)
				vdp.registers[regn] = vdp.valueRead
				// log.Printf("VDPS: %v\n", vdp.registers)
				vdp.updateRegisters()
				return
			} else {
				vdp.writeToVRAM = (val&0x40 != 0)
				val &= 0xBF
				vdp.pointerVRAM = 0
				vdp.pointerVRAM |= uint16(vdp.valueRead)
				vdp.pointerVRAM |= uint16(val) << 8
				return
			}
		}

	case ad == 0x98:
		// Writing to VRAM
		//log.Printf("Writing to VRAM: %04x -> %02x", vdp.pointerVRAM, val)
		vdp.vram[vdp.pointerVRAM] = val
		vdp.pointerVRAM++
		return

	}
	log.Fatalf("Not implemented: VDP: Out(%02x, %02x)", ad, val)
}

func (vdp *Vdp) readPort(ad byte) byte {
	switch {
	case ad == 0x98:
		// Reading from VRAM
		//log.Printf("Reading from VRAM: %04x", vdp.pointerVRAM)
		r := vdp.vram[vdp.pointerVRAM]
		vdp.pointerVRAM++
		return r

	case ad == 0x99:
		// Reading status register
		// TODO: mirar-ho bé....
		var r = vdp.statusReg
		vdp.statusReg &= 0x7F // Clear frame flag
		return r
	}

	log.Fatalf("Not implemented: VDP: In(%02x)", ad)
	return 0
}

func (vdp *Vdp) updateBuffer() {
	if !vdp.screenEnabled {
		return
	}
	nameTable := vdp.vram[(uint16(vdp.registers[2]) << 10):]
	patTable := vdp.vram[(uint16(vdp.registers[4]) << 11):]
	colorTable := vdp.vram[(uint16(vdp.registers[3]) << 6):]
	switch {
	case vdp.screenMode == SCREEN0:
		// Render SCREEN0 (40x24)
		color1 := int((vdp.registers[7] & 0xF0) >> 4)
		color2 := int((vdp.registers[7] & 0x0F))
		for y := 0; y < 24; y++ {
			for x := 0; x < 40; x++ {
				vdp.drawPatternsS0(x*8, y*8, int(nameTable[x+y*40])*8, patTable, color1, color2)
			}
		}
		break

	case vdp.screenMode == SCREEN1:
		// Render SCREEN1 (32x24)
		for y := 0; y < 24; y++ {
			for x := 0; x < 32; x++ {
				pat := int(nameTable[x+y*32])
				vdp.drawPatternsS1(x*8, y*8, pat*8, patTable, colorTable[pat/8])
			}
		}
		vdp.drawSprites()
		break

	case vdp.screenMode == SCREEN2:
		// Render SCREEN2
		// Pattern table: 0000H to 17FFH
		patTable := vdp.vram[(uint16(vdp.registers[4]&0x04) << 11):]
		colorTable := vdp.vram[(uint16(vdp.registers[3]&0x80) << 6):]
		for y := 0; y < 24; y++ {
			for x := 0; x < 32; x++ {
				pat := int(nameTable[x+y*32])
				vdp.drawPatternsS2(x*8, y*8, pat*8, patTable, colorTable)
			}
		}
		vdp.drawSprites()
		break

	case vdp.screenMode == SCREEN3:
		// Render SCREEN3
		log.Println("Drawing in screen3 not implemented yet")
		break

	default:
		panic("RenderScreen: impossible mode")

	}
}

func (vdp *Vdp) drawPatternsS0(x, y int, pt int, patTable []byte, color1, color2 int) {
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
func (vdp *Vdp) drawPatternsS1(x, y int, pt int, patTable []byte, color byte) {
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

func (vdp *Vdp) drawPatternsS2(x, y int, pt int, patTable []byte, colorTable []byte) {
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

func (vdp *Vdp) drawSprites() {
	// Sprite name table: 1B00H to 1B7FH
	// Sprite pattern table: 3800H to 3FFFH
	sprTable := vdp.vram[(uint16(vdp.registers[5]) << 7):]
	sprPatTable := vdp.vram[(uint16(vdp.registers[6]) << 11):]
	magnif := (vdp.registers[1] & 0x01) != 0
	spr16x16 := (vdp.registers[1] & 0x02) != 0
	for i, j := 0, 0; i < 32; i, j = i+1, j+4 {
		ypos := int(sprTable[j])
		if ypos == 0xd0 {
			// Ignore all sprites
			return
		}
		xpos := int(sprTable[j+1])
		patn := sprTable[j+2]
		ec := (sprTable[j+3] & 0x80) != 0
		if ec {
			xpos -= 32
		}
		color := int(sprTable[j+3] & 0x0F)
		if !spr16x16 {
			patt := sprPatTable[uint16(patn)*8:]
			vdp.drawSpr(magnif, xpos, ypos, patt, ec, color)
		} else {
			patt := sprPatTable[uint16((patn>>2))*8*4:]
			vdp.drawSpr(magnif, xpos, ypos, patt, ec, color)
			vdp.drawSpr(magnif, xpos, ypos+8, patt[8:], ec, color)
			vdp.drawSpr(magnif, xpos+8, ypos, patt[16:], ec, color)
			vdp.drawSpr(magnif, xpos+8, ypos+8, patt[24:], ec, color)
		}
	}
}

// TODO: sprite magnification not implemented
func (vdp *Vdp) drawSpr(magnif bool, xpos, ypos int, patt []byte, ec bool, color int) {
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
