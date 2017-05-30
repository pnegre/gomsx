package main

import "os"

type Memory struct {
	data [0x10000]byte
}

func NewMemory(romFile string) *Memory {
	mem := new(Memory)
	f, err := os.Open(romFile)
	if err != nil {
		panic(err)
	}

	f.Read(mem.data[:32768]) //32k
	f.Close()

	return mem
}

func (self *Memory) ReadByte(address uint16) byte {
	return self.ReadByteInternal(address)
}

// ReadByteInternal reads a byte from address without taking
// into account contention.
func (self *Memory) ReadByteInternal(address uint16) byte {
	return self.data[address]
}

// WriteByte writes a byte at address taking into account
// contention.
func (self *Memory) WriteByte(address uint16, value byte) {
	self.WriteByteInternal(address, value)
}

// WriteByteInternal writes a byte at address without taking
// into account contention.
func (self *Memory) WriteByteInternal(address uint16, value byte) {
	self.data[address] = value
}

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
