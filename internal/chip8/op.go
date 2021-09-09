package chip8

type OpType int

const (
	OpType_CLS OpType = iota
	OpType_RET
	OpType_JP_addr
	OpType_CALL_addr
	OpType_SE_Vx_byte
	OpType_SNE_Vx_byte
	OpType_SE_Vx_Vy
	OpType_LD_Vx_byte
	OpType_ADD_Vx_byte
	OpType_LD_Vx_Vy
	OpType_OR_Vx_Vy
	OpType_AND_Vx_Vy
	OpType_XOR_Vx_Vy
	OpType_ADD_Vx_Vy
	OpType_SUB_Vx_Vy
	OpType_SHR_Vx_Vy
	OpType_SUBN_Vx_Vy
	OpType_SHL_Vx_Vy
	OpType_SNE_Vx_Vy
	OpType_LD_I_addr
	OpType_JP_V0_addr
	OpType_RND_Vx_byte
	OpType_DRW_Vx_Vy_nibble
	OpType_SKP_Vx
	OpType_SKNP_Vx
	OpType_LD_Vx_DT
	OpType_LD_Vx_K
	OpType_LD_DT_Vx
	OpType_LD_ST_Vx
	OpType_ADD_I_Vx
	OpType_LD_F_Vx
	OpType_LD_B_Vx
	OpType_LD_I_Vx
	OpType_LD_Vx_I

	OpType_Unknown OpType = -1
)

