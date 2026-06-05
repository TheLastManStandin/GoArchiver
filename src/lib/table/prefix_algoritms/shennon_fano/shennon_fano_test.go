package shennon_fano

import (
	"archiver/src/lib/table/prefix_algoritms"
	"reflect"
	"testing"
)

func Test_bestDividerPosition(t *testing.T) {
	type args struct {
		codes []prefix_algoritms.Code
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				codes: []prefix_algoritms.Code{
					prefix_algoritms.Code{
						Quantity: 3,
					},
					prefix_algoritms.Code{
						Quantity: 2,
					},
					prefix_algoritms.Code{
						Quantity: 1,
					},
					prefix_algoritms.Code{
						Quantity: 1,
					},
				},
			},
			want: 1,
		},
		{
			name: "case 2 one element",
			args: args{
				codes: []prefix_algoritms.Code{
					prefix_algoritms.Code{
						Quantity: 3,
					},
				},
			},
			want: 0,
		},
		{
			name: "case 3 two element",
			args: args{
				codes: []prefix_algoritms.Code{
					prefix_algoritms.Code{
						Quantity: 3,
					},
					prefix_algoritms.Code{
						Quantity: 3,
					},
				},
			},
			want: 1,
		},
		{
			name: "case 4 three element",
			args: args{
				codes: []prefix_algoritms.Code{
					prefix_algoritms.Code{
						Quantity: 2,
					},
					prefix_algoritms.Code{
						Quantity: 1,
					},
					prefix_algoritms.Code{
						Quantity: 1,
					},
				},
			},
			want: 1,
		},
		{
			name: "case 5",
			args: args{
				codes: []prefix_algoritms.Code{
					prefix_algoritms.Code{
						Quantity: 2,
					},
					prefix_algoritms.Code{
						Quantity: 2,
					},
					prefix_algoritms.Code{
						Quantity: 1,
					},
					prefix_algoritms.Code{
						Quantity: 1,
					},
					prefix_algoritms.Code{
						Quantity: 1,
					},
					prefix_algoritms.Code{
						Quantity: 1,
					},
				},
			},
			want: 2,
		},
		{
			name: "case 5 (need rightmost)",
			args: args{
				codes: []prefix_algoritms.Code{
					prefix_algoritms.Code{
						Quantity: 1,
					},
					prefix_algoritms.Code{
						Quantity: 1,
					},
					prefix_algoritms.Code{
						Quantity: 1,
					},
				},
			},
			want: 1,
		},
		{
			name: "case 6",
			args: args{
				codes: []prefix_algoritms.Code{
					prefix_algoritms.Code{
						Quantity: 2,
					},
					prefix_algoritms.Code{
						Quantity: 2,
					},
					prefix_algoritms.Code{
						Quantity: 1,
					},
					prefix_algoritms.Code{
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
		codes []prefix_algoritms.Code
	}
	tests := []struct {
		name string
		args args
		want []prefix_algoritms.Code
	}{
		{
			name: "case 1",
			args: args{
				codes: []prefix_algoritms.Code{
					prefix_algoritms.Code{
						Char:     'a',
						Quantity: 3,
					},
					prefix_algoritms.Code{
						Char:     'b',
						Quantity: 2,
					},
					prefix_algoritms.Code{
						Char:     'c',
						Quantity: 1,
					},
				},
			},
			want: []prefix_algoritms.Code{
				prefix_algoritms.Code{
					Char:     'a',
					Quantity: 3,
					Bits:     0,
					Size:     1,
				},
				prefix_algoritms.Code{
					Char:     'b',
					Quantity: 2,
					Bits:     2,
					Size:     2,
				},
				prefix_algoritms.Code{
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
		want prefix_algoritms.EncodingTable
	}{
		{
			name: "case 1",
			args: args{
				stat: "abbbcc",
			},
			want: prefix_algoritms.EncodingTable{
				'a': prefix_algoritms.Code{
					Char:     'a',
					Quantity: 1,
					Bits:     3,
					Size:     2,
				},
				'b': prefix_algoritms.Code{
					Char:     'b',
					Quantity: 3,
					Bits:     0,
					Size:     1,
				},
				'c': prefix_algoritms.Code{
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
			want: prefix_algoritms.EncodingTable{
				'a': prefix_algoritms.Code{
					Char:     'a',
					Quantity: 1,
					Bits:     1,
					Size:     1,
				},
				'b': prefix_algoritms.Code{
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
			want: prefix_algoritms.EncodingTable{
				'a': prefix_algoritms.Code{
					Char:     'a',
					Quantity: 1,
					Bits:     0,
					Size:     1,
				},
				'b': prefix_algoritms.Code{
					Char:     'b',
					Quantity: 1,
					Bits:     2,
					Size:     2,
				},
				'c': prefix_algoritms.Code{
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
			if got := build(prefix_algoritms.NewCharStat(tt.args.stat)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("build() = %v, want %v", got, tt.want)
			}
		})
	}
}
