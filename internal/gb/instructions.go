package gb

import (
	"fmt"
	"math"
)

type Instruction struct {
	Name string
	// Exec executes a CPU instruction that can access registers and memory.
	// If an instruction needs an operand it reads it from the memory address pointed
	// to by PC and increments PC afterwards.
	Exec func(mem *Memory, reg *Registers)
}

// ###### 8-Bit Loads ######

// LD_B_n loads an 8-bit immediate value into B.
func LD_B_n(mem *Memory, reg *Registers) {
	reg.B = mem.Read8(reg.PC)
	reg.PC += 1
}

// LD_C_n loads an 8-bit immediate value into C.
func LD_C_n(mem *Memory, reg *Registers) {
	reg.C = mem.Read8(reg.PC)
	reg.PC += 1
}

// LD_D_n loads an 8-bit immediate value into D.
func LD_D_n(mem *Memory, reg *Registers) {
	reg.D = mem.Read8(reg.PC)
	reg.PC += 1
}

// LD_E_n loads an 8-bit immediate value into E.
func LD_E_n(mem *Memory, reg *Registers) {
	reg.E = mem.Read8(reg.PC)
	reg.PC += 1
}

// LD_H_n loads an 8-bit immediate value into H.
func LD_H_n(mem *Memory, reg *Registers) {
	reg.H = mem.Read8(reg.PC)
	reg.PC += 1
}

// LD_L_n loads an 8-bit immediate value into L.
func LD_L_n(mem *Memory, reg *Registers) {
	reg.L = mem.Read8(reg.PC)
	reg.PC += 1
}

// LD_A_A loads the value of A into A.
func LD_A_A(mem *Memory, reg *Registers) {
	//reg.A = reg.A
}

// LD_A_B loads the value of B into A.
func LD_A_B(mem *Memory, reg *Registers) {
	reg.A = reg.B
}

// LD_A_C loads the value of C into A.
func LD_A_C(mem *Memory, reg *Registers) {
	reg.A = reg.C
}

// LD_A_D loads the value of D into A.
func LD_A_D(mem *Memory, reg *Registers) {
	reg.A = reg.D
}

// LD_A_E loads the value of E into A.
func LD_A_E(mem *Memory, reg *Registers) {
	reg.A = reg.E
}

// LD_A_H loads the value of H into A.
func LD_A_H(mem *Memory, reg *Registers) {
	reg.A = reg.H
}

// LD_A_L loads the value of L into A.
func LD_A_L(mem *Memory, reg *Registers) {
	reg.A = reg.L
}

// LD_A_pBC loads the value pointed by BC into A.
func LD_A_pBC(mem *Memory, reg *Registers) {
	addr := reg.BC()
	reg.A = mem.Read8(addr)
}

// LD_A_pDE loads the value pointed by DE into A.
func LD_A_pDE(mem *Memory, reg *Registers) {
	addr := reg.DE()
	reg.A = mem.Read8(addr)
}

// LD_A_pHL loads the value pointed by HL into A.
func LD_A_pHL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	reg.A = mem.Read8(addr)
}

// LD_A_pnn loads the value pointed by the 16-bit immedite value into A.
func LD_A_pnn(mem *Memory, reg *Registers) {
	addr := mem.Read16(reg.PC)
	reg.PC += 2
	reg.A = mem.Read8(addr)
}

// LD_A_n loads an 8-bit immediate value into A.
func LD_A_n(mem *Memory, reg *Registers) {
	reg.A = mem.Read8(reg.PC)
	reg.PC += 1
}

// LD_B_A loads the value of A into B.
func LD_B_A(mem *Memory, reg *Registers) {
	reg.B = reg.A
}

// LD_C_A loads the value of A into C.
func LD_C_A(mem *Memory, reg *Registers) {
	reg.C = reg.A
}

// LD_D_A loads the value of A into D.
func LD_D_A(mem *Memory, reg *Registers) {
	reg.D = reg.A
}

// LD_E_A loads the value of A into E.
func LD_E_A(mem *Memory, reg *Registers) {
	reg.E = reg.A
}

// LD_H_A loads the value of A into H.
func LD_H_A(mem *Memory, reg *Registers) {
	reg.H = reg.A
}

// LD_L_A loads the value of A into L.
func LD_L_A(mem *Memory, reg *Registers) {
	reg.L = reg.A
}

// LD_pBC_ loads the value of A into the address pointed by BC.
func LD_pBC_A(mem *Memory, reg *Registers) {
	addr := reg.BC()
	mem.Write8(addr, reg.A)
}

