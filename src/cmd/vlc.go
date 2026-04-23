package cmd

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}

const packedExtension = ".vlc"

var ErrEmptyPath = errors.New("please specify a file to pack")

func pack(_ *cobra.Command, args []string) {
	if (len(args) != 1) || (args[0] == "") {
		handleError(ErrEmptyPath)
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleError(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handleError(err)
	}

	//packed := Encode(data)
	packed := string(data) // TODO: remove

	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleError(err)
	}
}

func packedFileName(path string) string {
	fileName := filepath.Base(path)
	ext := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, ext)

	return baseName + packedExtension
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
