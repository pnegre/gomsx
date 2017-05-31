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
		return sound_readPort(ad)
	}

	log.Fatalf("ReadPort: %02x\n", ad)
	return 0
}

func (self *Ports) WritePort(address uint16, b byte) {
	ad := byte(address & 0xFF)
	switch {
	case ad >= 0xa8 && ad <= 0xab:
		ppi_writePort(ad, b)
		return

	case ad >= 0xa0 && ad <= 0xa2:
		sound_writePort(ad, b)
		return

	case ad >= 0x90 && ad <= 0x91:
		// Printer. Do nothing
		return

	case ad >= 0x98 && ad <= 0x9b:
		vdp_writePort(ad, b)
		return
	}

	log.Fatalf("Writeport: %02x -> %02x\n", ad, b)
}

func (self *Ports) ReadPortInternal(address uint16, contend bool) byte {
	panic("ReadPortInternal")
}

func (self *Ports) WritePortInternal(address uint16, b byte, contend bool) {
	panic("WritePortInternal")
}

func (self *Ports) ContendPortPreio(address uint16) {
	panic("ContendPortPreio")

}

func (self *Ports) ContendPortPostio(address uint16) {
	panic("ContendPortPostio")
}
