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
	"strconv"
	"strings"
)

type Session struct {
	ctrl        *CtrlConnManager
	data        *DataConnManager
	basePath    string
	currentPath string
	mode        TransferMode
	done        chan struct{}
}

func NewSession(conn net.Conn, basePath string, done chan struct{}) *Session {
	ctrl := NewCtrlConnManager(conn)
	ctrl.Run()
	s := &Session{
		ctrl:        ctrl,
		basePath:    basePath,
		currentPath: "/",
		mode:        BIN,
		done:        done,
	}
	go func() {
		<-ctrl.done
		s.Close()
	}()
	return s
}

func (s *Session) Login() {
	s.SendCtrl(ReadyForUser, "my go ftp server ready")

	userseq := strings.Split(s.RecvCtrl(), " ")
	if !strings.EqualFold(userseq[0], "USER") {
		s.SendCtrl(SyntaxError, "Invalid Sequence.\n")
		return
	}

	pass := users[userseq[1]]
	s.SendCtrl(NeedPassword, "Send Password.")

	passseq := strings.Split(s.RecvCtrl(), " ")
	if !strings.EqualFold(passseq[0], "PASS") {
		s.SendCtrl(SyntaxError, "Invalid Sequence")
		return
	}

	if pass != passseq[1] {
		s.SendCtrl(NotLoggedIn, "Authentication failed")
		return
	}
	s.SendCtrl(UserLoggedIn, "Login Successful")
}

func (s *Session) Handle(command string, arg string) error {
	switch command {
	case "SYST":
		s.SendCtrl(SystemType, "UNIX Type: L8")
	case "FEAT":
		s.SendCtrl(SystemStatusReply, "End.")
	case "PWD":
		s.SendCtrl(Created, fmt.Sprintf("\"%s\" is the current directory.", s.CurrentPath()))
	case "PORT":
		network := strings.Split(arg, ",")
		host := strings.Join(network[0:4], ".")

		base, err := strconv.Atoi(network[4])
		if err != nil {
			log.Fatal(err)
		}
		p, err := strconv.Atoi(network[5])
		if err != nil {
			log.Fatal(err)
		}
		port := base*256 + p
		s.OpenDataConn(host, port)
		s.SendCtrl(OK, "PORT command successful")
	case "LIST":
		s.SendCtrl(TransferStarting, "start.")
		s.Ls()
		s.SendCtrl(CloseDataConnection, "Transfer complete")
	case "CWD":
		s.Cd(arg)
		s.SendCtrl(RequestedCompleted, fmt.Sprintf("%s is a current directory.", s.CurrentPath()))
	case "TYPE":
		s.SetType(FromCode(arg))
		s.SendCtrl(OK, fmt.Sprintf("Type set to: %s", s.Type()))
	case "SIZE":
		size := s.Size(arg)
		s.SendCtrl(FileStat, fmt.Sprintf("%d", size))
	case "RETR":
		s.SendCtrl(TransferStarting, "start.")
		s.SendFile(arg)
		s.SendCtrl(CloseDataConnection, "Transfer complete")
	case "STOR":
		s.SendCtrl(TransferStarting, "start.")
		s.RecvFile(arg)
		s.SendCtrl(CloseDataConnection, "Transfer complete")
	case "QUIT":
		s.SendCtrl(ServiceClosingControlConnection, "bye")
	default:
		s.SendCtrl(NotImplemented, "Not Implemented")
	}
	return nil
}

func (s *Session) Close() {
	close(s.done)
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

func (s *Session) WaitCtrl() <-chan string {
	return s.ctrl.out
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