var (
	OpRaw_CLS              uint16 = 0x00E0
	OpRaw_RET              uint16 = 0x00EE
	OpRaw_JP_addr          uint16 = 0x1000
	OpRaw_CALL_addr        uint16 = 0x2000
	OpRaw_SE_Vx_byte       uint16 = 0x3000
	OpRaw_SNE_Vx_byte      uint16 = 0x4000
	OpRaw_SE_Vx_Vy         uint16 = 0x5000
	OpRaw_LD_Vx_byte       uint16 = 0x6000
	OpRaw_ADD_Vx_byte      uint16 = 0x7000
	OpRaw_LD_Vx_Vy         uint16 = 0x8000
	OpRaw_OR_Vx_Vy         uint16 = 0x8001
	OpRaw_AND_Vx_Vy        uint16 = 0x8002
	OpRaw_XOR_Vx_Vy        uint16 = 0x8003
	OpRaw_ADD_Vx_Vy        uint16 = 0x8004
	OpRaw_SUB_Vx_Vy        uint16 = 0x8005
	OpRaw_SHR_Vx_Vy        uint16 = 0x8006
	OpRaw_SUBN_Vx_Vy       uint16 = 0x8007
	OpRaw_SHL_Vx_Vy        uint16 = 0x800E
	OpRaw_SNE_Vx_Vy        uint16 = 0x9000
	OpRaw_LD_I_addr        uint16 = 0xA000
	OpRaw_JP_V0_addr       uint16 = 0xB000
	OpRaw_RND_Vx_byte      uint16 = 0xC000
	OpRaw_DRW_Vx_Vy_nibble uint16 = 0xD000
	OpRaw_SKP_Vx           uint16 = 0xE09E
	OpRaw_SKNP_Vx          uint16 = 0xE0A1
	OpRaw_LD_Vx_DT         uint16 = 0xF007
	OpRaw_LD_Vx_K          uint16 = 0xF00A
	OpRaw_LD_DT_Vx         uint16 = 0xF015
	OpRaw_LD_ST_Vx         uint16 = 0xF018
	OpRaw_ADD_I_Vx         uint16 = 0xF01E
	OpRaw_LD_F_Vx          uint16 = 0xF029
	OpRaw_LD_B_Vx          uint16 = 0xF033
	OpRaw_LD_I_Vx          uint16 = 0xF055
	OpRaw_LD_Vx_I          uint16 = 0xF065

	OpRaw_Clear_CLS              uint16 = 0x0000
	OpRaw_Clear_RET              uint16 = 0x0000
	OpRaw_Clear_JP_addr          uint16 = 0x0FFF
	OpRaw_Clear_CALL_addr        uint16 = 0x0FFF
	OpRaw_Clear_SE_Vx_byte       uint16 = 0x0FFF
	OpRaw_Clear_SNE_Vx_byte      uint16 = 0x0FFF
	OpRaw_Clear_SE_Vx_Vy         uint16 = 0x0FF0
	OpRaw_Clear_LD_Vx_byte       uint16 = 0x0FFF
	OpRaw_Clear_ADD_Vx_byte      uint16 = 0x0FFF
	OpRaw_Clear_LD_Vx_Vy         uint16 = 0x0FF0
	OpRaw_Clear_OR_Vx_Vy         uint16 = 0x0FF0
	OpRaw_Clear_AND_Vx_Vy        uint16 = 0x0FF0
	OpRaw_Clear_XOR_Vx_Vy        uint16 = 0x0FF0
	OpRaw_Clear_ADD_Vx_Vy        uint16 = 0x0FF0
	OpRaw_Clear_SUB_Vx_Vy        uint16 = 0x0FF0
	OpRaw_Clear_SHR_Vx_Vy        uint16 = 0x0FF0
	OpRaw_Clear_SUBN_Vx_Vy       uint16 = 0x0FF0
	OpRaw_Clear_SHL_Vx_Vy        uint16 = 0x0FF0
	OpRaw_Clear_SNE_Vx_Vy        uint16 = 0x0FF0
	OpRaw_Clear_LD_I_addr        uint16 = 0x0FFF
	OpRaw_Clear_JP_V0_addr       uint16 = 0x0FFF
	OpRaw_Clear_RND_Vx_byte      uint16 = 0x0FFF
	OpRaw_Clear_DRW_Vx_Vy_nibble uint16 = 0x0FFF
	OpRaw_Clear_SKP_Vx           uint16 = 0x0F00
	OpRaw_Clear_SKNP_Vx          uint16 = 0x0F00
	OpRaw_Clear_LD_Vx_DT         uint16 = 0x0F00
	OpRaw_Clear_LD_Vx_K          uint16 = 0x0F00
	OpRaw_Clear_LD_DT_Vx         uint16 = 0x0F00
	OpRaw_Clear_LD_ST_Vx         uint16 = 0x0F00
	OpRaw_Clear_ADD_I_Vx         uint16 = 0x0F00
	OpRaw_Clear_LD_F_Vx          uint16 = 0x0F00
	OpRaw_Clear_LD_B_Vx          uint16 = 0x0F00
	OpRaw_Clear_LD_I_Vx          uint16 = 0x0F00
	OpRaw_Clear_LD_Vx_I          uint16 = 0x0F00

	OpRaw_Zero uint16 = 0x0000
)

type OpTypeMatch struct {
	Raw   uint16
	Clear uint16
	Type  OpType
}

func (o *OpTypeMatch) IsMatch(v uint16) bool {
	return (v &^ o.Clear) == o.Raw
}

