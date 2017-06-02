package main

import "os"

type Memory struct {
	data [0x10000]byte
}

func NewMemory() *Memory {
	mem := new(Memory)
	return mem
}

func (self *Memory) loadFromFile(romFile string) error {
	f, err := os.Open(romFile)
	if err != nil {
		return err
	}

	f.Read(self.data[:32768]) //32k
	f.Close()
	return nil
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
