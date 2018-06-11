package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/pnegre/gogame"
	"github.com/pnegre/gomsx/z80"
)

const (
	SYSTEMROMFILE = "cbios_main_msx1.rom"
	// 60Hz -> Interval de 16mseg
	INTERVAL = 16
	// EL z80 va a 3.58 Mhz. Cada 16mseg passen 57280 cicles
	CYCLESPERFRAME = 60000
)

type MSX struct {
	cpuz80 *z80.Z80
	vdp    *Vdp
	memory *Memory
	ppi    *PPI
	psg    *PSG
}

func main() {
	runtime.LockOSThread() // Assure SDL works...
	var cart string
	var systemRom string
	var quality bool
	var frameInterval int
	flag.StringVar(&cart, "cart", "", "ROM in SLOT 1")
	flag.StringVar(&systemRom, "sys", SYSTEMROMFILE, "System file")
	flag.BoolVar(&quality, "quality", true, "Best quality rendering")
	flag.IntVar(&frameInterval, "fint", INTERVAL, "Frame interval in milliseconds")
	flag.Parse()

	if flag.NArg() > 0 {
		flag.Usage()
		os.Exit(1)
	}

	ppi := NewPPI()
	memory := NewMemory(ppi)
	memory.loadBiosBasic(systemRom)

	if cart != "" {
		memory.loadRom(cart, 1)
	}

	psg := NewPSG()
	vdp := NewVdp()
	ports := &Ports{vdp: vdp, ppi: ppi, psg: psg}
	cpuZ80 := z80.NewZ80(memory, ports)
	cpuZ80.Reset()
	cpuZ80.SetPC(0)
	msx := &MSX{cpuz80: cpuZ80, vdp: vdp, memory: memory, ppi: ppi, psg: psg}

	if errg := graphics_init(quality); errg != nil {
		log.Printf("Error initalizing graphics: %v", errg)
	}
	sound_init(psg)
	defer graphics_quit()
	defer sound_quit()
	avgFPS := mainLoop(msx, frameInterval)
	log.Printf("Avg FPS: %.2f\n", avgFPS)
}

func mainLoop(msx *MSX, frameInterval int) float64 {
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
				cpuFrame(msx)
			}
			lag -= updateInterval
		}

		if quit := gogame.SlurpEvents(); quit == true {
			break
		}

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

func cpuFrame(msx *MSX) {
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
