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
import "github.com/pnegre/gogame"

const FREQUENCY = 22000

var psg_regs [16]byte
var psg_regNext byte

var sound_tones [3]*ToneGenerator
var sound_device *gogame.AudioDevice

// var sound_noise *Noise

func psg_init() {
	sound_device, _ = gogame.NewAudioDevice(FREQUENCY)
	sound_tones[0] = NewToneGenerator()
	sound_tones[1] = NewToneGenerator()
	sound_tones[2] = NewToneGenerator()
	// sound_noise = NewNoise()
	sound_device.SetCallback(sound_callback)
	// gogame.RegisterSoundCallback(sound_callback)
	sound_device.Start()
}

func sound_callback(data []int16) {
	for i := 0; i < len(data); i++ {
		data[i] = 0
	}
	sound_tones[0].feedSamples(data)
	sound_tones[1].feedSamples(data)
	sound_tones[2].feedSamples(data)

	for i := 0; i < len(data); i++ {
		if data[i] > 32760 {
			data[i] = 32760
		}
		if data[i] < -32760 {
			data[i] = -32760
		}
	}
}

func psg_quit() {
	sound_device.Stop()
	sound_device.Close()
}

func psg_loadCassette(fileName string) {
	log.Println("PSG: Loading Cassette file:", fileName)
}

func psg_writePort(ad byte, val byte) {
	switch {
	case ad == 0xa0:
		// Register write port
		psg_regNext = val
		return

	case ad == 0xa1:
		// Write value to port
		psg_regs[psg_regNext] = val
		if psg_regNext < 14 {
			for i := 0; i < 3; i++ {
				sound_doTones(i)
			}

			// TODO: sound_doNoises()
		}
		return
	}

	log.Fatalf("Sound, not implemented: out(%02x,%02x)", ad, val)
}

func psg_readPort(ad byte) byte {
	switch {
	case ad == 0xa2:
		// Read value from port
		if psg_regNext == 0x0e {
			// joystick triggers i cassete input
			bitCass := cassete_getNextBit() << 7
			// Per ara ho posem a 1 (no moviment de joystick)
			return 0x3f | bitCass
		}
		if psg_regNext == 0x0f {
			// PSG port 15 (joystick select)
			// TODO: millorar
			return 0
		}
		return psg_regs[psg_regNext]
	}

	log.Fatalf("Sound, not implemented: in(%02x)", ad)
	return 0
}

func cassete_getNextBit() byte {
	// log.Println("NextByte")
	return 0
}

// TODO: envelopes
func sound_doTones(chn int) {
	freq := (int(psg_regs[chn*2+1]&0x0f) << 8) | int(psg_regs[chn*2])
	envelopeEnabled := (psg_regs[8+chn] & 0x10) != 0
	if freq > 0 {
		realFreq := float32(111861) / float32(freq)
		if envelopeEnabled {
			// envFreq := (uint16(psg_regs[12]) << 8) | uint16(psg_regs[11])
			// envShape := psg_regs[13] & 0x0F
			// sound_tones[chn].setEnvelope(envFreq, envShape)
		} else {
			volume := float32(psg_regs[8+chn] & 0x0F)
			sound_tones[chn].setParameters(realFreq, volume)
		}
	}
	sound_tones[chn].activate((psg_regs[7] & (0x01 << uint(chn))) == 0)
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
