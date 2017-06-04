package main

import "github.com/pnegre/gogame"
import "log"

var data [][]int

func init() {
	// Mirar http://map.grauw.nl/articles/keymatrix.php
	data = [][]int{
		{
			gogame.K_7, gogame.K_6, gogame.K_5, gogame.K_4, gogame.K_3, gogame.K_2, gogame.K_1, gogame.K_0, // 7 &	6 ^	5 %	4 $	3 #	2 @	1 !	0 )
		},
		{
			gogame.K_SEMICOLON, gogame.K_RBRACKET, gogame.K_LBRACKET, gogame.K_F11, gogame.K_EQUALS, gogame.K_MINUS, gogame.K_9, gogame.K_8, // ; :	] }	[ {	\ ¦	= +	- _	9 (	8 *
		},
		{
			gogame.K_B, gogame.K_A, gogame.K_F11, gogame.K_SLASH, gogame.K_PERIOD, gogame.K_COMMA, gogame.K_F11, gogame.K_F11, // B	A	DEAD	/ ?		. >		, <		` ~		' "
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
			gogame.K_F3, gogame.K_F2, gogame.K_F1, gogame.K_F11, gogame.K_F11, gogame.K_F11, gogame.K_LCTRL, gogame.K_LSHIFT, // F3 F F1 CODE CAPS GRAPH CTRL SHIFT
		},
		{
			gogame.K_RETURN, gogame.K_F11, gogame.K_BACK, gogame.K_F11, gogame.K_F11, gogame.K_F11, gogame.K_F5, gogame.K_F4, // RET SELECT BS STOP TAB ESC F5 F4
		},
		{
			gogame.K_RIGHT, gogame.K_DOWN, gogame.K_UP, gogame.K_LEFT, gogame.K_DEL, gogame.K_INS, gogame.K_HOME, gogame.K_SPACE, // → ↓ ↑ ← DEL INS HOME SPACE
		},
		{
			gogame.K_F11, gogame.K_F11, gogame.K_F11, gogame.K_F11, gogame.K_F11, gogame.K_F11, gogame.K_F11, gogame.K_F11, // NUM4	NUM3	NUM2	NUM1	NUM0	NUM/	NUM+	NUM*
		},
		{
			gogame.K_F11, gogame.K_F11, gogame.K_F11, gogame.K_F11, gogame.K_F11, gogame.K_F11, gogame.K_F11, gogame.K_F11, // NUM.	NUM,	NUM-	NUM9	NUM8	NUM7	NUM6	NUM5
		},
	}
}

func keyMatrix(row byte) (result byte) {
	result = 0xff
	if row < 11 {
		for i := 0; i < 8; i++ {
			if gogame.IsKeyPressed(data[row][i]) {
				result &= ^byte(1 << byte(7-i))
				// log.Printf("Pressed key %d %d\n", row, i)
			}
		}
	} else {
		log.Fatalln("KeyMatrix: Tried to scan row > 10")
	}
	return
}
