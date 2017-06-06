package main

import "github.com/pnegre/gogame"
import "log"

var colors []*gogame.Color
var graphics_tex256 *gogame.Texture
var graphics_tex320 *gogame.Texture
var graphics_ActiveTexture *gogame.Texture

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

func graphics_init() {
	if err := gogame.Init(WINTITLE, WIN_W, WIN_H); err != nil {
		log.Fatal(err)
	}
	gogame.SetLogicalSize(WIN_W, WIN_H)
	graphics_tex256 = gogame.NewEmptyTexture(256, 192)
	graphics_tex256.SetDimensions(WIN_W, WIN_H)
	graphics_tex320 = gogame.NewEmptyTexture(320, 192)
	graphics_tex320.SetDimensions(WIN_W, WIN_H)
	graphics_ActiveTexture = graphics_tex256
}

func graphics_quit() {
	graphics_tex256.Destroy()
	graphics_tex320.Destroy()
	gogame.Quit()
}

func graphics_lock() {
	graphics_ActiveTexture.Lock()
}

func graphics_unlock() {
	graphics_ActiveTexture.Unlock()
}

func graphics_render() {
	gogame.RenderClear()
	graphics_ActiveTexture.Blit(0, 0)
	gogame.RenderPresent()
}

func graphics_drawPixel(x, y int, color int) {
	graphics_ActiveTexture.Pixel(x, y, colors[color])
	//gogame.DrawPixel(x, y, colors[color])
}

func graphics_setLogicalResolution() {
	switch vdp_screenMode {
	case SCREEN0:
		graphics_ActiveTexture = graphics_tex320
		return
	case SCREEN2:
		graphics_ActiveTexture = graphics_tex256
		return
	case SCREEN1:
		graphics_ActiveTexture = graphics_tex256
		return
	}
	panic("setLogicalResolution: mode not supported")
}
