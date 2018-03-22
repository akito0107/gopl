package tar

import (
	"log"

	"archive/tar"
	"io"

	"os"

	"bytes"

	"bufio"

	"github.com/akito0107/gopl/ch10/ex02/archive"
)

const tarHeader00 = "\x75\x73\x74\x61\x72\x00" // untar.00
const tarHeader20 = "\x75\x73\x74\x61\x72\x20" // untar.

func init() {
	checker := func(r *bufio.Reader) (bool, error) {
		b, err := r.Peek(263)
		if err != nil {
			return false, err
		}
		b = b[len(b)-6:] // magic byteだけ取ってくる
		if string(b) == tarHeader00 || string(b) == tarHeader20 {
			return true, nil
		}
		return false, nil
	}
	archive.RegisterFormat("tar", checker, Decode)
}

type Reader struct {
	inner *tar.Reader
	file  *os.File
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
		if err != nil {
			log.Fatal(err)
		}

		var buf bytes.Buffer
		if _, err := io.Copy(&buf, r.inner); err != nil {
			log.Fatal(err)
		}
		f := &File{header: hdr, body: &buf}
		files = append(files, f)
	}
	return files
}

func (r *Reader) Close() error {
	return r.file.Close()
}

func Decode(filename string) (archive.Reader, error) {
	log.Printf("decoding tar %s\n", filename)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := tar.NewReader(file)
	return &Reader{
		inner: reader,
		file:  file,
	}, nil
}
