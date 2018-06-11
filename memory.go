package main

import "log"
import "io/ioutil"

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
	ppi        *PPI
}

func NewMemory(ppi *PPI) *Memory {
	mem := new(Memory)
	mem.mapper = nil
	mem.slotMapper = -1
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			mem.canWrite[i][j] = true
		}
	}
	mem.ppi = ppi
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

func (self *Memory) loadBiosBasic(fname string) {
	buffer, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	// Load BIOS
	self.load(buffer, 0, 0)
	if len(buffer) > 0x4000 {
		// Load BASIC, if present
		self.load(buffer[0x4000:], 1, 0)
	}
}

func (self *Memory) loadRom(fname string, slot int) {
	buffer, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	switch getCartType(buffer) {
	case KONAMI4:
		log.Printf("Loading ROM %s to slot 1 as type KONAMI4\n", fname)
		mapper := NewMapperKonami4(buffer)
		self.setMapper(mapper, slot)
		return

	case KONAMI5:
		log.Printf("Loading ROM %s to slot 1 as type KONAMI5\n", fname)
		mapper := NewMapperKonami5(buffer)
		self.setMapper(mapper, slot)
		return

	case ASCII8KB:
		log.Printf("Loading ROM %s to slot 1 as type ASCII8KB\n", fname)
		mapper := NewMapperASCII8(buffer)
		self.setMapper(mapper, slot)
		return

	case NORMAL:
		log.Println("Cartridge is type NORMAL")

	case UNKNOWN:
		log.Println("Cartridge is type UNKNOWN")
	}

	log.Printf("Trying to load as a standard cartridge...\n")

	npages := len(buffer) / 0x4000
	switch npages {
	case 1:
		// Load ROM to page 1, slot 1
		// TODO: mirrored????
		log.Printf("Loading ROM %s to slot 1 (16KB)\n", fname)
		self.load(buffer, 1, slot)
	case 2:
		// Load ROM to slot 1. Mirrored pg1&pg2 <=> pg3&pg4
		log.Printf("Loading ROM %s to slot 1 (32KB)\n", fname)
		self.load(buffer, 0, slot)
		self.load(buffer, 1, slot)
		self.load(buffer[0x4000:], 2, slot)
		self.load(buffer[0x4000:], 3, slot)
	case 4:
		log.Printf("Loading ROM %s to slot 1 (64KB)\n", fname)
		self.load(buffer, 0, slot)
		self.load(buffer[0x4000:], 1, slot)
		self.load(buffer[0x8000:], 2, slot)
		self.load(buffer[0xC000:], 3, slot)
	default:
		panic("ROM size not supported")
	}

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
	slot := self.ppi.pgSlots[page]

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
	slot := self.ppi.pgSlots[page]

	if self.mapper != nil && self.slotMapper == slot && (page == 1 || page == 2) {
		self.mapper.writeByte(address, value)
		return
	}

	if self.canWrite[page][slot] {
		delta := address - page*0x4000
		self.contents[page][slot][delta] = value
	}
}
