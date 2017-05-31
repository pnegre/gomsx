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
