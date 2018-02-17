package main

import (
	"fmt"
	"log"
	"net"
)

type Session struct {
	ctrl *ConnManager
	data *ConnManager
}

func NewSession(conn net.Conn) *Session {
	ctrl := NewConnManager(conn)
	ctrl.Run()
	return &Session{
		ctrl: ctrl,
	}
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

func (s *Session) SendData(mes string) {
	s.data.Send(mes)
}
