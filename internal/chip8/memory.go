package chip8

type Memory struct {
	v [4096]uint8
}

func (o *Memory) Read8(addr uint16) uint8 {
	return o.v[addr]
}

func (o *Memory) Read16(addr uint16) uint16 {
	return (uint16(o.v[addr]) << 8) | uint16(o.v[addr+1])
}

func (o *Memory) Write(addr uint16, value uint8) {
	o.v[addr] = value
}

func (o *Memory) WriteProgram(program []uint8) {
	for i, b := range program {
		var addr uint16 = 0x200 + uint16(i)
		o.Write(addr, b)
	}
}
