package main

import "testing"
import "math/rand"

func TestMem001(t *testing.T) {
	r := rand.New(rand.NewSource(99))
	memory := NewMemory()
	ppi_slots = 0x14
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

func TestMem002(t *testing.T) {
	memory := NewMemory()
	loadBiosBasic(memory, ROMFILE)
	loadRom(memory, "games/loderunner.rom")
	ppi_slots = 0x14

	check := func(address uint16, val byte) {
		if memory.ReadByte(address) != val {
			t.Errorf("NORL: %02x\n", memory.ReadByte(address))
		}
	}

	check(0x3ffe, 0x09)
	check(0x3fff, 0x18)
	check(0x4000, 0x41)
	check(0x4001, 0x42)
	check(0x7ffe, 0x22)
	check(0x7fff, 0x22)
	check(0x8000, 0x52)
	check(0x8001, 0x00)
	check(0x8002, 0x05)
	check(0x8003, 0x00)
}
