package gb

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Debugger struct {
	mem        *Memory
	reg        *Registers
	breaks     map[uint]Breakpoint
	breakCount uint
	stepMode   bool
}

func NewDebugger(mem *Memory, reg *Registers) *Debugger {
	return &Debugger{
		mem:        mem,
		reg:        reg,
		breaks:     make(map[uint]Breakpoint),
		breakCount: 0,
		stepMode:   true,
	}
}

type Breakpoint struct {
	id      uint
	addr    uint16
	enabled bool
}

func (d *Debugger) Cycle() {
	if shouldBreak, bp := d.shouldBreakAt(d.reg.PC); shouldBreak {
		fmt.Printf("Breakpoint %v at 0x%04X\n", bp.id, d.reg.PC)
		d.processInput()
	} else if d.stepMode {
		d.processInput()
	}
}

var emptyPattern = regexp.MustCompile(`^\s*$`)
var continuePattern = regexp.MustCompile(`^(c|continue)$`)
var registersPattern = regexp.MustCompile(`^(i r|info registers)$`)
var memPattern = regexp.MustCompile(`^(i m|info mem) ([0-9a-fA-F]{1,4}) ([0-9a-fA-F]{1,4})$`)
var killPattern = regexp.MustCompile(`^(k|kill)$`)
var listBreaksPattern = regexp.MustCompile(`^(i b|info breakpoints)$`)
var addBreakPattern = regexp.MustCompile(`^(b|break) ([0-9a-fA-F]{1,4})$`)
var deleteBreakPattern = regexp.MustCompile(`^(d|delete) (\d+)$`)
var enableBreakPattern = regexp.MustCompile(`^(en|enable) (\d+)$`)
var disableBreakPattern = regexp.MustCompile(`^(dis|disable) (\d+)$`)
var disassemblePattern = regexp.MustCompile(`^(disas|disassemble)$`)
var stepPattern = regexp.MustCompile(`^(s|step)$`)

func (d *Debugger) processInput() {
	for {
		fmt.Printf("> ")
		cmd := d.readInput()
		cmd = strings.TrimSpace(cmd)
		if continuePattern.MatchString(cmd) {
			d.stepMode = false
			break
		} else if stepPattern.MatchString(cmd) {
			d.stepMode = true
			break
		} else if registersPattern.MatchString(cmd) {
			d.printRegisters()
		} else if matches := memPattern.FindStringSubmatch(cmd); len(matches) == 4 {
			start, _ := strconv.ParseUint(matches[2], 16, 16)
			size, _ := strconv.ParseUint(matches[3], 16, 16)
			cols := uint16(16)
			d.printMemory(uint16(start), uint16(size), cols)
		} else if listBreaksPattern.MatchString(cmd) {
			d.listBreakpoints()
		} else if matches := addBreakPattern.FindStringSubmatch(cmd); len(matches) == 3 {
			addr, _ := strconv.ParseUint(matches[2], 16, 16)
			d.addBreakpoint(uint16(addr))
		} else if matches := deleteBreakPattern.FindStringSubmatch(cmd); len(matches) == 3 {
			id, _ := strconv.ParseUint(matches[2], 10, 0)
			d.deleteBreakpoint(uint(id))
		} else if matches := enableBreakPattern.FindStringSubmatch(cmd); len(matches) == 3 {
			id, _ := strconv.ParseUint(matches[2], 10, 0)
			d.enableBreakpoint(uint(id), true)
		} else if matches := disableBreakPattern.FindStringSubmatch(cmd); len(matches) == 3 {
			id, _ := strconv.ParseUint(matches[2], 10, 0)
			d.enableBreakpoint(uint(id), false)
		} else if killPattern.MatchString(cmd) {
			fmt.Println("Exiting program...")
			os.Exit(0)
		} else if disassemblePattern.MatchString(cmd) {
			d.disassemble(d.reg.PC)
		} else if !emptyPattern.MatchString(cmd) {
			fmt.Printf("Unknown or invalid command: %v\n", cmd)
		}
	}
}

