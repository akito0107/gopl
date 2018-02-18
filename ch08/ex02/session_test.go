package main

import (
	"net"
	"testing"

	"bytes"

	"github.com/akito0107/gopl/ch08/ex02/ftptest"
)

func TestSession_Handle_SYST(t *testing.T) {
	block := make(chan struct{})
	conn := &ftptest.StubConn{
		ReadBlock: block,
	}
	done := make(chan struct{})
	s := NewSession(conn, "/", done)
	s.Handle("SYST", "")
	ex := "215 UNIX Type: L8\n"
	close(block)

	if !bytes.Equal(conn.WriteBuf.Bytes(), []byte(ex)) {
		t.Errorf("Must Be equal %s / %s \n", conn.WriteBuf.String(), ex)
	}

}

func TestSession_Handle_FEAT(t *testing.T) {
	block := make(chan struct{})
	conn := &ftptest.StubConn{
		ReadBlock: block,
	}
	done := make(chan struct{})
	s := NewSession(conn, "/", done)
	s.Handle("FEAT", "")
	ex := "211 End.\n"

	close(block)

	if !bytes.Equal(conn.WriteBuf.Bytes(), []byte(ex)) {
		t.Errorf("Must Be equal %s / %s \n", conn.WriteBuf.String(), ex)
	}
}

func TestSession_Handle_CWD_PWD(t *testing.T) {
	t.Run("pwd", func(t *testing.T) {
		block := make(chan struct{})
		conn := &ftptest.StubConn{
			ReadBlock: block,
		}
		done := make(chan struct{})
		s := NewSession(conn, "/hoge", done)
		s.Handle("PWD", "")
		ex := "257 \"/hoge\" is the current directory.\n"

		close(block)

		if !bytes.Equal(conn.WriteBuf.Bytes(), []byte(ex)) {
			t.Errorf("Must Be equal %s / %s \n", conn.WriteBuf.String(), ex)
		}
	})

	t.Run("cd and pwd", func(t *testing.T) {
		block := make(chan struct{})
		conn := &ftptest.StubConn{
			ReadBlock: block,
		}
		done := make(chan struct{})

		s := NewSession(conn, "/hoge", done)
		s.Handle("PWD", "")
		s.Handle("CWD", "/fuga")
		s.Handle("PWD", "")
		ex := "257 \"/hoge\" is the current directory.\n"
		ex = ex + "250 \"/hoge/fuga\" is the current directory.\n"
		ex = ex + "257 \"/hoge/fuga\" is the current directory.\n"

		close(block)

		if !bytes.Equal(conn.WriteBuf.Bytes(), []byte(ex)) {
			t.Errorf("Must Be equal %s / %s \n", conn.WriteBuf.String(), ex)
		}
	})
}

func TestSession_Handle_PORT(t *testing.T) {
	block := make(chan struct{})
	conn := &ftptest.StubConn{
		ReadBlock: block,
	}
	done := make(chan struct{})
	s := NewSession(conn, "/", done)
	mock := &ftptest.MockOpener{}
	called := false
	mock.MockOpen = func(host string, port int) (net.Conn, error) {
		if host != "127.0.0.1" {
			t.Errorf("host must be 127.0.0.1 but %s \n", host)
		}
		if port != 2570 {
			t.Errorf("posrt must be 2570 but %d\n", port)
		}
		called = true
		b := make(chan struct{})
		defer func() {
			close(b)
		}()
		return &ftptest.StubConn{ReadBlock: b}, nil
	}

	s.opener = mock

	s.Handle("PORT", "127,0,0,1,10,10")

	ex := "200 PORT command successful.\n"

	close(block)

	if !bytes.Equal(conn.WriteBuf.Bytes(), []byte(ex)) {
		t.Errorf("Must Be equal %s / %s \n", conn.WriteBuf.String(), ex)
	}
	if !called {
		t.Error("opener must be called.")
	}
}
