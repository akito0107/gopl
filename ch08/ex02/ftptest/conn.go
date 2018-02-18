package ftptest

import (
	"bytes"
	"net"
	"time"
)

type StubConn struct {
	ReadBuf   bytes.Buffer
	WriteBuf  bytes.Buffer
	ReadBlock chan struct{}
}

func (m *StubConn) Read(b []byte) (n int, err error) {
	<-m.ReadBlock // workaround
	return m.ReadBuf.Read(b)
}

func (m *StubConn) Write(b []byte) (n int, err error) {
	return m.WriteBuf.Write(b)
}

func (StubConn) Close() error {
	return nil
}

func (StubConn) LocalAddr() net.Addr {
	panic("implement me")
}

func (StubConn) RemoteAddr() net.Addr {
	panic("implement me")
}

func (StubConn) SetDeadline(t time.Time) error {
	panic("implement me")
}

func (StubConn) SetReadDeadline(t time.Time) error {
	panic("implement me")
}

func (StubConn) SetWriteDeadline(t time.Time) error {
	panic("implement me")
}

type MockOpener struct {
	MockOpen func(host string, port int) (net.Conn, error)
}

func (m *MockOpener) Open(host string, port int) (net.Conn, error) {
	return m.MockOpen(host, port)
}