func (d *Debugger) shouldBreakAt(pc uint16) (bool, *Breakpoint) {
	for _, bp := range d.breaks {
		if bp.addr == pc && bp.enabled {
			return true, &bp
		}
	}
	return false, nil
}

func (d *Debugger) disassemble(addr uint16) {
	opCode := uint16(d.mem.Read8(addr))

	if opCode == opCodeExt {
		opCode = (opCode << 8) | uint16(d.mem.Read8(addr+1))
	}

	instr, ok := instruction[opCode]
	if !ok {
		fmt.Printf("Could not disassemble instruction at address 0x%04X\n", addr)
		return
	}

	fmt.Println("Address\tOpcode\tInstruction")
	fmt.Printf("0x%04X\t0x%04X\t%v\n", addr, opCode, instr.Name)
}

func (d *Debugger) listBreakpoints() {
	if len(d.breaks) > 0 {
		fmt.Println("Num\tEnb\tAddress")

		ids := make([]uint, 0, len(d.breaks))
		for id := range d.breaks {
			ids = append(ids, id)
		}
		sort.Slice(ids, func(i, j int) bool {
			return ids[i] < ids[j]
		})

		for _, id := range ids {
			enbStr := "y"
			bp := d.breaks[id]
			if !bp.enabled {
				enbStr = "n"
			}
			fmt.Printf("%v\t%v\t0x%04X\n", id, enbStr, bp.addr)
		}
	} else {
		fmt.Println("No Breakpoints")
	}
}

func (d *Debugger) addBreakpoint(addr uint16) {
	d.breakCount += 1
	id := d.breakCount
	d.breaks[id] = Breakpoint{
		id:      id,
		addr:    addr,
		enabled: true,
	}
}

func (d *Debugger) deleteBreakpoint(id uint) {
	delete(d.breaks, id)
}

func (d *Debugger) enableBreakpoint(id uint, enabled bool) {
	bp, ok := d.breaks[id]
	if !ok {
		fmt.Printf("Unknown breakpoint: %v\n", id)
		return
	}
	bp.enabled = enabled
	d.breaks[id] = bp
}

func (d *Debugger) printMemory(start, size, cols uint16) {
	// Accessing invalid memory might panic. Catch this here.
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\nError: %v\n", r)
		}
	}()

	if size == 0 {
		return
	}
	fmt.Printf("ADDR ")
	for col := uint16(0); col < cols; col += 1 {
		fmt.Printf("%02X ", col)
	}
	fmt.Printf("    \n")

	for i := uint16(0); i < size; i += 1 {
		if (i % cols) == 0 {
			fmt.Printf("%04X ", start+i)
		}
		fmt.Printf("%02X ", d.mem.Read8(start+i))
		if ((i + 1) % cols) == 0 {
			fmt.Printf("%04X\n", (i + 1 - cols))
		}
	}
	fmt.Printf("\n")
}

func (d *Debugger) printRegisters() {
	fmt.Printf("A: 0x%02X | F: 0x%02X\n", d.reg.A, d.reg.F)
	fmt.Printf("B: 0x%02X | C: 0x%02X\n", d.reg.B, d.reg.C)
	fmt.Printf("D: 0x%02X | E: 0x%02X\n", d.reg.D, d.reg.E)
	fmt.Printf("H: 0x%02X | L: 0x%02X\n", d.reg.H, d.reg.L)
	fmt.Printf("SP: 0x%04X\n", d.reg.SP)
	fmt.Printf("PC: 0x%04X\n", d.reg.PC)

	fmt.Printf("Z N H C\n")
	var z uint8 = 0
	if d.reg.IsFlagSet(zeroFlag) {
		z = 1
	}
	var n uint8 = 0
	if d.reg.IsFlagSet(subtractFlag) {
		n = 1
	}
	var h uint8 = 0
	if d.reg.IsFlagSet(halfCarryFlag) {
		h = 1
	}
	var c uint8 = 0
	if d.reg.IsFlagSet(carryFlag) {
		c = 1
	}
	fmt.Printf("%1b %1b %1b %1b\n", z, n, h, c)
}

func (d *Debugger) readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
