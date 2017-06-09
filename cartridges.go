package main

import "crypto/sha1"
import "log"
import "fmt"

const (
	UNKNOWN = iota
	KONAMI4
	KONAMI5
	ASCII8KB
	ASCII16KB
	RTYPE
)

func getCartType(data []byte) int {
	hash := fmt.Sprintf("%x", sha1.Sum(data))
	log.Printf("Hash: %s\n", hash)
	switch hash {
	case "e31ac6520e912c27ce96431a1dfb112bf71cb7b9":
		// Nemesis 1
		return KONAMI4
	case "ab30cdeaacbdf14e6366d43d881338178fc665cb":
		// Nemesis 2
		return KONAMI5
	case "709fb35338f21897e275237cc4c5615d0a5c2753":
		// Batman
		return ASCII8KB
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
	contents []byte
	numBanks byte
	sels     [4]int
}

func NewMapperKonami5(data []byte) Mapper {
	m := new(MapperKonami5)
	for i := 0; i < 4; i++ {
		m.sels[i] = i
	}
	m.contents = data
	m.numBanks = byte(len(data) / 8192)
	return m
}

func (self *MapperKonami5) readByte(address uint16) byte {
	address -= 0x4000
	place := address / 0x2000
	realMem := self.contents[self.sels[place]*0x2000:]
	delta := address - 0x2000*place
	return realMem[delta]
}

func (self *MapperKonami5) writeByte(address uint16, value byte) {
	switch {
	case address >= 0x5000 && address <= 0x57ff:
		self.sels[0] = int(value % self.numBanks)
		return
	case address >= 0x7000 && address <= 0x77ff:
		self.sels[1] = int(value % self.numBanks)
		return
	case address >= 0x9000 && address <= 0x97ff:
		self.sels[2] = int(value % self.numBanks)
		return
	case address >= 0xb000 && address <= 0xb7ff:
		self.sels[3] = int(value % self.numBanks)
		return
	}

	// address -= 0x4000
	// place := address / 0x2000
	// realMem := self.contents[self.sels[place]*0x2000:]
	// delta := address - 0x2000*place
	// realMem[delta] = value
}

type MapperASCII8 struct {
	contents []byte
	numBanks byte
	sels     [4]int
}

func NewMapperASCII8(data []byte) Mapper {
	m := new(MapperASCII8)
	for i := 0; i < 4; i++ {
		m.sels[i] = i
	}
	m.contents = data
	m.numBanks = byte(len(data) / 8192)
	return m
}

func (self *MapperASCII8) readByte(address uint16) byte {
	address -= 0x4000
	place := address / 0x2000
	realMem := self.contents[self.sels[place]*0x2000:]
	delta := address - 0x2000*place
	return realMem[delta]
}

func (self *MapperASCII8) writeByte(address uint16, value byte) {
	switch {
	case address >= 0x6000 && address <= 0x67ff:
		self.sels[0] = int(value % self.numBanks)
		return
	case address >= 0x6800 && address <= 0x6fff:
		self.sels[1] = int(value % self.numBanks)
		return
	case address >= 0x7000 && address <= 0x77ff:
		self.sels[2] = int(value % self.numBanks)
		return
	case address >= 0x7800 && address <= 0x7fff:
		self.sels[3] = int(value % self.numBanks)
		return
	}
}
