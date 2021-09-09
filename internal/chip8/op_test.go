package chip8

import (
	"reflect"
	"testing"
)

func TestGetOpType(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want OpType
	}{
		{
			name: "CLS",
			args: args{
				v: 0x00E0,
			},
			want: OpType_CLS,
		},
		{
			name: "RET",
			args: args{
				v: 0x00EE,
			},
			want: OpType_RET,
		},
		{
			name: "JP addr",
			args: args{
				v: 0x1000,
			},
			want: OpType_JP_addr,
		},
		{
			name: "CALL addr",
			args: args{
				v: 0x2000,
			},
			want: OpType_CALL_addr,
		},
		{
			name: "SE Vx, byte",
			args: args{
				v: 0x3000,
			},
			want: OpType_SE_Vx_byte,
		},
		{
			name: "SNE Vx, byte",
			args: args{
				v: 0x4000,
			},
			want: OpType_SNE_Vx_byte,
		},
		{
			name: "SE Vx, Vy",
			args: args{
				v: 0x5000,
			},
			want: OpType_SE_Vx_Vy,
		},
		{
			name: "LD Vx, byte",
			args: args{
				v: 0x6000,
			},
			want: OpType_LD_Vx_byte,
		},
		{
			name: "ADD Vx, byte",
			args: args{
				v: 0x7000,
			},
			want: OpType_ADD_Vx_byte,
		},
		{
			name: "LD Vx, Vy",
			args: args{
				v: 0x8000,
			},
			want: OpType_LD_Vx_Vy,
		},
		{
			name: "OR Vx, Vy",
			args: args{
				v: 0x8001,
			},
			want: OpType_OR_Vx_Vy,
		},
		{
			name: "AND Vx, Vy",
			args: args{
				v: 0x8002,
			},
			want: OpType_AND_Vx_Vy,
		},
		{
			name: "XOR Vx, Vy",
			args: args{
				v: 0x8003,
			},
			want: OpType_XOR_Vx_Vy,
		},
		{
			name: "ADD Vx, Vy",
			args: args{
				v: 0x8004,
			},
			want: OpType_ADD_Vx_Vy,
		},
		{
			name: "SUB Vx, Vy",
			args: args{
				v: 0x8005,
			},
			want: OpType_SUB_Vx_Vy,
		},
		{
			name: "SHR Vx, Vy",
			args: args{
				v: 0x8006,
			},
			want: OpType_SHR_Vx_Vy,
		},
		{
			name: "SUBN Vx, Vy",
			args: args{
				v: 0x8007,
			},
			want: OpType_SUBN_Vx_Vy,
		},
		{
			name: "SHL Vx, Vy",
			args: args{
				v: 0x800E,
			},
			want: OpType_SHL_Vx_Vy,
		},
		{
			name: "SNE Vx, Vy",
			args: args{
				v: 0x9000,
			},
			want: OpType_SNE_Vx_Vy,
		},
		{
			name: "LD I, addr",
			args: args{
				v: 0xA000,
			},
			want: OpType_LD_I_addr,
		},
		{
			name: "JP V0, addr",
			args: args{
				v: 0xB000,
			},
			want: OpType_JP_V0_addr,
		},
		{
			name: "RND Vx, byte",
			args: args{
				v: 0xC000,
			},
			want: OpType_RND_Vx_byte,
		},
		{
			name: "DRW Vx, Vy, nibble",
			args: args{
				v: 0xD000,
			},
			want: OpType_DRW_Vx_Vy_nibble,
		},
		{
			name: "SKP Vx",
			args: args{
				v: 0xE09E,
			},
			want: OpType_SKP_Vx,
		},
		{
			name: "SKNP Vx",
			args: args{
				v: 0xE0A1,
			},
			want: OpType_SKNP_Vx,
		},
		{
			name: "LD Vx, DT",
			args: args{
				v: 0xF007,
			},
			want: OpType_LD_Vx_DT,
		},
		{
			name: "LD Vx, K",
			args: args{
				v: 0xF00A,
			},
			want: OpType_LD_Vx_K,
		},
		{
			name: "LD DT, Vx",
			args: args{
				v: 0xF015,
			},
			want: OpType_LD_DT_Vx,
		},
		{
			name: "LD ST, Vx",
			args: args{
				v: 0xF018,
			},
			want: OpType_LD_ST_Vx,
		},
		{
			name: "ADD I, Vx",
			args: args{
				v: 0xF01E,
			},
			want: OpType_ADD_I_Vx,
		},
		{
			name: "LD F, Vx",
			args: args{
				v: 0xF029,
			},
			want: OpType_LD_F_Vx,
		},
		{
			name: "LD B, Vx",
			args: args{
				v: 0xF033,
			},
			want: OpType_LD_B_Vx,
		},
		{
			name: "LD [I], Vx",
			args: args{
				v: 0xF055,
			},
			want: OpType_LD_I_Vx,
		},
		{
			name: "LD Vx, [I]",
			args: args{
				v: 0xF065,
			},
			want: OpType_LD_Vx_I,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOpType(tt.args.v); got != tt.want {
				t.Errorf("GetOpType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_CLS(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_CLS
	}{
		{
			name: "success",
			args: args{
				v: 0x00E0,
			},
			want: Op_CLS{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_CLS(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_CLS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_RET(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_RET
	}{
		{
			name: "success",
			args: args{
				v: 0x00EE,
			},
			want: Op_RET{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_RET(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_RET() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_JP_addr(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_JP_addr
	}{
		{
			name: "success",
			args: args{
				v: 0x1123,
			},
			want: Op_JP_addr{
				Addr: 0x0123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_JP_addr(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_JP_addr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_CALL_addr(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_CALL_addr
	}{
		{
			name: "success",
			args: args{
				v: 0x2123,
			},
			want: Op_CALL_addr{
				Addr: 0x0123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_CALL_addr(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_CALL_addr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_SE_Vx_byte(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_SE_Vx_byte
	}{
		{
			name: "success",
			args: args{
				v: 0x3112,
			},
			want: Op_SE_Vx_byte{
				Vx:   0x01,
				Byte: 0x12,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_SE_Vx_byte(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_SE_Vx_byte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_SNE_Vx_byte(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_SNE_Vx_byte
	}{
		{
			name: "success",
			args: args{
				v: 0x4112,
			},
			want: Op_SNE_Vx_byte{
				Vx:   0x01,
				Byte: 0x12,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_SNE_Vx_byte(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_SNE_Vx_byte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_SE_Vx_Vy(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_SE_Vx_Vy
	}{
		{
			name: "success",
			args: args{
				v: 0x5120,
			},
			want: Op_SE_Vx_Vy{
				Vx: 0x01,
				Vy: 0x02,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_SE_Vx_Vy(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_SE_Vx_Vy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_LD_Vx_byte(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_LD_Vx_byte
	}{
		{
			name: "success",
			args: args{
				v: 0x6112,
			},
			want: Op_LD_Vx_byte{
				Vx:   0x01,
				Byte: 0x12,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_LD_Vx_byte(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_LD_Vx_byte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_ADD_Vx_byte(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_ADD_Vx_byte
	}{
		{
			name: "success",
			args: args{
				v: 0x7112,
			},
			want: Op_ADD_Vx_byte{
				Vx:   0x01,
				Byte: 0x12,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_ADD_Vx_byte(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_ADD_Vx_byte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_LD_Vx_Vy(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_LD_Vx_Vy
	}{
		{
			name: "success",
			args: args{
				v: 0x8120,
			},
			want: Op_LD_Vx_Vy{
				Vx: 0x01,
				Vy: 0x02,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_LD_Vx_Vy(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_LD_Vx_Vy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_OR_Vx_Vy(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_OR_Vx_Vy
	}{
		{
			name: "success",
			args: args{
				v: 0x8121,
			},
			want: Op_OR_Vx_Vy{
				Vx: 0x01,
				Vy: 0x02,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_OR_Vx_Vy(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_OR_Vx_Vy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_AND_Vx_Vy(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_AND_Vx_Vy
	}{
		{
			name: "success",
			args: args{
				v: 0x8122,
			},
			want: Op_AND_Vx_Vy{
				Vx: 0x01,
				Vy: 0x02,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_AND_Vx_Vy(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_AND_Vx_Vy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_XOR_Vx_Vy(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_XOR_Vx_Vy
	}{
		{
			name: "success",
			args: args{
				v: 0x8123,
			},
			want: Op_XOR_Vx_Vy{
				Vx: 0x01,
				Vy: 0x02,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_XOR_Vx_Vy(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_XOR_Vx_Vy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_ADD_Vx_Vy(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_ADD_Vx_Vy
	}{
		{
			name: "success",
			args: args{
				v: 0x8124,
			},
			want: Op_ADD_Vx_Vy{
				Vx: 0x01,
				Vy: 0x02,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_ADD_Vx_Vy(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_ADD_Vx_Vy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_SUB_Vx_Vy(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_SUB_Vx_Vy
	}{
		{
			name: "success",
			args: args{
				v: 0x8125,
			},
			want: Op_SUB_Vx_Vy{
				Vx: 0x01,
				Vy: 0x02,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_SUB_Vx_Vy(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_SUB_Vx_Vy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_SHR_Vx_Vy(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_SHR_Vx_Vy
	}{
		{
			name: "success",
			args: args{
				v: 0x8126,
			},
			want: Op_SHR_Vx_Vy{
				Vx: 0x01,
				Vy: 0x02,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_SHR_Vx_Vy(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_SHR_Vx_Vy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_SUBN_Vx_Vy(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_SUBN_Vx_Vy
	}{
		{
			name: "success",
			args: args{
				v: 0x8127,
			},
			want: Op_SUBN_Vx_Vy{
				Vx: 0x01,
				Vy: 0x02,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_SUBN_Vx_Vy(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_SUBN_Vx_Vy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_SHL_Vx_Vy(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_SHL_Vx_Vy
	}{
		{
			name: "success",
			args: args{
				v: 0x812E,
			},
			want: Op_SHL_Vx_Vy{
				Vx: 0x01,
				Vy: 0x02,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_SHL_Vx_Vy(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_SHL_Vx_Vy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_SNE_Vx_Vy(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_SNE_Vx_Vy
	}{
		{
			name: "success",
			args: args{
				v: 0x9120,
			},
			want: Op_SNE_Vx_Vy{
				Vx: 0x01,
				Vy: 0x02,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_SNE_Vx_Vy(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_SNE_Vx_Vy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_LD_I_addr(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_LD_I_addr
	}{
		{
			name: "success",
			args: args{
				v: 0xA123,
			},
			want: Op_LD_I_addr{
				Addr: 0x0123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_LD_I_addr(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_LD_I_addr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_JP_V0_addr(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_JP_V0_addr
	}{
		{
			name: "success",
			args: args{
				v: 0xB123,
			},
			want: Op_JP_V0_addr{
				Addr: 0x0123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_JP_V0_addr(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_JP_V0_addr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_RND_Vx_byte(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_RND_Vx_byte
	}{
		{
			name: "success",
			args: args{
				v: 0xC112,
			},
			want: Op_RND_Vx_byte{
				Vx:   0x01,
				Byte: 0x12,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_RND_Vx_byte(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_RND_Vx_byte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_DRW_Vx_Vy_nibble(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_DRW_Vx_Vy_nibble
	}{
		{
			name: "success",
			args: args{
				v: 0xD121,
			},
			want: Op_DRW_Vx_Vy_nibble{
				Vx:     0x01,
				Vy:     0x02,
				Nibble: 0x01,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_DRW_Vx_Vy_nibble(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_DRW_Vx_Vy_nibble() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_SKP_Vx(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_SKP_Vx
	}{
		{
			name: "success",
			args: args{
				v: 0xE19E,
			},
			want: Op_SKP_Vx{
				Vx: 0x01,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_SKP_Vx(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_SKP_Vx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_SKNP_Vx(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_SKNP_Vx
	}{
		{
			name: "success",
			args: args{
				v: 0xE1A1,
			},
			want: Op_SKNP_Vx{
				Vx: 0x01,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_SKNP_Vx(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_SKNP_Vx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_LD_Vx_DT(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_LD_Vx_DT
	}{
		{
			name: "success",
			args: args{
				v: 0xF107,
			},
			want: Op_LD_Vx_DT{
				Vx: 0x01,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_LD_Vx_DT(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_LD_Vx_DT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_LD_Vx_K(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_LD_Vx_K
	}{
		{
			name: "success",
			args: args{
				v: 0xF10A,
			},
			want: Op_LD_Vx_K{
				Vx: 0x01,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_LD_Vx_K(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_LD_Vx_K() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_LD_DT_Vx(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_LD_DT_Vx
	}{
		{
			name: "success",
			args: args{
				v: 0xF115,
			},
			want: Op_LD_DT_Vx{
				Vx: 0x01,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_LD_DT_Vx(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_LD_DT_Vx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_LD_ST_Vx(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_LD_ST_Vx
	}{
		{
			name: "success",
			args: args{
				v: 0xF118,
			},
			want: Op_LD_ST_Vx{
				Vx: 0x01,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_LD_ST_Vx(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_LD_ST_Vx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_ADD_I_Vx(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_ADD_I_Vx
	}{
		{
			name: "success",
			args: args{
				v: 0xF11E,
			},
			want: Op_ADD_I_Vx{
				Vx: 0x01,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_ADD_I_Vx(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_ADD_I_Vx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_LD_F_Vx(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_LD_F_Vx
	}{
		{
			name: "success",
			args: args{
				v: 0xF129,
			},
			want: Op_LD_F_Vx{
				Vx: 0x01,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_LD_F_Vx(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_LD_F_Vx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_LD_B_Vx(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_LD_B_Vx
	}{
		{
			name: "success",
			args: args{
				v: 0xF133,
			},
			want: Op_LD_B_Vx{
				Vx: 0x01,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_LD_B_Vx(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_LD_B_Vx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_LD_I_Vx(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_LD_I_Vx
	}{
		{
			name: "success",
			args: args{
				v: 0xF155,
			},
			want: Op_LD_I_Vx{
				Vx: 0x01,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_LD_I_Vx(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_LD_I_Vx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOp_LD_Vx_I(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Op_LD_Vx_I
	}{
		{
			name: "success",
			args: args{
				v: 0xF165,
			},
			want: Op_LD_Vx_I{
				Vx: 0x01,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOp_LD_Vx_I(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOp_LD_Vx_I() = %v, want %v", got, tt.want)
			}
		})
	}
}
