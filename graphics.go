package main

import "github.com/pnegre/gogame"

var colors []*gogame.Color

func init() {
	colors = []*gogame.Color{
		&gogame.Color{255, 255, 255, 0},      // Transparent
		&gogame.Color{0, 0, 0, 255},          // Black
		&gogame.Color{0x20, 0xc8, 0x40, 255}, // Green
		&gogame.Color{0x58, 0xd8, 0x78, 255}, // Light Green
		&gogame.Color{0x50, 0x50, 0xe8, 255}, // Dark Blue
		&gogame.Color{0x78, 0x70, 0xf7, 255}, // Light Blue
		&gogame.Color{0xd0, 0x50, 0x48, 255}, // Dark Red
		&gogame.Color{0x40, 0xe8, 0xf0, 255}, // Cyan
		&gogame.Color{0xf7, 0x50, 0x50, 255}, // Red
		&gogame.Color{0xf7, 0x78, 0x78, 255}, // Bright Red
		&gogame.Color{0xd0, 0xc0, 0x50, 255}, // Yellow
		&gogame.Color{0xe0, 0xc8, 0x80, 255}, // Light Yellow
		&gogame.Color{0x20, 0xb0, 0x38, 255}, // Dark Green
		&gogame.Color{0xc8, 0x58, 0xb8, 255}, // Purple
		&gogame.Color{0xc8, 0xc8, 0xc8, 255}, // Gray
		&gogame.Color{0xf7, 0xf7, 0xf7, 255}, // White
	}
}

func graphics_setLogicalResolution() {
	switch vdp_screenMode {
	case SCREEN0:
		gogame.SetLogicalSize(320, 192)
		return
	case SCREEN2:
		gogame.SetLogicalSize(256, 192)
		return
	case SCREEN1:
		gogame.SetLogicalSize(256, 192)
		return
	}
	println(vdp_screenMode)
	panic("setLogicalResolution: mode not supported")
}

func graphics_renderScreen() {
	if !vdp_screenEnabled {
		return
	}
	vdp_setFrameFlag()
	gogame.RenderClear()
	switch {
	case vdp_screenMode == SCREEN0:
		// Render SCREEN0 (40x24)
		// Pattern table: 0x0800 - 0x0FFF
		// Name table: 0x0000 - 0x03BF
		color1 := colors[(vdp_registers[7]&0xF0)>>4]
		color2 := colors[(vdp_registers[7] & 0x0F)]
		patTable := vdp_VRAM[0x800 : 0xFFF+1]
		nameTable := vdp_VRAM[0x000 : 0x03BF+1]
		for y := 0; y < 24; y++ {
			for x := 0; x < 40; x++ {
				graphics_drawPatternS0(x*8, y*8, int(nameTable[x+y*40])*8, patTable, color1, color2)
			}
		}
		break

	case vdp_screenMode == SCREEN1:
		// Render SCREEN1 (32x24)
		// Pattern table: 0x0000 - 0x07FF
		// Name table: 0x1800 - 0x1AFF
		// Color table: 0x2000 - 0x201F.
		patTable := vdp_VRAM[0x0000 : 0x07FF+1]
		nameTable := vdp_VRAM[0x1800 : 0x1AFF+1]
		colorTable := vdp_VRAM[0x2000 : 0x201F+1]
		for y := 0; y < 24; y++ {
			for x := 0; x < 32; x++ {
				pat := int(nameTable[x+y*32])
				graphics_drawPatternS1(x*8, y*8, pat*8, patTable, colorTable[pat/8])
			}
		}
		break

	case vdp_screenMode == SCREEN2:
		// Render SCREEN2
		// Pattern table: 0000H to 17FFH
		// Name table: 1800H to 1AFFH
		// Color table: 2000H to 37FFH
		patTable := vdp_VRAM[0x0000 : 0x17FF+1]
		nameTable := vdp_VRAM[0x1800 : 0x1AFF+1]
		colorTable := vdp_VRAM[0x2000 : 0x37FF+1]
		for y := 0; y < 24; y++ {
			for x := 0; x < 32; x++ {
				pat := int(nameTable[x+y*32])
				graphics_drawPatternS2(x*8, y*8, pat*8, patTable, colorTable)
			}
		}
		break

	case vdp_screenMode == SCREEN3:
		// Render SCREEN3
		break

	default:
		panic("RenderScreen: impossible mode")

	}
	gogame.RenderPresent()
}

func graphics_drawPatternS0(x, y int, pt int, patTable []byte, color1, color2 *gogame.Color) {
	var mask byte
	for i := 0; i < 8; i++ {
		b := patTable[i+pt]
		xx := 0
		for mask = 0x80; mask > 0; mask >>= 1 {
			if mask&b != 0 {
				gogame.DrawPixel(x+xx, y+i, color1)
			} else {
				gogame.DrawPixel(x+xx, y+i, color2)
			}
			xx++
		}
	}
}
func graphics_drawPatternS1(x, y int, pt int, patTable []byte, color byte) {
	color1 := colors[(color&0xF0)>>4]
	color2 := colors[color&0x0F]
	var mask byte
	for i := 0; i < 8; i++ {
		b := patTable[i+pt]
		xx := 0
		for mask = 0x80; mask > 0; mask >>= 1 {
			if mask&b != 0 {
				gogame.DrawPixel(x+xx, y+i, color1)
			} else {
				gogame.DrawPixel(x+xx, y+i, color2)
			}
			xx++
		}
	}
}

func graphics_drawPatternS2(x, y int, pt int, patTable []byte, colorTable []byte) {
	var mask byte
	for i := 0; i < 8; i++ {
		b := patTable[i+pt]
		color := colorTable[i+pt]
		color1 := colors[(color&0xF0)>>4]
		color2 := colors[color&0x0F]
		xx := 0
		for mask = 0x80; mask > 0; mask >>= 1 {
			if mask&b != 0 {
				gogame.DrawPixel(x+xx, y+i, color1)
			} else {
				gogame.DrawPixel(x+xx, y+i, color2)
			}
			xx++
		}
	}
}
