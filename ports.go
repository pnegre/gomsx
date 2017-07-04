package main

import "log"

type Ports struct {
}

func (self *Ports) ReadPort(address uint16) byte {
	ad := byte(address & 0xFF)
	switch {
	case ad >= 0xa8 && ad <= 0xab:
		return ppi_readPort(ad)

	case ad >= 0xa0 && ad <= 0xa2:
		return psg_readPort(ad)

	case ad >= 0x98 && ad <= 0x9b:
		return vdp_readPort(ad)
	}

	log.Printf("ReadPort: %02x\n", ad)
	return 0
}

func (self *Ports) WritePort(address uint16, b byte) {
	ad := byte(address & 0xFF)
	switch {
	case ad >= 0xa8 && ad <= 0xab:
		ppi_writePort(ad, b)
		return

	case ad >= 0xa0 && ad <= 0xa2:
		psg_writePort(ad, b)
		return

	case ad >= 0x90 && ad <= 0x91:
		// Printer. Do nothing
		return

	case ad >= 0x98 && ad <= 0x9b:
		vdp_writePort(ad, b)
		return

	case ad >= 0x00 && ad <= 0x01:
		// MIDI / Sensor Kid
		return
	}

	log.Printf("Writeport: %02x -> %02x\n", ad, b)
}
