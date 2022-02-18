package gb

type Registers struct {
	A  uint8
	B  uint8
	C  uint8
	D  uint8
	E  uint8
	F  uint8
	L  uint8
	H  uint8
	SP uint16
	PC uint16
}

func (r *Registers) SetFlags(flags uint8, set bool) {
	if set {
		r.F |= flags
	} else {
		r.F &= ^flags
	}
}

func (r *Registers) IsFlagSet(flag uint8) bool {
	return (r.F & flag) != 0
}

func (r *Registers) HL() uint16 {
	return (uint16(r.H) << 8) | uint16(r.L)
}

func (r *Registers) Reset() {
	r.A = 0
	r.B = 0
	r.C = 0
	r.D = 0
	r.E = 0
	r.F = 0
	r.L = 0
	r.H = 0
	r.SP = 0
	r.PC = 0
}

const zeroFlag uint8 = 1 << 7      // Z
const subtractFlag uint8 = 1 << 6  // N
const halfCarryFlag uint8 = 1 << 5 // H
const carryFlag uint8 = 1 << 4     // C
