package vlc

import (
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
		want string
	}{
		{
			name: "default",
			args: args{
				str: "My name is Ted",
			},
			want: "20 30 3C 18 77 4A E4 4D 28",
		},
		{
			name: "with exclamations",
			args: args{
				str: "Hello my Friends!!!",
			},
			want: "20 E9 24 C7 0C 0E 40 88 4D 81 54 82 08 20 82 00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.str); got != tt.want {
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
				str: "20 E9 24 C7 0C 0E 40 88 4D 81 54 82 08 20 82 00",
			},
			want: "Hello my Friends!!!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.args.str); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
