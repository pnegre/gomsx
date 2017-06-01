package main

// Mètodes no implementats //////////////////////////
/////////////////////////////////////////////////////
/////////////////////////////////////////////////////
// Follow contention methods. Leave unimplemented if you don't
// care about memory contention.

// ContendRead increments the Tstates counter by time as a
// result of a memory read at the given address.
func (self *Memory) ContendRead(address uint16, time int) {
	//panic("ContendRead not implemented")
}

func (self *Memory) ContendReadNoMreq(address uint16, time int) {
	//panic("ContendReadNoMreq not implemented")
}
func (self *Memory) ContendReadNoMreq_loop(address uint16, time int, count uint) {
	///panic("ContendReadNoMreq_loop not implemented")
}

func (self *Memory) ContendWriteNoMreq(address uint16, time int) {
	//panic("ContendWriteNoMreq not implemented")
}
func (self *Memory) ContendWriteNoMreq_loop(address uint16, time int, count uint) {
	//panic("ContendWriteNoMreq_loop not implemented")
}

func (self *Memory) Read(address uint16) byte {
	panic("Memory.Read Not implemented")
}

func (self *Memory) Write(address uint16, value byte, protectROM bool) {
	panic("Memory.Write Not implemented")
}

// Data returns the memory content.
func (self *Memory) Data() []byte {
	return nil
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
