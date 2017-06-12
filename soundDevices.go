package main

import "github.com/pnegre/gogame"
import "log"

type Tone struct {
	dev      *gogame.ToneGenerator
	volume   float32
	freq     int
	active   bool
	envFreq  uint16
	envShape byte
}

func NewTone() *Tone {
	sd := new(Tone)
	var err error
	if sd.dev, err = gogame.NewToneGenerator(gogame.GENERATOR_TYPE_TONE); err != nil {
		panic("Error creating tone generator!")
	}

	sd.active = false
	return sd
}

func (self *Tone) setParameters(freq int, vol float32) {
	if self.volume != vol || self.freq != freq {
		self.volume = vol
		self.freq = freq
		self.dev.SetAmplitude(vol)
		self.dev.SetFreq(freq)
	}
}

func (self *Tone) setEnvelope(envFreq uint16, envShape byte) {
	// TODO: implementar...
	if envFreq != self.envFreq || envShape != self.envShape {
		log.Printf("Set envelope: %d %d\n", envFreq, envShape)
		self.envShape = envShape
		self.envFreq = envFreq
	}
}

func (self *Tone) activate(act bool) {
	if self.active != act {
		self.active = act
		if act {
			self.dev.Start()
		} else {
			self.dev.Stop()
		}
	}
}

type Noise struct {
	dev    *gogame.ToneGenerator
	volume float32
	freq   int
	active bool
}

func NewNoise() *Noise {
	sd := new(Noise)
	var err error
	if sd.dev, err = gogame.NewToneGenerator(gogame.GENERATOR_TYPE_NOISE); err != nil {
		panic("Error creating tone generator!")
	}

	sd.active = false
	return sd
}

func (self *Noise) activate(act bool) {
	if self.active != act {
		self.active = act
		if act {
			self.dev.Start()
		} else {
			self.dev.Stop()
		}
	}
}

func (self *Noise) setParameters(freq int, vol float32) {
	if self.volume != vol || self.freq != freq {
		self.volume = vol
		self.freq = freq
		self.dev.SetAmplitude(vol)
		self.dev.SetFreq(freq)
	}
}
