package chip8

type Stack struct {
	v  [16]uint16
	sp uint8
}

func (o *Stack) Push(value uint16) {
	o.sp += 1
	o.v[o.sp] = value
}

func (o *Stack) Pop() uint16 {
	value := o.v[o.sp]
	o.sp -= 1
	return value
}
