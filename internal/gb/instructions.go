package gb

// Instruction is a single CPU instruction that cane access registers and memory.
// If an instruction needs an operand it reads it from the memory address pointed
// to by PC and increments PC afterwards.
type Instruction func(*Memory, *Registers)

// ###### 16-Bit Loads ######

// Load 16 bit value into SP.
func LD_SP_nn(mem *Memory, reg *Registers) {
	nn := mem.Read16(reg.PC)
	reg.SP = nn
	reg.PC += 2
}
