package main

import "github.com/remogatto/z80"
import "github.com/pnegre/gogame"
import "log"

import "time"
import "os"
import "bufio"
import "flag"

const (
	WINTITLE = "gomsx"
	WIN_W    = 800
	WIN_H    = 600
	ROMFILE  = "msx1.rom"
	// 50Hz -> Interval de 20Mseg
	FRAMETIME = 20
	FPS       = 60
	// EL z80 executa devers 580000 instr per segon
	// (Un "frame" són 16mseg, per tant executa 9280 instr. per frame)
	INSTRPERFRAME = 9280
)

func main() {
	flag.Parse()
	memory := NewMemory()
	loadBiosBasic(memory, ROMFILE)
	if flag.NArg() == 1 {
		rom := flag.Args()[0]
		loadRom(memory, rom)
	}
	ports := new(Ports)
	cpuZ80 := z80.NewZ80(memory, ports)
	cpuZ80.Reset()
	cpuZ80.SetPC(0)

	if err := gogame.Init(WINTITLE, WIN_W, WIN_H); err != nil {
		log.Fatal(err)
	}
	defer gogame.Quit()

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

		gogame.RenderClear()
		vdp_updateBuffer()
		gogame.RenderPresent()

		gogame.Delay(1)
		nframes++
	}
	delta := (time.Now().UnixNano() - startTime) / int64(time.Second)
	log.Printf("Avg FPS: %.2f\n", float64(nframes)/float64(delta))
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

func loadRom(memory *Memory, fname string) {
	buffer, err := readFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	npages := len(buffer) / 0x4000
	switch npages {
	case 1:
		// Load ROM to page 1, slot 1
		// TODO: mirrored????
		memory.load(buffer, 1, 1)
	case 2:
		// Load ROM to slot 1. Mirrored pg1&pg2 <=> pg3&pg4
		memory.load(buffer, 0, 1)
		memory.load(buffer, 1, 1)
		memory.load(buffer[0x4000:], 2, 1)
		memory.load(buffer[0x4000:], 3, 1)
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
