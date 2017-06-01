package main

import "github.com/remogatto/z80"
import "github.com/pnegre/gogame"
import "log"
import "time"

const (
	WINTITLE  = "gomsx"
	WIN_W     = 800
	WIN_H     = 600
	MSX_W     = 320
	MSX_H     = 192
	ROMFILE   = "msx1.rom"
	NANOS_SCR = 20000000 // 50Hz -> Interval de 20Mseg
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
	logAssembler := false
	for {
		for delta < NANOS_SCR {
			if logAssembler {
				pc := cpuZ80.PC()
				instr, _, _ := z80.Disassemble(memory, pc, 0)
				log.Printf("%04x: %s\n", pc, instr)
			}
			delta = time.Now().UnixNano() - lastTm
			cpuZ80.DoOpcode()
		}
		delta -= NANOS_SCR
		lastTm = time.Now().UnixNano()

		if quit := gogame.SlurpEvents(); quit == true {
			break
		}

		graphics_renderScreen()
		vdp_setFrameFlag()
		if vdp_enabledInterrupts {
			cpuZ80.Interrupt()
		}
		gogame.Delay(1)
	}
}
