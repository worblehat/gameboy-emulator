package gb

import "fmt"

type CPU struct {
	mem Memory
	reg Registers
}

func NewCPU(bootROMPath string) (*CPU, error) {

	bootROM, err := loadBootROM(bootROMPath)
	if err != nil {
		return nil, err
	}

	cpu := &CPU{
		mem: Memory{bootROM: bootROM},
		reg: Registers{},
	}
	return cpu, nil
}

func (c *CPU) Run() {
	c.reset()

	for {
		opCode := uint16(c.mem.Read8(c.reg.PC))
		instrAddr := c.reg.PC
		c.reg.PC += 1

		var instr Instruction
		if opCode == opCodeExt {
			opCode = (opCode << 8) | uint16(c.mem.Read8(c.reg.PC))
			c.reg.PC += 1
		}

		instr = instruction[opCode]
		if instr == nil {
			panic(fmt.Sprintf(
				"Fetched unknown op code 0x%X from address 0x%x",
				opCode, instrAddr))
		}

		instr(&c.mem, &c.reg)
	}
}

func (c *CPU) reset() {
	c.reg.Reset()
}

const opCodeExt uint16 = 0xCB

var instruction = map[uint16]Instruction{
	0x01:   LD_BC_nn,
	0x02:   LD_pBC_A,
	0x04:   INC_B,
	0x0C:   INC_C,
	0x0A:   LD_A_pBC,
	0x11:   LD_DE_nn,
	0x12:   LD_pDE_A,
	0x14:   INC_D,
	0x1A:   LD_A_pDE,
	0x1C:   INC_E,
	0x20:   JR_NZ_n,
	0x21:   LD_HL_nn,
	0x24:   INC_H,
	0x28:   JR_Z_n,
	0x2C:   INC_L,
	0x30:   JR_NC_n,
	0x31:   LD_SP_nn,
	0x32:   LDD_pHL_A,
	0x34:   INC_pHL,
	0x38:   JR_C_n,
	0x3C:   INC_A,
	0x3E:   LD_A_n,
	0x47:   LD_B_A,
	0x4F:   LD_C_A,
	0x57:   LD_D_A,
	0x5F:   LD_E_A,
	0x67:   LD_H_A,
	0x6F:   LD_L_A,
	0x77:   LD_pHL_A,
	0x78:   LD_A_B,
	0x79:   LD_A_C,
	0x7A:   LD_A_D,
	0x7B:   LD_A_E,
	0x7C:   LD_A_H,
	0x7D:   LD_A_L,
	0x7E:   LD_A_pHL,
	0x7F:   LD_A_A,
	0xAF:   XOR_A_A,
	0xA8:   XOR_A_B,
	0xA9:   XOR_A_C,
	0xAA:   XOR_A_D,
	0xAB:   XOR_A_E,
	0xAC:   XOR_A_H,
	0xAD:   XOR_A_L,
	0xAE:   XOR_A_pHL,
	0xEA:   LD_pnn_A,
	0xE0:   LDD_IO_n_A,
	0xE2:   LDD_IO_C_A,
	0xEE:   XOR_A_n,
	0xF0:   LDD_A_IO_n,
	0xFA:   LD_A_pnn,
	0xF2:   LDD_A_IO_C,
	0xCB40: BIT_0_B,
	0xCB41: BIT_0_C,
	0xCB42: BIT_0_D,
	0xCB43: BIT_0_E,
	0xCB44: BIT_0_H,
	0xCB45: BIT_0_L,
	0xCB46: BIT_0_pHL,
	0xCB47: BIT_0_A,
	0xCB48: BIT_1_B,
	0xCB49: BIT_1_C,
	0xCB4A: BIT_1_D,
	0xCB4B: BIT_1_E,
	0xCB4C: BIT_1_H,
	0xCB4D: BIT_1_L,
	0xCB4E: BIT_1_pHL,
	0xCB4F: BIT_1_A,
	0xCB50: BIT_2_B,
	0xCB51: BIT_2_C,
	0xCB52: BIT_2_D,
	0xCB53: BIT_2_E,
	0xCB54: BIT_2_H,
	0xCB55: BIT_2_L,
	0xCB56: BIT_2_pHL,
	0xCB57: BIT_2_A,
	0xCB58: BIT_3_B,
	0xCB59: BIT_3_C,
	0xCB5A: BIT_3_D,
	0xCB5B: BIT_3_E,
	0xCB5C: BIT_3_H,
	0xCB5D: BIT_3_L,
	0xCB5E: BIT_3_pHL,
	0xCB5F: BIT_3_A,
	0xCB60: BIT_4_B,
	0xCB61: BIT_4_C,
	0xCB62: BIT_4_D,
	0xCB63: BIT_4_E,
	0xCB64: BIT_4_H,
	0xCB65: BIT_4_L,
	0xCB66: BIT_4_pHL,
	0xCB67: BIT_4_A,
	0xCB68: BIT_5_B,
	0xCB69: BIT_5_C,
	0xCB6A: BIT_5_D,
	0xCB6B: BIT_5_E,
	0xCB6C: BIT_5_H,
	0xCB6D: BIT_5_L,
	0xCB6E: BIT_5_pHL,
	0xCB6F: BIT_5_A,
	0xCB70: BIT_6_B,
	0xCB71: BIT_6_C,
	0xCB72: BIT_6_D,
	0xCB73: BIT_6_E,
	0xCB74: BIT_6_H,
	0xCB75: BIT_6_L,
	0xCB76: BIT_6_pHL,
	0xCB77: BIT_6_A,
	0xCB78: BIT_7_B,
	0xCB79: BIT_7_C,
	0xCB7A: BIT_7_D,
	0xCB7B: BIT_7_E,
	0xCB7C: BIT_7_H,
	0xCB7D: BIT_7_L,
	0xCB7E: BIT_7_pHL,
	0xCB7F: BIT_7_A,
}
