package archive

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
)

type format struct {
	name, magic string
	decode      func(string) (Reader, error)
}

type Config struct{}

type Reader interface {
	Files() []File
	SetConfig(*Config) error
}

type File interface {
	FileInfo() os.FileInfo
	Open() (io.ReadCloser, error)
}

var formats []format

func RegisterFormat(name, magic string, decode func(string) (Reader, error)) {
	formats = append(formats, format{name, magic, decode})
}

func Decode(filename string) (Reader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)

	for _, format := range formats {
		size := len(format.magic)
		b, err := reader.Peek(size)
		if err != nil {
			return nil, err
		}
		if string(b) == format.magic {
			log.Printf("decoding %s format\n", format.name)
			return format.decode(filename)
		}
	}

	return nil, errors.New("unsupported format")
}
