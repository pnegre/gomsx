package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

func main() {
	f, err := os.Open("timingsMSX.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	r.Comma = ';'
	rcs, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	work(rcs[1:])
}

var tplText = `
package z80

func InitMSXTimings() {
	{{ range .}}
	/* {{.Instr}} */
	timingsMSX[{{.Opc}}] = {{.Tim}}
	{{ end }}
}
`

func work(lines [][]string) {
	tp, _ := template.New("test").Parse(tplText)
	type dt struct {
		Instr string
		Tim   string
		Opc   string
	}

	var data []dt

	for _, l := range lines {
		instr, tim, opc := l[0], getTiming(l[1]), getOpcode(l[2])
		if tim == "" {
			tim = "0"
		}
		data = append(data, dt{instr, tim, opc})
	}

	tp.Execute(os.Stdout, data)
}

func getOpcode(s string) string {
	s = strings.ToLower(s)
	ar := strings.Split(s, " ")
	switch len(ar) {
	case 1:
		return "0x" + ar[0]
	case 2:
		if ar[0] == "cb" {
			return "SHIFT_0xCB+" + "0x" + ar[1]
		} else if ar[0] == "ed" {
			return "SHIFT_0xED+" + "0x" + ar[1]
		} else if ar[0] == "dd" {
			return "SHIFT_0xDD+" + "0x" + ar[1]
		} else if ar[0] == "fd" {
			return "SHIFT_0xFD+" + "0x" + ar[1]
		}
		panic("Opcode 2byte not valid")
	case 3:
		if ar[0] == "dd" {
			return "SHIFT_0xDDCB+" + "0x" + ar[2]
		} else if ar[0] == "fd" {
			return "SHIFT_0xDDCB+" + "0x" + ar[2]
		}
		panic("Opcode 3byte not valid")
	default:
		log.Fatalf("Error getOpcode: %v", ar)
	}
	return ""
}

func getTiming(s string) string {
	// TODO: en les operacions on hi ha dos timings possibles, fem la mitjana
	// Això no és ideal i s'ha de canviar...
	ar := strings.Split(s, "/")
	switch len(ar) {
	case 1:
		return s
	case 2:
		i1, _ := strconv.Atoi(ar[0])
		i2, _ := strconv.Atoi(ar[1])
		return fmt.Sprintf("%d", (i1+i2)/2)
	}
	log.Fatal("Error timing")
	return ""
}
