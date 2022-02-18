package gb

import (
	"fmt"
	"math"
)

// Instruction is a single CPU instruction that cane access registers and memory.
// If an instruction needs an operand it reads it from the memory address pointed
// to by PC and increments PC afterwards.
type Instruction func(*Memory, *Registers)

// ###### 8-Bit Loads ######

// LDD_HL_A loads the value of A to address HL and decrements HL.
func LDD_HL_A(mem *Memory, reg *Registers) {
	addr := reg.HL()
	mem.Write8(addr, reg.A)
	addr -= 1
	reg.H = uint8(addr >> 8)
	reg.L = uint8(addr)
}

// ###### 16-Bit Loads ######

// LD_BC_nn loads a 16 bit immediate value into BC.
func LD_BC_nn(mem *Memory, reg *Registers) {
	reg.C = mem.Read8(reg.PC)
	reg.B = mem.Read8(reg.PC + 1)
	reg.PC += 2
}

// LD_DE_nn loads a 16 bit immediate value into DE.
func LD_DE_nn(mem *Memory, reg *Registers) {
	reg.E = mem.Read8(reg.PC)
	reg.D = mem.Read8(reg.PC + 1)
	reg.PC += 2
}

// LD_HL_nn loads a 16 bit immediate value into HL.
func LD_HL_nn(mem *Memory, reg *Registers) {
	reg.L = mem.Read8(reg.PC)
	reg.H = mem.Read8(reg.PC + 1)
	reg.PC += 2
}

// LD_SP_nn loads a 16 bit immediate value into SP.
func LD_SP_nn(mem *Memory, reg *Registers) {
	reg.SP = mem.Read16(reg.PC)
	reg.PC += 2
}

// ###### Logical Operations ######

// XOR_A_A xors A with A and puts result into A.
func XOR_A_A(mem *Memory, reg *Registers) {
	xorA(reg.A, reg)
}

// XOR_A_B xors A with B and puts result into A.
func XOR_A_B(mem *Memory, reg *Registers) {
	xorA(reg.B, reg)
}

// XOR_A_C xors A with C and puts result into A.
func XOR_A_C(mem *Memory, reg *Registers) {
	xorA(reg.C, reg)
}

// XOR_A_D xors A with D and puts result into A.
func XOR_A_D(mem *Memory, reg *Registers) {
	xorA(reg.D, reg)
}

// XOR_A_E xors A with E and puts result into A.
func XOR_A_E(mem *Memory, reg *Registers) {
	xorA(reg.E, reg)
}

// XOR_A_H xors A with H and puts result into A.
func XOR_A_H(mem *Memory, reg *Registers) {
	xorA(reg.H, reg)
}

// XOR_A_L xors A with L and puts result into A.
func XOR_A_L(mem *Memory, reg *Registers) {
	xorA(reg.L, reg)
}

// XOR_A_HL xors A with the value pointed by HL and puts result into A.
func XOR_A_HL(mem *Memory, reg *Registers) {
	addr := (uint16(reg.H) << 8) | uint16(reg.L)
	xorA(mem.Read8(addr), reg)
}

// XOR_A_n xors A with the 8 bit immediate value and puts result into A.
func XOR_A_n(mem *Memory, reg *Registers) {
	n := mem.Read8(reg.PC)
	reg.PC += 1
	xorA(n, reg)
}

// ###### Single-Bit Operations ######

// BIT_0_A tests bit 0 in register A.
func BIT_0_A(mem *Memory, reg *Registers) {
	testBit(reg.A, 0, reg)
}

// BIT_1_A tests bit 1 in register A.
func BIT_1_A(mem *Memory, reg *Registers) {
	testBit(reg.A, 1, reg)
}

// BIT_2_A tests bit 2 in register A.
func BIT_2_A(mem *Memory, reg *Registers) {
	testBit(reg.A, 2, reg)
}

// BIT_3_A tests bit 3 in register A.
func BIT_3_A(mem *Memory, reg *Registers) {
	testBit(reg.A, 3, reg)
}

