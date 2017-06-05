package main

import "log"

type Memory struct {
	page0 [4][0x4000]byte
	page1 [4][0x4000]byte
	page2 [4][0x4000]byte
	page3 [4][0x4000]byte
	ffff  byte
}

func NewMemory() *Memory {
	mem := new(Memory)
	return mem
}

// Loads 16k (one page)
func (self *Memory) load(data []byte, page, slot int) {
	switch page {
	case 0:
		copy(self.page0[slot][:], data[:0x4000])
		return
	case 1:
		copy(self.page1[slot][:], data[:0x4000])
		return
	case 2:
		copy(self.page2[slot][:], data[:0x4000])
		return
	case 3:
		copy(self.page3[slot][:], data[:0x4000])
		return
	}
}

func (self *Memory) ReadByte(address uint16) byte {
	return self.ReadByteInternal(address)
}

// ReadByteInternal reads a byte from address without taking
// into account contention.
func (self *Memory) ReadByteInternal(address uint16) byte {
	if address == 0xffff {
		log.Printf("Get secondary memory mapper\n")
		return self.ffff
	}
	pg0Slot := ppi_slots & 0x03
	pg1Slot := (ppi_slots & 0x0C) >> 2
	pg2Slot := (ppi_slots & 0x30) >> 4
	pg3Slot := (ppi_slots & 0xC0) >> 6
	page := address / 0x4000
	switch page {
	case 0:
		return self.page0[pg0Slot][address]
	case 1:
		return self.page1[pg1Slot][address-0x4000]
	case 2:
		return self.page2[pg2Slot][address-0x8000]
	case 3:
		return self.page3[pg3Slot][address-0xC000]
	}
	panic("Tried to read impossible memory location")
}

// WriteByte writes a byte at address taking into account
// contention.
func (self *Memory) WriteByte(address uint16, value byte) {
	self.WriteByteInternal(address, value)
}

// WriteByteInternal writes a byte at address without taking
// into account contention.
func (self *Memory) WriteByteInternal(address uint16, value byte) {
	if address == 0xffff {
		log.Printf("Set secondary memory mapper: %02x\n", value)
		self.ffff = value
		return
	}
	pg0Slot := ppi_slots & 0x03
	pg1Slot := (ppi_slots & 0x0C) >> 2
	pg2Slot := (ppi_slots & 0x30) >> 4
	pg3Slot := (ppi_slots & 0xC0) >> 6
	page := address / 0x4000
	switch page {
	case 0:
		self.page0[pg0Slot][address] = value
		return
	case 1:
		self.page1[pg1Slot][address-0x4000] = value
		return
	case 2:
		self.page2[pg2Slot][address-0x8000] = value
		return
	case 3:
		self.page3[pg3Slot][address-0xC000] = value
		return
	}
	panic("Tried to write impossible memory location")
}
