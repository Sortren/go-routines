package files

import (
	"log"
	"os"
)

type FileProvider struct {
	filename string
}

func NewFileProvider(filename string) *FileProvider {
	return &FileProvider{filename: filename}
}

func (p FileProvider) File() (*os.File, func(file *os.File)) {
	file, err := os.Open(p.filename)

	closeFunc := func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	return file, closeFunc
}
