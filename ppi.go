package main

import "log"

var ppi_slots uint8

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
		ppi_slots = val
		return
	}

	log.Fatalf("PPI: not implemented: %02x -> %02x", ad, val)
}
