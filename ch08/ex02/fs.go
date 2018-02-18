package main

import (
	"io"
	"os"
)

type FSManager interface {
	Open(filepath string) (io.ReadWriteCloser, error)
	Stat(filepath string) (os.FileInfo, error)
	Create(filepath string) (io.ReadWriteCloser, error)
}

type fsManager struct{}

func DefaultFS() FSManager {
	return &fsManager{}
}

func (*fsManager) Open(filepath string) (io.ReadWriteCloser, error) {
	return os.Open(filepath)
}

func (*fsManager) Stat(filepath string) (os.FileInfo, error) {
	return os.Stat(filepath)
}

func (*fsManager) Create(filepath string) (io.ReadWriteCloser, error) {
	return os.Create(filepath)
}