// BIT_4_A tests bit 4 in register A.
func BIT_4_A(mem *Memory, reg *Registers) {
	testBit(reg.A, 4, reg)
}

// BIT_5_A tests bit 5 in register A.
func BIT_5_A(mem *Memory, reg *Registers) {
	testBit(reg.A, 5, reg)
}

// BIT_6_A tests bit 6 in register A.
func BIT_6_A(mem *Memory, reg *Registers) {
	testBit(reg.A, 6, reg)
}

// BIT_7_A tests bit 7 in register A.
func BIT_7_A(mem *Memory, reg *Registers) {
	testBit(reg.A, 7, reg)
}

// BIT_0_B tests bit 0 in register B.
func BIT_0_B(mem *Memory, reg *Registers) {
	testBit(reg.B, 0, reg)
}

// BIT_1_B tests bit 1 in register B.
func BIT_1_B(mem *Memory, reg *Registers) {
	testBit(reg.B, 1, reg)
}

// BIT_2_B tests bit 2 in register B.
func BIT_2_B(mem *Memory, reg *Registers) {
	testBit(reg.B, 2, reg)
}

// BIT_3_B tests bit 3 in register B.
func BIT_3_B(mem *Memory, reg *Registers) {
	testBit(reg.B, 3, reg)
}

// BIT_4_B tests bit 4 in register B.
func BIT_4_B(mem *Memory, reg *Registers) {
	testBit(reg.B, 4, reg)
}

// BIT_5_B tests bit 5 in register B.
func BIT_5_B(mem *Memory, reg *Registers) {
	testBit(reg.B, 5, reg)
}

// BIT_6_B tests bit 6 in register B.
func BIT_6_B(mem *Memory, reg *Registers) {
	testBit(reg.B, 6, reg)
}

// BIT_7_B tests bit 7 in register B.
func BIT_7_B(mem *Memory, reg *Registers) {
	testBit(reg.B, 7, reg)
}

// BIT_0_C tests bit 0 in register C.
func BIT_0_C(mem *Memory, reg *Registers) {
	testBit(reg.C, 0, reg)
}

// BIT_1_C tests bit 1 in register C.
func BIT_1_C(mem *Memory, reg *Registers) {
	testBit(reg.C, 1, reg)
}

// BIT_2_C tests bit 2 in register C.
func BIT_2_C(mem *Memory, reg *Registers) {
	testBit(reg.C, 2, reg)
}

// BIT_3_C tests bit 3 in register C.
func BIT_3_C(mem *Memory, reg *Registers) {
	testBit(reg.C, 3, reg)
}

// BIT_4_C tests bit 4 in register C.
func BIT_4_C(mem *Memory, reg *Registers) {
	testBit(reg.C, 4, reg)
}

// BIT_5_C tests bit 5 in register C.
func BIT_5_C(mem *Memory, reg *Registers) {
	testBit(reg.C, 5, reg)
}

// BIT_6_C tests bit 6 in register C.
func BIT_6_C(mem *Memory, reg *Registers) {
	testBit(reg.C, 6, reg)
}

// BIT_7_C tests bit 7 in register C.
func BIT_7_C(mem *Memory, reg *Registers) {
	testBit(reg.C, 7, reg)
}

// BIT_0_D tests bit 0 in register D.
func BIT_0_D(mem *Memory, reg *Registers) {
	testBit(reg.D, 0, reg)
}

// BIT_1_D tests bit 1 in register D.
func BIT_1_D(mem *Memory, reg *Registers) {
	testBit(reg.D, 1, reg)
}

// BIT_2_D tests bit 2 in register D.
func BIT_2_D(mem *Memory, reg *Registers) {
	testBit(reg.D, 2, reg)
}

// BIT_3_D tests bit 3 in register D.
func BIT_3_D(mem *Memory, reg *Registers) {
	testBit(reg.D, 3, reg)
}

// BIT_4_D tests bit 4 in register D.
func BIT_4_D(mem *Memory, reg *Registers) {
	testBit(reg.D, 4, reg)
}

