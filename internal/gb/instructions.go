package gb

// Instruction is a single CPU instruction that cane access registers and memory.
// If an instruction needs an operand it reads it from the memory address pointed
// to by PC and increments PC afterwards.
type Instruction func(*Memory, *Registers)

// ###### 16-Bit Loads ######

// LD_BC_nn loads a 16 bit value into BC.
func LD_BC_nn(mem *Memory, reg *Registers) {
	reg.C = mem.Read8(reg.PC)
	reg.B = mem.Read8(reg.PC + 1)
	reg.PC += 2
}

// LD_DE_nn loads a 16 bit value into DE.
func LD_DE_nn(mem *Memory, reg *Registers) {
	reg.E = mem.Read8(reg.PC)
	reg.D = mem.Read8(reg.PC + 1)
	reg.PC += 2
}

// LD_HL_nn loads a 16 bit value into HL.
func LD_HL_nn(mem *Memory, reg *Registers) {
	reg.L = mem.Read8(reg.PC)
	reg.H = mem.Read8(reg.PC + 1)
	reg.PC += 2
}

// LD_SP_nn loads a 16 bit value into SP.
func LD_SP_nn(mem *Memory, reg *Registers) {
	reg.SP = mem.Read16(reg.PC)
	reg.PC += 2
}
