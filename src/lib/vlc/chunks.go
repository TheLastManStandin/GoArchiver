package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk

type HexChunks []HexChunk

type BinaryChunk string

type HexChunk string

const chunkSize = 8

const hexChunkSeparator = " "

func (chunks BinaryChunks) toMonolitStr() string {
	var monolitStr strings.Builder

	for _, chunk := range chunks {
		monolitStr.WriteString(string(chunk))
	}

	return monolitStr.String()
}

func DecodeStrToHexChunks(str string) HexChunks {
	parts := strings.Split(str, hexChunkSeparator)

	chunks := make(HexChunks, 0, len(parts))

	for _, chunk := range parts {
		if chunk == "" {
			continue
		}
		chunk = strings.TrimSpace(chunk)
		chunks = append(chunks, HexChunk(chunk))
	}

	return chunks
}

func (chunks HexChunks) toBinary() BinaryChunks {
	res := make(BinaryChunks, 0, len(chunks))

	for _, chunk := range chunks {
		res = append(res, chunk.toBinary())
	}

	return res
}

func (chunk HexChunk) toBinary() BinaryChunk {
	val, err := strconv.ParseUint(string(chunk), 16, chunkSize)
	if err != nil {
		panic(err)
	}

	res := fmt.Sprintf("%08b", val)

	return BinaryChunk(res)
}

func (chunks HexChunks) toStr() string {
	switch len(chunks) {
	case 0:
		return ""
	case 1:
		return string(chunks[0])
	}

	res := strings.Builder{}

	res.WriteString(string(chunks[0]))

	for _, chunk := range chunks[1:] {
		res.WriteString(hexChunkSeparator)
		res.WriteString(string(chunk))
	}

	return res.String()
}

func (chunks BinaryChunks) toHex() HexChunks {
	hexChunks := make(HexChunks, 0, len(chunks))

	for _, chunk := range chunks {
		hexChunk := chunk.toHex()
		hexChunks = append(hexChunks, hexChunk)
	}

	return hexChunks
}

func (chunk BinaryChunk) toHex() HexChunk {
	val, err := strconv.ParseUint(string(chunk), 2, chunkSize)
	if err != nil {
		panic(err)
	}

	res := strings.ToUpper(fmt.Sprintf("%x", val))

	if len(res) < 2 {
		res = "0" + res
	}

	return HexChunk(res)
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
