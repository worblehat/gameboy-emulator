package gb

// Instruction is a single CPU instruction that cane access registers and memory.
// If an instruction needs an operand it reads it from the memory address pointed
// to by PC and increments PC afterwards.
type Instruction func(*Memory, *Registers)
