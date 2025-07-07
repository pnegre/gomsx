package main

import "github.com/pnegre/gogame"

const (
	WINTITLE = "gomsx"
	WIN_W    = 800
	WIN_H    = 600
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

var graphics_pixels256 [256 * 192]int
var graphics_pixels320 [320 * 192]int
var graphics_mode int = MODE256

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
	if err = gogame.Init(WINTITLE, WIN_W, WIN_H); err != nil {
		return err
	}
	gogame.SetLogicalSize(WIN_W, WIN_H)
	if quality {
		gogame.SetScaleQuality(1)
	}
	graphics_tex320, err = gogame.NewEmptyTexture(MSX_W1, MSX_H)
	if err != nil {
		return err
	}
	graphics_tex320.SetDimensions(WIN_W, WIN_H)
	graphics_tex256, err = gogame.NewEmptyTexture(MSX_W2, MSX_H)
	if err != nil {
		return err
	}
	graphics_tex256.SetDimensions(WIN_W, WIN_H)

	// Initialize pixel buffers
	for i := 0; i < MSX_W1*MSX_H; i++ {
		graphics_pixels320[i] = 0 // Transparent
	}
	for i := 0; i < MSX_W2*MSX_H; i++ {
		graphics_pixels256[i] = 0 // Transparent
	}
	return nil
}

func graphics_quit() {
	graphics_tex256.Destroy()
	graphics_tex320.Destroy()
	gogame.Quit()
}

func updatePixels(tex *gogame.Texture, pixels []int, width, height int) {
	tex.Lock()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var color = colors[pixels[y*width+x]]
			tex.Pixel(x, y, &color)
		}
	}
	tex.Unlock()
}

func graphics_render() {
	gogame.RenderClear()
	var tex *gogame.Texture
	if graphics_mode == MODE320 {
		tex = graphics_tex320
		updatePixels(tex, graphics_pixels320[:], MSX_W1, MSX_H)
	} else if graphics_mode == MODE256 {
		tex = graphics_tex256
		updatePixels(tex, graphics_pixels256[:], MSX_W2, MSX_H)
	} else {
		panic("render: mode not supported")
	}
	tex.Blit(0, 0)
	gogame.RenderPresent()
}

func graphics_drawPixel(x, y int, color int) {
	if graphics_mode == MODE320 {
		if x < 0 || x >= MSX_W1 || y < 0 || y >= MSX_H {
			return
		}
		graphics_pixels320[y*MSX_W1+x] = color
		return
	} else if graphics_mode == MODE256 {
		if x < 0 || x >= MSX_W2 || y < 0 || y >= MSX_H {
			return
		}
		graphics_pixels256[y*MSX_W2+x] = color
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
