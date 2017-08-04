package main

type ToneGenerator struct {
	amp    float32
	freq   float32
	count  int
	active bool
	wform  [32]int16
	index  float32
}

var SQWAVE []byte

func init() {
	SQWAVE = []byte{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
	}
}
func NewToneGenerator() *ToneGenerator {
	sd := new(ToneGenerator)
	sd.updateWaveform(SQWAVE)
	return sd
}

func (self *ToneGenerator) updateWaveform(data []byte) {
	// Set waveform of tone generator
	// TODO: implementar
	for i := 0; i < len(data); i++ {
		self.wform[i] = int16(data[i]) - 127
	}
}

func (self *ToneGenerator) setVolume(volume float32) {
	self.amp = volume / 2
}

func (self *ToneGenerator) setFrequency(freq float32) {
	self.freq = freq
}

func (self *ToneGenerator) activate(par bool) {
	self.active = par
}

func (self *ToneGenerator) feedSamples(data []int16) {
	if !self.active || self.freq == 0 || self.amp == 0 {
		return
	}

	if self.freq > FREQUENCY/2 {
		return
	}

	nsamples := FREQUENCY / (self.freq)
	delta := float32(32) / nsamples
	for i := 0; i < len(data); i++ {
		data[i] += int16(float32(self.wform[int(self.index)]) * self.amp)
		self.index += delta
		for self.index >= 32 {
			self.index -= 32
		}
	}

}
