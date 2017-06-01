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
			// Manipulate directly register C
			bitn := (val & 0x0F) >> 1
			vl := val & 0x01
			//log.Printf("PPI: manipulate regC: Set bit %d to %d\n", bitn, vl)
			if vl == 1 {
				ppi_regc |= (0x01 << (bitn + 1))
			} else {
				ppi_regc &= ^(0x01 << (bitn + 1))
			}
			return
			//panic("PPI Write command register")
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

	case ad == 0xaa:
		return ppi_regc

	case ad == 0xa9:
		return ppi_keyboardMatrix()
	}

	log.Fatalf("PPI: not implemented: in(%02x)", ad)
	return 0
}

func ppi_keyboardMatrix() byte {
	// Mirar http://map.grauw.nl/articles/keymatrix.php
	row := ppi_regc & 0x0F
	switch row {
	case 2:
		return 0xbf
	default:
		return 0xff
	}
}
