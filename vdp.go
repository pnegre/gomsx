package main

import "log"

var vdp_regValue byte
var vdp_writeState = 0
var registers [27]byte

func vdp_writePort(ad byte, val byte) {
    log.Printf("VDP: Out(%02x, %02x)\n", ad, val)
    switch {
    case ad == 0x99:
        if (vdp_writeState == 0) {
            vdp_regValue = val;
            vdp_writeState = 1;
            return
        } else {
            vdp_writeState = 0
            // Bit 7 must be 1 for write
            if (val & 0x80 != 0) {
                regn := val - 127
                registers[regn] = vdp_regValue
                return
            } else {
                log.Fatalf("NOOOO")
            }
        }

    }
	log.Fatalf("Not implemented: VDP: Out(%02x, %02x)", ad, val)
}
