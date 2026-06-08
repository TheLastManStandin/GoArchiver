package cmd

import (
	"archiver/src/lib/compression"
	"archiver/src/lib/compression/algorithms/huffman"
	"archiver/src/lib/compression/algorithms/shennon_fano"
	"fmt"
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
		encoder = shennon_fano.New()
	case "huffman":
		encoder = huffman.New()
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

	err = os.WriteFile(packedFileName(filepath.Dir(filePath), filePath), packed, 0644)
	if err != nil {
		handleError(err)
	}
}

// packedFileName генерирует уникальное имя файла, добавляя суффикс, если файл уже существует.
// dir - папка, в которой проверяется существование файла.
// path - исходный путь к файлу.
func packedFileName(dir string, path string) string {
	fileName := filepath.Base(path)
	//ext := filepath.Ext(fileName)
	//baseName := strings.TrimSuffix(fileName, ext)

	// Формируем начальное имя: "имя.ext.pack"
	// Если нужно убрать оригинальное расширение, замените fileName на baseName
	fullName := fileName + packedExtension
	fullPath := filepath.Join(dir, fullName)

	// Проверяем, существует ли файл. Если да — ищем свободный номер.
	counter := 1
	for {
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			// Файл не существует, имя свободно
			break
		}

		// Если файл существует, добавляем счетчик перед .pack
		// Пример: "имя.ext_1.pack"
		fullName = fmt.Sprintf("%s_%d%s", fileName, counter, packedExtension)
		fullPath = filepath.Join(dir, fullName)
		counter++
	}

	return fullName
}

func init() {
	rootCmd.AddCommand(packCmd)

	// also change unpack
	packCmd.Flags().StringP("method", "m", "", "compression methods: \n\thuffman\n\tshennon_fano")

	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
