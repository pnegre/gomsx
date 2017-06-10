package main

/*

	Exemple en msx basic:

	sound 0,105         // Set tone A frequency
	sound 1,0
	sound 7, &B10111110 // enable tone generator A
	sound 8, &B00001111 // Set amplitude for channel A

*/

import "github.com/pnegre/gogame"
import "log"

var sound_regs [16]byte
var sound_regNext byte

var sound_freqA int
var sound_volA int

type SoundDevice struct {
	dev    *gogame.SoundDevice
	volume int
	freq   int
}

var sound_devices [3]*SoundDevice

func NewSoundDevice() *SoundDevice {
	sd := new(SoundDevice)
	sd.dev, _ = gogame.NewSoundDevice()
	sd.dev.Start()
	return sd
}

func (self *SoundDevice) setParameters(freq int, vol int) {
	if self.volume != vol || self.freq != freq {
		self.volume = vol
		self.freq = freq
		self.dev.SetAmplitude(vol)
		self.dev.SetFreq(freq)
	}
}

func sound_init() {
	sound_devices[0] = NewSoundDevice()
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
	fa := (int(sound_regs[1]&0x0f) << 8) | int(sound_regs[0])
	va := int(sound_regs[8] & 0x0F)
	if fa > 0 {
		realFreqA := 111861 / fa
		sound_devices[0].setParameters(realFreqA, va)
	}
}
