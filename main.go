package main

import "github.com/remogatto/z80"
import "github.com/pnegre/gogame"
import "log"

import "time"
import "os"
import "bufio"

const (
	WINTITLE        = "gomsx"
	WIN_W           = 800
	WIN_H           = 600
	ROMFILE         = "msx1.rom"
	FRAMETIME       = 20 // 50Hz -> Interval de 20Mseg
	TSTATESPERFRAME = 71400
)

func main() {
	if err := gogame.Init(WINTITLE, WIN_W, WIN_H); err != nil {
		log.Fatal(err)
	}
	defer gogame.Quit()

	memory := NewMemory()
	loadROMS(memory)
	ports := new(Ports)
	cpuZ80 := z80.NewZ80(memory, ports)
	cpuZ80.Reset()
	cpuZ80.SetPC(0)
	log.Println("Beginning simulation...")
	logAssembler := false
	for {
		tme := millis()
		// EL z80 executa devers 580000 instr per segon
		// Un "frame" són 20mseg, per tant executa 11600 instr. per frame
		for i := 0; i < 11600; i++ {
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
	}
}

func millis() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func loadROMS(memory *Memory) {
	buffer, err := readFile(ROMFILE)
	if err != nil {
		log.Fatal(err)
	}
	memory.load(buffer, 0, 0)
	memory.load(buffer[0x4000:], 1, 0)

	// buffer, err = readFile("caos.rom")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// memory.load(buffer, 1, 1)
	// memory.load(buffer[0x4000:], 2, 1)
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
