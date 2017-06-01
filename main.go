package main

import "github.com/remogatto/z80"
import "github.com/pnegre/gogame"
import "log"
import "time"

const (
	WINTITLE     = "gomsx"
	WIN_W        = 800
	WIN_H        = 600
	MSX_W        = 320
	MSX_H        = 192
	ROMFILE      = "msx1.rom"
	logAssembler = false
	NANOS_SCR    = 100000000
)

func main() {
	if err := gogame.Init(WINTITLE, WIN_W, WIN_H); err != nil {
		log.Fatal(err)
	}
	gogame.SetLogicalSize(MSX_W, MSX_H)
	defer gogame.Quit()

	memory := NewMemory(ROMFILE)
	ports := new(Ports)
	cpuZ80 := z80.NewZ80(memory, ports)
	cpuZ80.Reset()
	cpuZ80.SetPC(0)
	log.Println("Beginning simulation...")
	lastTm := time.Now().UnixNano()
	delta := int64(0)
	for {
		if logAssembler {
			pc := cpuZ80.PC()
			instr, _, _ := z80.Disassemble(memory, pc, 0)
			log.Printf("%04x: %s\n", pc, instr)
		}
		cpuZ80.DoOpcode()

		if quit := gogame.SlurpEvents(); quit == true {
			break
		}

		delta = time.Now().UnixNano() - lastTm
		if delta > NANOS_SCR {
			graphics_renderScreen()
			lastTm = time.Now().UnixNano()
		}

	}
}
