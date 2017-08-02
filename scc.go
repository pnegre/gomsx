package main

import "log"

type SCCChannel struct {
	waveform  []int
	volume    int
	frequency int
	on        bool
}

var scc_channels []*SCCChannel

func init() {
	scc_channels = make([]*SCCChannel, 5)
	for i := 0; i < len(scc_channels); i++ {
		scc_channels[i] = NewSCCChannel()
	}
}

func NewSCCChannel() *SCCChannel {
	return &SCCChannel{}
}

func (self *SCCChannel) feedSamples(data []int16) {
	if !self.on {
		return
	}
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
	case n >= 0x80 && n < 0x82:
		// Frequency channel 1
	case n >= 0x82 && n < 0x84:
		// Frequency channel 2
	case n >= 0x84 && n < 0x86:
		// Frequency channel 3
	case n >= 0x86 && n < 0x88:
		// Frequency channel 4
	case n >= 0x88 && n < 0x8a:
		// Frequency channel 5
	case n == 0x8a:
		// Volume channel 1
	case n == 0x8b:
		// Volume channel 2
	case n == 0x8c:
		// Volume channel 3
	case n == 0x8d:
		// Volume channel 4
	case n == 0x8e:
		// Volume channel 5
	case n == 0x8f:
		// ON/OFF switch channel 1 to 5
	default:
		log.Printf("SCC: %d -> %d\n", n, b)
	}
}
