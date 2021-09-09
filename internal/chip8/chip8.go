package chip8

import (
	"fmt"
	"math/rand"
	"time"
)

type Chip8 struct {
	// Pseudo-registers
	pc PC

	register Register

	stack Stack

	memory Memory

	display Display

	rand *rand.Rand
}

func (o *Chip8) Reset(program []uint8) {
	// Seed random number generator
	o.rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	// Initialize memory
	{
		var addr uint16 = 0x0000
		for _, item := range RomSlice {
			for _, b := range item {
				o.memory.Write(addr, b)
				addr++
			}
		}
	}

	// Write program to memory
	o.memory.WriteProgram(program)

	// Set PC to 0x200 (start of most Chip-8 programs)
	o.pc.Set(0x200)
}

func (o *Chip8) DumpDisplay() [64][32]bool {
	return o.display.Dump()
}

func (o *Chip8) Tick() error {
	pc := o.pc.Get()
	opRaw := o.memory.Read16(pc)

	fmt.Printf("PC %x INS %x\n", pc-0x200, opRaw)

	if opRaw == 0x0000 {
		return nil
	}

	opType := GetOpType(opRaw)
	switch opType {
	case OpType_CLS:
		op := GetOp_CLS(opRaw)
		o.RunOp_CLS(op)
	case OpType_RET:
		op := GetOp_RET(opRaw)
		o.RunOp_RET(op)
	case OpType_JP_addr:
		op := GetOp_JP_addr(opRaw)
		o.RunOp_JP_addr(op)
	case OpType_CALL_addr:
		op := GetOp_CALL_addr(opRaw)
		o.RunOp_CALL_addr(op)
	case OpType_SE_Vx_byte:
		op := GetOp_SE_Vx_byte(opRaw)
		o.RunOp_SE_Vx_byte(op)
	case OpType_SNE_Vx_byte:
		op := GetOp_SNE_Vx_byte(opRaw)
		o.RunOp_SNE_Vx_byte(op)
	case OpType_SE_Vx_Vy:
		op := GetOp_SE_Vx_Vy(opRaw)
		o.RunOp_SE_Vx_Vy(op)
	case OpType_LD_Vx_byte:
		op := GetOp_LD_Vx_byte(opRaw)
		o.RunOp_LD_Vx_byte(op)
	case OpType_ADD_Vx_byte:
		op := GetOp_ADD_Vx_byte(opRaw)
		o.RunOp_ADD_Vx_byte(op)
	case OpType_LD_Vx_Vy:
		op := GetOp_LD_Vx_Vy(opRaw)
		o.RunOp_LD_Vx_Vy(op)
	case OpType_OR_Vx_Vy:
		op := GetOp_LD_Vx_Vy(opRaw)
		o.RunOp_LD_Vx_Vy(op)
	case OpType_AND_Vx_Vy:
		op := GetOp_AND_Vx_Vy(opRaw)
		o.RunOp_AND_Vx_Vy(op)
	case OpType_XOR_Vx_Vy:
		op := GetOp_XOR_Vx_Vy(opRaw)
		o.RunOp_XOR_Vx_Vy(op)
	case OpType_ADD_Vx_Vy:
		op := GetOp_ADD_Vx_Vy(opRaw)
		o.RunOp_ADD_Vx_Vy(op)
	case OpType_SUB_Vx_Vy:
		op := GetOp_SUB_Vx_Vy(opRaw)
		o.RunOp_SUB_Vx_Vy(op)
	case OpType_SHR_Vx_Vy:
		op := GetOp_SHR_Vx_Vy(opRaw)
		o.RunOp_SHR_Vx_Vy(op)
	case OpType_SUBN_Vx_Vy:
		op := GetOp_SUBN_Vx_Vy(opRaw)
		o.RunOp_SUBN_Vx_Vy(op)
	case OpType_SHL_Vx_Vy:
		op := GetOp_SHL_Vx_Vy(opRaw)
		o.RunOp_SHL_Vx_Vy(op)
	case OpType_SNE_Vx_Vy:
		op := GetOp_SNE_Vx_Vy(opRaw)
		o.RunOp_SNE_Vx_Vy(op)
	case OpType_LD_I_addr:
		op := GetOp_LD_I_addr(opRaw)
		o.RunOp_LD_I_addr(op)
	case OpType_JP_V0_addr:
		op := GetOp_JP_V0_addr(opRaw)
		o.RunOp_JP_V0_addr(op)
	case OpType_RND_Vx_byte:
		op := GetOp_RND_Vx_byte(opRaw)
		o.RunOp_RND_Vx_byte(op)
	case OpType_DRW_Vx_Vy_nibble:
		op := GetOp_DRW_Vx_Vy_nibble(opRaw)
		o.RunOp_DRW_Vx_Vy_nibble(op)
	case OpType_SKP_Vx:
		op := GetOp_SKP_Vx(opRaw)
		o.RunOp_SKP_Vx(op)
	case OpType_SKNP_Vx:
		op := GetOp_SKNP_Vx(opRaw)
		o.RunOp_SKNP_Vx(op)
	case OpType_LD_Vx_DT:
		op := GetOp_LD_Vx_DT(opRaw)
		o.RunOp_LD_Vx_DT(op)
	case OpType_LD_Vx_K:
		op := GetOp_LD_Vx_K(opRaw)
		o.RunOp_LD_Vx_K(op)
	case OpType_LD_DT_Vx:
		op := GetOp_LD_DT_Vx(opRaw)
		o.RunOp_LD_DT_Vx(op)
	case OpType_LD_ST_Vx:
		op := GetOp_LD_ST_Vx(opRaw)
		o.RunOp_LD_ST_Vx(op)
	case OpType_ADD_I_Vx:
		op := GetOp_ADD_I_Vx(opRaw)
		o.RunOp_ADD_I_Vx(op)
	case OpType_LD_F_Vx:
		op := GetOp_LD_F_Vx(opRaw)
		o.RunOp_LD_F_Vx(op)
	case OpType_LD_B_Vx:
		op := GetOp_LD_B_Vx(opRaw)
		o.RunOp_LD_B_Vx(op)
	case OpType_LD_I_Vx:
		op := GetOp_LD_I_Vx(opRaw)
		o.RunOp_LD_I_Vx(op)
	case OpType_LD_Vx_I:
		op := GetOp_LD_Vx_I(opRaw)
		o.RunOp_LD_Vx_I(op)
	default:
		return fmt.Errorf("unknown op type for instruction %x", opRaw)
	}

	o.pc.Increment()

	return nil
}

