package vlc

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type encodingTable map[rune]string

type BinaryChunks []string

type BinaryChunk string

const chunkSize = 8

func Encode(str string) string {
	// prepare text H -> !h
	preparedStr := prepareText(str)

	// encode text to binary: some text -> 10011010010110
	binaryStr := encodeBinary(preparedStr)

	// slise text to 10110011 01011011 10010010 10010110
	chunks := splitByChunk(binaryStr, chunkSize)

	if len(chunks) == 0 { // TODO: deleate
		return ""
	}

	// modifi to 3A F0 D3
	return ""
}

func splitByChunk(str string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(str)
	chunksCount := strLen / chunkSize

	if chunksCount == 0 && len(str) == 0 {
		return BinaryChunks{}
	}

	if strLen%chunkSize != 0 {
		chunksCount++
	}

	res := make(BinaryChunks, 0, chunksCount)
	var buf strings.Builder

	for i, ch := range str {
		buf.WriteRune(ch)
		if (i+1)%chunkSize == 0 {
			res = append(res, buf.String())
			buf.Reset()
		}
	}

	if bufLen := len(buf.String()); bufLen != 0 {
		buf.WriteString(strings.Repeat("0", chunkSize-bufLen))
		res = append(res, buf.String())
	}

	return res
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
		} else {
			buf.WriteRune(ch)
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
