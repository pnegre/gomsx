package main

import "github.com/remogatto/z80"

func main() {
    memory := new(Memory);
    ports := new(Ports);
    z80 := z80.NewZ80(memory, ports)
    z80.Reset()
    z80.SetPC(0)
    z80.DoOpcode()
    println(z80.PC())
}
