package main

import "testing"
import "math/rand"

func TestMem001(t *testing.T) {
	r := rand.New(rand.NewSource(99))
	memory := NewMemory()
	ppi_slots = 0x41
	data := make([]byte, 0x10000)
	for i := 0; i < len(data); i++ {
		data[i] = byte(r.Uint32() % 255)
		memory.WriteByte(uint16(i), data[i])
	}
	for i := 0; i < len(data); i++ {
		if memory.ReadByte(uint16(i)) != data[i] {
			t.Errorf("Error reading memory!")
		}
	}
}
