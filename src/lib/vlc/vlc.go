package vlc

import (
	"strings"
	"unicode"
)

type encodingTable map[rune]string

func Encode(str string) string {
	// prepare text H -> !h
	preparedStr := prepareText(str)

	// encode text to binary: some text -> 10011010010110
	binaryStr := encodeBinary(preparedStr)

	// slice text to 10110011 01011011 10010010 10010110
	chunks := splitByChunk(binaryStr, chunkSize)

	// modify to 3A F0 D3
	hexChunks := chunks.toHex()

	return hexChunks.toStr()
}

func Decode(str string) string {
	binaryString := DecodeStrToHexChunks(str).toBinary().toMonolitStr()

	// actual decoding
	dt := getEncodingTable().DecodeTree()
	decodedStr := dt.Decode(binaryString)

	result := unprepareText(decodedStr)

	return result
}

// Prepares text for encoding: changes upper case letters to:
// ! + lower case letter
// i.g.: Hello WoRld -> !hello !wo!rld
func prepareText(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else if ch == '!' {
			buf.WriteRune('!')
			buf.WriteRune('!')
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}

// Removes escaping from string
// i.g.: !hello !world!! -> Hello World!
func unprepareText(str string) string {
	var buf strings.Builder

	for i := 0; i < len(str); i++ {
		if str[i] == '!' {
			i++
			buf.WriteRune(unicode.ToUpper(rune(str[i])))
		} else {
			buf.WriteRune(rune(str[i]))
		}
	}

	return buf.String()
}

// encodeBinary encodes string to binary codes string
// without spaces.
// i.g.: !hello !wo!rld -> 1101011101010001...
func encodeBinary(str string) string {
	var buf strings.Builder
	tbl := getEncodingTable()

	for _, ch := range str {
		buf.WriteString(bin(ch, tbl))
	}

	return buf.String()
}

func bin(c rune, tbl encodingTable) string {
	res, ok := tbl[c]
	if !ok {
		panic("unknown character: " + string(c))
	}
	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}
