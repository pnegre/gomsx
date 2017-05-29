package main

type Memory struct {
    data [0x10000]byte
}


func (self Memory) ReadByte(address uint16) byte {
    return 0;
}

// ReadByteInternal reads a byte from address without taking
// into account contention.
func (self Memory) ReadByteInternal(address uint16) byte {
    return 0;
}

// WriteByte writes a byte at address taking into account
// contention.
func (self Memory) WriteByte(address uint16, value byte) {

}

// WriteByteInternal writes a byte at address without taking
// into account contention.
func (self Memory) WriteByteInternal(address uint16, value byte) {

}

// Follow contention methods. Leave unimplemented if you don't
// care about memory contention.

// ContendRead increments the Tstates counter by time as a
// result of a memory read at the given address.
func (self Memory) ContendRead(address uint16, time int) {

}

func (self Memory) ContendReadNoMreq(address uint16, time int) {

}
func (self Memory) ContendReadNoMreq_loop(address uint16, time int, count uint) {

}

func (self Memory) ContendWriteNoMreq(address uint16, time int) {

}
func (self Memory) ContendWriteNoMreq_loop(address uint16, time int, count uint) {

}

func (self Memory) Read(address uint16) byte {
    panic("Not implemented")
}

func (self Memory) Write(address uint16, value byte, protectROM bool) {
    panic("Not implemented")
}

// Data returns the memory content.
func (self Memory) Data() []byte {
    return nil;
}
