package main

import "github.com/pnegre/gogame"

var colors []*gogame.Color

func init() {
	colors = []*gogame.Color{
		&gogame.Color{0, 0, 0, 255},          // Transparent
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

func graphics_drawPixel(x, y int, color int) {
	gogame.DrawPixel(x, y, colors[color])
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
