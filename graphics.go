package main

import (
	"log"

	"github.com/pnegre/gogame"
)

const (
	WINTITLE = "gomsx"
	MSX_W1   = 320
	MSX_W2   = 256
	MSX_H    = 192
)

const (
	MODE256 = iota
	MODE320
)

var colors []gogame.Color
var graphics_tex256 *gogame.Texture
var graphics_tex320 *gogame.Texture

var graphics_pixels256 [256 * 192 * 3]byte
var graphics_pixels320 [320 * 192 * 3]byte
var graphics_mode int = MODE256

var win_h, win_w int

func init() {
	colors = []gogame.Color{
		{R: 0x00, G: 0x00, B: 0x00, A: 255}, // Transparent
		{R: 0x00, G: 0x00, B: 0x00, A: 255}, // Black
		{R: 0x20, G: 0xc8, B: 0x40, A: 255}, // Green
		{R: 0x58, G: 0xd8, B: 0x78, A: 255}, // Light Green
		{R: 0x50, G: 0x50, B: 0xe8, A: 255}, // Dark Blue
		{R: 0x78, G: 0x70, B: 0xf7, A: 255}, // Light Blue
		{R: 0xd0, G: 0x50, B: 0x48, A: 255}, // Dark Red
		{R: 0x40, G: 0xe8, B: 0xf0, A: 255}, // Cyan
		{R: 0xf7, G: 0x50, B: 0x50, A: 255}, // Red
		{R: 0xf7, G: 0x78, B: 0x78, A: 255}, // Bright Red
		{R: 0xd0, G: 0xc0, B: 0x50, A: 255}, // Yellow
		{R: 0xe0, G: 0xc8, B: 0x80, A: 255}, // Light Yellow
		{R: 0x20, G: 0xb0, B: 0x38, A: 255}, // Dark Green
		{R: 0xc8, G: 0x58, B: 0xb8, A: 255}, // Purple
		{R: 0xc8, G: 0xc8, B: 0xc8, A: 255}, // Gray
		{R: 0xf7, G: 0xf7, B: 0xf7, A: 255}, // White
	}
}

func graphics_init(quality bool) error {
	var err error

	if err = gogame.InitSDL(); err != nil {
		return err
	}

	w, h := gogame.GetDesktopDisplayResolution()
	log.Println("Desktop resolution: ", w, "x", h)

	win_h = (h * 70) / 100
	win_w = (win_h * 4) / 3 // 4:3 aspect ratio
	log.Println("Window resolution: ", win_w, "x", win_h)

	if err = gogame.Init(WINTITLE, win_w, win_h); err != nil {
		return err
	}
	gogame.SetLogicalSize(win_w, win_h)
	if quality {
		gogame.SetScaleQuality(1)
	}
	graphics_tex320, err = gogame.NewEmptyTexture(MSX_W1, MSX_H)
	if err != nil {
		return err
	}
	graphics_tex256, err = gogame.NewEmptyTexture(MSX_W2, MSX_H)
	if err != nil {
		return err
	}

	return nil
}

func graphics_quit() {
	graphics_tex256.Destroy()
	graphics_tex320.Destroy()
	gogame.Quit()
}

func graphics_render() {
	gogame.RenderClear()
	var tex *gogame.Texture
	if graphics_mode == MODE320 {
		tex = graphics_tex320
		tex.Update(graphics_pixels320[:])
	} else if graphics_mode == MODE256 {
		tex = graphics_tex256
		tex.Update(graphics_pixels256[:])
	} else {
		panic("render: mode not supported")
	}
	rect := gogame.Rect{X: 0, Y: 0, W: win_w, H: win_h}
	tex.BlitRect(&rect)
	gogame.RenderPresent()
}

func graphics_drawPixel(x, y int, color int) {
	if graphics_mode == MODE320 {
		if x < 0 || x >= MSX_W1 || y < 0 || y >= MSX_H {
			return
		}
		gcolor := colors[color]
		delta := 3 * (y*MSX_W1 + x)
		graphics_pixels320[delta] = gcolor.R
		graphics_pixels320[delta+1] = gcolor.G
		graphics_pixels320[delta+2] = gcolor.B
		return
	} else if graphics_mode == MODE256 {
		if x < 0 || x >= MSX_W2 || y < 0 || y >= MSX_H {
			return
		}
		gcolor := colors[color]
		delta := 3 * (y*MSX_W2 + x)
		graphics_pixels256[delta] = gcolor.R
		graphics_pixels256[delta+1] = gcolor.G
		graphics_pixels256[delta+2] = gcolor.B
		return
	}
	panic("drawPixel: mode not supported")
}

func graphics_setLogicalResolution(scrMode int) {
	switch scrMode {
	case SCREEN0:
		graphics_mode = MODE320
		return
	case SCREEN2:
		graphics_mode = MODE256
		return
	case SCREEN1:
		graphics_mode = MODE256
		return
	}
	panic("setLogicalResolution: mode not supported")
}
