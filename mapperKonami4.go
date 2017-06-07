package main

type MapperKonami4 struct {
	contents []byte
	sel1     int
	sel2     int
	sel3     int
	sel4     int
}

func NewMapperKonami4() Mapper {
	m := new(MapperKonami4)
	m.sel1 = 0
	m.sel2 = 1
	m.sel3 = 2
	m.sel4 = 3
	return m
}

func (self *MapperKonami4) load(data []byte) {
	self.contents = make([]byte, len(data))
	copy(self.contents, data)
}

func (self *MapperKonami4) readByte(address uint16) byte {
	address -= 0x4000
	place := address / 0x2000
	var delta uint16
	var realMem []byte
	switch place {
	case 0:
		realMem = self.contents[self.sel1*0x2000:]
		delta = address
	case 1:
		realMem = self.contents[self.sel2*0x2000:]
		delta = address - 0x2000
	case 2:
		realMem = self.contents[self.sel3*0x2000:]
		delta = address - 0x4000
	case 3:
		realMem = self.contents[self.sel4*0x2000:]
		delta = address - 0x6000
	default:
		panic("Read mapper: impossible")
	}

	return realMem[delta]
}

func (self *MapperKonami4) writeByte(address uint16, value byte) {
	address -= 0x4000
	place := address / 0x2000
	switch place {
	case 0:
		return
	case 1:
		self.sel2 = int(value)
	case 2:
		self.sel3 = int(value)
	case 3:
		self.sel4 = int(value)
	default:
		panic("Write mapper: impossible")
	}
}
