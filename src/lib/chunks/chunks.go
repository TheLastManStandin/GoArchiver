package chunks

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk

type BinaryChunk string

const ChunkSize = 8

func (chunks BinaryChunks) ToMonolithStr() string {
	var monolithStr strings.Builder

	for _, chunk := range chunks {
		monolithStr.WriteString(string(chunk))
	}

	return monolithStr.String()
}

func DecodeStrToBinChunks(data []byte) BinaryChunks {
	chunks := make(BinaryChunks, 0, len(data))

	for _, code := range data {
		chunks = append(chunks, NewBinChunk(code))
	}

	return chunks
}

func NewBinChunk(code byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", code))
}

func (chunks BinaryChunks) ToBytes() []byte {
	res := make([]byte, 0, len(chunks))

	for _, chunk := range chunks {
		res = append(res, chunk.toByte())
	}

	return res
}

func (chunk BinaryChunk) toByte() byte {
	num, err := strconv.ParseUint(string(chunk), 2, ChunkSize)

	if err != nil {
		panic("cant parse bin chunk" + err.Error())
	}

	return byte(num)
}

func SplitByChunk(str string, chunkSize int) BinaryChunks {
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
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if bufLen := len(buf.String()); bufLen != 0 {
		buf.WriteString(strings.Repeat("0", chunkSize-bufLen))
		res = append(res, BinaryChunk(buf.String()))
	}

	return res
}
