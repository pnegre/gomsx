package main

/*

	Exemple en msx basic:

	sound 0,105          // Set tone A frequency
	sound 1,0
	sound 7, &B10111110  // enable tone generator A
	sound 8, &B00010000  // Set amplitude for channel A (envelope)

	sound 11, 0          // Frequency for envelope
	sound 12, 20

	sound 13, &B00001000 // Select envelope "1000"


*/

import "log"

var sound_regs [16]byte
var sound_regNext byte

var sound_freqA int
var sound_volA int

var sound_tones [3]*Tone
var sound_noise *Noise

func sound_init() {
	sound_tones[0] = NewTone()
	sound_tones[1] = NewTone()
	sound_tones[2] = NewTone()
	sound_noise = NewNoise()
}

func sound_quit() {
	for i := 0; i < 3; i++ {
		sound_tones[i].dev.Close()
	}
}

func sound_writePort(ad byte, val byte) {
	switch {
	case ad == 0xa0:
		// Register write port
		sound_regNext = val
		return

	case ad == 0xa1:
		// Write value to port
		sound_regs[sound_regNext] = val
		if sound_regNext < 14 {
			sound_work()
		}
		return
	}

	log.Fatalf("Sound, not implemented: out(%02x,%02x)", ad, val)
}

func sound_readPort(ad byte) byte {
	switch {
	case ad == 0xa2:
		// Read value from port
		if sound_regNext == 0x0e {
			// joystick triggers.
			// Per ara ho posem a 1 (no moviment de joystick)
			return 0x3f
		}
		if sound_regNext == 0x0f {
			// PSG port 15 (joystick select)
			// TODO: millorar
			return 0
		}
		return sound_regs[sound_regNext]
	}

	log.Fatalf("Sound, not implemented: in(%02x)", ad)
	return 0
}

func sound_work() {
	// log.Println(sound_regs)
	for i := 0; i < 3; i++ {
		sound_doTones(i)
	}

	sound_doNoises()
}

func sound_doNoises() {
	if (sound_regs[7] & 0x38) == 0x38 {
		sound_noise.activate(false)
	} else {
		freq := int(sound_regs[6] & 0x1F)

		if freq > 0 {
			realFreq := 111861 / freq

			var vol float32 = 0
			if (sound_regs[7] & 0x20) == 0 {
				v := float32(sound_regs[8] & 0x0F)
				if v > vol {
					vol = v
				}
			}

			if (sound_regs[7] & 0x10) == 0 {
				v := float32(sound_regs[9] & 0x0F)
				if v > vol {
					vol = v
				}
			}

			if (sound_regs[7] & 0x04) == 0 {
				v := float32(sound_regs[10] & 0x0F)
				if v > vol {
					vol = v
				}
			}

			sound_noise.setParameters(realFreq, vol)
			sound_noise.activate(true)
		}
	}
}

// TODO: envelopes
func sound_doTones(chn int) {
	freq := (int(sound_regs[chn*2+1]&0x0f) << 8) | int(sound_regs[chn*2])
	envelopeEnabled := (sound_regs[8+chn] & 0x10) != 0
	if freq > 0 {
		realFreq := 111861 / freq
		if envelopeEnabled {
			envFreq := (uint16(sound_regs[12]) << 8) | uint16(sound_regs[11])
			envShape := sound_regs[13] & 0x0F
			sound_tones[chn].setEnvelope(envFreq, envShape)
		} else {
			volume := float32(sound_regs[8+chn] & 0x0F)
			volume /= 16
			volume *= 4
			sound_tones[chn].setParameters(realFreq, volume)
		}
	}
	sound_tones[chn].activate((sound_regs[7] & (0x01 << uint(chn))) == 0)
}
