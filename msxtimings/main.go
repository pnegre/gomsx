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

func InitMSXTimings {
	{{ range .}}
	/* {{.Instr}} */
	TimingsMSX[{{.Opc}}] = {{.Tim}}
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
	ar := strings.Split(s, " ")
	switch len(ar) {
	case 1:
		return "0x0000" + ar[0]
	case 2:
		return "0x00" + ar[0] + ar[1]
	case 3:
		return "0x" + ar[0] + ar[1] + ar[2]
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
