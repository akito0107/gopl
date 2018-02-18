package main

import (
	"net"
	"testing"

	"bytes"

	"os"

	"io"

	"github.com/akito0107/gopl/ch08/ex02/ftptest"
)

func TestSession_Handle(t *testing.T) {
	t.Run("SYST", func(t *testing.T) {
		handleTestHelper(t, "/", "SYST", "", "215 UNIX Type: L8\n")
	})
	t.Run("FEAT", func(t *testing.T) {
		handleTestHelper(t, "/", "FEAT", "", "211 End.\n")
	})
	t.Run("PWD", func(t *testing.T) {
		handleTestHelper(t, "/hoge", "PWD", "/hoge", "257 \"/hoge\" is the current directory.\n")
	})
	t.Run("CWD", func(t *testing.T) {
		s, conn, block := newTestSession("/hoge")
		defer tearDown(block)
		s.Handle("PWD", "")
		s.Handle("CWD", "/fuga")
		s.Handle("PWD", "")
		ex := "257 \"/hoge\" is the current directory.\n"
		ex = ex + "250 \"/hoge/fuga\" is the current directory.\n"
		ex = ex + "257 \"/hoge/fuga\" is the current directory.\n"
		compareMessage(t, conn.WriteBuf.Bytes(), []byte(ex))
	})
	t.Run("PORT", func(t *testing.T) {
		s, conn, block := newTestSession("/hoge")
		defer tearDown(block)
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
		compareMessage(t, conn.WriteBuf.Bytes(), []byte(ex))
		if !called {
			t.Error("opener must be called.")
		}
	})
	t.Run("SIZE", func(t *testing.T) {
		s, conn, block := newTestSession("/hoge")
		defer tearDown(block)
		mock := &ftptest.MockFS{}
		called := false
		mock.StatMock = func(filepath string) (os.FileInfo, error) {
			if filepath != "/hoge/test.txt" {
				t.Errorf("filepath must be /hoge/text.txt but %s", filepath)
			}
			f := &ftptest.StubFileInfo{
				FileName: "test.txt",
				FileSize: 10,
			}
			called = true
			return f, nil
		}
		s.fs = mock

		s.Handle("SIZE", "test.txt")
		ex := "213 10\n"
		compareMessage(t, conn.WriteBuf.Bytes(), []byte(ex))

		if !called {
			t.Error("opener must be called.")
		}
	})
	t.Run("RETR", func(t *testing.T) {
		s, conn, block := newTestSession("/")
		defer tearDown(block)

		f := "test file"
		mf := &ftptest.MockFS{}
		fsCalled := false
		mf.OpenMock = func(filepath string) (io.ReadWriteCloser, error) {
			fsCalled = true
			file := bytes.NewBufferString(f)
			return &ftptest.StubFile{ReadBuf: file}, nil
		}
		s.fs = mf

		dataConn := &ftptest.StubConn{}
		s.data = NewDataConnManager(dataConn)
		s.data.Run()

		s.Handle("RETR", "test.txt")
		ex := "125 start.\n"
		ex += "226 Transfer complete\n"
		compareMessage(t, conn.WriteBuf.Bytes(), []byte(ex))
		compareMessage(t, dataConn.WriteBuf.Bytes(), []byte(f))

		if !fsCalled {
			t.Error("opener must be called.")
		}
	})

	t.Run("STOR", func(t *testing.T) {
		s, conn, block := newTestSession("/hoge")
		defer tearDown(block)

		mf := &ftptest.MockFS{}
		fsCalled := false
		var dist bytes.Buffer
		mf.CreateMock = func(filepath string) (io.ReadWriteCloser, error) {
			if filepath != "/hoge/test.txt" {
				t.Errorf("filepath must be /hoge/test.txt but %s", filepath)
			}
			fsCalled = true
			return &ftptest.StubFile{WriteBuf: &dist}, nil
		}
		s.fs = mf

		f := "test file"
		b := make(chan struct{})
		dataConn := &ftptest.StubConn{
			ReadBuf:   *bytes.NewBufferString(f),
			ReadBlock: b,
		}
		s.data = NewDataConnManager(dataConn)
		s.data.Run()

		close(b)
		s.Handle("STOR", "test.txt")
		ex := "125 start.\n"
		ex += "226 Transfer complete\n"
		compareMessage(t, conn.WriteBuf.Bytes(), []byte(ex))
		compareMessage(t, dist.Bytes(), []byte(f))

		if !fsCalled {
			t.Error("opener must be called.")
		}
	})

}

func newTestSession(basePath string) (*Session, *ftptest.StubConn, chan struct{}) {
	block := make(chan struct{})
	conn := &ftptest.StubConn{
		ReadBlock: block,
	}
	done := make(chan struct{})
	s := NewSession(conn, basePath, done)
	return s, conn, done
}

func compareMessage(t *testing.T, act, ex []byte) {
	t.Helper()
	if !bytes.Equal(ex, act) {
		t.Errorf("Must Be equal %s / %s \n", string(act), string(ex))
	}
}

func tearDown(block chan struct{}) {
	close(block)
}

func handleTestHelper(t *testing.T, basePath, command, arg, ex string) {
	t.Helper()
	s, conn, block := newTestSession(basePath)
	defer tearDown(block)
	s.Handle(command, arg)
	compareMessage(t, conn.WriteBuf.Bytes(), []byte(ex))
}
