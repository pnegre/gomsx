package main

import "github.com/pnegre/gogame"

func graphics_renderScreen() {
	if !vdp_screenEnabled {
		return
	}

	gogame.RenderClear()
	switch {
	case vdp_screenMode == SCREEN0:
		// Render SCREEN0 (40x24)
		// Pattern table: 0x0800 - 0x0FFF
		// Name table: 0x0000 - 0x03BF
		patTable := vdp_VRAM[0x800 : 0xFFF+1]
		nameTable := vdp_VRAM[0x000 : 0x03BF+1]
		for y := 0; y < 5; y++ {
			for x := 0; x < 40; x++ {
				graphics_drawPattern(x*8, y*8, int(nameTable[x+y*40])*8, patTable)
			}
		}
		break

	case vdp_screenMode == SCREEN1:
		// Render SCREEN1 (32x24)
		// Pattern table: 0x0000 - 0x07FF
		// Name table: 0x1800 - 0x1AFF
		patTable := vdp_VRAM[0x0000 : 0x07FF+1]
		nameTable := vdp_VRAM[0x1800 : 0x1AFF+1]
		for y := 0; y < 24; y++ {
			for x := 0; x < 32; x++ {
				graphics_drawPattern(x*8, y*8, int(nameTable[x+y*32])*8, patTable)
			}
		}
		break

	case vdp_screenMode == SCREEN2:
		// Render SCREEN2
		break

	case vdp_screenMode == SCREEN3:
		// Render SCREEN3
		break

	default:
		panic("RenderScreen: impossible mode")

	}
	gogame.RenderPresent()
}

func graphics_drawPattern(x, y int, pt int, patTable []byte) {
	var mask byte
	for i := 0; i < 8; i++ {
		b := patTable[i+pt]
		xx := 0
		for mask = 0x80; mask > 0; mask >>= 1 {
			if mask&b != 0 {
				gogame.DrawPixel(x+xx, y+i, gogame.COLOR_WHITE)
			}
			xx++
		}
	}
}
