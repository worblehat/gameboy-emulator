package gb

type CPU struct {
	mem Memory
}

func NewCPU(romPath string) (*CPU, error) {

	rom, err := loadROM(romPath)
	if err != nil {
		return nil, err
	}

	cpu := &CPU{
		mem: Memory{rom: rom},
	}
	return cpu, nil
}

func (c *CPU) Run() {
}