// LD_pDE_ loads the value of A into the address pointed by DE.
func LD_pDE_A(mem *Memory, reg *Registers) {
	addr := reg.DE()
	mem.Write8(addr, reg.A)
}

// LD_pHL_ loads the value of A into the address pointed by HL.
func LD_pHL_A(mem *Memory, reg *Registers) {
	addr := reg.HL()
	mem.Write8(addr, reg.A)
}

// LD_pnn_A loads the value of A into the address pointed by the 16-bit immedite value.
func LD_pnn_A(mem *Memory, reg *Registers) {
	addr := mem.Read16(reg.PC)
	reg.PC += 2
	mem.Write8(addr, reg.A)
}

// LDI_pHL_A loads the value of A to address HL and increments HL.
func LDI_pHL_A(mem *Memory, reg *Registers) {
	addr := reg.HL()
	mem.Write8(addr, reg.A)
	addr += 1
	reg.SetHL(addr)
}

// LDI_A_pHL loads the value pointed by HL into A and increments HL.
func LDI_A_pHL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	reg.A = mem.Read8(addr)
	addr += 1
	reg.SetHL(addr)
}

// LDD_pHL_A loads the value of A to address HL and decrements HL.
func LDD_pHL_A(mem *Memory, reg *Registers) {
	addr := reg.HL()
	mem.Write8(addr, reg.A)
	addr -= 1
	reg.SetHL(addr)
}

// LDD_A_pHL loads the value pointed by HL into A and decrements HL.
func LDD_A_pHL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	reg.A = mem.Read8(addr)
	addr -= 1
	reg.SetHL(addr)
}

// LD_IO_C_A loads the value of A to address 0xFF00 + C (I/O memory).
func LD_IO_C_A(mem *Memory, reg *Registers) {
	addr := 0xff00 + uint16(reg.C)
	mem.Write8(addr, reg.A)
}

// LD_A_IO_C loads the value at 0xFF00 + C (I/O memory) into A.
func LD_A_IO_C(mem *Memory, reg *Registers) {
	addr := 0xff00 + uint16(reg.C)
	reg.A = mem.Read8(addr)
}

// LD_IO_n_A loads the value of A to address 0xFF00 + immediate value (I/O memory).
func LD_IO_n_A(mem *Memory, reg *Registers) {
	addr := 0xff00 + uint16(mem.Read8(reg.PC))
	reg.PC += 1
	mem.Write8(addr, reg.A)
}

// LD_A_IO_n loads the value at 0xFF00 + immediate value (I/O memory) into A.
func LD_A_IO_n(mem *Memory, reg *Registers) {
	addr := 0xff00 + uint16(mem.Read8(reg.PC))
	reg.PC += 1
	reg.A = mem.Read8(addr)
}

// ###### 16-Bit Loads ######

// LD_BC_nn loads a 16 bit immediate value into BC.
func LD_BC_nn(mem *Memory, reg *Registers) {
	reg.SetBC(mem.Read16(reg.PC))
	reg.PC += 2
}

// LD_DE_nn loads a 16 bit immediate value into DE.
func LD_DE_nn(mem *Memory, reg *Registers) {
	reg.SetDE(mem.Read16(reg.PC))
	reg.PC += 2
}

// LD_HL_nn loads a 16 bit immediate value into HL.
func LD_HL_nn(mem *Memory, reg *Registers) {
	reg.SetHL(mem.Read16(reg.PC))
	reg.PC += 2
}

// LD_SP_nn loads a 16 bit immediate value into SP.
func LD_SP_nn(mem *Memory, reg *Registers) {
	reg.SP = mem.Read16(reg.PC)
	reg.PC += 2
}

// PUSH_AF pushes AF onto the stack.
func PUSH_AF(mem *Memory, reg *Registers) {
	value := reg.AF()
	pushOntoStack(value, mem, reg)
}

// PUSH_BC pushes BC onto the stack.
func PUSH_BC(mem *Memory, reg *Registers) {
	value := reg.BC()
	pushOntoStack(value, mem, reg)
}

// PUSH_DE pushes DE onto the stack.
func PUSH_DE(mem *Memory, reg *Registers) {
	value := reg.DE()
	pushOntoStack(value, mem, reg)
}

// PUSH_HL pushes HL onto the stack.
func PUSH_HL(mem *Memory, reg *Registers) {
	value := reg.HL()
	pushOntoStack(value, mem, reg)
}

