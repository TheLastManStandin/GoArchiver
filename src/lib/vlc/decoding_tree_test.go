package vlc

import (
	"reflect"
	"testing"
)

func Test_encodingTable_DecodeTree(t *testing.T) {
	tests := []struct {
		name string
		ec   encodingTable
		want DecodingTree
	}{
		{
			name: "default",
			ec: encodingTable{
				'a': "11",
				'b': "1001",
				'z': "0101",
			},
			want: DecodingTree{
				zero: &DecodingTree{
					one: &DecodingTree{
						zero: &DecodingTree{
							one: &DecodingTree{
								val: 'z',
							},
						},
					},
				},
				one: &DecodingTree{
					one: &DecodingTree{
						val: 'a',
					},
					zero: &DecodingTree{
						zero: &DecodingTree{
							one: &DecodingTree{
								val: 'b',
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ec.DecodeTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
					one: &DecodingTree{
						zero: &DecodingTree{
							one: &DecodingTree{
								val: 'z',
							},
						},
					},
				},
				one: &DecodingTree{
					one: &DecodingTree{
						val: 'a',
					},
					zero: &DecodingTree{
						zero: &DecodingTree{
							one: &DecodingTree{
								val: 'b',
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
				val:  tt.fields.val,
				zero: tt.fields.zero,
				one:  tt.fields.one,
			}
			if got := dt.Decode(tt.args.str); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