// BIT_5_D tests bit 5 in register D.
func BIT_5_D(mem *Memory, reg *Registers) {
	testBit(reg.D, 5, reg)
}

// BIT_6_D tests bit 6 in register D.
func BIT_6_D(mem *Memory, reg *Registers) {
	testBit(reg.D, 6, reg)
}

// BIT_7_D tests bit 7 in register D.
func BIT_7_D(mem *Memory, reg *Registers) {
	testBit(reg.D, 7, reg)
}

// BIT_0_E tests bit 0 in register E.
func BIT_0_E(mem *Memory, reg *Registers) {
	testBit(reg.E, 0, reg)
}

// BIT_1_E tests bit 1 in register E.
func BIT_1_E(mem *Memory, reg *Registers) {
	testBit(reg.E, 1, reg)
}

// BIT_2_E tests bit 2 in register E.
func BIT_2_E(mem *Memory, reg *Registers) {
	testBit(reg.E, 2, reg)
}

// BIT_3_E tests bit 3 in register E.
func BIT_3_E(mem *Memory, reg *Registers) {
	testBit(reg.E, 3, reg)
}

// BIT_4_E tests bit 4 in register E.
func BIT_4_E(mem *Memory, reg *Registers) {
	testBit(reg.E, 4, reg)
}

// BIT_5_E tests bit 5 in register E.
func BIT_5_E(mem *Memory, reg *Registers) {
	testBit(reg.E, 5, reg)
}

// BIT_6_E tests bit 6 in register E.
func BIT_6_E(mem *Memory, reg *Registers) {
	testBit(reg.E, 6, reg)
}

// BIT_7_E tests bit 7 in register E.
func BIT_7_E(mem *Memory, reg *Registers) {
	testBit(reg.E, 7, reg)
}

// BIT_0_H tests bit 0 in register H.
func BIT_0_H(mem *Memory, reg *Registers) {
	testBit(reg.H, 0, reg)
}

// BIT_1_H tests bit 1 in register H.
func BIT_1_H(mem *Memory, reg *Registers) {
	testBit(reg.H, 1, reg)
}

// BIT_2_H tests bit 2 in register H.
func BIT_2_H(mem *Memory, reg *Registers) {
	testBit(reg.H, 2, reg)
}

// BIT_3_H tests bit 3 in register H.
func BIT_3_H(mem *Memory, reg *Registers) {
	testBit(reg.H, 3, reg)
}

// BIT_4_H tests bit 4 in register H.
func BIT_4_H(mem *Memory, reg *Registers) {
	testBit(reg.H, 4, reg)
}

// BIT_5_H tests bit 5 in register H.
func BIT_5_H(mem *Memory, reg *Registers) {
	testBit(reg.H, 5, reg)
}

// BIT_6_H tests bit 6 in register H.
func BIT_6_H(mem *Memory, reg *Registers) {
	testBit(reg.H, 6, reg)
}

// BIT_7_H tests bit 7 in register H.
func BIT_7_H(mem *Memory, reg *Registers) {
	testBit(reg.H, 7, reg)
}

// BIT_0_L tests bit 0 in register L.
func BIT_0_L(mem *Memory, reg *Registers) {
	testBit(reg.L, 0, reg)
}

// BIT_1_L tests bit 1 in register L.
func BIT_1_L(mem *Memory, reg *Registers) {
	testBit(reg.L, 1, reg)
}

// BIT_2_L tests bit 2 in register L.
func BIT_2_L(mem *Memory, reg *Registers) {
	testBit(reg.L, 2, reg)
}

// BIT_3_L tests bit 3 in register L.
func BIT_3_L(mem *Memory, reg *Registers) {
	testBit(reg.L, 3, reg)
}

// BIT_4_L tests bit 4 in register L.
func BIT_4_L(mem *Memory, reg *Registers) {
	testBit(reg.L, 4, reg)
}

// BIT_5_L tests bit 5 in register L.
func BIT_5_L(mem *Memory, reg *Registers) {
	testBit(reg.L, 5, reg)
}

