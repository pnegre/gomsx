package main

import "regexp"

const (
	UNKNOWN = iota
	KONAMI4
	KONAMI5
	ASCII8KB
	ASCII16KB
	RTYPE
)

// TODO: implementar bé aquesta rutina...
func getCartType(fname string, data []byte) int {
	if match, _ := regexp.MatchString("nemesis1.rom", fname); match {
		return KONAMI4
	}
	return UNKNOWN
}

type MapperKonami4 struct {
	contents []byte
	sels     [4]int
}

func NewMapperKonami4(data []byte) Mapper {
	m := new(MapperKonami4)
	for i := 0; i < 4; i++ {
		m.sels[i] = i
	}
	m.contents = data
	return m
}

func (self *MapperKonami4) readByte(address uint16) byte {
	address -= 0x4000
	place := address / 0x2000
	realMem := self.contents[self.sels[place]*0x2000:]
	delta := address - 0x2000*place
	return realMem[delta]
}

func (self *MapperKonami4) writeByte(address uint16, value byte) {
	address -= 0x4000
	place := address / 0x2000
	if place == 0 {
		return
	}
	self.sels[place] = int(value)
}
