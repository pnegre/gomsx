package main

import (
	"flag"
	"log"
	"os"
	"runtime"

	"github.com/pnegre/gomsx/z80"
)

const (
	SYSTEMROMFILE = "cbios_main_msx1.rom"
	// 60Hz -> Interval de 16mseg
	INTERVAL = 16
	// EL z80 va a 3.58 Mhz. Cada 16mseg passen 57280 cicles
	CYCLESPERFRAME = 60000
)

func main() {
	runtime.LockOSThread() // Assure SDL works...
	var cart string
	var systemRom string
	var quality bool
	var frameInterval int
	var mtype string
	flag.StringVar(&cart, "cart", "", "ROM in SLOT 1")
	flag.StringVar(&systemRom, "sys", SYSTEMROMFILE, "System file")
	flag.BoolVar(&quality, "quality", true, "Best quality rendering")
	flag.IntVar(&frameInterval, "fint", INTERVAL, "Frame interval in milliseconds")
	flag.StringVar(&mtype, "mtype", "", "Mapper type (KONAMI4...)")
	flag.Parse()

	if flag.NArg() > 0 {
		flag.Usage()
		os.Exit(1)
	}

	ppi := NewPPI()
	memory := NewMemory(ppi)
	memory.loadBiosBasic(systemRom)

	if cart != "" {
		memory.loadRom(cart, 1, mtype)
	}

	psg := NewPSG()
	vdp := NewVdp()
	ports := &Ports{vdp: vdp, ppi: ppi, psg: psg}
	cpuZ80 := z80.NewZ80(memory, ports)
	cpuZ80.Reset()
	cpuZ80.SetPC(0)
	msx := &MSX{cpuz80: cpuZ80, vdp: vdp, memory: memory, ppi: ppi, psg: psg}

	if errg := graphics_init(quality); errg != nil {
		log.Fatalf("Error initalizing graphics: %v", errg.Error())
	}
	defer graphics_quit()

	if errs := sound_init(psg); errs != nil {
		log.Fatalf("Error intializing sound: %v", errs.Error())
	}
	defer sound_quit()

	avgFPS := msx.mainLoop(frameInterval)
	log.Printf("Avg FPS: %.2f\n", avgFPS)
}
