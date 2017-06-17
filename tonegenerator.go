package main

type ToneGenerator struct {
	amp    float32
	freq   float32
	count  int
	active bool
}

func NewToneGenerator() *ToneGenerator {
	sd := new(ToneGenerator)
	return sd
}

func (self *ToneGenerator) setParameters(freq float32, volume float32) {
	self.freq = freq
	self.amp = volume
}

func (self *ToneGenerator) activate(par bool) {
	self.active = par
}

func (self *ToneGenerator) feedSamples(data []float32) {
	if !self.active || self.freq == 0 || self.amp == 0 {
		return
	}

	if self.freq > FREQUENCY/2 {
		return
	}
	/*

				Tret de EMULIB


				if(WaveCH[J].Freq>=SndRate/2) break;
		          K=0x10000*WaveCH[J].Freq/SndRate;
		          L1=WaveCH[J].Count;

				  for(I=0;I<Samples;I++,L1+=K)
		          {
		            L2 = L1+K;
		            A1 = L1&0x8000? 127:-128;
		            if((L1^L2)&0x8000)
		              A1=A1*(0x8000-(L1&0x7FFF)-(L2&0x7FFF))/K;
		            Wave[I]+=A1*V;
		          }
		          WaveCH[J].Count=L1&0xFFFF;

	*/

	K := int(0x10000 * self.freq / FREQUENCY)
	L1 := self.count
	var A1 float32
	for i := 0; i < len(data); i, L1 = i+1, L1+K {
		L2 := L1 + K
		if L1&0x8000 != 0 {
			A1 = 127
		} else {
			A1 = -128
		}
		if (L1^L2)&0x8000 != 0 {
			A1 = A1 * (0x8000 - float32(L1&0x7FFF) - float32(L2&0x7FFF)) / float32(K)
		}
		data[i] += A1 * self.amp * 0.0005

		// data[i] = self.amp * float32(math.Sin(float64(self.v*2*math.Pi*self.freq)))
		// self.v += INVFREQUENCY
		// if self.j > self.period {
		// 	self.v -= float32(self.j) * INVFREQUENCY
		// 	self.j = 0
		// } else {
		// 	self.j++
		// }
	}
	self.count = L1 & 0xFFFF

}
