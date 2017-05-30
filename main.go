package main

import "github.com/remogatto/z80"

func main() {
    memory := NewMemory("hb-501p_basic-bios.rom")
    ports := new(Ports);
    cpuZ80 := z80.NewZ80(memory, ports)
    cpuZ80.Reset()
    cpuZ80.SetPC(0)
    cpuZ80.DoOpcode()
    println(cpuZ80.PC())
}
