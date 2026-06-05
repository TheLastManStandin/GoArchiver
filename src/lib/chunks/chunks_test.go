package chunks

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
			if got := SplitByChunk(tt.args.str, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunk() = %v, want %v", got, tt.want)
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
			if got := tt.chunks.ToMonolithStr(); got != tt.want {
				t.Errorf("ToMonolithStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeStrToBinChunks(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "default",
			args: args{
				data: []byte{
					32, 48, 21, 5, 62,
				},
			},
			want: BinaryChunks{
				"00100000",
				"00110000",
				"00010101",
				"00000101",
				"00111110",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeStrToBinChunks(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeStrToBinChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
