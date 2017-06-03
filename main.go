package main

import "github.com/remogatto/z80"
import "github.com/pnegre/gogame"
import "log"

import "time"
import "os"
import "bufio"
import "flag"

const (
	WINTITLE      = "gomsx"
	WIN_W         = 800
	WIN_H         = 600
	ROMFILE       = "msx1.rom"
	FRAMETIME     = 20    // 50Hz -> Interval de 20Mseg
	INSTRPERFRAME = 11600 // EL z80 executa devers 580000 instr per segon (Un "frame" són 20mseg, per tant executa 11600 instr. per frame)
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
	var nframes uint64 = 0
	startTime := millis()
	for {
		tme := millis()
		for i := 0; i < INSTRPERFRAME; i++ {
			if logAssembler {
				pc := cpuZ80.PC()
				instr, _, _ := z80.Disassemble(memory, pc, 0)
				log.Printf("%04x: %s\n", pc, instr)
			}
			cpuZ80.DoOpcode()
		}

		if gogame.IsKeyPressed(gogame.K_ESC) {
			logAssembler = true
		}

		if quit := gogame.SlurpEvents(); quit == true {
			break
		}

		graphics_renderScreen()
		if vdp_enabledInterrupts {
			cpuZ80.Interrupt()
		}
		tme = FRAMETIME - (millis() - tme)
		if tme > 0 {
			gogame.Delay(int(tme))
		}
		nframes++
	}
	log.Println(nframes)
	// log.Println(startTime)
	// log.Println(millis())
	// log.Println(millis() - startTime)
	fps := float64(nframes) / (float64(millis()-startTime) / 1000)
	log.Printf("Average FPS: %f\n", fps)
}

func millis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
	// return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
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
	memory.load(buffer, 1, 1)
	memory.load(buffer[0x4000:], 2, 1)
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