// POP_AF pops two bytes off stack into AF.
func POP_AF(mem *Memory, reg *Registers) {
	value := popOffStack(mem, reg)
	reg.SetAF(value)
}

// POP_BC pops two bytes off stack into BC.
func POP_BC(mem *Memory, reg *Registers) {
	value := popOffStack(mem, reg)
	reg.SetBC(value)
}

// POP_DE pops two bytes off stack into DE.
func POP_DE(mem *Memory, reg *Registers) {
	value := popOffStack(mem, reg)
	reg.SetDE(value)
}

// POP_HL pops two bytes off stack into HL.
func POP_HL(mem *Memory, reg *Registers) {
	value := popOffStack(mem, reg)
	reg.SetHL(value)
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

// XOR_A_pHL xors A with the value pointed by HL and puts result into A.
func XOR_A_pHL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	xorA(mem.Read8(addr), reg)
}

// XOR_A_n xors A with the 8 bit immediate value and puts result into A.
func XOR_A_n(mem *Memory, reg *Registers) {
	n := mem.Read8(reg.PC)
	reg.PC += 1
	xorA(n, reg)
}

// ###### 8-Bit Arithmetic Operations ######

// INC_A increments register A.
func INC_A(mem *Memory, reg *Registers) {
	increment(&reg.A, reg)
}

// INC_B increments register B.
func INC_B(mem *Memory, reg *Registers) {
	increment(&reg.B, reg)
}

// INC_C increments register C.
func INC_C(mem *Memory, reg *Registers) {
	increment(&reg.C, reg)
}

// INC_D increments register D.
func INC_D(mem *Memory, reg *Registers) {
	increment(&reg.D, reg)
}

// INC_E increments register E.
func INC_E(mem *Memory, reg *Registers) {
	increment(&reg.E, reg)
}

// INC_H increments register H.
func INC_H(mem *Memory, reg *Registers) {
	increment(&reg.H, reg)
}

// INC_L increments register L.
func INC_L(mem *Memory, reg *Registers) {
	increment(&reg.L, reg)
}

// INC_pHL increments the value pointed to by HL.
func INC_pHL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	value := mem.Read8(addr)
	increment(&value, reg)
	mem.Write8(addr, value)
}

// DEC_A decrements register A.
func DEC_A(mem *Memory, reg *Registers) {
	decrement(&reg.A, reg)
}

// DEC_B decrements register B.
func DEC_B(mem *Memory, reg *Registers) {
	decrement(&reg.B, reg)
}

// DEC_C decrements register C.
func DEC_C(mem *Memory, reg *Registers) {
	decrement(&reg.C, reg)
}

// DEC_D decrements register D.
func DEC_D(mem *Memory, reg *Registers) {
	decrement(&reg.D, reg)
}

// DEC_E decrements register E.
func DEC_E(mem *Memory, reg *Registers) {
	decrement(&reg.E, reg)
}

// DEC_H decrements register H.
func DEC_H(mem *Memory, reg *Registers) {
	decrement(&reg.H, reg)
}

// DEC_L decrements register L.
func DEC_L(mem *Memory, reg *Registers) {
	decrement(&reg.L, reg)
}

// DEC_pHL decrements the value pointed to by HL.
func DEC_pHL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	value := mem.Read8(addr)
	decrement(&value, reg)
	mem.Write8(addr, value)
}

// CP_A compares A with A.
func CP_A(mem *Memory, reg *Registers) {
	subtract(reg.A, reg.A, reg)
}

// CP_B compares A with B.
func CP_B(mem *Memory, reg *Registers) {
	subtract(reg.A, reg.B, reg)
}

// CP_C compares A with C.
func CP_C(mem *Memory, reg *Registers) {
	subtract(reg.A, reg.C, reg)
}

// CP_D compares A with D.
func CP_D(mem *Memory, reg *Registers) {
	subtract(reg.A, reg.D, reg)
}

// CP_E compares A with E.
func CP_E(mem *Memory, reg *Registers) {
	subtract(reg.A, reg.E, reg)
}

// CP_H compares A with H.
func CP_H(mem *Memory, reg *Registers) {
	subtract(reg.A, reg.H, reg)
}

// CP_L compares A with L.
func CP_L(mem *Memory, reg *Registers) {
	subtract(reg.A, reg.L, reg)
}

// CP_pHL compares A with the value pointed by HL.
func CP_pHL(mem *Memory, reg *Registers) {
	value := mem.Read8(reg.HL())
	subtract(reg.A, value, reg)
}

