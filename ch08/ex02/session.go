package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"path/filepath"
)

type Session struct {
	ctrl        *ConnManager
	data        *ConnManager
	basePath    string
	currentPath string
}

func NewSession(conn net.Conn, basePath string) *Session {
	ctrl := NewConnManager(conn)
	ctrl.Run()
	return &Session{
		ctrl:        ctrl,
		basePath:    basePath,
		currentPath: "/",
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
	data := NewConnManager(conn)
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

func (s *Session) RecvData() string {
	return s.data.Recv()
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