var OpTypeMatchSlice = []OpTypeMatch{
	{
		Raw:   OpRaw_CLS,
		Clear: OpRaw_Clear_CLS,
		Type:  OpType_CLS,
	},
	{
		Raw:   OpRaw_RET,
		Clear: OpRaw_Clear_RET,
		Type:  OpType_RET,
	},
	{
		Raw:   OpRaw_JP_addr,
		Clear: OpRaw_Clear_JP_addr,
		Type:  OpType_JP_addr,
	},
	{
		Raw:   OpRaw_CALL_addr,
		Clear: OpRaw_Clear_CALL_addr,
		Type:  OpType_CALL_addr,
	},
	{
		Raw:   OpRaw_SE_Vx_byte,
		Clear: OpRaw_Clear_SE_Vx_byte,
		Type:  OpType_SE_Vx_byte,
	},
	{
		Raw:   OpRaw_SNE_Vx_byte,
		Clear: OpRaw_Clear_SNE_Vx_byte,
		Type:  OpType_SNE_Vx_byte,
	},
	{
		Raw:   OpRaw_SE_Vx_Vy,
		Clear: OpRaw_Clear_SE_Vx_Vy,
		Type:  OpType_SE_Vx_Vy,
	},
	{
		Raw:   OpRaw_LD_Vx_byte,
		Clear: OpRaw_Clear_LD_Vx_byte,
		Type:  OpType_LD_Vx_byte,
	},
	{
		Raw:   OpRaw_ADD_Vx_byte,
		Clear: OpRaw_Clear_ADD_Vx_byte,
		Type:  OpType_ADD_Vx_byte,
	},
	{
		Raw:   OpRaw_LD_Vx_Vy,
		Clear: OpRaw_Clear_LD_Vx_Vy,
		Type:  OpType_LD_Vx_Vy,
	},
	{
		Raw:   OpRaw_OR_Vx_Vy,
		Clear: OpRaw_Clear_OR_Vx_Vy,
		Type:  OpType_OR_Vx_Vy,
	},
	{
		Raw:   OpRaw_AND_Vx_Vy,
		Clear: OpRaw_Clear_AND_Vx_Vy,
		Type:  OpType_AND_Vx_Vy,
	},
	{
		Raw:   OpRaw_XOR_Vx_Vy,
		Clear: OpRaw_Clear_XOR_Vx_Vy,
		Type:  OpType_XOR_Vx_Vy,
	},
	{
		Raw:   OpRaw_ADD_Vx_Vy,
		Clear: OpRaw_Clear_ADD_Vx_Vy,
		Type:  OpType_ADD_Vx_Vy,
	},
	{
		Raw:   OpRaw_SUB_Vx_Vy,
		Clear: OpRaw_Clear_SUB_Vx_Vy,
		Type:  OpType_SUB_Vx_Vy,
	},
	{
		Raw:   OpRaw_SHR_Vx_Vy,
		Clear: OpRaw_Clear_SHR_Vx_Vy,
		Type:  OpType_SHR_Vx_Vy,
	},
	{
		Raw:   OpRaw_SUBN_Vx_Vy,
		Clear: OpRaw_Clear_SUBN_Vx_Vy,
		Type:  OpType_SUBN_Vx_Vy,
	},
	{
		Raw:   OpRaw_SHL_Vx_Vy,
		Clear: OpRaw_Clear_SHL_Vx_Vy,
		Type:  OpType_SHL_Vx_Vy,
	},
	{
		Raw:   OpRaw_SNE_Vx_Vy,
		Clear: OpRaw_Clear_SNE_Vx_Vy,
		Type:  OpType_SNE_Vx_Vy,
	},
	{
		Raw:   OpRaw_LD_I_addr,
		Clear: OpRaw_Clear_LD_I_addr,
		Type:  OpType_LD_I_addr,
	},
	{
		Raw:   OpRaw_JP_V0_addr,
		Clear: OpRaw_Clear_JP_V0_addr,
		Type:  OpType_JP_V0_addr,
	},
	{
		Raw:   OpRaw_RND_Vx_byte,
		Clear: OpRaw_Clear_RND_Vx_byte,
		Type:  OpType_RND_Vx_byte,
	},
	{
		Raw:   OpRaw_DRW_Vx_Vy_nibble,
		Clear: OpRaw_Clear_DRW_Vx_Vy_nibble,
		Type:  OpType_DRW_Vx_Vy_nibble,
	},
	{
		Raw:   OpRaw_SKP_Vx,
		Clear: OpRaw_Clear_SKP_Vx,
		Type:  OpType_SKP_Vx,
	},
	{
		Raw:   OpRaw_SKNP_Vx,
		Clear: OpRaw_Clear_SKNP_Vx,
		Type:  OpType_SKNP_Vx,
	},
	{
		Raw:   OpRaw_LD_Vx_DT,
		Clear: OpRaw_Clear_LD_Vx_DT,
		Type:  OpType_LD_Vx_DT,
	},
	{
		Raw:   OpRaw_LD_Vx_K,
		Clear: OpRaw_Clear_LD_Vx_K,
		Type:  OpType_LD_Vx_K,
	},
	{
		Raw:   OpRaw_LD_DT_Vx,
		Clear: OpRaw_Clear_LD_DT_Vx,
		Type:  OpType_LD_DT_Vx,
	},
	{
		Raw:   OpRaw_LD_ST_Vx,
		Clear: OpRaw_Clear_LD_ST_Vx,
		Type:  OpType_LD_ST_Vx,
	},
	{
		Raw:   OpRaw_ADD_I_Vx,
		Clear: OpRaw_Clear_ADD_I_Vx,
		Type:  OpType_ADD_I_Vx,
	},
	{
		Raw:   OpRaw_LD_F_Vx,
		Clear: OpRaw_Clear_LD_F_Vx,
		Type:  OpType_LD_F_Vx,
	},
	{
		Raw:   OpRaw_LD_B_Vx,
		Clear: OpRaw_Clear_LD_B_Vx,
		Type:  OpType_LD_B_Vx,
	},
	{
		Raw:   OpRaw_LD_I_Vx,
		Clear: OpRaw_Clear_LD_I_Vx,
		Type:  OpType_LD_I_Vx,
	},
	{
		Raw:   OpRaw_LD_Vx_I,
		Clear: OpRaw_Clear_LD_Vx_I,
		Type:  OpType_LD_Vx_I,
	},
}

