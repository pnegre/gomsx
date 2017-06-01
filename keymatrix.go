package main

import "github.com/pnegre/gogame"

var data [][]int

func init() {
	// Mirar http://map.grauw.nl/articles/keymatrix.php
	data = [][]int{
		{
			gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, // 7 &	6 ^	5 %	4 $	3 #	2 @	1 !	0 )
		},
		{
			gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, // ; :	] }	[ {	\ ¦	= +	- _	9 (	8 *
		},
		{
			gogame.K_B, gogame.K_A, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, // B	A	DEAD	/ ?	. >	, <	` ~	' "
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
		{
			gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, // F3 F F1 CODE CAPS GRAPH CTRL SHIFT
		},
		{
			gogame.K_RETURN, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, // RET SELECT BS STOP TAB ESC F5 F4
		},
		{
			gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_SPACE, // → ↓ ↑ ← DEL INS HOME SPACE
		},
		{
			gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, // NUM4	NUM3	NUM2	NUM1	NUM0	NUM/	NUM+	NUM*
		},
		{
			gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, gogame.K_UP, // NUM.	NUM,	NUM-	NUM9	NUM8	NUM7	NUM6	NUM5
		},
	}
}

func keyMatrix(row byte) (result byte) {
	result = 0xff
	if row < 11 {
		for i := 0; i < 8; i++ {
			if gogame.IsKeyPressed(data[row][i]) {
				result &= ^byte(1 << byte(7-i))
			}
		}
	} else {
		panic("KeyMatrix: Tried to scan row > 10")
	}
	return
}
