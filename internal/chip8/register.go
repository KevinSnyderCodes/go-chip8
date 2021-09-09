package chip8

type Register struct {
	v  [16]uint8
	i  uint16
	dt uint8
	st uint8
}

func (o *Register) ReadV(n uint8) uint8 {
	return o.v[n]
}

func (o *Register) WriteV(n uint8, value uint8) {
	o.v[n] = value
}

func (o *Register) ReadI() uint16 {
	return o.i
}

func (o *Register) WriteI(value uint16) {
	o.i = value
}

func (o *Register) ReadDT() uint8 {
	return o.dt
}

func (o *Register) WriteDT(value uint8) {
	o.dt = value
}

func (o *Register) ReadST() uint8 {
	return o.st
}

func (o *Register) WriteST(value uint8) {
	o.st = value
}