func (o *Chip8) RunOp_CLS(v Op_CLS) {
	o.display.Clear()
}

func (o *Chip8) RunOp_RET(v Op_RET) {
	value := o.stack.Pop()
	o.pc.Set(value)
}

func (o *Chip8) RunOp_JP_addr(v Op_JP_addr) {
	o.pc.Set(v.Addr)
}

func (o *Chip8) RunOp_CALL_addr(v Op_CALL_addr) {
	pc := o.pc.Get()
	o.stack.Push(pc)
	o.pc.Set(v.Addr)
}

func (o *Chip8) RunOp_SE_Vx_byte(v Op_SE_Vx_byte) {
	value := o.register.ReadV(v.Vx)
	if value == v.Byte {
		o.pc.Increment()
	}
}

func (o *Chip8) RunOp_SNE_Vx_byte(v Op_SNE_Vx_byte) {
	value := o.register.ReadV(v.Vx)
	if value != v.Byte {
		o.pc.Increment()
	}
}

func (o *Chip8) RunOp_SE_Vx_Vy(v Op_SE_Vx_Vy) {
	valueX := o.register.ReadV(v.Vx)
	valueY := o.register.ReadV(v.Vy)
	if valueX == valueY {
		o.pc.Increment()
	}
}

func (o *Chip8) RunOp_LD_Vx_byte(v Op_LD_Vx_byte) {
	o.register.WriteV(v.Vx, v.Byte)
}

func (o *Chip8) RunOp_ADD_Vx_byte(v Op_ADD_Vx_byte) {
	value := o.register.ReadV(v.Vx)
	value += v.Byte
	o.register.WriteV(v.Vx, value)
}

func (o *Chip8) RunOp_LD_Vx_Vy(v Op_LD_Vx_Vy) {
	value := o.register.ReadV(v.Vy)
	o.register.WriteV(v.Vx, value)
}

func (o *Chip8) RunOp_OR_Vx_Vy(v Op_OR_Vx_Vy) {
	valueX := o.register.ReadV(v.Vx)
	valueY := o.register.ReadV(v.Vy)
	value := valueX | valueY
	o.register.WriteV(v.Vx, value)
}

func (o *Chip8) RunOp_AND_Vx_Vy(v Op_AND_Vx_Vy) {
	valueX := o.register.ReadV(v.Vx)
	valueY := o.register.ReadV(v.Vy)
	value := valueX & valueY
	o.register.WriteV(v.Vx, value)
}

func (o *Chip8) RunOp_XOR_Vx_Vy(v Op_XOR_Vx_Vy) {
	valueX := o.register.ReadV(v.Vx)
	valueY := o.register.ReadV(v.Vy)
	value := valueX ^ valueY
	o.register.WriteV(v.Vx, value)
}

func (o *Chip8) RunOp_ADD_Vx_Vy(v Op_ADD_Vx_Vy) {
	valueX := o.register.ReadV(v.Vx)
	valueY := o.register.ReadV(v.Vy)
	value := valueX + valueY
	o.register.WriteV(v.Vx, value)

	if (valueX & valueY & 0b10000000) == 0b10000000 {
		o.register.WriteV(0x0F, 1)
	}
}

func (o *Chip8) RunOp_SUB_Vx_Vy(v Op_SUB_Vx_Vy) {
	valueX := o.register.ReadV(v.Vx)
	valueY := o.register.ReadV(v.Vy)
	value := valueX - valueY
	o.register.WriteV(v.Vx, value)

	if valueX > valueY {
		o.register.WriteV(0x0F, 1)
	}
}

