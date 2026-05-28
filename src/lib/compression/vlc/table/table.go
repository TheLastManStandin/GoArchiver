package table

import "strings"

type Generator interface {
	NewTable(text string) EncodingTable
}

type decodingTree struct {
	Val  rune
	Zero *decodingTree
	One  *decodingTree
}

type EncodingTable map[rune]string

func (ec EncodingTable) decodingTree() *decodingTree {
	res := decodingTree{}

	for ch, code := range ec {
		res.add(ch, code)
	}

	return &res
}

func (dt *decodingTree) add(addCh rune, code string) {
	pos := dt
	for _, ch := range code {
		switch ch {
		case '0':
			if pos.Zero == nil {
				pos.Zero = &decodingTree{}
			}
			pos = pos.Zero
		case '1':
			if pos.One == nil {
				pos.One = &decodingTree{}
			}
			pos = pos.One
		}
	}
	pos.Val = addCh
}

func (dt *decodingTree) Decode(str string) string {
	res := strings.Builder{}
	pos := dt

	for _, bite := range str {
		switch bite {
		case '1':
			pos = pos.One
		case '0':
			pos = pos.Zero
		}

		if pos.Val != 0 {
			res.WriteRune(pos.Val)
			pos = dt
		}
	}

	return res.String()
}