// CP_n compares A with an 8-bit immediate value.
func CP_n(mem *Memory, reg *Registers) {
	value := mem.Read8(reg.PC)
	subtract(reg.A, value, reg)
	reg.PC += 1
}

// ###### 16-Bit Arithmetic Operations ######

// INC_BC increments register BC.
func INC_BC(mem *Memory, reg *Registers) {
	value := reg.BC()
	value += 1
	reg.SetBC(value)
}

// INC_DE increments register DE.
func INC_DE(mem *Memory, reg *Registers) {
	value := reg.DE()
	value += 1
	reg.SetDE(value)
}

// INC_HL increments register HL.
func INC_HL(mem *Memory, reg *Registers) {
	value := reg.HL()
	value += 1
	reg.SetHL(value)
}

// INC_SP increments register SP.
func INC_SP(mem *Memory, reg *Registers) {
	reg.SP += 1
}

// DEC_BC decrements register BC.
func DEC_BC(mem *Memory, reg *Registers) {
	value := reg.BC()
	value -= 1
	reg.SetBC(value)
}

// DEC_DE decrements register DE.
func DEC_DE(mem *Memory, reg *Registers) {
	value := reg.DE()
	value -= 1
	reg.SetDE(value)
}

// DEC_HL decrements register HL.
func DEC_HL(mem *Memory, reg *Registers) {
	value := reg.HL()
	value -= 1
	reg.SetHL(value)
}

// DEC_SP decrements register SP.
func DEC_SP(mem *Memory, reg *Registers) {
	reg.SP -= 1
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

// BIT_0_pHL tests bit 0 of the value pointed by HL.
func BIT_0_pHL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	testBit(mem.Read8(addr), 0, reg)
}

// BIT_1_pHL tests bit 1 of the value pointed by HL.
func BIT_1_pHL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	testBit(mem.Read8(addr), 1, reg)
}

// BIT_2_pHL tests bit 2 of the value pointed by HL.
func BIT_2_pHL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	testBit(mem.Read8(addr), 2, reg)
}

// BIT_3_pHL tests bit 3 of the value pointed by HL.
func BIT_3_pHL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	testBit(mem.Read8(addr), 3, reg)
}

// BIT_4_pHL tests bit 4 of the value pointed by HL.
func BIT_4_pHL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	testBit(mem.Read8(addr), 4, reg)
}

// BIT_5_pHL tests bit 5 of the value pointed by HL.
func BIT_5_pHL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	testBit(mem.Read8(addr), 5, reg)
}

// BIT_6_pHL tests bit 6 of the value pointed by HL.
func BIT_6_pHL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	testBit(mem.Read8(addr), 6, reg)
}

// BIT_7_pHL tests bit 7 of the value pointed by HL.
func BIT_7_pHL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	testBit(mem.Read8(addr), 7, reg)
}

// ###### Absolute Jumps ######

// JP_nn jumps to the address pointed by a 16-bit immediate value.
func JP_nn(mem *Memory, reg *Registers) {
	jumpToImmediateValue(mem, reg)
}

// JP_NZ_nn jumps to the address pointed by a 16-bit immediate value
// if the zero flag is not set.
func JP_NZ_nn(mem *Memory, reg *Registers) {
	if !reg.IsFlagSet(zeroFlag) {
		jumpToImmediateValue(mem, reg)
	} else {
		reg.PC += 2
	}
}

// JP_Z_nn jumps to the address pointed by a 16-bit immediate value
// if the zero flag is set.
func JP_Z_nn(mem *Memory, reg *Registers) {
	if reg.IsFlagSet(zeroFlag) {
		jumpToImmediateValue(mem, reg)
	} else {
		reg.PC += 2
	}
}

// JP_NC_nn jumps to the address pointed by a 16-bit immediate value
// if the carry flag is not set.
func JP_NC_nn(mem *Memory, reg *Registers) {
	if !reg.IsFlagSet(carryFlag) {
		jumpToImmediateValue(mem, reg)
	} else {
		reg.PC += 2
	}
}

// JP_pHL jumps to the address pointed by a HL.
func JP_pHL(mem *Memory, reg *Registers) {
	reg.PC = reg.HL()
}

// JP_C_nn jumps to the address pointed by a 16-bit immediate value
// if the carry flag is set.
func JP_C_nn(mem *Memory, reg *Registers) {
	if reg.IsFlagSet(carryFlag) {
		jumpToImmediateValue(mem, reg)
	} else {
		reg.PC += 2
	}
}

