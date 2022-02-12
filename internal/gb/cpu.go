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
	c.reg.Reset()
}

const opCodeExt uint8 = 0xCB

var instruction = map[uint8]Instruction{
	0x31: LD_SP_nn,
}

var extendedInstruction = map[uint8]Instruction{}
