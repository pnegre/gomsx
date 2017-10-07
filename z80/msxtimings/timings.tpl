package z80

// http://map.grauw.nl/resources/z80instr.php

func initMSXTimings() {
    for i:=0; i<1536; i++ {
        timingsMSX[i] = func(z80 *Z80) uint64 { return 0 }
    }
	{{ range .}}
	/* {{.Instr}} */
	timingsMSX[{{.Opc}}] = func(z80 *Z80) uint64 { return {{.Tim}} }
	{{ end }}

    // CALL C, nn
    timingsMSX[0xdc] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_C) != 0 {
            return 18
        }
        return 11
    }

    // CALL M, nn
    timingsMSX[0xfc] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_S) != 0 {
            return 18
        }
        return 11
    }

    // CALL NC, nn
    timingsMSX[0xd4] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_C) == 0 {
            return 18
        }
        return 11
    }

    // CALL NZ, nn
    timingsMSX[0xc4] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_Z) == 0 {
            return 18
        }
        return 11
    }

    // CALL P, nn
    timingsMSX[0xf4] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_S) == 0 {
            return 18
        }
        return 11
    }

    // CALL PE, nn
    timingsMSX[0xec] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_P) != 0 {
            return 18
        }
        return 11
    }

    // CALL PO, nn
    timingsMSX[0xe4] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_P) == 0 {
            return 18
        }
        return 11
    }

    // CALL Z, nn
    timingsMSX[0xe4] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_Z) != 0 {
            return 18
        }
        return 11
    }

    // DJNZ o
    timingsMSX[0x10] = func(z80 *Z80) uint64 {
        if (z80.B != 0) {
            return 14
        }
        return 9
    }

    // JR C, nn
    timingsMSX[0x38] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_C) != 0 {
            return 13
        }
        return 8
    }

    // JR NC, nn
    timingsMSX[0x30] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_C) == 0 {
            return 13
        }
        return 8
    }

    // JR Z, nn
    timingsMSX[0x28] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_Z) != 0 {
            return 13
        }
        return 8
    }

    // JR NZ, nn
    timingsMSX[0x20] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_Z) == 0 {
            return 13
        }
        return 8
    }

    // RET C
    timingsMSX[0xd8] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_C) != 0 {
            return 12
        }
        return 6
    }

    // RET M
    timingsMSX[0xf8] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_S) != 0 {
            return 12
        }
        return 6
    }

    // RET NC
    timingsMSX[0xd0] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_C) == 0 {
            return 12
        }
        return 6
    }

    // RET NZ
    timingsMSX[0xc0] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_Z) == 0 {
            return 12
        }
        return 6
    }

    // RET P
    timingsMSX[0xf0] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_S) == 0 {
            return 12
        }
        return 6
    }

    // RET PE
    timingsMSX[0xe8] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_P) != 0 {
            return 12
        }
        return 6
    }

    // RET PO
    timingsMSX[0xe0] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_P) == 0 {
            return 12
        }
        return 6
    }

     // RET Z
    timingsMSX[0xc8] = func(z80 *Z80) uint64 {
        if (z80.F & FLAG_Z) != 0 {
            return 12
        }
        return 6
    }


}
