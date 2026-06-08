package cmd

import (
	"archiver/src/lib/compression"
	"archiver/src/lib/compression/algorithms/huffman"
	"archiver/src/lib/compression/algorithms/shennon_fano"
	"fmt"

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

	err = os.WriteFile(unpackedFileName(filepath.Dir(filePath), filePath), []byte(unpacked), 0644)
	if err != nil {
		handleError(err)
	}
}

// unpackedFileName генерирует уникальное имя для распакованного файла.
// dir - папка, в которой проверяется существование файла.
// path - исходный путь к упакованному файлу (например, "file.txt.pack").
func unpackedFileName(dir string, path string) string {
	fileName := filepath.Base(path)
	ext := filepath.Ext(fileName)

	// Получаем имя без крайнего расширения (например, из "file.txt.pack" получим "file.txt")
	baseName := strings.TrimSuffix(fileName, ext)

	// Если у исходного файла не было расширения, можно добавить дефолтное
	if filepath.Ext(baseName) == "" && unpackedDefExtension != "" {
		baseName = baseName + unpackedDefExtension
	}

	// Разделяем базовое имя и его собственное расширение для правильной вставки счетчика
	// Пример для "file.txt": finalExt = ".txt", finalBase = "file"
	finalExt := filepath.Ext(baseName)
	finalBase := strings.TrimSuffix(baseName, finalExt)

	// Формируем начальное имя для проверки
	fullName := baseName
	fullPath := filepath.Join(dir, fullName)

	// Проверяем существование файла на диске
	counter := 1
	for {
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			// Файл не существует, имя свободно
			break
		}

		// Если файл существует, вставляем счетчик ПЕРЕД расширением файла
		// Пример: из "file.txt" делаем "file_1.txt"
		fullName = fmt.Sprintf("%s_%d%s", finalBase, counter, finalExt)
		fullPath = filepath.Join(dir, fullName)
		counter++
	}

	return fullName
}

func init() {
	rootCmd.AddCommand(unpackCmd)

	// also change pack
	unpackCmd.Flags().StringP("method", "m", "", "compression methods: \n\thuffman\n\tshennon_fano")
}
