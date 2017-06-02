package main

import "github.com/remogatto/z80"
import "github.com/pnegre/gogame"
import "log"
import "time"

const (
	WINTITLE  = "gomsx"
	WIN_W     = 800
	WIN_H     = 600
	ROMFILE   = "msx1.rom"
	NANOS_SCR = 20000000 // 50Hz -> Interval de 20Mseg
)

func main() {
	if err := gogame.Init(WINTITLE, WIN_W, WIN_H); err != nil {
		log.Fatal(err)
	}
	defer gogame.Quit()

	memory := NewMemory()
	memory.loadFromFile(ROMFILE)
	ports := new(Ports)
	cpuZ80 := z80.NewZ80(memory, ports)
	cpuZ80.Reset()
	cpuZ80.SetPC(0)
	log.Println("Beginning simulation...")
	lastTm := time.Now().UnixNano()
	delta := int64(0)
	logAssembler := false
	for {
		for i := 0; i < 500; i++ {
			if logAssembler {
				pc := cpuZ80.PC()
				instr, _, _ := z80.Disassemble(memory, pc, 0)
				log.Printf("%04x: %s\n", pc, instr)
			}
			cpuZ80.DoOpcode()
		}

		delta = time.Now().UnixNano() - lastTm
		if delta > NANOS_SCR {
			if quit := gogame.SlurpEvents(); quit == true {
				break
			}

			graphics_renderScreen()
			if vdp_enabledInterrupts {
				cpuZ80.Interrupt()
			}
			lastTm = time.Now().UnixNano()
			gogame.Delay(1)
		}
	}
}