func GetOpType(v uint16) OpType {
	for _, item := range OpTypeMatchSlice {
		if item.IsMatch(v) {
			return item.Type
		}
	}
	return OpType_Unknown
}

func GetAddr(v uint16) uint16 {
	return v &^ 0xF000
}

func GetNibble(v uint16) uint8 {
	return uint8(v &^ 0xFFF0)
}

func GetVx(v uint16) uint8 {
	return uint8((v &^ 0xF0FF) >> 8)
}

func GetVy(v uint16) uint8 {
	return uint8((v &^ 0xFF0F) >> 4)
}

func GetByte(v uint16) uint8 {
	return uint8(v &^ 0xFF00)
}

type Op_CLS struct{}

func GetOp_CLS(v uint16) Op_CLS {
	return Op_CLS{}
}

type Op_RET struct{}

func GetOp_RET(v uint16) Op_RET {
	return Op_RET{}
}

type Op_JP_addr struct {
	Addr uint16
}

func GetOp_JP_addr(v uint16) Op_JP_addr {
	return Op_JP_addr{
		Addr: GetAddr(v),
	}
}

type Op_CALL_addr struct {
	Addr uint16
}

func GetOp_CALL_addr(v uint16) Op_CALL_addr {
	return Op_CALL_addr{
		Addr: GetAddr(v),
	}
}

type Op_SE_Vx_byte struct {
	Vx   uint8
	Byte uint8
}

func GetOp_SE_Vx_byte(v uint16) Op_SE_Vx_byte {
	return Op_SE_Vx_byte{
		Vx:   GetVx(v),
		Byte: GetByte(v),
	}
}

type Op_SNE_Vx_byte struct {
	Vx   uint8
	Byte uint8
}

func GetOp_SNE_Vx_byte(v uint16) Op_SNE_Vx_byte {
	return Op_SNE_Vx_byte{
		Vx:   GetVx(v),
		Byte: GetByte(v),
	}
}

type Op_SE_Vx_Vy struct {
	Vx uint8
	Vy uint8
}

