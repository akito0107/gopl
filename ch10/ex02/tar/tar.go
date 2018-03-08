package tar

import (
	"log"

	"archive/tar"
	"io"

	"os"

	"bytes"

	"github.com/akito0107/gopl/ch10/ex02/archive"
)

const tarHeader = "\x75\x73\x74\x61\x72"

func init() {
	archive.RegisterFormat("tar", tarHeader, Decode)
}

type Reader struct {
	inner *tar.Reader
}

type File struct {
	header *tar.Header
	body   io.Reader
}

func (f *File) FileInfo() os.FileInfo {
	return f.header.FileInfo()
}

func (f *File) Open() (io.Reader, error) {
	return f.body, nil
}

func (r *Reader) Files() []archive.File {
	var files []archive.File
	for {
		hdr, err := r.inner.Next()
		if err == io.EOF {
			break
		}
		var buf *bytes.Buffer
		if _, err := buf.ReadFrom(r.inner); err != nil {
			log.Fatal(err)
		}
		f := &File{header: hdr, body: buf}
		files = append(files, f)
	}
	return files
}

func Decode(filename string) (archive.Reader, error) {
	log.Printf("decoding tar %s\n", filename)
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := tar.NewReader(file)
	return &Reader{
		inner: reader,
	}, nil
}
