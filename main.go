package main

import "github.com/remogatto/z80"
import "github.com/pnegre/gogame"
import "log"

import "time"
import "os"
import "bufio"
import "flag"

const (
	ROMFILE = "msx1.rom"
	// 60Hz -> Interval de 16Mseg
	FRAMETIME = 16
	FPS       = 60
	// EL z80 executa devers 580000 instr per segon
	// (Un "frame" són 16mseg, per tant executa 9280 instr. per frame)
	INSTRPERFRAME = 6280
)

func main() {
	flag.Parse()
	memory := NewMemory()
	loadBiosBasic(memory, ROMFILE)
	if flag.NArg() == 1 {
		rom := flag.Args()[0]
		loadRom(memory, rom, 1) // Load to slot 1
	}
	ports := new(Ports)
	cpuZ80 := z80.NewZ80(memory, ports)
	cpuZ80.Reset()
	cpuZ80.SetPC(0)

	graphics_init()
	sound_init()
	defer graphics_quit()
	defer sound_quit()
	avgFPS := mainLoop(memory, cpuZ80)
	log.Printf("Avg FPS: %.2f\n", avgFPS)
}

func mainLoop(memory *Memory, cpuZ80 *z80.Z80) float64 {
	log.Println("Beginning simulation...")
	logAssembler := false
	var currentTime, elapsedTime, lag int64
	updateInterval := int64(time.Second) / int64(FPS)
	previousTime := time.Now().UnixNano()

	startTime := time.Now().UnixNano()
	nframes := 0
	for {
		currentTime = time.Now().UnixNano()
		elapsedTime = currentTime - previousTime
		previousTime = currentTime
		lag += elapsedTime
		for lag >= updateInterval {
			cpuFrame(cpuZ80, memory, logAssembler)
			lag -= updateInterval
		}

		if gogame.IsKeyPressed(gogame.K_F7) {
			logAssembler = true
		}

		if quit := gogame.SlurpEvents(); quit == true {
			break
		}

		graphics_lock()
		vdp_updateBuffer()
		graphics_unlock()
		graphics_render()

		nframes++
	}
	delta := (time.Now().UnixNano() - startTime) / int64(time.Second)
	return float64(nframes) / float64(delta)
}

func cpuFrame(cpuZ80 *z80.Z80, memory *Memory, logAssembler bool) {
	for i := 0; i < INSTRPERFRAME; i++ {
		if logAssembler {
			pc := cpuZ80.PC()
			instr, _, _ := z80.Disassemble(memory, pc, 0)
			log.Printf("%04x: %s\n", pc, instr)
		}
		cpuZ80.DoOpcode()
	}
	if vdp_enabledInterrupts {
		vdp_setFrameFlag()
		cpuZ80.Interrupt()
	}
}

func loadBiosBasic(memory *Memory, fname string) {
	buffer, err := readFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	memory.load(buffer, 0, 0)
	memory.load(buffer[0x4000:], 1, 0)
}

func loadRom(memory *Memory, fname string, slot int) {
	buffer, err := readFile(fname)
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
	}

	log.Printf("Failed to identify cartridge as MegaROM. Trying to load as a standard cartridge...\n")

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

func readFile(fname string) ([]byte, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	stats, statsErr := f.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(f)
	_, err = bufr.Read(bytes)

	return bytes, err
}
