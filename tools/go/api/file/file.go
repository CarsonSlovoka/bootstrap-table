package file

import (
	"os"
	"path/filepath"
)

func MustFile(filePath string) *os.File {
	filePath, err := filepath.Abs(filePath)
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		panic(err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	return file
}
