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

var sound_devices [3]*SoundDevice

func sound_init() {
	sound_devices[0] = NewSoundDevice()
	sound_devices[1] = NewSoundDevice()
	sound_devices[2] = NewSoundDevice()
}

func sound_quit() {
	for i := 0; i < 3; i++ {
		sound_devices[i].dev.Close()
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

type SoundDevice struct {
	dev    *gogame.ToneGenerator
	volume int
	freq   int
	active bool
}

func NewSoundDevice() *SoundDevice {
	sd := new(SoundDevice)
	var err error
	if sd.dev, err = gogame.NewToneGenerator(); err != nil {
		panic("Error creating tone generator!")
	}

	sd.active = false
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

func (self *SoundDevice) activate(act bool) {
	if self.active != act {
		self.active = act
		if act {
			self.dev.Start()
		} else {
			self.dev.Stop()
		}
	}
}

func sound_work() {
	// log.Println(sound_regs)
	for i := 0; i < 3; i++ {
		sound_workChannel(i)
	}
}

// TODO: envelopes
func sound_workChannel(chn int) {
	fa := (int(sound_regs[chn*2+1]&0x0f) << 8) | int(sound_regs[chn*2])
	va := int(sound_regs[8+chn] & 0x0F)
	if fa > 0 {
		realFreq := 111861 / fa
		sound_devices[chn].setParameters(realFreq, va)
	}
	sound_devices[chn].activate((sound_regs[7] & (0x01 << uint(chn))) == 0)
}
