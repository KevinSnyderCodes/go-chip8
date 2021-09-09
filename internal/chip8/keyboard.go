package chip8

type Keyboard struct {
	v [16]bool
}

func (o *Keyboard) Get(n uint8) bool {
	return o.v[n]
}
