package chip8

type PC struct {
	v uint16
}

func (o *PC) Get() uint16 {
	return o.v
}

func (o *PC) Set(value uint16) {
	o.v = value
}

func (o *PC) Increment() {
	o.v += 2
}
