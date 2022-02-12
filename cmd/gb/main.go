package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/worblehat/Gameboy-Emulator/internal/gb"
)

func main() {
	romPath := flag.String("rom", "", "Path to a file with ROM content")
	flag.Parse()

	if *romPath == "" {
		fmt.Println("Error: No rom file provided on command line")
		flag.PrintDefaults()
		os.Exit(1)
	}

	cpu, err := gb.NewCPU(*romPath)
	if err != nil {
		fmt.Printf("Error: Could not initialize CPU (%v)\n", err)
		os.Exit(2)
	}
	cpu.Run()
}