// ###### Relative Jumps ######

// JP_n adds an 8 bit signed immediate value to current PC and jumps to it.
func JP_n(mem *Memory, reg *Registers) {
	relJumpByImmediateValue(mem, reg)
}

// If the zero flag is not set JP_NZ_n adds an 8 bit signed immediate
// value to current PC and jumps to it.
func JP_NZ_n(mem *Memory, reg *Registers) {
	if !reg.IsFlagSet(zeroFlag) {
		relJumpByImmediateValue(mem, reg)
	} else {
		reg.PC += 1
	}
}

// If the zero flag is set JP_Z_n adds an 8 bit signed immediate
// value to current PC and jumps to it.
func JP_Z_n(mem *Memory, reg *Registers) {
	if reg.IsFlagSet(zeroFlag) {
		relJumpByImmediateValue(mem, reg)
	} else {
		reg.PC += 1
	}
}

// If the carry flag is not set JP_NC_n adds an 8 bit signed immediate
// value to current PC and jumps to it.
func JP_NC_n(mem *Memory, reg *Registers) {
	if !reg.IsFlagSet(carryFlag) {
		relJumpByImmediateValue(mem, reg)
	} else {
		reg.PC += 1
	}
}

// If the carry flag is set JP_C_n adds an 8 bit signed immediate
// value to current PC and jumps to it.
func JP_C_n(mem *Memory, reg *Registers) {
	if reg.IsFlagSet(carryFlag) {
		relJumpByImmediateValue(mem, reg)
	} else {
		reg.PC += 1
	}
}

// ###### Calls ######

// CALL_nn calls the address pointed by the immediate value.
func CALL_nn(mem *Memory, reg *Registers) {
	callImmediateValue(mem, reg)
}

// CALL_NZ_nn calls the address pointed by the immediate value if zero flag is not set.
func CALL_NZ_nn(mem *Memory, reg *Registers) {
	if !reg.IsFlagSet(zeroFlag) {
		callImmediateValue(mem, reg)
	} else {
		reg.PC += 2
	}
}

// CALL_Z_nn calls the address pointed by the immediate value if zero flag is set.
func CALL_Z_nn(mem *Memory, reg *Registers) {
	if reg.IsFlagSet(zeroFlag) {
		callImmediateValue(mem, reg)
	} else {
		reg.PC += 2
	}
}

// CALL_NC_nn calls the address pointed by the immediate value if carry flag is not set.
func CALL_NC_nn(mem *Memory, reg *Registers) {
	if !reg.IsFlagSet(carryFlag) {
		callImmediateValue(mem, reg)
	} else {
		reg.PC += 2
	}
}

// CALL_C_nn calls the address pointed by the immediate value if carry flag is set.
func CALL_C_nn(mem *Memory, reg *Registers) {
	if reg.IsFlagSet(carryFlag) {
		callImmediateValue(mem, reg)
	} else {
		reg.PC += 2
	}
}

// ###### Returns ######

// RET returns to the address on top of the stack.
func RET(mem *Memory, reg *Registers) {
	returnFromStack(mem, reg)
}

// RET_NZ returns to the address on top of the stack if zero flag is not set.
func RET_NZ(mem *Memory, reg *Registers) {
	if !reg.IsFlagSet(zeroFlag) {
		returnFromStack(mem, reg)
	}
}

// RET_Z returns to the address on top of the stack if zero flag is set.
func RET_Z(mem *Memory, reg *Registers) {
	if reg.IsFlagSet(zeroFlag) {
		returnFromStack(mem, reg)
	}
}

// RET_NC returns to the address on top of the stack if carry flag is not set.
func RET_NC(mem *Memory, reg *Registers) {
	if !reg.IsFlagSet(carryFlag) {
		returnFromStack(mem, reg)
	}
}

// RET_C returns to the address on top of the stack if carry flag is set.
func RET_C(mem *Memory, reg *Registers) {
	if reg.IsFlagSet(carryFlag) {
		returnFromStack(mem, reg)
	}
}

// ###### Rotates and Shifts ######

// RLA rotates A to the left (with Carry flag -> Bit 0, Bit 7 -> Carry flag).
func RLA(mem *Memory, reg *Registers) {
	RL_A(mem, reg)
}

// RL_A rotates A to the left (with Carry flag -> Bit 0, Bit 7 -> Carry flag).
func RL_A(mem *Memory, reg *Registers) {
	rotateLeftThroughCarry(&reg.A, reg)
}

