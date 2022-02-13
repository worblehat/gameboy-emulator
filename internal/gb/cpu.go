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
		opCode := uint16(c.mem.Read8(c.reg.PC))
		c.reg.PC += 1

		var instr Instruction
		if opCode == opCodeExt {
			opCode = (opCode << 8) | uint16(c.mem.Read8(c.reg.PC))
			c.reg.PC += 1
		}

		instr = instruction[opCode]
		if instr == nil {
			panic(fmt.Sprintf("Fetched unknown op code %X", opCode))
		}

		instr(&c.mem, &c.reg)
	}
}

func (c *CPU) reset() {
	c.reg.Reset()
}

const opCodeExt uint16 = 0xCB

var instruction = map[uint16]Instruction{
	0x01: LD_BC_nn,
	0x11: LD_DE_nn,
	0x21: LD_HL_nn,
	0x31: LD_SP_nn,
	0x32: LDD_HL_A,
	0xAF: XOR_A_A,
	0xA8: XOR_A_B,
	0xA9: XOR_A_C,
	0xAA: XOR_A_D,
	0xAB: XOR_A_E,
	0xAC: XOR_A_H,
	0xAD: XOR_A_L,
	0xAE: XOR_A_HL,
	0xEE: XOR_A_n,
}
