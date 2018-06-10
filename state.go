package main

import "container/ring"
import "github.com/pnegre/gomsx/z80"

type stateDataT struct {
	cpuBackup *z80.Z80
	memory    *Memory
	vdp       *Vdp
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
	data.memory = msx.memory.saveState()

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
	msx.memory.restoreState(data.memory)

	// Restore VDP
	msx.vdp.restoreState(data.vdp)

	sr := ring.New(1)
	state_ring.Link(sr)
	state_ring = state_ring.Next()
}
