package vlc

import "testing"

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
