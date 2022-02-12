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

const zeroFlag uint8 = 1 << 7
const subtractFlag uint8 = 1 << 6
const halfCarryFlag uint8 = 1 << 5
const carryFlag uint8 = 1 << 4
