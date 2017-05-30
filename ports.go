package main

type Ports struct {
	a int
}

func (self *Ports) ReadPort(address uint16) byte {
	panic("ReadPort")
}

func (self *Ports) WritePort(address uint16, b byte) {
	ad := byte(address & 0xFF)
	switch {
	case ad >= 0xa8 && ad <= 0xab:
		panic("PPI")
	}

	panic("WritePort")
}

func (self *Ports) ReadPortInternal(address uint16, contend bool) byte {
	panic("ReadPortInternal")
}

func (self *Ports) WritePortInternal(address uint16, b byte, contend bool) {
	panic("WritePortInternal")
}

func (self *Ports) ContendPortPreio(address uint16) {
	panic("ContendPortPreio")

}

func (self *Ports) ContendPortPostio(address uint16) {
	panic("ContendPortPostio")
}
