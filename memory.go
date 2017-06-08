package main

import "log"

type Mapper interface {
	readByte(address uint16) byte
	writeByte(address uint16, value byte)
}

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
	// if address == 0xffff {
	// 	// log.Printf("Get secondary memory mapper\n")
	// 	return self.ffff
	// }

	pgSlots := []byte{
		ppi_slots & 0x03,
		(ppi_slots & 0x0C) >> 2,
		(ppi_slots & 0x30) >> 4,
		(ppi_slots & 0xC0) >> 6,
	}
	// pg0Slot := ppi_slots & 0x03
	// pg1Slot := (ppi_slots & 0x0C) >> 2
	// pg2Slot := (ppi_slots & 0x30) >> 4
	// pg3Slot := (ppi_slots & 0xC0) >> 6

	if self.mapper != nil && address >= 0x4000 && address <= 0xBFFF {
		if address < 0x8000 && self.slotMapper == int(pgSlots[1]) {
			return self.mapper.readByte(address)
		}
		if address < 0xC000 && self.slotMapper == int(pgSlots[2]) {
			return self.mapper.readByte(address)
		}
	}

	page := address / 0x4000
	delta := address - page*0x4000
	return self.contents[page][pgSlots[page]][delta]
	// switch page {
	// case 0:
	// 	return self.contents[0][pg0Slot][address]
	// case 1:
	// 	return self.contents[1][pg1Slot][address-0x4000]
	// case 2:
	// 	return self.contents[2][pg2Slot][address-0x8000]
	// case 3:
	// 	return self.contents[3][pg3Slot][address-0xC000]
	// }
	// panic("Tried to read impossible memory location")
}

// WriteByte writes a byte at address taking into account
// contention.
func (self *Memory) WriteByte(address uint16, value byte) {
	self.WriteByteInternal(address, value)
}

// WriteByteInternal writes a byte at address without taking
// into account contention.
func (self *Memory) WriteByteInternal(address uint16, value byte) {
	// if address == 0xffff {
	// 	// log.Printf("Set secondary memory mapper: %02x\n", value)
	// 	self.ffff = value
	// 	return
	// }
	pg0Slot := ppi_slots & 0x03
	pg1Slot := (ppi_slots & 0x0C) >> 2
	pg2Slot := (ppi_slots & 0x30) >> 4
	pg3Slot := (ppi_slots & 0xC0) >> 6

	if self.mapper != nil && address >= 0x4000 && address <= 0xBFFF {
		if address < 0x8000 && self.slotMapper == int(pg1Slot) {
			self.mapper.writeByte(address, value)
			return
		}
		if address < 0xC000 && self.slotMapper == int(pg2Slot) {
			self.mapper.writeByte(address, value)
			return
		}
	}

	page := address / 0x4000
	switch page {
	case 0:
		self.contents[0][pg0Slot][address] = value
		return
	case 1:
		self.contents[1][pg1Slot][address-0x4000] = value
		return
	case 2:
		self.contents[2][pg2Slot][address-0x8000] = value
		return
	case 3:
		self.contents[3][pg3Slot][address-0xC000] = value
		return
	}
	panic("Tried to write impossible memory location")
}
