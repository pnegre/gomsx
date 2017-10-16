package main

import (
	"flag"
	"io/ioutil"
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

	memory := NewMemory()
	loadBiosBasic(memory, systemRom)

	if cart != "" {
		loadRom(memory, cart, 1)
	}

	ports := new(Ports)
	cpuZ80 := z80.NewZ80(memory, ports)
	cpuZ80.Reset()
	cpuZ80.SetPC(0)

	if errg := graphics_init(quality); errg != nil {
		log.Printf("Error initalizing graphics: %v", errg)
	}
	psg_init()
	defer graphics_quit()
	defer psg_quit()
	avgFPS := mainLoop(memory, cpuZ80, frameInterval)
	log.Printf("Avg FPS: %.2f\n", avgFPS)

}

func mainLoop(memory *Memory, cpuZ80 *z80.Z80, frameInterval int) float64 {
	log.Println("Beginning simulation...")
	state_init()
	var currentTime, elapsedTime, lag int64
	updateInterval := int64(time.Millisecond) * int64(frameInterval)
	previousTime := time.Now().UnixNano()

	startTime := time.Now().UnixNano()
	nframes := 0
	for {
		currentTime = time.Now().UnixNano()
		elapsedTime = currentTime - previousTime
		previousTime = currentTime
		lag += elapsedTime
		for lag >= updateInterval {
			cpuFrame(cpuZ80, memory)
			lag -= updateInterval
		}

		if quit := gogame.SlurpEvents(); quit == true {
			break
		}

		graphics_lock()
		vdp_updateBuffer()
		graphics_unlock()
		graphics_render()

		if nframes%(60*2) == 0 {
			state_save(cpuZ80, memory)
		}

		if gogame.IsKeyPressed(gogame.K_F12) {
			state_revert(cpuZ80, memory)
		}

		nframes++
	}
	delta := (time.Now().UnixNano() - startTime) / int64(time.Second)
	return float64(nframes) / float64(delta)
}

func cpuFrame(cpuZ80 *z80.Z80, memory *Memory) {
	cpuZ80.Cycles %= CYCLESPERFRAME
	for cpuZ80.Cycles < CYCLESPERFRAME {
		cpuZ80.DoOpcode()
	}

	if vdp_enabledInterrupts {
		vdp_setFrameFlag()
		cpuZ80.Interrupt()
	}
}

func loadBiosBasic(memory *Memory, fname string) {
	buffer, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	// Load BIOS
	memory.load(buffer, 0, 0)
	if len(buffer) > 0x4000 {
		// Load BASIC, if present
		memory.load(buffer[0x4000:], 1, 0)
	}
}

func loadRom(memory *Memory, fname string, slot int) {
	buffer, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	switch getCartType(buffer) {
	case KONAMI4:
		log.Printf("Loading ROM %s to slot 1 as type KONAMI4\n", fname)
		mapper := NewMapperKonami4(buffer)
		memory.setMapper(mapper, slot)
		return

	case KONAMI5:
		log.Printf("Loading ROM %s to slot 1 as type KONAMI5\n", fname)
		mapper := NewMapperKonami5(buffer)
		memory.setMapper(mapper, slot)
		return

	case ASCII8KB:
		log.Printf("Loading ROM %s to slot 1 as type ASCII8KB\n", fname)
		mapper := NewMapperASCII8(buffer)
		memory.setMapper(mapper, slot)
		return

	case NORMAL:
		log.Println("Cartridge is type NORMAL")

	case UNKNOWN:
		log.Println("Cartridge is type UNKNOWN")
	}

	log.Printf("Trying to load as a standard cartridge...\n")

	npages := len(buffer) / 0x4000
	switch npages {
	case 1:
		// Load ROM to page 1, slot 1
		// TODO: mirrored????
		log.Printf("Loading ROM %s to slot 1 (16KB)\n", fname)
		memory.load(buffer, 1, slot)
	case 2:
		// Load ROM to slot 1. Mirrored pg1&pg2 <=> pg3&pg4
		log.Printf("Loading ROM %s to slot 1 (32KB)\n", fname)
		memory.load(buffer, 0, slot)
		memory.load(buffer, 1, slot)
		memory.load(buffer[0x4000:], 2, slot)
		memory.load(buffer[0x4000:], 3, slot)
	case 4:
		log.Printf("Loading ROM %s to slot 1 (64KB)\n", fname)
		memory.load(buffer, 0, slot)
		memory.load(buffer[0x4000:], 1, slot)
		memory.load(buffer[0x8000:], 2, slot)
		memory.load(buffer[0xC000:], 3, slot)
	default:
		panic("ROM size not supported")
	}

}
