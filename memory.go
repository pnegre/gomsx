package main

import "log"

type Mapper interface {
	readByte(address uint16) byte
	writeByte(address uint16, value byte)
}

// TODO: Secondary mapper (0xFFFF)
type Memory struct {
	contents   [4][4][0x4000]byte
	canWrite   [4][4]bool
	mapper     Mapper
	slotMapper int
}

func NewMemory() *Memory {
	mem := new(Memory)
	mem.mapper = nil
	mem.slotMapper = -1
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			mem.canWrite[i][j] = true
		}
	}
	return mem
}

func (self *Memory) saveState() *Memory {
	m := Memory{}
	m = *self
	return &m
}

func (self *Memory) restoreState(m *Memory) {
	*self = *m
}

// Loads 16k (one page)
func (self *Memory) load(data []byte, page, slot int) {
	copy(self.contents[page][slot][:], data[:0x4000])
	self.canWrite[page][slot] = false
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
	page := address / 0x4000
	slot := ppi_pgSlots[page]

	if self.mapper != nil && self.slotMapper == slot && (page == 1 || page == 2) {
		return self.mapper.readByte(address)
	}

	delta := address - page*0x4000
	return self.contents[page][slot][delta]
}

// WriteByte writes a byte at address taking into account
// contention.
func (self *Memory) WriteByte(address uint16, value byte) {
	self.WriteByteInternal(address, value)
}

// WriteByteInternal writes a byte at address without taking
// into account contention.
func (self *Memory) WriteByteInternal(address uint16, value byte) {
	page := address / 0x4000
	slot := ppi_pgSlots[page]

	if self.mapper != nil && self.slotMapper == slot && (page == 1 || page == 2) {
		self.mapper.writeByte(address, value)
		return
	}

	if self.canWrite[page][slot] {
		delta := address - page*0x4000
		self.contents[page][slot][delta] = value
	}
}
