package huffman

import (
	"archiver/src/lib/table/prefix_algoritms"
	"reflect"
	"testing"
)

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

func Test_build(t *testing.T) {
	type args struct {
		stat prefix_algoritms.CharStat
	}
	tests := []struct {
		name string
		args args
		want encodingTable
	}{
		{
			name: "case1",
			args: args{
				stat: prefix_algoritms.CharStat{
					'r': 1,
					'!': 1,
					'p': 2,
					'o': 2,
					' ': 2,
					'b': 3,
					'e': 4,
				},
			},
			want: encodingTable{
				'b': code{
					Char:     'b',
					Quantity: 3,
					Bits:     0,
					Size:     2,
				},
				'o': code{
					Char:     'o',
					Quantity: 2,
					Bits:     2,
					Size:     3,
				},
				'e': code{
					Char:     'e',
					Quantity: 4,
					Bits:     3,
					Size:     2,
				},
				'p': code{
					Char:     'p',
					Quantity: 2,
					Bits:     5,
					Size:     3,
				},
				' ': code{
					Char:     ' ',
					Quantity: 2,
					Bits:     3,
					Size:     3,
				},
				'r': code{
					Char:     'r',
					Quantity: 1,
					Bits:     8,
					Size:     4,
				},
				'!': code{
					Char:     '!',
					Quantity: 1,
					Bits:     9,
					Size:     4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := build(tt.args.stat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("build() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_getHuffmanBinTree(t *testing.T) {
//	type args struct {
//		codes []code
//	}
//	tests := []struct {
//		name string
//		args args
//		want binTree
//	}{
//		{
//			name: "case1",
//			args: args{
//				codes: []code{
//					code{
//						Char:     'r',
//						Quantity: 1,
//					},
//					code{
//						Char:     '!',
//						Quantity: 1,
//					},
//					code{
//						Char:     'p',
//						Quantity: 2,
//					},
//					code{
//						Char:     'o',
//						Quantity: 2,
//					},
//					code{
//						Char:     ' ',
//						Quantity: 2,
//					},
//					code{
//						Char:     'b',
//						Quantity: 3,
//					},
//					code{
//						Char:     'e',
//						Quantity: 4,
//					},
//				},
//			},
//			want: binTree{
//				zero: &binTree{
//					zero: &binTree{
//						val: code{
//							Char: 'b',
//						},
//					},
//					one: &binTree{
//						one: &binTree{
//							val: code{
//								Char: ' ',
//							},
//						},
//						zero: &binTree{
//							val: code{
//								Char: 'o',
//							},
//						},
//					},
//				},
//				one: &binTree{
//					one: &binTree{
//						val: code{
//							Char: 'e',
//						},
//					},
//					zero: &binTree{
//						one: &binTree{
//							val: code{
//								Char: 'p',
//							},
//						},
//						zero: &binTree{
//							one: &binTree{
//								val: code{
//									Char: '!',
//								},
//							},
//							zero: &binTree{
//								val: code{
//									Char: 'r',
//								},
//							},
//						},
//					},
//				},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := getHuffmanBinTree(tt.args.codes); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("getHuffmanBinTree() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
