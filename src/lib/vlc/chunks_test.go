package vlc

import (
	"reflect"
	"testing"
)

func Test_splitByChunk(t *testing.T) {
	type args struct {
		str       string
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "default",
			args: args{
				str:       "100100010010110101010100",
				chunkSize: 8,
			},
			want: BinaryChunks{
				"10010001",
				"00101101",
				"01010100",
			},
		},
		{
			name: "nil",
			args: args{
				str:       "",
				chunkSize: 8,
			},
			want: BinaryChunks{},
		},
		{
			name: "not equal",
			args: args{
				str:       "01101",
				chunkSize: 8,
			},
			want: BinaryChunks{
				"01101000",
			},
		},
		{
			name: "not equal 2",
			args: args{
				str:       "011010101001",
				chunkSize: 8,
			},
			want: BinaryChunks{
				"01101010",
				"10010000",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunk(tt.args.str, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_toHex(t *testing.T) {
	tests := []struct {
		name  string
		chunk BinaryChunks
		want  HexChunks
	}{
		{
			name:  "case 1",
			chunk: BinaryChunks{"10011010", "00000010"},
			want:  HexChunks{"9A", "02"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.chunk.toHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeStrToHexChunks(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want HexChunks
	}{
		{
			name: "default",
			args: args{
				str: "20 E9 24 C7 0C",
			},
			want: HexChunks{
				"20",
				"E9",
				"24",
				"C7",
				"0C",
			},
		},
		{
			name: "default2",
			args: args{
				str: "  20 E9 24   C7 0C ",
			},
			want: HexChunks{
				"20",
				"E9",
				"24",
				"C7",
				"0C",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeStrToHexChunks(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeStrToHexChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunks_toBinary(t *testing.T) {
	tests := []struct {
		name   string
		chunks HexChunks
		want   BinaryChunks
	}{
		{
			name:   "default",
			chunks: HexChunks{"9A", "02"},
			want:   BinaryChunks{"10011010", "00000010"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.chunks.toBinary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_toMonolitStr(t *testing.T) {
	tests := []struct {
		name   string
		chunks BinaryChunks
		want   string
	}{
		{
			name:   "default",
			chunks: BinaryChunks{"10011010", "00000010"},
			want:   "1001101000000010",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.chunks.toMonolitStr(); got != tt.want {
				t.Errorf("toMonolitStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
