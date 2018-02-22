package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"path/filepath"
	"strconv"
	"strings"
)

type SessionController struct {
	ctrl        *CtrlConnManager
	data        *DataConnManager
	basePath    string
	currentPath string
	mode        TransferMode
	done        chan struct{}
	opener      DataConnOpener
	fs          FSManager
}

func NewSessionController(conn net.Conn, basePath string, done chan struct{}) *SessionController {
	ctrl := NewCtrlConnManager(conn)
	ctrl.Run()
	s := &SessionController{
		ctrl:        ctrl,
		basePath:    basePath,
		currentPath: "/",
		mode:        BIN,
		done:        done,
		opener:      DefaultDataConnOpener(),
		fs:          DefaultFS(),
	}
	go func() {
		<-ctrl.done
		s.Close()
	}()
	return s
}

func (s *SessionController) Login() {
	s.SendCtrl(ReadyForUser, "my go ftp server ready (default: user)")

	userseq := strings.Split(s.RecvCtrl(), " ")
	if !strings.EqualFold(userseq[0], "USER") {
		s.SendCtrl(SyntaxError, "Invalid Sequence.\n")
		return
	}

	pass := users[userseq[1]]
	s.SendCtrl(NeedPassword, "Send Password. (default: 1234)")

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

func (s *SessionController) Handle(command string, arg string) error {
	switch command {
	case "SYST":
		if err := s.SendCtrl(SystemType, "UNIX Type: L8"); err != nil {
			s.CloseWithError(err)
			return err
		}
	case "FEAT":
		if err := s.SendCtrl(SystemStatusReply, "End."); err != nil {
			s.CloseWithError(err)
			return err
		}
	case "PWD":
		if err := s.SendCtrl(Created, fmt.Sprintf("\"%s\" is the current directory.", s.CurrentPath())); err != nil {
			s.CloseWithError(err)
			return err
		}
	case "PORT":
		network := strings.Split(arg, ",")
		host := strings.Join(network[0:4], ".")

		base, err := strconv.Atoi(network[4])
		if err != nil {
			s.CloseWithError(err)
			return err
		}
		p, err := strconv.Atoi(network[5])
		if err != nil {
			s.CloseWithError(err)
			return err
		}
		port := base*256 + p
		if err := s.OpenDataConn(host, port); err != nil {
			s.CloseWithError(err)
			return err
		}
		if err := s.SendCtrl(OK, "PORT command successful."); err != nil {
			s.CloseWithError(err)
			return err
		}
	case "LIST":
		if err := s.SendCtrl(TransferStarting, "start."); err != nil {
			s.CloseWithError(err)
			return err
		}
		if err := s.Ls(); err != nil {
			s.CloseWithError(err)
			return err
		}
		if err := s.SendCtrl(CloseDataConnection, "Transfer complete"); err != nil {
			s.CloseWithError(err)
			return err
		}
	case "CWD":
		s.Cd(arg)
		if err := s.SendCtrl(RequestedCompleted, fmt.Sprintf("\"%s\" is the current directory.", s.CurrentPath())); err != nil {
			s.CloseWithError(err)
			return err
		}
	case "TYPE":
		s.SetType(FromCode(arg))
		if err := s.SendCtrl(OK, fmt.Sprintf("Type set to: %s", s.Type())); err != nil {
			s.CloseWithError(err)
			return err
		}
	case "SIZE":
		size, err := s.Size(arg)
		if err != nil {
			s.CloseWithError(err)
			return err
		}
		if err := s.SendCtrl(FileStat, fmt.Sprintf("%d", *size)); err != nil {
			s.CloseWithError(err)
			return err
		}
	case "RETR":
		if err := s.SendCtrl(TransferStarting, "start."); err != nil {
			s.CloseWithError(err)
			return err
		}
		if err := s.SendFile(arg); err != nil {
			s.CloseWithError(err)
			return err
		}
		if err := s.SendCtrl(CloseDataConnection, "Transfer complete"); err != nil {
			s.CloseWithError(err)
			return err
		}
	case "STOR":
		if err := s.SendCtrl(TransferStarting, "start."); err != nil {
			s.CloseWithError(err)
			return err
		}
		if err := s.RecvFile(arg); err != nil {
			s.CloseWithError(err)
			return err
		}
		if err := s.SendCtrl(CloseDataConnection, "Transfer complete"); err != nil {
			s.CloseWithError(err)
			return err
		}
	case "QUIT":
		if err := s.SendCtrl(ServiceClosingControlConnection, "bye"); err != nil {
			s.CloseWithError(err)
			return err
		}
	default:
		if err := s.SendCtrl(NotImplemented, "Not Implemented"); err != nil {
			s.CloseWithError(err)
			return err
		}
	}
	return nil
}

func (s *SessionController) Close() {
	close(s.done)
}

func (s *SessionController) CurrentPath() string {
	return filepath.Join(s.basePath, s.currentPath)
}

func (s *SessionController) OpenDataConn(host string, port int) error {
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("connecting to: %s \n", addr)
	conn, err := s.opener.Open(host, port)
	if err != nil {
		return err
	}
	data := NewDataConnManager(conn)
	data.Run()

	s.data = data

	return nil
}

func (s *SessionController) RecvCtrl() string {
	return s.ctrl.Recv()
}

func (s *SessionController) WaitCtrl() <-chan string {
	return s.ctrl.out
}

func (s *SessionController) SendCtrl(code int, mes string) error {
	return s.ctrl.SendMessage(code, mes)
}

func (s *SessionController) RecvFile(filename string) error {
	fi, err := s.fs.Create(filepath.Join(s.CurrentPath(), filename))
	if err != nil {
		return err
	}
	defer fi.Close()
	s.data.CopyFromConn(fi)

	return nil
}

func (s *SessionController) SendData(r io.Reader) error {
	return s.data.SendBin(r)
}

func (s *SessionController) CloseData() {
	s.data.Close()
}

// TODO add test
func (s *SessionController) Ls() error {
	files, err := ioutil.ReadDir(s.CurrentPath())
	if err != nil {
		return err
	}
	for _, f := range files {
		p := fmt.Sprintf("%s\t%s\t%s\r\n", f.Mode(), f.ModTime(), f.Name())
		if err := s.SendData(bytes.NewBufferString(p)); err != nil {
			return err
		}
	}

	s.CloseData()

	return nil
}

func (s *SessionController) Cd(cwd string) {
	s.currentPath = filepath.Join(s.currentPath, cwd)
	log.Printf("cd: %s \n", s.currentPath)
}

// TODO
func (s *SessionController) SetType(t TransferMode) {
	s.mode = t
}

func (s *SessionController) Type() TransferMode {
	return s.mode
}

func (s *SessionController) Size(filename string) (*int, error) {
	fi, err := s.fs.Stat(filepath.Join(s.CurrentPath(), filename))
	if err != nil {
		return nil, err
	}
	res := int(fi.Size())
	return &res, nil
}

func (s *SessionController) SendFile(filename string) error {
	fi, err := s.fs.Open(filepath.Join(s.CurrentPath(), filename))
	if err != nil {
		return err
	}
	defer fi.Close()
	if err := s.SendData(fi); err != nil {
		return err
	}
	s.CloseData()

	return nil
}

func (s *SessionController) CloseWithError(e error) {
	s.SendCtrl(RequestedActionNotTaken, fmt.Sprintf("Unknown Error: %v", e))
	s.SendCtrl(ServiceClosingControlConnection, "shutdown")
	s.Close()
}
