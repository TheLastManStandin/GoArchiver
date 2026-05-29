package vlc

import (
	"archiver/src/lib/compression"
	"archiver/src/lib/compression/chunks"
	"archiver/src/lib/compression/table/vlc"
	"strings"
	"unicode"
)

type EncoderDecoder struct {
}

func New() EncoderDecoder {
	return EncoderDecoder{}
}

func (_ EncoderDecoder) Encode(str string) []byte {
	// prepare text H -> !h
	preparedStr := prepareText(str)

	// encode text to binary: some text -> 10011010010110
	binaryStr := compression.EncodeBinary(preparedStr, vlc.GetEncodingTable())

	// slice text to 10110011 01011011 10010010 10010110
	binChunks := chunks.SplitByChunk(binaryStr, chunks.ChunkSize)

	return binChunks.ToBytes()
}

func (_ EncoderDecoder) Decode(encodedData []byte) string {
	binaryString := chunks.DecodeStrToBinChunks(encodedData).ToMonolithStr()

	// actual decoding
	dt := vlc.GetEncodingTable().DecodingTree()
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
