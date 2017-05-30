package main

import "github.com/remogatto/z80"
import "log"

const logAssembler = 0

func main() {
	memory := NewMemory("hb-501p_basic-bios1.rom")
	ports := new(Ports)
	cpuZ80 := z80.NewZ80(memory, ports)
	cpuZ80.Reset()
	cpuZ80.SetPC(0)
	log.Println("Beginning simulation...")
	for {
		if logAssembler != 0 {
			pc := cpuZ80.PC()
			instr, _, _ := z80.Disassemble(memory, pc, 0)
			log.Printf("%04x: %s\n", pc, instr)
		}
		cpuZ80.DoOpcode()

	}
}
