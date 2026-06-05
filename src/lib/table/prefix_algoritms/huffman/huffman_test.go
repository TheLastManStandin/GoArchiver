package huffman

import (
	"reflect"
	"testing"
)

func Test_assignCodes(t *testing.T) {
	type args struct {
		codes []code
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				codes: []code{
					{
						Char:     'a',
						Quantity: 3,
					},
					{
						Char:     'b',
						Quantity: 2,
					},
					{
						Char:     'c',
						Quantity: 6,
					},
					{
						Char:     'd',
						Quantity: 1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assignCodes(tt.args.codes)
		})
	}
}

func Test_insertNewBinTree(t *testing.T) {
	type args struct {
		binTrees   []binTree
		newBinTree binTree
	}
	tests := []struct {
		name string
		args args
		want []binTree
	}{
		{
			name: "case1",
			args: args{
				binTrees: []binTree{
					//binTree{
					//	priority: 2,
					//},
					//binTree{
					//	priority: 3,
					//},
					binTree{
						priority: 5,
					},
				},
				newBinTree: binTree{
					priority: 5,
				},
			},
			want: []binTree{
				binTree{
					priority: 5,
				},
				binTree{
					priority: 5,
				},
			},
		},
		{
			name: "case2",
			args: args{
				binTrees: []binTree{
					//binTree{
					//	priority: 5,
					//},
					//binTree{
					//	priority: 14,
					//},
					binTree{
						priority: 29,
					},
					binTree{
						priority: 39,
					},
				},
				newBinTree: binTree{
					priority: 19,
				},
			},
			want: []binTree{
				binTree{
					priority: 19,
				},
				binTree{
					priority: 29,
				},
				binTree{
					priority: 39,
				},
			},
		},
		{
			name: "case3",
			args: args{
				binTrees: []binTree{
					//binTree{
					//	priority: 27,
					//},
					//binTree{
					//	priority: 28,
					//},
					binTree{
						priority: 29,
					},
					binTree{
						priority: 39,
					},
				},
				newBinTree: binTree{
					priority: 55,
				},
			},
			want: []binTree{
				binTree{
					priority: 29,
				},
				binTree{
					priority: 39,
				},
				binTree{
					priority: 55,
				},
			},
		},
		{
			name: "case4",
			args: args{
				binTrees: []binTree{
					//binTree{
					//	priority: 7,
					//},
					//binTree{
					//	priority: 28,
					//},
					binTree{
						priority: 29,
					},
					binTree{
						priority: 39,
					},
				},
				newBinTree: binTree{
					priority: 35,
				},
			},
			want: []binTree{
				binTree{
					priority: 29,
				},
				binTree{
					priority: 35,
				},
				binTree{
					priority: 39,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertNewBinTree(tt.args.binTrees, tt.args.newBinTree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertNewBinTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
