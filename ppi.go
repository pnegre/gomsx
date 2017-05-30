package main

import "log"

var ppi_slots uint8
var ppi_regc uint8

func ppi_writePort(ad byte, val byte) {
	switch {
	case ad == 0xab:
		if val&0x80 != 0 {
			log.Println("PPI initialization")
			// PPI initialization
			return
		} else {
			panic("PPI Write command register")
		}

	case ad == 0xa8:
		// TODO: manage slots
		ppi_slots = val
		return

	case ad == 0xaa:
		ppi_regc = val
		return
	}

	log.Fatalf("PPI: not implemented: out(%02x,%02x)", ad, val)
}

func ppi_readPort(ad byte) byte {
	switch {
	case ad == 0xa8:
		return ppi_slots
	}

	log.Fatalf("PPI: not implemented: in(%02x)", ad)
	return 0
}
