package shennon_fano

import (
	"archiver/src/lib/chunks"
	"archiver/src/lib/compression/algorithms"
	"archiver/src/lib/table/prefix_algoritms/shennon_fano"
)

type EncoderDecoder struct {
}

func New() EncoderDecoder {
	return EncoderDecoder{}
}

func (_ EncoderDecoder) Encode(str string) []byte {
	fileTable := shennon_fano.NewGenerator().NewTable(str)

	return algorithms.BuildEncodedFile(str, fileTable)
}

func (_ EncoderDecoder) Decode(encodedData []byte) string {
	tbl, data := algorithms.ParseEncodedData(encodedData)

	binaryString := chunks.DecodeStrToBinChunks(data).ToMonolithStr()
	decodedStr := tbl.DecodingTree().Decode(binaryString)

	return decodedStr
}
