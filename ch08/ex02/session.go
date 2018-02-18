package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
)

type TransferMode int

const (
	ASCII TransferMode = iota
	BIN
	UNKNOWN
)

func (t TransferMode) String() string {
	switch t {
	case ASCII:
		return "Ascii"
	case BIN:
		return "Binary"
	default:
		return "UNKNOWN"
	}
}

func FromCode(code string) TransferMode {
	switch code {
	case "A":
		return ASCII
	case "I":
		return BIN
	default:
		return UNKNOWN
	}
}

type Session struct {
	ctrl        *CtrlConnManager
	data        *DataConnManager
	basePath    string
	currentPath string
	mode        TransferMode
}

func NewSession(conn net.Conn, basePath string) *Session {
	ctrl := NewCtrlConnManager(conn)
	ctrl.Run()
	return &Session{
		ctrl:        ctrl,
		basePath:    basePath,
		currentPath: "/",
		mode:        BIN,
	}
}

func (s *Session) CurrentPath() string {
	return filepath.Join(s.basePath, s.currentPath)
}

func (s *Session) OpenDataConn(host string, port int) error {
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("connecting to: %s \n", addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	data := NewDataConnManager(conn)
	data.Run()

	s.data = data

	return nil
}

func (s *Session) RecvCtrl() string {
	return s.ctrl.Recv()
}

func (s *Session) SendCtrl(code int, mes string) {
	s.ctrl.SendMessage(code, mes)
}

func (s *Session) RecvFile(filename string) {
	fi, err := os.Create(filepath.Join(s.CurrentPath(), filename))
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()
	s.data.CopyFromConn(fi)
}

func (s *Session) SendData(r io.Reader) {
	s.data.SendBin(r)
}

func (s *Session) CloseData() {
	s.data.Close()
}

func (s *Session) Ls() {
	files, err := ioutil.ReadDir(s.CurrentPath())
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		p := fmt.Sprintf("%s\t%s\t%s\n", f.Mode(), f.ModTime(), f.Name())
		s.SendData(bytes.NewBufferString(p))
	}

	s.CloseData()
}

func (s *Session) Cd(cwd string) {
	s.currentPath = filepath.Join(s.currentPath, cwd)
	log.Printf("cd: %s \n", s.currentPath)
}

// TODO
func (s *Session) SetType(t TransferMode) {
	s.mode = t
}

func (s *Session) Type() TransferMode {
	return s.mode
}

func (s *Session) Size(filename string) int {
	fi, err := os.Stat(filepath.Join(s.CurrentPath(), filename))
	if err != nil {
		log.Fatal(err)
	}
	return int(fi.Size())
}

func (s *Session) SendFile(filename string) {
	fi, err := os.Open(filepath.Join(s.CurrentPath(), filename))
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()
	s.SendData(fi)
	s.CloseData()
}
