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

func (r *Registers) SetFlags(flags uint8) {
	r.F |= flags
}

func (r *Registers) ClearFlags(flags uint8) {
	r.F &= ^flags
}

func (r *Registers) IsFlagSet(flag uint8) bool {
	return (r.F & flag) != 0
}

const zeroFlag uint8 = 1 << 7
const subtractFlag uint8 = 1 << 6
const halfCarryFlag uint8 = 1 << 5
const carryFlag uint8 = 1 << 4
