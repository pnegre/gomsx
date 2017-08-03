package main

import "log"

const (
	SCC_NCHANNELS = 5
)

type SCCChannel struct {
	waveform  []int
	volume    byte
	frequency int
	on        bool
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

	case n >= 0x8a && n < 0x8f:
		// Volume
		nch := n - 0x8a
		scc_channels[nch].volume = b

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
