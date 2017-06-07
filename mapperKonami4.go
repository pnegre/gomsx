package main

type MapperKonami4 struct {
	contents []byte
	sels     [4]int
}

func NewMapperKonami4() Mapper {
	m := new(MapperKonami4)
	for i := 0; i < 4; i++ {
		m.sels[i] = i
	}
	return m
}

func (self *MapperKonami4) load(data []byte) {
	self.contents = make([]byte, len(data))
	copy(self.contents, data)
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
