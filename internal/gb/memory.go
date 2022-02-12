package gb

import (
	"fmt"
	"os"
)

const romSize = 256

type Memory struct {
	rom [romSize]byte
}

func (m *Memory) Read(addr uint16) byte {
	// Only ROM for now
	return m.rom[addr]
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
