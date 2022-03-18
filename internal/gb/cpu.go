package gb

import "fmt"

type CPU struct {
	mem *Memory
	reg Registers
}

func NewCPU(mem *Memory) *CPU {
	return &CPU{
		mem: mem,
		reg: Registers{},
	}
}

func (c *CPU) Run(withDebugger bool) {
	c.reset()

	dbg := NewDebugger(c.mem, &c.reg)
	dbg.Enabled = withDebugger

	for {
		dbg.Cycle()

		opCode := uint16(c.mem.Read8(c.reg.PC))
		instrAddr := c.reg.PC
		c.reg.PC += 1

		if opCode == opCodeExt {
			opCode = (opCode << 8) | uint16(c.mem.Read8(c.reg.PC))
			c.reg.PC += 1
		}

		instr, ok := instruction[opCode]
		if !ok {
			panic(fmt.Sprintf(
				"Fetched unknown op code 0x%X from address 0x%04X",
				opCode, instrAddr))
		}

		instr.Exec(c.mem, &c.reg)
		fmt.Printf("Executed 0x%04X [%v] at 0x%04X. Next instruction at 0x%04X\n",
			opCode, instr.Name, instrAddr, c.reg.PC)
	}
}

func (c *CPU) reset() {
	c.reg.Reset()
}