// RL_B rotates B to the left (with Carry flag -> Bit 0, Bit 7 -> Carry flag).
func RL_B(mem *Memory, reg *Registers) {
	rotateLeftThroughCarry(&reg.B, reg)
}

// RL_C rotates C to the left (with Carry flag -> Bit 0, Bit 7 -> Carry flag).
func RL_C(mem *Memory, reg *Registers) {
	rotateLeftThroughCarry(&reg.C, reg)
}

// RL_D rotates D to the left (with Carry flag -> Bit 0, Bit 7 -> Carry flag).
func RL_D(mem *Memory, reg *Registers) {
	rotateLeftThroughCarry(&reg.D, reg)
}

// RL_E rotates E to the left (with Carry flag -> Bit 0, Bit 7 -> Carry flag).
func RL_E(mem *Memory, reg *Registers) {
	rotateLeftThroughCarry(&reg.E, reg)
}

// RL_H rotates H to the left (with Carry flag -> Bit 0, Bit 7 -> Carry flag.
func RL_H(mem *Memory, reg *Registers) {
	rotateLeftThroughCarry(&reg.H, reg)
}

// RL_L rotates L to the left (with Carry flag -> Bit 0, Bit 7 -> Carry flag).
func RL_L(mem *Memory, reg *Registers) {
	rotateLeftThroughCarry(&reg.L, reg)
}

// RL_pHL rotates the value pointed by HL to the left (with Carry flag -> Bit 0, Bit 7 -> Carry flag).
func RL_pHL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	value := mem.Read8(addr)
	rotateLeftThroughCarry(&value, reg)
	mem.Write8(addr, value)
}

// RRA rotates A to the right (with Carry flag -> Bit 7, Bit 0 -> Carry flag).
func RRA(mem *Memory, reg *Registers) {
	RR_A(mem, reg)
}

// RR_A rotates A to the right (with Carry flag -> Bit 7, Bit 0 -> Carry flag).
func RR_A(mem *Memory, reg *Registers) {
	rotateRightThroughCarry(&reg.A, reg)
}

// RR_B rotates B to the right (with Carry flag -> Bit 7, Bit 0 -> Carry flag).
func RR_B(mem *Memory, reg *Registers) {
	rotateRightThroughCarry(&reg.B, reg)
}

// RR_C rotates C to the right (with Carry flag -> Bit 7, Bit 0 -> Carry flag).
func RR_C(mem *Memory, reg *Registers) {
	rotateRightThroughCarry(&reg.C, reg)
}

// RR_D rotates D to the right (with Carry flag -> Bit 7, Bit 0 -> Carry flag).
func RR_D(mem *Memory, reg *Registers) {
	rotateRightThroughCarry(&reg.D, reg)
}

// RR_E rotates E to the right (with Carry flag -> Bit 7, Bit 0 -> Carry flag).
func RR_E(mem *Memory, reg *Registers) {
	rotateRightThroughCarry(&reg.E, reg)
}

// RR_H rotates H to the right (with Carry flag -> Bit 7, Bit 0 -> Carry flag).
func RR_H(mem *Memory, reg *Registers) {
	rotateRightThroughCarry(&reg.H, reg)
}

// RR_L rotates L to the right (with Carry flag -> Bit 7, Bit 0 -> Carry flag).
func RR_L(mem *Memory, reg *Registers) {
	rotateRightThroughCarry(&reg.L, reg)
}

// RR_pHL rotates the value pointed by HL to the right (with Carry flag -> Bit 0, Bit 7 -> Carry flag).
func RR_pHL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	value := mem.Read8(addr)
	rotateRightThroughCarry(&value, reg)
	mem.Write8(addr, value)
}

// RLCA rotates A to the left (with Bit 7 -> Carry flag and Bit 0).
func RLCA(mem *Memory, reg *Registers) {
	RLC_A(mem, reg)
}

// RLC_A rotates A to the left (with Bit 7 -> Carry flag and Bit 0).
func RLC_A(mem *Memory, reg *Registers) {
	rotateLeft(&reg.A, reg)
}

// RLC_B rotates B to the left (with Bit 7 -> Carry flag and Bit 0).
func RLC_B(mem *Memory, reg *Registers) {
	rotateLeft(&reg.B, reg)
}

// RLC_C rotates C to the left (with Bit 7 -> Carry flag and Bit 0).
func RLC_C(mem *Memory, reg *Registers) {
	rotateLeft(&reg.C, reg)
}

