package vlc

import (
	"reflect"
	"testing"
)

func Test_prepareText(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "default",
			str:  "Hello, WorLd",
			want: "!hello, !wor!ld",
		},
		{
			name: "only big symbols",
			str:  "HL",
			want: "!h!l",
		},
		{
			name: "default with exclamation point",
			str:  "Hello, WorLd!",
			want: "!hello, !wor!ld!",
		},
		{
			name: "only exclamation point",
			str:  "!",
			want: "!",
		},
		{
			name: "exclamation point and big symbols",
			str:  "!H!E",
			want: "!!h!!e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareText(tt.str); got != tt.want {
				t.Errorf("prepareText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeBinary(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				str: "!ta q",
			},
			want: "001000100101111000000000001",
		},
		{
			name: "nil",
			args: args{
				str: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBinary(tt.args.str); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
