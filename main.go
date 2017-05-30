package main

import "github.com/remogatto/z80"
import "fmt"

func main() {
	memory := NewMemory("hb-501p_basic-bios1.rom")
	ports := new(Ports)
	cpuZ80 := z80.NewZ80(memory, ports)
	cpuZ80.Reset()
	cpuZ80.SetPC(0)
	for {
		cpuZ80.DoOpcode()
		fmt.Printf("PC = %04x\n", cpuZ80.PC())
	}
}
