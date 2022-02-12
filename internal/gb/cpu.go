package gb

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
