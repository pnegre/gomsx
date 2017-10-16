package main

import "github.com/pnegre/gomsx/z80"

var state_cpuBackup *z80.Z80
var state_memContents [4][4][0x4000]byte

func state_init() {
	state_cpuBackup = new(z80.Z80)
}

func state_save(cpu *z80.Z80, mem *Memory) {
	cpu.SaveState(state_cpuBackup)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 0x4000; k++ {
				state_memContents[i][j][k] = mem.contents[i][j][k]
			}
		}
	}
}

func state_revert(cpu *z80.Z80, mem *Memory) {
	cpu.RestoreState(state_cpuBackup)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 0x4000; k++ {
				mem.contents[i][j][k] = state_memContents[i][j][k]
			}
		}
	}
}
