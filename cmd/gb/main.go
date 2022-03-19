package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/worblehat/Gameboy-Emulator/internal/gb"
)

func main() {
	bootROMPath := flag.String("boot-rom", "", "Path to a file with the boot ROM.")
	cartROMPath := flag.String("cartridge-rom", "", "Path to a file with a cartridge ROM.")
	withDebugger := flag.Bool("debug", false, "Enable debuger.")
	withTrace := flag.Bool("trace", false, "Print instructions on stdout as they are executed.")
	flag.Parse()

	if *bootROMPath == "" {
		fmt.Println("Error: No boot ROM file provided on command line")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *cartROMPath == "" {
		fmt.Println("Error: No cartridge ROM file provided on command line")
		flag.PrintDefaults()
		os.Exit(1)
	}

	bootROM, err := loadBootROM(*bootROMPath)
	if err != nil {
		fmt.Printf("Error: Could not load boot ROM from file %v (%v)\n", *bootROMPath, err)
		os.Exit(2)
	}

	cartROM0, err := loadCartROM0(*cartROMPath)
	if err != nil {
		fmt.Printf("Error: Could not load cartridge ROM from file %v (%v)\n", *cartROMPath, err)
		os.Exit(3)
	}

	memory := gb.NewMemory(bootROM, cartROM0)

	cpu := gb.NewCPU(memory)
	cpu.Run(*withDebugger, *withTrace)
}

func loadBootROM(romPath string) ([gb.BootROMSize]byte, error) {
	var rom [gb.BootROMSize]byte

	content, err := os.ReadFile(romPath)
	if err != nil {
		return rom, err
	}

	if len(content) > len(rom) {
		return rom, fmt.Errorf(
			"provided ROM file of size %vB is too large for boot ROM", len(content))
	}

	copy(rom[:], content[:gb.BootROMSize])

	return rom, nil
}

func loadCartROM0(romPath string) ([gb.CartROM0Size]byte, error) {
	var rom [gb.CartROM0Size]byte

	file, err := os.Open(romPath)
	if err != nil {
		return rom, err
	}
	defer file.Close()

	n, err := file.Read(rom[:])

	if err != nil {
		return rom, err
	}
	if n != len(rom) {
		return rom, errors.New("provided cartridge ROM file is to small")
	}

	return rom, nil
}
