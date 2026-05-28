package shennon_fano

import "testing"

func Test_bestDividerPosition(t *testing.T) {
	type args struct {
		codes []code
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				codes: []code{
					code{
						Quantity: 3,
					},
					code{
						Quantity: 2,
					},
					code{
						Quantity: 1,
					},
					code{
						Quantity: 1,
					},
				},
			},
			want: 1,
		},
		{
			name: "case 2 one element",
			args: args{
				codes: []code{
					code{
						Quantity: 3,
					},
				},
			},
			want: 0,
		},
		{
			name: "case 3 two element",
			args: args{
				codes: []code{
					code{
						Quantity: 3,
					},
					code{
						Quantity: 3,
					},
				},
			},
			want: 1,
		},
		{
			name: "case 4 three element",
			args: args{
				codes: []code{
					code{
						Quantity: 2,
					},
					code{
						Quantity: 1,
					},
					code{
						Quantity: 1,
					},
				},
			},
			want: 1,
		},
		{
			name: "case 5",
			args: args{
				codes: []code{
					code{
						Quantity: 2,
					},
					code{
						Quantity: 2,
					},
					code{
						Quantity: 1,
					},
					code{
						Quantity: 1,
					},
					code{
						Quantity: 1,
					},
					code{
						Quantity: 1,
					},
				},
			},
			want: 2,
		},
		{
			name: "case 5 (need rightmost)",
			args: args{
				codes: []code{
					code{
						Quantity: 1,
					},
					code{
						Quantity: 1,
					},
					code{
						Quantity: 1,
					},
				},
			},
			want: 1,
		},
		{
			name: "case 6",
			args: args{
				codes: []code{
					code{
						Quantity: 2,
					},
					code{
						Quantity: 2,
					},
					code{
						Quantity: 1,
					},
					code{
						Quantity: 1,
					},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestDividerPosition(tt.args.codes); got != tt.want {
				t.Errorf("bestDividerPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