// RLC_D rotates D to the left (with Bit 7 -> Carry flag and Bit 0).
func RLC_D(mem *Memory, reg *Registers) {
	rotateLeft(&reg.D, reg)
}

// RLC_E rotates E to the left (with Bit 7 -> Carry flag and Bit 0).
func RLC_E(mem *Memory, reg *Registers) {
	rotateLeft(&reg.E, reg)
}

// RLC_H rotates H to the left (with Bit 7 -> Carry flag and Bit 0).
func RLC_H(mem *Memory, reg *Registers) {
	rotateLeft(&reg.H, reg)
}

// RLC_L rotates L to the left (with Bit 7 -> Carry flag and Bit 0).
func RLC_L(mem *Memory, reg *Registers) {
	rotateLeft(&reg.L, reg)
}

// RLC_pHL rotates the value pointed by HL to the left  (with Bit 7 -> Carry flag and Bit 0).
func RLC_pHL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	value := mem.Read8(addr)
	rotateLeft(&value, reg)
	mem.Write8(addr, value)
}

// RRCA rotates A to the right (with Bit 7 -> Carry flag and Bit 0).
func RRCA(mem *Memory, reg *Registers) {
	RRC_A(mem, reg)
}

// RRC_A rotates A to the right (with Bit 7 -> Carry flag and Bit 0).
func RRC_A(mem *Memory, reg *Registers) {
	rotateRight(&reg.A, reg)
}

// RRC_B rotates B to the right (with Bit 7 -> Carry flag and Bit 0).
func RRC_B(mem *Memory, reg *Registers) {
	rotateRight(&reg.B, reg)
}

// RRC_C rotates C to the right (with Bit 7 -> Carry flag and Bit 0).
func RRC_C(mem *Memory, reg *Registers) {
	rotateRight(&reg.C, reg)
}

// RRC_D rotates D to the right (with Bit 7 -> Carry flag and Bit 0).
func RRC_D(mem *Memory, reg *Registers) {
	rotateRight(&reg.D, reg)
}

// RRC_E rotates E to the right (with Bit 7 -> Carry flag and Bit 0).
func RRC_E(mem *Memory, reg *Registers) {
	rotateRight(&reg.E, reg)
}

// RRC_H rotates H to the right (with Bit 7 -> Carry flag and Bit 0).
func RRC_H(mem *Memory, reg *Registers) {
	rotateRight(&reg.H, reg)
}

// RRC_L rotates L to the right (with Bit 7 -> Carry flag and Bit 0).
func RRC_L(mem *Memory, reg *Registers) {
	rotateRight(&reg.L, reg)
}

// RRC_pHL rotates the value pointed by HL to the right  (with Bit 7 -> Carry flag and Bit 0).
func RRC_pHL(mem *Memory, reg *Registers) {
	addr := reg.HL()
	value := mem.Read8(addr)
	rotateRight(&value, reg)
	mem.Write8(addr, value)
}

// ###### Common functions used by the instructions ######

// xorA xors register A with the given reigster or memory value and puts the result into A.
func xorA(value uint8, reg *Registers) {
	reg.A = reg.A ^ value
	reg.SetFlags(subtractFlag|halfCarryFlag|carryFlag, false)
	reg.SetFlags(zeroFlag, reg.A == 0)
}

// testBit checks if the given bit is set in the given register or memory value.
func testBit(value uint8, bit uint8, reg *Registers) {
	isSet := (value & (1 << bit)) != 0
	reg.SetFlags(zeroFlag, !isSet)
	reg.SetFlags(halfCarryFlag, true)
	reg.SetFlags(subtractFlag, false)
}

// increment increments the given register or memory value.
func increment(value *uint8, reg *Registers) {
	*value += 1
	reg.SetFlags(halfCarryFlag, (*value&0x0F) == 0)
	reg.SetFlags(subtractFlag, false)
	reg.SetFlags(zeroFlag, *value == 0)
}

// decrement decrements the given register or memory value.
func decrement(value *uint8, reg *Registers) {
	reg.SetFlags(halfCarryFlag, (*value&0x0F) == 0)
	*value -= 1
	reg.SetFlags(subtractFlag, true)
	reg.SetFlags(zeroFlag, *value == 0)
}

// jumpToImmediateValue performs an absolute jump to the address pointed
// by an 16 bit immediate value.
func jumpToImmediateValue(mem *Memory, reg *Registers) {
	reg.PC = mem.Read16(reg.PC)
}

