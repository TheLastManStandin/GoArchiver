package cmd

import (
	"archiver/src/lib/compression"
	"archiver/src/lib/compression/algorithms"
	//"archiver/src/lib/compression/algorithms/vlc"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
	Run:   pack,
}

const packedExtension = ".vlc"

var ErrEmptyPath = errors.New("please specify a file to pack")

func pack(cmd *cobra.Command, args []string) {
	var encoder compression.Encoder

	if (len(args) != 1) || (args[0] == "") {
		handleError(ErrEmptyPath)
	}
	filePath := args[0]

	method := cmd.Flag("method").Value.String()

	switch method {
	//case "vlc":
	//	encoder = vlc.New()
	case "shennon_fano":
		encoder = algorithms.New()
	default:
		cmd.PrintErr("Unsupported method: " + method)
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

	packed := encoder.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), packed, 0644)
	if err != nil {
		handleError(err)
	}
}

func packedFileName(path string) string {
	fileName := filepath.Base(path)
	//ext := filepath.Ext(fileName)
	//baseName := strings.TrimSuffix(fileName, ext)
	baseName := fileName

	return baseName + packedExtension
}

func init() {
	rootCmd.AddCommand(packCmd)

	// also change unpack
	packCmd.Flags().StringP("method", "m", "", "compression methods: \n\tvlc\n\tshennon_fano")

	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
