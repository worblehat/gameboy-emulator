package gb

import (
	"fmt"
	"os"
)

const romSize = 256
const vramSize = 0x2000

type Memory struct {
	rom  [romSize]byte
	vram [vramSize]byte
}

func (m *Memory) Read8(addr uint16) uint8 {
	if addr >= 0x0000 && addr < 0x00FF {
		return m.rom[addr]
	} else if addr >= 0x8000 && addr < 0xA000 {
		return m.vram[addr-0x8000]
	}
	panic(fmt.Sprintf("Read from unknown memory address %X", addr))
}

func (m *Memory) Read16(addr uint16) uint16 {
	// Little endian
	loByte := uint16(m.Read8(addr))
	hiByte := uint16(m.Read8(addr + 1))
	return (hiByte << 8) | loByte
}

func (m *Memory) Write8(addr uint16, val uint8) {
	if addr >= 0x8000 && addr < 0xA000 {
		m.vram[addr-0x8000] = val
	} else {
		panic(fmt.Sprintf("Write to unknown memory address %X", addr))
	}
}

func (m *Memory) Write16(addr uint16, val uint16) {
	// Little endian
	loByte := uint8(val)
	hiByte := uint8(val >> 8)
	m.Write8(addr, loByte)
	m.Write8(addr+1, hiByte)
}

func loadROM(romPath string) ([romSize]byte, error) {
	var rom [romSize]byte

	content, err := os.ReadFile(romPath)
	if err != nil {
		return rom, err
	}

	if len(content) > len(rom) {
		return rom, fmt.Errorf(
			"provided ROM file of size %vB is too large for ROM", len(content))
	}

	copy(rom[:], content[:romSize])

	return rom, nil
}