// relJumpByImmediateValue performs a relative jump by adding an 8 bit signed
// immediate value to the current PC.
func relJumpByImmediateValue(mem *Memory, reg *Registers) {
	n := int8(mem.Read8(reg.PC))
	reg.PC += 1
	relJump(n, reg)
}

// relJump performs a relative jump by adding n to the curren PC.
func relJump(n int8, reg *Registers) {
	newPC := int32(reg.PC) + int32(n)
	if newPC < 0 || math.MaxUint16 < newPC {
		panic(fmt.Sprintf(
			"invalid relative jump (current PC: 0x%x, jump: %v)",
			reg.PC, n))
	}
	reg.PC = uint16(newPC)
}

// callImmediateValue calls the address pointed by a 16-bit immediate value by
// pushing the address of the next instruction onto the stack and
// jumping to the address pointed by the 16-bit imemdiate value.
func callImmediateValue(mem *Memory, reg *Registers) {
	pushOntoStack(reg.PC+2, mem, reg)
	reg.PC = mem.Read16(reg.PC)
}

// returnFromStack pops a 16-bit address of the stack and jumps to it.
func returnFromStack(mem *Memory, reg *Registers) {
	reg.PC = popOffStack(mem, reg)
}

// pushOntoStack pushes a 16-bit value onto the stack.
func pushOntoStack(value uint16, mem *Memory, reg *Registers) {
	reg.SP -= 2
	mem.Write16(reg.SP, value)
}

// popOffStack pops a 16-bit value off the stack.
func popOffStack(mem *Memory, reg *Registers) uint16 {
	value := mem.Read16(reg.SP)
	reg.SP += 2
	return value
}

// rotateLeft rotates the given register or memory value to the left by one bit.
// Copies bit 7 to bit 0 and to the the carry flag.
func rotateLeft(value *uint8, reg *Registers) {
	carry := (*value & 0b10000000) != 0

	*value = *value << 1
	if carry {
		*value = *value | 0b00000001
	} else {
		*value &= ^uint8(0b00000001)
	}

	reg.SetFlags(carryFlag, carry)
	reg.SetFlags(subtractFlag|halfCarryFlag, false)
	reg.SetFlags(zeroFlag, *value == 0)
}

// rotateRight rotates the given register or memory value to the right by one bit.
// Copies bit 0 to bit 7 and to the the carry flag.
func rotateRight(value *uint8, reg *Registers) {
	carry := (*value & 0b00000001) != 0

	*value = *value >> 1
	if carry {
		*value = *value | 0b10000000
	} else {
		*value &= ^uint8(0b10000000)
	}

	reg.SetFlags(carryFlag, carry)
	reg.SetFlags(subtractFlag|halfCarryFlag, false)
	reg.SetFlags(zeroFlag, *value == 0)
}

// rotateLeftThroughCarry rotates the given register or memory value to the left by one bit.
// Copies the carry flag to bit 0 and bit 7 to the carry flag.
func rotateLeftThroughCarry(value *uint8, reg *Registers) {
	oldCarry := reg.IsFlagSet(carryFlag)
	newCarry := (*value & 0b10000000) != 0

	*value = *value << 1
	if oldCarry {
		*value = *value | 0b00000001
	} else {
		*value &= ^uint8(0b00000001)
	}

	reg.SetFlags(carryFlag, newCarry)
	reg.SetFlags(subtractFlag|halfCarryFlag, false)
	reg.SetFlags(zeroFlag, *value == 0)
}

// rotateRightThroughCarry rotates the given register or memory value to the right by one bit.
// Copies the carry flag to bit 7 and bit 0 to the carry flag.
func rotateRightThroughCarry(value *uint8, reg *Registers) {
	oldCarry := reg.IsFlagSet(carryFlag)
	newCarry := (*value & 0b00000001) != 0

	*value = *value >> 1
	if oldCarry {
		*value = *value | 0b10000000
	} else {
		*value &= ^uint8(0b10000000)
	}

	reg.SetFlags(carryFlag, newCarry)
	reg.SetFlags(subtractFlag|halfCarryFlag, false)
	reg.SetFlags(zeroFlag, *value == 0)
}

// subtract subtracts subtrahend from minuend and returns the result.
func subtract(minuend uint8, subtrahend uint8, reg *Registers) uint8 {
	reg.SetFlags(subtractFlag, true)
	reg.SetFlags(zeroFlag, subtrahend == minuend)
	reg.SetFlags(carryFlag, subtrahend > minuend)
	reg.SetFlags(halfCarryFlag, (subtrahend&0x0f) > (minuend&0x0f))
	return minuend - subtrahend
}
