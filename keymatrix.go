package main

import "github.com/pnegre/gogame"

func keyMatrix(row byte) (result byte) {
	// Mirar http://map.grauw.nl/articles/keymatrix.php
	data := [][]int{
		{
			gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP,
		},
		{
			gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP,
		},
		{
			gogame.K_B, gogame.K_A, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP,
		},
		{
			gogame.K_J, gogame.K_I, gogame.K_H, gogame.K_G, gogame.K_F, gogame.K_E, gogame.K_D, gogame.K_C,
		},
		{
			gogame.K_R, gogame.K_Q, gogame.K_P, gogame.K_O, gogame.K_N, gogame.K_M, gogame.K_L, gogame.K_K,
		},
		{
			gogame.K_Z, gogame.K_Y, gogame.K_X, gogame.K_W, gogame.K_V, gogame.K_U, gogame.K_T, gogame.K_S,
		},
	}

	result = 0xff
	switch {
	case row < 6:
		for i := 0; i < 8; i++ {
			if gogame.IsKeyPressed(data[row][i]) {
				result &= ^byte(1 << byte(7-i))
			}
		}

	}
	return
}
