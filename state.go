package main

import "container/ring"
import "github.com/pnegre/gomsx/z80"

type stateDataT struct {
	cpuBackup   *z80.Z80
	memContents [4][4][0x4000]byte
	vdp         struct {
		screenEnabled     bool
		screenMode        int
		valueRead         byte
		writeState        int
		enabledInterrupts bool
		registers         [10]byte
		writeToVRAM       bool
		VRAM              [0x10000]byte
		pointerVRAM       uint16
		statusReg         byte
	}
}

func newStateData() *stateDataT {
	sd := new(stateDataT)
	sd.cpuBackup = new(z80.Z80)
	return sd
}

const NSTATEDATA = 5

var state_data [NSTATEDATA]stateDataT
var state_current int = 0
var state_ring *ring.Ring

func state_init() {
	state_ring = ring.New(NSTATEDATA)
}

func state_save(cpu *z80.Z80, mem *Memory) {
	var data *stateDataT
	if state_ring.Value == nil {
		data = newStateData()
		state_ring.Value = data
	} else {
		data = state_ring.Value.(*stateDataT)
	}

	// Save CPU state
	cpu.SaveState(data.cpuBackup)

	// Save RAM
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 0x4000; k++ {
				data.memContents[i][j][k] = mem.contents[i][j][k]
			}
		}
	}

	// Save VDP
	state_saveVDP(data)

	// Advance state
	state_ring = state_ring.Next()
}

func state_revert(cpu *z80.Z80, mem *Memory) {
	state_ring = state_ring.Move(-NSTATEDATA)
	// state_current = (state_current + NSTATEDATA - 1) % NSTATEDATA
	data := state_ring.Value.(*stateDataT)

	// Restore CPU state
	cpu.RestoreState(data.cpuBackup)

	// Restore RAM
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 0x4000; k++ {
				mem.contents[i][j][k] = data.memContents[i][j][k]
			}
		}
	}

	// Restore VDP
	state_restoreVDP(data)
}

func state_saveVDP(data *stateDataT) {
	data.vdp.screenEnabled = vdp_screenEnabled
	data.vdp.screenMode = vdp_screenMode
	data.vdp.valueRead = vdp_valueRead
	data.vdp.writeState = vdp_writeState
	data.vdp.enabledInterrupts = vdp_enabledInterrupts
	data.vdp.writeToVRAM = vdp_writeToVRAM
	data.vdp.pointerVRAM = vdp_pointerVRAM
	data.vdp.statusReg = vdp_statusReg

	for i := 0; i < 10; i++ {
		data.vdp.registers[i] = vdp_registers[i]
	}

	for i := 0; i < 0x10000; i++ {
		data.vdp.VRAM[i] = vdp_VRAM[i]
	}
}

func state_restoreVDP(data *stateDataT) {
	vdp_screenEnabled = data.vdp.screenEnabled
	vdp_screenMode = data.vdp.screenMode
	vdp_valueRead = data.vdp.valueRead
	vdp_writeState = data.vdp.writeState
	vdp_enabledInterrupts = data.vdp.enabledInterrupts
	vdp_writeToVRAM = data.vdp.writeToVRAM
	vdp_pointerVRAM = data.vdp.pointerVRAM
	vdp_statusReg = data.vdp.statusReg

	for i := 0; i < 10; i++ {
		vdp_registers[i] = data.vdp.registers[i]
	}

	for i := 0; i < 0x10000; i++ {
		vdp_VRAM[i] = data.vdp.VRAM[i]
	}
}
