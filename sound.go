package main

import "log"

var sound_regs [16]byte
var sound_regNext byte

func sound_writePort(ad byte, val byte) {
	switch {
	case ad == 0xa0:
		// Register write port
		sound_regNext = val
		return

	case ad == 0xa1:
		// Write value to port
		sound_regs[sound_regNext] = val
		return
	}

	log.Fatalf("Sound, not implemented: out(%02x,%02x)", ad, val)
}

func sound_readPort(ad byte) byte {
	switch {
	case ad == 0xa2:
		// Read value from port
		return sound_regs[sound_regNext]
	}

	log.Fatalf("Sound, not implemented: in(%02x)", ad)
	return 0
}
