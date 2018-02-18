package ftptest

import (
	"bytes"
	"io"
	"os"
	"time"
)

type MockFS struct {
	OpenMock   func(string) (io.ReadWriteCloser, error)
	StatMock   func(string) (os.FileInfo, error)
	CreateMock func(string) (io.ReadWriteCloser, error)
}

func (m *MockFS) Open(filepath string) (io.ReadWriteCloser, error) {
	return m.OpenMock(filepath)
}

func (m *MockFS) Stat(filepath string) (os.FileInfo, error) {
	return m.StatMock(filepath)
}

func (m *MockFS) Create(filepath string) (io.ReadWriteCloser, error) {
	return m.CreateMock(filepath)
}

type StubFileInfo struct {
	FileName string
	FileSize int64
}

func (m *StubFileInfo) Name() string {
	return m.FileName
}

func (m *StubFileInfo) Size() int64 {
	return m.FileSize
}

func (StubFileInfo) Mode() os.FileMode {
	return os.ModePerm
}

func (StubFileInfo) ModTime() time.Time {
	panic("implement me")
}

func (StubFileInfo) IsDir() bool {
	panic("implement me")
}

func (StubFileInfo) Sys() interface{} {
	panic("implement me")
}

type StubFile struct {
	ReadBuf  *bytes.Buffer
	WriteBuf *bytes.Buffer
}

func (s *StubFile) Read(p []byte) (n int, err error) {
	return s.ReadBuf.Read(p)
}

func (s *StubFile) Write(p []byte) (n int, err error) {
	return s.WriteBuf.Write(p)
}

func (StubFile) Close() error {
	return nil
}
