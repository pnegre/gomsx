package main

import "container/ring"
import "github.com/pnegre/gomsx/z80"
import "log"

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

const NSTATEDATA = 15

var state_ring *ring.Ring

func state_init() {
	state_ring = ring.New(NSTATEDATA)
	for i := 0; i < NSTATEDATA; i++ {
		state_ring.Value = nil
		state_ring = state_ring.Next()
	}
}

func state_save(msx *MSX) {
	var data *stateDataT
	data = newStateData()
	state_ring.Value = data
	state_ring = state_ring.Next()

	// Save CPU state
	msx.cpuz80.SaveState(data.cpuBackup)

	// Save RAM
	data.memory = msx.memory.saveState()

	// Save VDP
	data.vdp = msx.vdp.saveState()
}

func get_revert_data() *stateDataT {
	state_ring = state_ring.Prev()

	count := 0
	for state_ring.Value == nil && count < NSTATEDATA {
		state_ring = state_ring.Prev()
		count++
	}
	if count >= NSTATEDATA {
		return nil
	}

	return state_ring.Value.(*stateDataT)
}

func state_revert(msx *MSX) {
	data := get_revert_data()
	if data == nil {
		return
	}

	log.Println("Reverting state...")

	// Restore CPU state
	msx.cpuz80.RestoreState(data.cpuBackup)

	// Restore RAM
	msx.memory.restoreState(data.memory)

	// Restore VDP
	msx.vdp.restoreState(data.vdp)
}
