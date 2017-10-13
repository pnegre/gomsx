package main

import (
	"testing"

	"github.com/pnegre/gomsx/z80"
)

func Test1(t *testing.T) {
	// LD A, 0 (8 cicles)
	// HALT    (5 cicles)
	ar := []byte{0x3e, 0x00, 0x76}
	nc := checkCycles(ar)
	if nc != 13 {
		t.Errorf("ncycles: %d", nc)
	}
}

func Test2(t *testing.T) {
	// LD A, 0  (8 cicles)
	// JP Z, 1  (11 cicles)
	// HALT     (5 cicles)
	ar := []byte{0x3e, 0x00, 0xca, 0x00, 0x00, 0x76}
	nc := checkCycles(ar)
	if nc != 24 {
		t.Errorf("ncycles: %d", nc)
	}
}

func Test3(t *testing.T) {
	//     LD B, 5  (8 cicles)
	// xx: INC A    (5 cicles)
	//     DJNZ xx  (14/9 cicles)
	//     HALT     (5 cicles)
	ar := []byte{0x06, 0x05, 0x3c, 0x10, 0xfd, 0x76}
	nc := checkCycles(ar)
	if nc != 103 {
		t.Errorf("ncycles: %d", nc)
	}
}

func checkCycles(ar []byte) int {
	memory := NewMemory()
	ports := new(Ports)
	cpuZ80 := z80.NewZ80(memory, ports)
	cpuZ80.Reset()
	cpuZ80.SetPC(0)
	cpuZ80.Cycles = 0

	for i := 0; i < len(ar); i++ {
		memory.WriteByte(uint16(i), ar[i])
	}

	for !cpuZ80.Halted {
		cpuZ80.DoOpcode()
	}

	return int(cpuZ80.Cycles)
}