func GetOp_SE_Vx_Vy(v uint16) Op_SE_Vx_Vy {
	return Op_SE_Vx_Vy{
		Vx: GetVx(v),
		Vy: GetVy(v),
	}
}

type Op_LD_Vx_byte struct {
	Vx   uint8
	Byte uint8
}

func GetOp_LD_Vx_byte(v uint16) Op_LD_Vx_byte {
	return Op_LD_Vx_byte{
		Vx:   GetVx(v),
		Byte: GetByte(v),
	}
}

type Op_ADD_Vx_byte struct {
	Vx   uint8
	Byte uint8
}

func GetOp_ADD_Vx_byte(v uint16) Op_ADD_Vx_byte {
	return Op_ADD_Vx_byte{
		Vx:   GetVx(v),
		Byte: GetByte(v),
	}
}

type Op_LD_Vx_Vy struct {
	Vx uint8
	Vy uint8
}

func GetOp_LD_Vx_Vy(v uint16) Op_LD_Vx_Vy {
	return Op_LD_Vx_Vy{
		Vx: GetVx(v),
		Vy: GetVy(v),
	}
}

type Op_OR_Vx_Vy struct {
	Vx uint8
	Vy uint8
}

func GetOp_OR_Vx_Vy(v uint16) Op_OR_Vx_Vy {
	return Op_OR_Vx_Vy{
		Vx: GetVx(v),
		Vy: GetVy(v),
	}
}

type Op_AND_Vx_Vy struct {
	Vx uint8
	Vy uint8
}

func GetOp_AND_Vx_Vy(v uint16) Op_AND_Vx_Vy {
	return Op_AND_Vx_Vy{
		Vx: GetVx(v),
		Vy: GetVy(v),
	}
}

type Op_XOR_Vx_Vy struct {
	Vx uint8
	Vy uint8
}

func GetOp_XOR_Vx_Vy(v uint16) Op_XOR_Vx_Vy {
	return Op_XOR_Vx_Vy{
		Vx: GetVx(v),
		Vy: GetVy(v),
	}
}

type Op_ADD_Vx_Vy struct {
	Vx uint8
	Vy uint8
}

func GetOp_ADD_Vx_Vy(v uint16) Op_ADD_Vx_Vy {
	return Op_ADD_Vx_Vy{
		Vx: GetVx(v),
		Vy: GetVy(v),
	}
}

type Op_SUB_Vx_Vy struct {
	Vx uint8
	Vy uint8
}

func GetOp_SUB_Vx_Vy(v uint16) Op_SUB_Vx_Vy {
	return Op_SUB_Vx_Vy{
		Vx: GetVx(v),
		Vy: GetVy(v),
	}
}

type Op_SHR_Vx_Vy struct {
	Vx uint8
	Vy uint8
}

func GetOp_SHR_Vx_Vy(v uint16) Op_SHR_Vx_Vy {
	return Op_SHR_Vx_Vy{
		Vx: GetVx(v),
		Vy: GetVy(v),
	}
}

type Op_SUBN_Vx_Vy struct {
	Vx uint8
	Vy uint8
}

func GetOp_SUBN_Vx_Vy(v uint16) Op_SUBN_Vx_Vy {
	return Op_SUBN_Vx_Vy{
		Vx: GetVx(v),
		Vy: GetVy(v),
	}
}

type Op_SHL_Vx_Vy struct {
	Vx uint8
	Vy uint8
}

func GetOp_SHL_Vx_Vy(v uint16) Op_SHL_Vx_Vy {
	return Op_SHL_Vx_Vy{
		Vx: GetVx(v),
		Vy: GetVy(v),
	}
}

type Op_SNE_Vx_Vy struct {
	Vx uint8
	Vy uint8
}

func GetOp_SNE_Vx_Vy(v uint16) Op_SNE_Vx_Vy {
	return Op_SNE_Vx_Vy{
		Vx: GetVx(v),
		Vy: GetVy(v),
	}
}

