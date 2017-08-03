package main

import "log"

const (
	SCC_NCHANNELS = 5
)

type SCCChannel struct {
	waveform   []int
	volume     float32
	frequency  float32
	sccRegFreq [2]byte
	on         bool
	count      int
}

var scc_channels []*SCCChannel

func init() {
	for i := 0; i < SCC_NCHANNELS; i++ {
		scc_channels = append(scc_channels, NewSCCChannel())
	}
}

func NewSCCChannel() *SCCChannel {
	return &SCCChannel{}
}

func (self *SCCChannel) feedSamples(data []int16) {
	if !self.on {
		return
	}

	if self.frequency > FREQUENCY/2 {
		return
	}

	// Based on code from EMULIB, Sound.c, function RenderAudio()

	K := int(0x10000 * self.frequency / FREQUENCY)
	L1 := self.count
	var A1 int16
	for i := 0; i < len(data); i, L1 = i+1, L1+K {
		L2 := L1 + K
		if L1&0x8000 != 0 {
			A1 = 127
		} else {
			A1 = -128
		}
		if (L1^L2)&0x8000 != 0 {
			A1 = A1 * int16((0x8000-(L1&0x7FFF)-(L2&0x7FFF))/K)
		}
		data[i] += int16(float32(A1) * self.volume)
	}
	self.count = L1 & 0xFFFF
}

func scc_feedSamples(data []int16) {
	for _, c := range scc_channels {
		c.feedSamples(data)
	}
}

func scc_write(n uint16, b byte) {
	switch {
	case n < 0x20:
		// Waveform channel 1
	case n >= 0x20 && n < 0x40:
		// Wafeform channel 2
	case n >= 0x40 && n < 0x60:
		// Wafeform channel 3
	case n >= 0x60 && n < 0x80:
		// Wafeform channel 4 & 5
	case n >= 0x80 && n < 0x8a:
		// Frequency
		nch := (n - 0x80) / 2
		scc_channels[nch].sccRegFreq[n%2] = b
		freq := int(scc_channels[nch].sccRegFreq[0]) | (int(scc_channels[nch].sccRegFreq[1]&0x0f) << 8)
		if freq > 0 {
			scc_channels[nch].frequency = float32(111861) / float32(freq)
		}

	case n >= 0x8a && n < 0x8f:
		// Volume
		nch := n - 0x8a
		scc_channels[nch].volume = float32(b)

	case n == 0x8f:
		// ON/OFF switch channel 1 to 5
		for i := 0; i < 5; i++ {
			var m byte = 0x01 << uint(i)
			if m&b != 0 {
				scc_channels[i].on = true
			} else {
				scc_channels[i].on = false
			}
		}
	default:
		log.Printf("SCC: %x -> %d\n", n, b)
	}
}
