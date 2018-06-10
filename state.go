package main

import "container/ring"
import "github.com/pnegre/gomsx/z80"

type stateDataT struct {
	cpuBackup   *z80.Z80
	memContents [4][4][0x4000]byte
	vdp         *Vdp
}

func newStateData() *stateDataT {
	sd := new(stateDataT)
	sd.cpuBackup = new(z80.Z80)
	return sd
}

const NSTATEDATA = 5

var state_ring *ring.Ring

func state_init() {
	state_ring = ring.New(1)
}

func state_save(msx *MSX) {
	var data *stateDataT
	if state_ring.Value == nil {
		data = newStateData()
		state_ring.Value = data
	} else {
		data = state_ring.Value.(*stateDataT)
	}

	// Save CPU state
	msx.cpuz80.SaveState(data.cpuBackup)

	// Save RAM
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 0x4000; k++ {
				data.memContents[i][j][k] = msx.memory.contents[i][j][k]
			}
		}
	}

	// Save VDP
	data.vdp = msx.vdp.saveState()

	// Advance state
	if state_ring.Len() < NSTATEDATA {
		sr := ring.New(1)
		state_ring.Link(sr)
	}
	state_ring = state_ring.Next()
}

func state_revert(msx *MSX) {
	state_ring = state_ring.Move(-(state_ring.Len() - 1))
	data := state_ring.Value.(*stateDataT)
	state_ring = ring.New(1)
	state_ring.Value = data

	// Restore CPU state
	msx.cpuz80.RestoreState(data.cpuBackup)

	// Restore RAM
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 0x4000; k++ {
				msx.memory.contents[i][j][k] = data.memContents[i][j][k]
			}
		}
	}

	// Restore VDP
	msx.vdp.restoreState(data.vdp)

	sr := ring.New(1)
	state_ring.Link(sr)
	state_ring = state_ring.Next()
}

func state_saveVDP(data *stateDataT) {
	// data.vdp.screenEnabled = vdp_screenEnabled
	// data.vdp.screenMode = vdp_screenMode
	// data.vdp.valueRead = vdp_valueRead
	// data.vdp.writeState = vdp_writeState
	// data.vdp.enabledInterrupts = vdp_enabledInterrupts
	// data.vdp.writeToVRAM = vdp_writeToVRAM
	// data.vdp.pointerVRAM = vdp_pointerVRAM
	// data.vdp.statusReg = vdp_statusReg

	// for i := 0; i < 10; i++ {
	// 	data.vdp.registers[i] = vdp_registers[i]
	// }

	// for i := 0; i < 0x10000; i++ {
	// 	data.vdp.VRAM[i] = vdp_VRAM[i]
	// }
}

func state_restoreVDP(data *stateDataT) {
	// vdp_screenEnabled = data.vdp.screenEnabled
	// vdp_screenMode = data.vdp.screenMode
	// vdp_valueRead = data.vdp.valueRead
	// vdp_writeState = data.vdp.writeState
	// vdp_enabledInterrupts = data.vdp.enabledInterrupts
	// vdp_writeToVRAM = data.vdp.writeToVRAM
	// vdp_pointerVRAM = data.vdp.pointerVRAM
	// vdp_statusReg = data.vdp.statusReg

	// for i := 0; i < 10; i++ {
	// 	vdp_registers[i] = data.vdp.registers[i]
	// }

	// for i := 0; i < 0x10000; i++ {
	// 	vdp_VRAM[i] = data.vdp.VRAM[i]
	// }
}
