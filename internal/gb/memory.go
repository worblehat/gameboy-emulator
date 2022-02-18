package gb

import (
	"fmt"
	"os"
)

const bootROMSize = 256
const vramSize = 0x2000
const ioMemSize = 0x80

type Memory struct {
	bootROM [bootROMSize]byte
	vram    [vramSize]byte
	ioMem   [ioMemSize]byte
}

func (m *Memory) Read8(addr uint16) uint8 {
	if addr >= 0x0000 && addr < 0x00FF {
		return m.bootROM[addr]
	} else if addr >= 0x8000 && addr < 0xA000 {
		return m.vram[addr-0x8000]
	} else if addr >= 0xFF00 && addr < 0xFF80 {
		return m.ioMem[addr-0xFF00]
	}
	panic(fmt.Sprintf("Read from unknown memory address 0x%X", addr))
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
	} else if addr >= 0xFF00 && addr < 0xFF80 {
		m.ioMem[addr-0xFF00] = val
	} else {
		panic(fmt.Sprintf("Write to unknown memory address 0x%X", addr))
	}
}

func (m *Memory) Write16(addr uint16, val uint16) {
	// Little endian
	loByte := uint8(val)
	hiByte := uint8(val >> 8)
	m.Write8(addr, loByte)
	m.Write8(addr+1, hiByte)
}

func loadBootROM(romPath string) ([bootROMSize]byte, error) {
	var rom [bootROMSize]byte

	content, err := os.ReadFile(romPath)
	if err != nil {
		return rom, err
	}

	if len(content) > len(rom) {
		return rom, fmt.Errorf(
			"provided ROM file of size %vB is too large for ROM", len(content))
	}

	copy(rom[:], content[:bootROMSize])

	return rom, nil
}
