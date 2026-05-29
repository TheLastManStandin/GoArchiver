package shennon_fano

import (
	"archiver/src/lib/compression"
	"archiver/src/lib/compression/chunks"
	"archiver/src/lib/compression/table"
	"archiver/src/lib/compression/table/shennon_fano"
	"bytes"
	"encoding/binary"
	"encoding/gob"
)

type EncoderDecoder struct {
}

func New() EncoderDecoder {
	return EncoderDecoder{}
}

func (_ EncoderDecoder) Encode(str string) []byte {
	fileTable := shennon_fano.NewGenerator().NewTable(str)

	return buildEncodedFile(str, fileTable)
}

func (_ EncoderDecoder) Decode(encodedData []byte) string {
	parseEncodedData(encodedData)
}

func parseEncodedData(data []byte) {
	
}

func buildEncodedFile(str string, fileTable table.EncodingTable) []byte {
	var buf bytes.Buffer

	encodedTable := encodeTable(fileTable)
	binaryStr := compression.EncodeBinary(str, fileTable)
	binChunks := chunks.SplitByChunk(binaryStr, chunks.ChunkSize)

	buf.Write(encodeInt(len(encodedTable)))
	buf.Write(encodeInt(len(str)))
	buf.Write(encodedTable)
	buf.Write(binChunks.ToBytes())

	return buf.Bytes()
}

func encodeInt(num int) []byte {
	res := make([]byte, 4)

	binary.BigEndian.PutUint32(res, uint32(num))

	return res
}

func encodeTable(fileTable table.EncodingTable) []byte {
	var tableBuf bytes.Buffer

	if err := gob.NewEncoder(&tableBuf).Encode(fileTable); err != nil {
		panic(err)
	}

	return tableBuf.Bytes()
}
