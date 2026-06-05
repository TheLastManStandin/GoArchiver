package cmd

import (
	"archiver/src/lib/compression"
	"archiver/src/lib/compression/algorithms/huffman"
	"archiver/src/lib/compression/algorithms/shennon_fano"

	//"archiver/src/lib/compression/algorithms/vlc"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file",
	Run:   unpack,
}

// TODO: take original extention
const unpackedDefExtension = ".txt"

func unpack(cmd *cobra.Command, args []string) {
	var decoder compression.Decoder
	if (len(args) != 1) || (args[0] == "") {
		handleError(ErrEmptyPath)
	}
	filePath := args[0]

	method := cmd.Flag("method").Value.String()

	switch method {
	//case "vlc":
	//	decoder = vlc.New()
	case "shennon_fano":
		decoder = shennon_fano.New()
	case "huffman":
		decoder = huffman.New()
	default:
		cmd.PrintErr("Unsupported method: " + method + "\n")
	}

	r, err := os.Open(filePath)
	if err != nil {
		handleError(err)
	}
	defer r.Close()

	// TODO: Read file by chunks, not fully
	data, err := io.ReadAll(r)
	if err != nil {
		handleError(err)
	}

	unpacked := decoder.Decode(data)

	err = os.WriteFile(unpackedFileName(filePath), []byte(unpacked), 0644)
	if err != nil {
		handleError(err)
	}
}

func unpackedFileName(path string) string {
	fileName := filepath.Base(path)
	ext := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, ext)

	return baseName
	//+ unpackedDefExtension
}

func init() {
	rootCmd.AddCommand(unpackCmd)

	// also change pack
	unpackCmd.Flags().StringP("method", "m", "", "compression methods: \n\tvlc\n\tshennon_fano")
}
