package prefix_algoritms

import (
	"archiver/src/lib/table"
	"fmt"
	"strings"
)

type CharStat map[rune]int
type EncodingTable map[rune]Code

type Code struct {
	Char     rune
	Quantity int
	Bits     uint32
	Size     int
}

func (et EncodingTable) Export() table.EncodingTable {
	res := make(table.EncodingTable)

	for i, v := range et {
		byteString := fmt.Sprintf("%b", v.Bits)

		if lenDiff := v.Size - len(byteString); lenDiff > 0 {
			byteString = strings.Repeat("0", lenDiff) + byteString
		}

		res[i] = byteString
	}

	return res
}

func NewCharStat(text string) CharStat {
	res := CharStat{}

	for _, ch := range text {
		res[ch]++
	}

	return res
}