func (o *Chip8) RunOp_SHR_Vx_Vy(v Op_SHR_Vx_Vy) {
	valueX := o.register.ReadV(v.Vx)
	value := valueX >> 1
	o.register.WriteV(v.Vx, value)

	if (valueX & 0b00000001) == 0b00000001 {
		o.register.WriteV(0x0F, 1)
	}
}

func (o *Chip8) RunOp_SUBN_Vx_Vy(v Op_SUBN_Vx_Vy) {
	valueX := o.register.ReadV(v.Vx)
	valueY := o.register.ReadV(v.Vy)
	value := valueY - valueX
	o.register.WriteV(v.Vx, value)

	if valueY > valueX {
		o.register.WriteV(0x0F, 1)
	}
}

func (o *Chip8) RunOp_SHL_Vx_Vy(v Op_SHL_Vx_Vy) {
	valueX := o.register.ReadV(v.Vx)
	value := valueX << 1
	o.register.WriteV(v.Vx, value)

	if (valueX & 0b10000000) == 0b10000000 {
		o.register.WriteV(0x0F, 1)
	}
}

func (o *Chip8) RunOp_SNE_Vx_Vy(v Op_SNE_Vx_Vy) {
	valueX := o.register.ReadV(v.Vx)
	valueY := o.register.ReadV(v.Vy)
	if valueX != valueY {
		o.pc.Increment()
	}
}

func (o *Chip8) RunOp_LD_I_addr(v Op_LD_I_addr) {
	o.register.WriteI(v.Addr)
}

func (o *Chip8) RunOp_JP_V0_addr(v Op_JP_V0_addr) {
	valueX := o.register.ReadV(0x00)
	value := v.Addr + uint16(valueX)
	o.pc.Set(value)
}

func (o *Chip8) RunOp_RND_Vx_byte(v Op_RND_Vx_byte) {
	value := uint8(o.rand.Uint32()) & v.Byte
	o.register.WriteV(v.Vx, value)
}

func (o *Chip8) RunOp_DRW_Vx_Vy_nibble(v Op_DRW_Vx_Vy_nibble) {
	var sprite [256]uint8
	for i := uint8(0); i < v.Nibble; i++ {
		addr := o.register.ReadI() + uint16(i)
		sprite[i] = o.memory.Read8(addr)
	}

	valueX := o.register.ReadV(v.Vx)
	valueY := o.register.ReadV(v.Vy)
	collision := o.display.Draw(valueX, valueY, v.Nibble, sprite)

	if collision {
		o.register.WriteV(0x0F, 1)
	}
}

func (o *Chip8) RunOp_SKP_Vx(v Op_SKP_Vx) {
	fmt.Println("NOT IMPLEMENTED")
	// TODO: Implement
}

func (o *Chip8) RunOp_SKNP_Vx(v Op_SKNP_Vx) {
	fmt.Println("NOT IMPLEMENTED")
	// TODO: Implement
}

func (o *Chip8) RunOp_LD_Vx_DT(v Op_LD_Vx_DT) {
	value := o.register.ReadDT()
	o.register.WriteV(v.Vx, value)
}

func (o *Chip8) RunOp_LD_Vx_K(v Op_LD_Vx_K) {
	fmt.Println("NOT IMPLEMENTED")
	// TODO: Implement
}

func (o *Chip8) RunOp_LD_DT_Vx(v Op_LD_DT_Vx) {
	value := o.register.ReadV(v.Vx)
	o.register.WriteDT(value)
}

func (o *Chip8) RunOp_LD_ST_Vx(v Op_LD_ST_Vx) {
	value := o.register.ReadV(v.Vx)
	o.register.WriteST(value)
}

func (o *Chip8) RunOp_ADD_I_Vx(v Op_ADD_I_Vx) {
	valueI := o.register.ReadI()
	valueX := o.register.ReadV(v.Vx)
	value := valueI + uint16(valueX)
	o.register.WriteI(value)
}

func (o *Chip8) RunOp_LD_F_Vx(v Op_LD_F_Vx) {
	fmt.Println("NOT IMPLEMENTED")
	// TODO: Implement
}

func (o *Chip8) RunOp_LD_B_Vx(v Op_LD_B_Vx) {
	fmt.Println("NOT IMPLEMENTED")
	// TODO: Implement
}

func (o *Chip8) RunOp_LD_I_Vx(v Op_LD_I_Vx) {
	valueI := o.register.ReadI()
	for n := uint8(0x00); n < 0x0F; n++ {
		value := o.register.ReadV(n)
		addr := valueI + uint16(n)
		o.memory.Write(addr, value)
	}
}

func (o *Chip8) RunOp_LD_Vx_I(v Op_LD_Vx_I) {
	valueI := o.register.ReadI()
	for n := uint8(0x00); n < 0x0F; n++ {
		addr := valueI + uint16(n)
		value := o.memory.Read8(addr)
		o.register.WriteV(n, value)
	}
}
