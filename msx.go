package main

import (
	"log"
	"time"

	"github.com/pnegre/gogame"
	"github.com/pnegre/gomsx/z80"
)

type MSX struct {
	cpuz80 *z80.Z80
	vdp    *Vdp
	memory *Memory
	ppi    *PPI
	psg    *PSG
}

func (msx *MSX) mainLoop(frameInterval int) float64 {
	log.Println("Beginning simulation...")
	state_init()
	var currentTime, elapsedTime, lag int64
	updateInterval := int64(time.Millisecond) * int64(frameInterval)
	previousTime := time.Now().UnixNano()

	startTime := time.Now().UnixNano()
	nframes := 0
	paused := false
	for {
		currentTime = time.Now().UnixNano()
		elapsedTime = currentTime - previousTime
		previousTime = currentTime
		lag += elapsedTime
		for lag >= updateInterval {
			if !paused {
				msx.cpuFrame()
			}
			lag -= updateInterval
		}

		if quit := gogame.SlurpEvents(); quit == true {
			break
		}

		checkFullScreen()

		graphics_lock()
		msx.vdp.updateBuffer()
		graphics_unlock()
		graphics_render()

		if !paused {
			if nframes%(60*2) == 0 {
				state_save(msx)
			}
		}

		if gogame.IsKeyPressed(gogame.K_F12) {
			state_revert(msx)
			paused = true
		}

		if gogame.IsKeyPressed(gogame.K_SPACE) {
			paused = false
		}

		nframes++
	}
	delta := (time.Now().UnixNano() - startTime) / int64(time.Second)
	return float64(nframes) / float64(delta)
}

func (msx *MSX) cpuFrame() {
	msx.cpuz80.Cycles %= CYCLESPERFRAME
	for msx.cpuz80.Cycles < CYCLESPERFRAME {
		if msx.cpuz80.Halted == true {
			break
		}
		msx.cpuz80.DoOpcode()
	}

	if msx.vdp.enabledInterrupts {
		msx.vdp.setFrameFlag()
		msx.cpuz80.Interrupt()
	}
}

// Closure
var checkFullScreen = func() func() {
	state := 0
	fullscr := false
	var last int64 = -1
	var t int64
	return func() {
		_, _, but := gogame.GetMouseState()
		pressed := but == gogame.MOUSE_BUTTON_LEFT
		switch state {
		case 0:
			if pressed {
				state = 1
				last = time.Now().UnixNano() / 1000000
			}
		case 1:
			if !pressed {
				state = 2
			}
		case 2:
			t = time.Now().UnixNano() / 1000000
			if t-last > 500 {
				state = 0
				break
			}
			if pressed {
				state = 3
			}
		case 3:
			t = time.Now().UnixNano() / 1000000
			if t-last > 1500 {
				state = 0
				break
			}
			if !pressed {
				fullscr = !fullscr
				gogame.SetFullScreen(fullscr)
				state = 0
			}
		}
	}
}()
