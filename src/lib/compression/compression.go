package compression

import (
	"archiver/src/lib/compression/table"
	"strings"
)

type Encoder interface {
	Encode(str string) []byte
}

type Decoder interface {
	Decode(str []byte) string
}

// encodeBinary encodes string to binary codes string
// without spaces.
// i.g.: !hello !wo!rld -> 1101011101010001...
func EncodeBinary(str string, tbl table.EncodingTable) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(bin(ch, tbl))
	}

	return buf.String()
}

func bin(c rune, tbl table.EncodingTable) string {
	res, ok := tbl[c]
	if !ok {
		panic("unknown character: " + string(c))
	}
	return res
}
