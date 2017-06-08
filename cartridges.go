package main

import "regexp"
import "log"

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
	if match, _ := regexp.MatchString("nemesis2.rom", fname); match {
		return KONAMI5
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

type MapperKonami5 struct {
	contents    []byte
	sels        [4]int
	sccActive   bool
	sccContents [0x800]byte
}

func NewMapperKonami5(data []byte) Mapper {
	m := new(MapperKonami5)
	for i := 0; i < 4; i++ {
		m.sels[i] = i
	}
	m.contents = data
	m.sccActive = false
	return m
}

func (self *MapperKonami5) readByte(address uint16) byte {
	if self.sccActive && address >= 0x9800 && address <= 0x9fff {
		return self.sccContents[address-0x9800]
	}
	if false {
		log.Printf("Konami5 read: %04x sels: %v\n", address, self.sels)
	}
	address -= 0x4000
	place := address / 0x2000
	realMem := self.contents[self.sels[place]*0x2000:]
	delta := address - 0x2000*place
	return realMem[delta]
}

func (self *MapperKonami5) writeByte(address uint16, value byte) {
	if self.sccActive && address >= 0x9800 && address <= 0x9fff {
		self.sccContents[address-0x9800] = value
	}
	switch {
	case address >= 0x5000 && address <= 0x57ff:
		self.sels[0] = int(value % 16)
		return
	case address >= 0x7000 && address <= 0x77ff:
		self.sels[1] = int(value % 16)
		return
	case address >= 0x9000 && address <= 0x97ff:
		if (value & 0x3f) == 0x3f {
			self.sccActive = true
			return
		}
		self.sels[2] = int(value % 16)
		return
	case address >= 0xb000 && address <= 0xb7ff:
		self.sels[3] = int(value % 16)
		return
	}

	address -= 0x4000
	place := address / 0x2000
	realMem := self.contents[self.sels[place]*0x2000:]
	delta := address - 0x2000*place
	realMem[delta] = value
}
