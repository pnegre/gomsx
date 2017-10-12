package main

import (
	"testing"

	"github.com/pnegre/gomsx/z80"
)

func Test1(t *testing.T) {
	memory := NewMemory()
	ports := new(Ports)
	cpuZ80 := z80.NewZ80(memory, ports)
	cpuZ80.Reset()
	cpuZ80.SetPC(0)

	memory.WriteByte(0, 0x3e)
	memory.WriteByte(1, 0)
	memory.WriteByte(2, 0x76)
	cpuZ80.Cycles = 0
	for !cpuZ80.Halted {
		cpuZ80.DoOpcode()
	}
	if cpuZ80.Cycles != 13 {
		t.Errorf("ncycles: %d", cpuZ80.Cycles)
	}
}
