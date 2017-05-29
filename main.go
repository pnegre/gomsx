package main

import "fmt"
import "github.com/remogatto/z80"

func main() {
    fmt.Println("hey, world");
    var memory Memory;
    var ports Ports;
    z80 := z80.NewZ80(memory, ports)
    z80.Reset()
    z80.SetPC(0)
    z80.DoOpcode()
    fmt.Println(z80.PC())
}
