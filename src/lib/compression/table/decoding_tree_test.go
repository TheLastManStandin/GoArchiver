package table

import (
	"testing"
)

func TestDecodingTree_Decode(t *testing.T) {
	type fields struct {
		val  rune
		zero *DecodingTree
		one  *DecodingTree
	}
	type args struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "default",
			fields: fields{
				zero: &DecodingTree{
					One: &DecodingTree{
						Zero: &DecodingTree{
							One: &DecodingTree{
								Val: 'z',
							},
						},
					},
				},
				one: &DecodingTree{
					One: &DecodingTree{
						Val: 'a',
					},
					Zero: &DecodingTree{
						Zero: &DecodingTree{
							One: &DecodingTree{
								Val: 'b',
							},
						},
					},
				},
			},
			args: args{
				str: "010111100111",
			},
			want: "zaba",
		},
		{
			name:   "nothing passed",
			fields: fields{},
			args: args{
				str: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dt := &DecodingTree{
				Val:  tt.fields.val,
				Zero: tt.fields.zero,
				One:  tt.fields.one,
			}
			if got := dt.Decode(tt.args.str); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_encodingTable_DecodeTree(t *testing.T) {
//	tests := []struct {
//		name string
//		ec   encodingTable
//		want decodingTree
//	}{
//		{
//			name: "ez",
//			ec: encodingTable{
//				'a': "01",
//			},
//			want: decodingTree{
//				Zero: &decodingTree{
//					One: &decodingTree{
//						Val: 'a',
//					},
//				},
//			},
//		},
//		{
//			name: "default",
//			ec: encodingTable{
//				'a': "11",
//				'b': "1001",
//				'z': "0101",
//			},
//			want: decodingTree{
//				Zero: &decodingTree{
//					One: &decodingTree{
//						Zero: &decodingTree{
//							One: &decodingTree{
//								Val: 'z',
//							},
//						},
//					},
//				},
//				One: &decodingTree{
//					One: &decodingTree{
//						Val: 'a',
//					},
//					Zero: &decodingTree{
//						Zero: &decodingTree{
//							One: &decodingTree{
//								Val: 'b',
//							},
//						},
//					},
//				},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.ec.decodingTree(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("decodingTree() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
