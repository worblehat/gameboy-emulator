package gb

import "fmt"

type CPU struct {
	mem Memory
	reg Registers
}

func NewCPU(romPath string) (*CPU, error) {

	rom, err := loadROM(romPath)
	if err != nil {
		return nil, err
	}

	cpu := &CPU{
		mem: Memory{rom: rom},
		reg: Registers{},
	}
	return cpu, nil
}

func (c *CPU) Run() {
	c.reset()

	for {
		opCode := c.mem.Read8(c.reg.PC)
		c.reg.PC += 1

		var instr Instruction
		if opCode == opCodeExt {
			opCode = c.mem.Read8(c.reg.PC + 1)
			c.reg.PC += 1
			instr = extendedInstruction[opCode]
		} else {
			instr = instruction[opCode]
		}

		if instr == nil {
			panic(fmt.Sprintf("Fetched unknown op code %X", opCode))
		}

		instr(&c.mem, &c.reg)
	}
}

func (c *CPU) reset() {
	c.reg.A = 0
	c.reg.B = 0
	c.reg.C = 0
	c.reg.D = 0
	c.reg.E = 0
	c.reg.F = 0
	c.reg.L = 0
	c.reg.H = 0
	c.reg.SP = 0
	c.reg.PC = 0
}

const opCodeExt uint8 = 0xCB

var instruction = map[uint8]Instruction{}

var extendedInstruction = map[uint8]Instruction{}