type Op_LD_I_addr struct {
	Addr uint16
}

func GetOp_LD_I_addr(v uint16) Op_LD_I_addr {
	return Op_LD_I_addr{
		Addr: GetAddr(v),
	}
}

type Op_JP_V0_addr struct {
	Addr uint16
}

func GetOp_JP_V0_addr(v uint16) Op_JP_V0_addr {
	return Op_JP_V0_addr{
		Addr: GetAddr(v),
	}
}

type Op_RND_Vx_byte struct {
	Vx   uint8
	Byte uint8
}

func GetOp_RND_Vx_byte(v uint16) Op_RND_Vx_byte {
	return Op_RND_Vx_byte{
		Vx:   GetVx(v),
		Byte: GetByte(v),
	}
}

type Op_DRW_Vx_Vy_nibble struct {
	Vx     uint8
	Vy     uint8
	Nibble uint8
}

func GetOp_DRW_Vx_Vy_nibble(v uint16) Op_DRW_Vx_Vy_nibble {
	return Op_DRW_Vx_Vy_nibble{
		Vx:     GetVx(v),
		Vy:     GetVy(v),
		Nibble: GetNibble(v),
	}
}

type Op_SKP_Vx struct {
	Vx uint8
}

func GetOp_SKP_Vx(v uint16) Op_SKP_Vx {
	return Op_SKP_Vx{
		Vx: GetVx(v),
	}
}

type Op_SKNP_Vx struct {
	Vx uint8
}

func GetOp_SKNP_Vx(v uint16) Op_SKNP_Vx {
	return Op_SKNP_Vx{
		Vx: GetVx(v),
	}
}

type Op_LD_Vx_DT struct {
	Vx uint8
}

func GetOp_LD_Vx_DT(v uint16) Op_LD_Vx_DT {
	return Op_LD_Vx_DT{
		Vx: GetVx(v),
	}
}

type Op_LD_Vx_K struct {
	Vx uint8
}

func GetOp_LD_Vx_K(v uint16) Op_LD_Vx_K {
	return Op_LD_Vx_K{
		Vx: GetVx(v),
	}
}

type Op_LD_DT_Vx struct {
	Vx uint8
}

func GetOp_LD_DT_Vx(v uint16) Op_LD_DT_Vx {
	return Op_LD_DT_Vx{
		Vx: GetVx(v),
	}
}

type Op_LD_ST_Vx struct {
	Vx uint8
}

func GetOp_LD_ST_Vx(v uint16) Op_LD_ST_Vx {
	return Op_LD_ST_Vx{
		Vx: GetVx(v),
	}
}

type Op_ADD_I_Vx struct {
	Vx uint8
}

func GetOp_ADD_I_Vx(v uint16) Op_ADD_I_Vx {
	return Op_ADD_I_Vx{
		Vx: GetVx(v),
	}
}

type Op_LD_F_Vx struct {
	Vx uint8
}

func GetOp_LD_F_Vx(v uint16) Op_LD_F_Vx {
	return Op_LD_F_Vx{
		Vx: GetVx(v),
	}
}

type Op_LD_B_Vx struct {
	Vx uint8
}

func GetOp_LD_B_Vx(v uint16) Op_LD_B_Vx {
	return Op_LD_B_Vx{
		Vx: GetVx(v),
	}
}

type Op_LD_I_Vx struct {
	Vx uint8
}

func GetOp_LD_I_Vx(v uint16) Op_LD_I_Vx {
	return Op_LD_I_Vx{
		Vx: GetVx(v),
	}
}

type Op_LD_Vx_I struct {
	Vx uint8
}

func GetOp_LD_Vx_I(v uint16) Op_LD_Vx_I {
	return Op_LD_Vx_I{
		Vx: GetVx(v),
	}
}

// TODO: Add Super Chip-48 instructions
