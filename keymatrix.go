package main

import "github.com/pnegre/gogame"

func keyMatrix(row byte) (result byte) {
	// Mirar http://map.grauw.nl/articles/keymatrix.php
	result = 0xff
	switch row {
	case 2:
		if gogame.IsKeyPressed(gogame.K_A) {
			result &= ^byte(1 << 6)
		}
	}
	return
}
