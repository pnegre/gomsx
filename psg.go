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

type PSG struct {
	regs        [16]byte
	regNext     byte
	bytesCass   []byte
	sound_tones [3]*ToneGenerator
}

func NewPSG() *PSG {
	psg := &PSG{}
	psg.sound_tones[0] = NewToneGenerator()
	psg.sound_tones[1] = NewToneGenerator()
	psg.sound_tones[2] = NewToneGenerator()
	return psg
}

func (psg *PSG) feedSamples(data []int16) {
	psg.sound_tones[0].feedSamples(data)
	psg.sound_tones[1].feedSamples(data)
	psg.sound_tones[2].feedSamples(data)
}

// func psg_loadCassette(fileName string) {
// 	var err error
// 	psg_bytesCass, err = ioutil.ReadFile(fileName)
// 	if err != nil {
// 		log.Println(err)
// 		psg_bytesCass = nil
// 	}
// 	log.Println("PSG: Loaded cassete:", fileName)
// }

func (psg *PSG) writePort(ad byte, val byte) {
	switch {
	case ad == 0xa0:
		// Register write port
		psg.regNext = val
		return

	case ad == 0xa1:
		// Write value to port
		psg.regs[psg.regNext] = val
		if psg.regNext < 14 {
			for i := 0; i < 3; i++ {
				psg.doTones(i)
			}

			// TODO: sound_doNoises()
		}
		return
	}

	log.Fatalf("Sound, not implemented: out(%02x,%02x)", ad, val)
}

func (psg *PSG) readPort(ad byte) byte {
	switch {
	case ad == 0xa2:
		// Read value from port
		if psg.regNext == 0x0e {
			// joystick triggers i cassete input
			bitCass := cassete_getNextBit() << 7
			// Per ara ho posem a 1 (no moviment de joystick)
			return 0x3f | bitCass
		}
		if psg.regNext == 0x0f {
			// PSG port 15 (joystick select)
			// TODO: millorar
			return 0
		}
		return psg.regs[psg.regNext]
	}

	log.Fatalf("Sound, not implemented: in(%02x)", ad)
	return 0
}

func cassete_getNextBit() byte {
	// log.Println("NextByte")
	return 0
}

// TODO: envelopes
func (psg *PSG) doTones(chn int) {
	freq := (int(psg.regs[chn*2+1]&0x0f) << 8) | int(psg.regs[chn*2])
	envelopeEnabled := (psg.regs[8+chn] & 0x10) != 0
	if freq > 0 {
		realFreq := float32(111861) / float32(freq)
		if envelopeEnabled {
			// envFreq := (uint16(psg.regs[12]) << 8) | uint16(psg.regs[11])
			// envShape := psg.regs[13] & 0x0F
			// sound_tones[chn].setEnvelope(envFreq, envShape)
		} else {
			volume := float32(psg.regs[8+chn] & 0x0F)
			psg.sound_tones[chn].setVolume(volume)
			psg.sound_tones[chn].setFrequency(realFreq)
		}
	}
	psg.sound_tones[chn].activate((psg.regs[7] & (0x01 << uint(chn))) == 0)
}

func sound_doNoises() {
	// if (psg_regs[7] & 0x38) == 0x38 {
	// 	// sound_noise.activate(false)
	// } else {
	// 	freq := int(psg_regs[6] & 0x1F)
	//
	// 	if freq > 0 {
	// 		realFreq := float32(111861) / float32(freq)
	//
	// 		var vol float32 = 0
	// 		if (psg_regs[7] & 0x20) == 0 {
	// 			v := float32(psg_regs[8] & 0x0F)
	// 			if v > vol {
	// 				vol = v
	// 			}
	// 		}
	//
	// 		if (psg_regs[7] & 0x10) == 0 {
	// 			v := float32(psg_regs[9] & 0x0F)
	// 			if v > vol {
	// 				vol = v
	// 			}
	// 		}
	//
	// 		if (psg_regs[7] & 0x04) == 0 {
	// 			v := float32(psg_regs[10] & 0x0F)
	// 			if v > vol {
	// 				vol = v
	// 			}
	// 		}
	//
	// 		// sound_noise.setParameters(realFreq, vol)
	// 		// sound_noise.activate(true)
	// 	}
	// }
}
