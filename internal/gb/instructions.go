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

// ###### Logical Operations ######

// XOR_A_A xors A with A and puts result into A.
func XOR_A_A(mem *Memory, reg *Registers) {
	//lint:ignore SA4000 does not make much sense but this is what the instruction does
	reg.A = reg.A ^ reg.A
	reg.SetFlags(subtractFlag|halfCarryFlag|carryFlag, false)
	reg.SetFlags(zeroFlag, reg.A == 0)
}

// XOR_A_B xors A with B and puts result into A.
func XOR_A_B(mem *Memory, reg *Registers) {
	reg.A = reg.A ^ reg.B
	reg.SetFlags(subtractFlag|halfCarryFlag|carryFlag, false)
	reg.SetFlags(zeroFlag, reg.A == 0)
}

// XOR_A_C xors A with C and puts result into A.
func XOR_A_C(mem *Memory, reg *Registers) {
	reg.A = reg.A ^ reg.C
	reg.SetFlags(subtractFlag|halfCarryFlag|carryFlag, false)
	reg.SetFlags(zeroFlag, reg.A == 0)
}

// XOR_A_D xors A with D and puts result into A.
func XOR_A_D(mem *Memory, reg *Registers) {
	reg.A = reg.A ^ reg.D
	reg.SetFlags(subtractFlag|halfCarryFlag|carryFlag, false)
	reg.SetFlags(zeroFlag, reg.A == 0)
}

// XOR_A_E xors A with E and puts result into A.
func XOR_A_E(mem *Memory, reg *Registers) {
	reg.A = reg.A ^ reg.E
	reg.SetFlags(subtractFlag|halfCarryFlag|carryFlag, false)
	reg.SetFlags(zeroFlag, reg.A == 0)
}

// XOR_A_H xors A with H and puts result into A.
func XOR_A_H(mem *Memory, reg *Registers) {
	reg.A = reg.A ^ reg.H
	reg.SetFlags(subtractFlag|halfCarryFlag|carryFlag, false)
	reg.SetFlags(zeroFlag, reg.A == 0)
}

// XOR_A_L xors A with L and puts result into A.
func XOR_A_L(mem *Memory, reg *Registers) {
	reg.A = reg.A ^ reg.L
	reg.SetFlags(subtractFlag|halfCarryFlag|carryFlag, false)
	reg.SetFlags(zeroFlag, reg.A == 0)
}

// XOR_A_HL xors A with the value pointed to by HL and puts result into A.
func XOR_A_HL(mem *Memory, reg *Registers) {
	addr := (uint16(reg.H) << 8) | uint16(reg.L)
	reg.A = reg.A ^ mem.Read8(addr)
	reg.SetFlags(subtractFlag|halfCarryFlag|carryFlag, false)
	reg.SetFlags(zeroFlag, reg.A == 0)
}

// XOR_A_n xors A with the value pointed to by PC and puts result into A.
func XOR_A_n(mem *Memory, reg *Registers) {
	reg.A = reg.A ^ mem.Read8(reg.PC)
	reg.PC += 1
	reg.SetFlags(subtractFlag|halfCarryFlag|carryFlag, false)
	reg.SetFlags(zeroFlag, reg.A == 0)
}
