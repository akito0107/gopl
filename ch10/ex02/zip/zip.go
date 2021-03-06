package zip

import (
	"archive/zip"
	"log"

	"io"
	"os"

	"bufio"

	"github.com/akito0107/gopl/ch10/ex02/archive"
)

const zipHeader = "\x50\x4b"

func init() {
	checker := func(r *bufio.Reader) (bool, error) {
		size := len(zipHeader)
		b, err := r.Peek(size)
		if err != nil {
			return false, err
		}
		if string(b) == zipHeader {
			return true, nil
		}
		return false, nil
	}
	archive.RegisterFormat("zip", checker, Decode)
}

type Reader struct {
	inner *zip.ReadCloser
}

type File struct {
	inner *zip.File
}

func (f *File) FileInfo() os.FileInfo {
	return f.inner.FileInfo()
}

func (f *File) Open() (io.Reader, error) {
	return f.inner.Open()
}

func (r *Reader) Files() []archive.File {
	var files []archive.File
	for _, f := range r.inner.File {
		files = append(files, &File{inner: f})
	}
	return files
}

func (r *Reader) Close() error {
	return r.inner.Close()
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
