package main

import "log"

type PPI struct {
	slots   uint8
	regc    uint8
	pgSlots [4]int
}

func NewPPI() *PPI {
	return &PPI{}
}

func (ppi *PPI) refreshSlotsValues() {
	ppi.pgSlots[0] = int(ppi.slots & 0x03)
	ppi.pgSlots[1] = int((ppi.slots & 0x0C) >> 2)
	ppi.pgSlots[2] = int((ppi.slots & 0x30) >> 4)
	ppi.pgSlots[3] = int((ppi.slots & 0xC0) >> 6)
}

func (ppi *PPI) writePort(ad byte, val byte) {
	switch {
	case ad == 0xab:
		if val&0x80 != 0 {
			log.Println("PPI initialization")
			// PPI initialization
			return
		} else {
			// Manipulate directly register C
			bitn := (val & 0x0F) >> 1
			//log.Printf("PPI: manipulate regC: Set bit %d to %d\n", bitn, vl)
			if (val & 0x01) != 0 {
				ppi.regc |= (0x01 << bitn)
			} else {
				ppi.regc &= ^(0x01 << bitn)
			}
			return
		}

	case ad == 0xa8:
		// if val != ppi.slots {
		// 	log.Printf("Set slots: %02x\n", val)
		// }
		ppi.slots = val
		ppi.refreshSlotsValues()
		return

	case ad == 0xaa:
		ppi.regc = val
		return
	}

	log.Fatalf("PPI: not implemented: out(%02x,%02x)", ad, val)
}

func (ppi *PPI) readPort(ad byte) byte {
	switch {
	case ad == 0xa8:
		// log.Printf("Get slots: %02x\n", ppi_slots)
		return ppi.slots

	case ad == 0xaa:
		return ppi.regc

	case ad == 0xa9:
		return keyMatrix(ppi.regc & 0x0f)
	}

	log.Fatalf("PPI: not implemented: in(%02x)", ad)
	return 0
}
