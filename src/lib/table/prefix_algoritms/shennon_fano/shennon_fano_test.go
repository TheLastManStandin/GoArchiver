package shennon_fano

import (
	"reflect"
	"testing"
)

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

func Test_assignCodes(t *testing.T) {
	type args struct {
		codes []code
	}
	tests := []struct {
		name string
		args args
		want []code
	}{
		{
			name: "case 1",
			args: args{
				codes: []code{
					code{
						Char:     'a',
						Quantity: 3,
					},
					code{
						Char:     'b',
						Quantity: 2,
					},
					code{
						Char:     'c',
						Quantity: 1,
					},
				},
			},
			want: []code{
				code{
					Char:     'a',
					Quantity: 3,
					Bits:     0,
					Size:     1,
				},
				code{
					Char:     'b',
					Quantity: 2,
					Bits:     2,
					Size:     2,
				},
				code{
					Char:     'c',
					Quantity: 1,
					Bits:     3,
					Size:     2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assignCodes(tt.args.codes)

			if !reflect.DeepEqual(tt.args.codes, tt.want) {
				t.Errorf("assignCodes() = %v, want %v", tt.args.codes, tt.want)
			}
		})
	}
}

func Test_build(t *testing.T) {
	type args struct {
		stat string
	}
	tests := []struct {
		name string
		args args
		want encodingTable
	}{
		{
			name: "case 1",
			args: args{
				stat: "abbbcc",
			},
			want: encodingTable{
				'a': code{
					Char:     'a',
					Quantity: 1,
					Bits:     3,
					Size:     2,
				},
				'b': code{
					Char:     'b',
					Quantity: 3,
					Bits:     0,
					Size:     1,
				},
				'c': code{
					Char:     'c',
					Quantity: 2,
					Bits:     2,
					Size:     2,
				},
			},
		},
		{
			name: "case 2",
			args: args{
				stat: "abbb",
			},
			want: encodingTable{
				'a': code{
					Char:     'a',
					Quantity: 1,
					Bits:     1,
					Size:     1,
				},
				'b': code{
					Char:     'b',
					Quantity: 3,
					Bits:     0,
					Size:     1,
				},
			},
		},
		{
			name: "case 3",
			args: args{
				stat: "abc",
			},
			want: encodingTable{
				'a': code{
					Char:     'a',
					Quantity: 1,
					Bits:     0,
					Size:     1,
				},
				'b': code{
					Char:     'b',
					Quantity: 1,
					Bits:     2,
					Size:     2,
				},
				'c': code{
					Char:     'c',
					Quantity: 1,
					Bits:     3,
					Size:     2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := build(newCharStat(tt.args.stat)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("build() = %v, want %v", got, tt.want)
			}
		})
	}
}
