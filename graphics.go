package main

import "github.com/pnegre/gogame"
import "log"

const (
	WINTITLE = "gomsx"
	WIN_W    = 800
	WIN_H    = 600
	MSX_W1   = 320
	MSX_W2   = 256
	MSX_H    = 192
)

var colors []*gogame.Color
var graphics_tex256 *gogame.Texture
var graphics_tex320 *gogame.Texture
var graphics_ActiveTexture *gogame.Texture

func init() {
	colors = []*gogame.Color{
		&gogame.Color{R: 0x00, G: 0x00, B: 0x00, A: 255}, // Transparent
		&gogame.Color{R: 0x00, G: 0x00, B: 0x00, A: 255}, // Black
		&gogame.Color{R: 0x20, G: 0xc8, B: 0x40, A: 255}, // Green
		&gogame.Color{R: 0x58, G: 0xd8, B: 0x78, A: 255}, // Light Green
		&gogame.Color{R: 0x50, G: 0x50, B: 0xe8, A: 255}, // Dark Blue
		&gogame.Color{R: 0x78, G: 0x70, B: 0xf7, A: 255}, // Light Blue
		&gogame.Color{R: 0xd0, G: 0x50, B: 0x48, A: 255}, // Dark Red
		&gogame.Color{R: 0x40, G: 0xe8, B: 0xf0, A: 255}, // Cyan
		&gogame.Color{R: 0xf7, G: 0x50, B: 0x50, A: 255}, // Red
		&gogame.Color{R: 0xf7, G: 0x78, B: 0x78, A: 255}, // Bright Red
		&gogame.Color{R: 0xd0, G: 0xc0, B: 0x50, A: 255}, // Yellow
		&gogame.Color{R: 0xe0, G: 0xc8, B: 0x80, A: 255}, // Light Yellow
		&gogame.Color{R: 0x20, G: 0xb0, B: 0x38, A: 255}, // Dark Green
		&gogame.Color{R: 0xc8, G: 0x58, B: 0xb8, A: 255}, // Purple
		&gogame.Color{R: 0xc8, G: 0xc8, B: 0xc8, A: 255}, // Gray
		&gogame.Color{R: 0xf7, G: 0xf7, B: 0xf7, A: 255}, // White
	}
}

func graphics_init(quality bool) {
	if err := gogame.Init(WINTITLE, WIN_W, WIN_H); err != nil {
		log.Fatal(err)
	}
	gogame.SetLogicalSize(WIN_W, WIN_H)
	if quality {
		gogame.SetScaleQuality(1)
	}
	graphics_tex320 = gogame.NewEmptyTexture(MSX_W1, MSX_H)
	if graphics_tex320 == nil {
		log.Fatal("Error creating texture")
	}
	graphics_tex320.SetDimensions(WIN_W, WIN_H)
	graphics_tex256 = gogame.NewEmptyTexture(MSX_W2, MSX_H)
	if graphics_tex256 == nil {
		log.Fatal("Error creating texture")
	}
	graphics_tex256.SetDimensions(WIN_W, WIN_H)
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
}

func graphics_setLogicalResolution(scrMode int) {
	switch scrMode {
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
