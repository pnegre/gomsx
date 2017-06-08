package main

import "log"

type Mapper interface {
	readByte(address uint16) byte
	writeByte(address uint16, value byte)
}

// TODO: Secondary mapper (0xFFFF)
type Memory struct {
	contents   [4][4][0x4000]byte
	mapper     Mapper
	slotMapper int
}

func NewMemory() *Memory {
	mem := new(Memory)
	mem.mapper = nil
	mem.slotMapper = -1
	return mem
}

// Loads 16k (one page)
func (self *Memory) load(data []byte, page, slot int) {
	copy(self.contents[page][slot][:], data[:0x4000])
}

func (self *Memory) setMapper(mapper Mapper, slot int) {
	log.Printf("Loading MegaROM in slot %d\n", slot)
	self.mapper = mapper
	self.slotMapper = slot
}

func (self *Memory) ReadByte(address uint16) byte {
	return self.ReadByteInternal(address)
}

// ReadByteInternal reads a byte from address without taking
// into account contention.
func (self *Memory) ReadByteInternal(address uint16) byte {
	pgSlots := ppi_getSlots()
	if self.mapper != nil && address >= 0x4000 && address <= 0xBFFF {
		if address < 0x8000 && self.slotMapper == pgSlots[1] {
			return self.mapper.readByte(address)
		}
		if address < 0xC000 && self.slotMapper == pgSlots[2] {
			return self.mapper.readByte(address)
		}
	}

	page := address / 0x4000
	delta := address - page*0x4000
	return self.contents[page][pgSlots[page]][delta]
}

// WriteByte writes a byte at address taking into account
// contention.
func (self *Memory) WriteByte(address uint16, value byte) {
	self.WriteByteInternal(address, value)
}

// WriteByteInternal writes a byte at address without taking
// into account contention.
func (self *Memory) WriteByteInternal(address uint16, value byte) {
	pgSlots := ppi_getSlots()
	if self.mapper != nil && address >= 0x4000 && address <= 0xBFFF {
		if address < 0x8000 && self.slotMapper == pgSlots[1] {
			self.mapper.writeByte(address, value)
			return
		}
		if address < 0xC000 && self.slotMapper == pgSlots[2] {
			self.mapper.writeByte(address, value)
			return
		}
	}

	page := address / 0x4000
	delta := address - page*0x4000
	self.contents[page][pgSlots[page]][delta] = value
}
