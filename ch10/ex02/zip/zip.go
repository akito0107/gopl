package zip

import (
	"archive/zip"
	"log"

	"github.com/akito0107/gopl/ch10/ex02/archive"
)

const zipHeader = "\x50\x4b"

func init() {
	archive.RegisterFormat("zip", zipHeader, Decode)
}

type Reader struct {
	inner *zip.ReadCloser
}

func (r *Reader) Files() []archive.File {
	var files []archive.File
	for _, f := range r.inner.File {
		files = append(files, f)
	}
	return files
}

func (Reader) SetConfig(*archive.Config) error {
	return nil
}

func Decode(filename string) (archive.Reader, error) {
	log.Printf("decoding zip %s\n", filename)
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}
	return &Reader{
		inner: reader,
	}, nil
}