// BIT_6_L tests bit 6 in register L.
func BIT_6_L(mem *Memory, reg *Registers) {
	testBit(reg.L, 6, reg)
}

// BIT_7_L tests bit 7 in register L.
func BIT_7_L(mem *Memory, reg *Registers) {
	testBit(reg.L, 7, reg)
}

// BIT_0_HL tests bit 0 of the value pointed by HL.
func BIT_0_HL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	testBit(mem.Read8(addr), 0, reg)
}

// BIT_1_HL tests bit 1 of the value pointed by HL.
func BIT_1_HL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	testBit(mem.Read8(addr), 1, reg)
}

// BIT_2_HL tests bit 2 of the value pointed by HL.
func BIT_2_HL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	testBit(mem.Read8(addr), 2, reg)
}

// BIT_3_HL tests bit 3 of the value pointed by HL.
func BIT_3_HL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	testBit(mem.Read8(addr), 3, reg)
}

// BIT_4_HL tests bit 4 of the value pointed by HL.
func BIT_4_HL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	testBit(mem.Read8(addr), 4, reg)
}

// BIT_5_HL tests bit 5 of the value pointed by HL.
func BIT_5_HL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	testBit(mem.Read8(addr), 5, reg)
}

// BIT_6_HL tests bit 6 of the value pointed by HL.
func BIT_6_HL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	testBit(mem.Read8(addr), 6, reg)
}

// BIT_7_HL tests bit 7 of the value pointed by HL.
func BIT_7_HL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	testBit(mem.Read8(addr), 7, reg)
}

// ###### Relative Jumps ######

// If the zero flag is not set JR_NZ_n adds an 8 bit signed immediate
// value to current PC and jumps to it.
func JR_NZ_n(mem *Memory, reg *Registers) {
	if !reg.IsFlagSet(zeroFlag) {
		relJumpByImmediateValue(mem, reg)
	}
}

// If the zero flag is set JR_Z_n adds an 8 bit signed immediate
// value to current PC and jumps to it.
func JR_Z_n(mem *Memory, reg *Registers) {
	if reg.IsFlagSet(zeroFlag) {
		relJumpByImmediateValue(mem, reg)
	}
}

// If the carry flag is not set JR_NC_n adds an 8 bit signed immediate
// value to current PC and jumps to it.
func JR_NC_n(mem *Memory, reg *Registers) {
	if !reg.IsFlagSet(carryFlag) {
		relJumpByImmediateValue(mem, reg)
	}
}

// If the carry flag is set JR_C_n adds an 8 bit signed immediate
// value to current PC and jumps to it.
func JR_C_n(mem *Memory, reg *Registers) {
	if reg.IsFlagSet(carryFlag) {
		relJumpByImmediateValue(mem, reg)
	}
}

// ###### Common functions used by the instructions ######

// xorA xors register A with the given value and puts the result into A.
func xorA(value uint8, reg *Registers) {
	reg.A = reg.A ^ value
	reg.SetFlags(subtractFlag|halfCarryFlag|carryFlag, false)
	reg.SetFlags(zeroFlag, reg.A == 0)
}

// testBit checks if the given bit is set in the given value.
func testBit(value uint8, bit uint8, reg *Registers) {
	isSet := (value & (1 << bit)) != 0
	reg.SetFlags(zeroFlag, !isSet)
	reg.SetFlags(halfCarryFlag, true)
	reg.SetFlags(subtractFlag, false)
}

// relJump performs a relative jump by adding an 8 bit signed immediate value
// to the current PC.
func relJumpByImmediateValue(mem *Memory, reg *Registers) {
	n := int8(mem.Read8(reg.PC))
	reg.PC += 1
	relJump(n, reg)
}

// relJump performs a relative jump by adding n to the curren PC.
func relJump(n int8, reg *Registers) {
	newpc := int32(reg.PC) + int32(n)
	if newpc < 0 || math.MaxUint16 < newpc {
		panic(fmt.Sprintf(
			"invalid relative jump (current PC: 0x%x, jump: %v)",
			reg.PC, n))
	}
	reg.PC += uint16(newpc)
}
