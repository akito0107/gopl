package archive

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
)

type format struct {
	name    string
	checker func(*bufio.Reader) (bool, error)
	decode  func(string) (Reader, error)
}

type Config struct{}

type Reader interface {
	Files() []File
	Close() error
}

type File interface {
	FileInfo() os.FileInfo
	Open() (io.Reader, error)
}

var formats []format

func RegisterFormat(name string, checker func(*bufio.Reader) (bool, error), decode func(string) (Reader, error)) {
	formats = append(formats, format{name, checker, decode})
}

func Decode(filename string) (Reader, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)

	for _, format := range formats {
		ok, err := format.checker(reader)
		if err != nil {
			log.Printf("error occured %v", err)
			continue
		}
		if ok {
			log.Printf("decoding %s format\n", format.name)
			return format.decode(filename)
		}
	}

	return nil, errors.New("unsupported format")
}
