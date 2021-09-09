package chip8

type Display struct {
	v [64][32]bool
}

func (o *Display) Get(x uint8, y uint8) bool {
	return o.v[x][y]
}

func (o *Display) Set(x uint8, y uint8, value bool) {
	o.v[x][y] = value
}

func (o *Display) Dump() [64][32]bool {
	return o.v
}

func (o *Display) Clear() {
	for x := 0; x < 64; x++ {
		for y := 0; y < 32; y++ {
			o.v[x][y] = false
		}
	}
}

func (o *Display) Draw(x uint8, y uint8, n uint8, sprite [256]uint8) (collision bool) {
	for iy := uint8(0); iy < n; iy++ {
		var and uint8 = 0b10000000
		for ix := uint8(0); ix < 8; ix++ {
			this_x := (x + ix) % 64
			this_y := (y + iy) % 32
			sprite_on := (sprite[iy] & and) == and
			display_on := o.Get(this_x, this_y)
			value := sprite_on != display_on
			if !value && display_on {
				collision = true
			}
			o.Set(this_x, this_y, value)
			and = and >> 1
		}
	}
	return
}
