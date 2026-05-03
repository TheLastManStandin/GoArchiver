package vlc

import "strings"

type DecodingTree struct {
	val  rune
	zero *DecodingTree
	one  *DecodingTree
}

func (dt *DecodingTree) Decode(str string) string {
	res := strings.Builder{}
	pos := dt

	for _, bite := range str {
		switch bite {
		case '1':
			pos = pos.one
		case '0':
			pos = pos.zero
		}

		if pos.val != 0 {
			res.WriteRune(pos.val)
			pos = dt
		}
	}

	return res.String()
}

func (ec encodingTable) DecodeTree() *DecodingTree {
	res := DecodingTree{}

	for ch, code := range ec {
		res.Add(ch, code)
	}

	return &res
}

func (dt *DecodingTree) Add(addCh rune, code string) {
	pos := dt
	for _, ch := range code {
		switch ch {
		case '0':
			if pos.zero == nil {
				pos.zero = &DecodingTree{}
			}
			pos = pos.zero
		case '1':
			if pos.one == nil {
				pos.one = &DecodingTree{}
			}
			pos = pos.one
		}
	}
	pos.val = addCh
}
