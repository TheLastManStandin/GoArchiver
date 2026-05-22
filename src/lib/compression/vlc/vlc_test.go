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
			want: "!hello, !wor!ld!!",
		},
		{
			name: "only exclamation point",
			str:  "!",
			want: "!!",
		},
		{
			name: "exclamation point and big symbols",
			str:  "!H!E",
			want: "!!!h!!!e",
		},
		{
			name: "exclamations",
			str:  "Hi!!!!",
			want: "!hi!!!!!!!!",
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

func TestEncode(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "default",
			args: args{
				str: "My name is Ted",
			},
			want: []byte{
				0x20,
				0x30,
				0x3C,
				0x18,
				0x77,
				0x4A,
				0xE4,
				0x4D,
				0x28,
			},
		},
		{
			name: "with exclamations",
			args: args{
				str: "Hello my Friends!!!",
			},
			want: []byte{
				0x20,
				0xE9,
				0x24,
				0xC7,
				0x0C,
				0x0E,
				0x40,
				0x88,
				0x4D,
				0x81,
				0x54,
				0x82,
				0x08,
				0x20,
				0x82,
				0x00,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := New()
			if got := encoder.Encode(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unprepareText(t *testing.T) {
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
				str: "!hello my !friends!!!!!!",
			},
			want: "Hello my Friends!!!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unprepareText(tt.args.str); got != tt.want {
				t.Errorf("unprepareText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		encodedText []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				encodedText: []byte{
					0x20,
					0xE9,
					0x24,
					0xC7,
					0x0C,
					0x0E,
					0x40,
					0x88,
					0x4D,
					0x81,
					0x54,
					0x82,
					0x08,
					0x20,
					0x82,
					0x00,
				},
			},
			want: "Hello my Friends!!!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoder := New()
			if got := decoder.Decode(tt.args.encodedText); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
