package main

import (
	"io"
	"net"
)

type DataConnManager struct {
	conn  net.Conn
	binIn chan io.Reader
	ack   chan struct{}
	done  chan struct{}
}

func NewDataConnManager(conn net.Conn) *DataConnManager {
	return &DataConnManager{
		conn:  conn,
		binIn: make(chan io.Reader),
		ack:   make(chan struct{}),
		done:  make(chan struct{}),
	}
}

func (c *DataConnManager) SendBin(r io.Reader) {
	c.binIn <- r
	<-c.ack
}

func (c *DataConnManager) CopyFromConn(w io.Writer) {
	io.Copy(w, c.conn)
}

func (c *DataConnManager) Run() {
	go func() {
		defer c.conn.Close()
		for {
			select {
			case r := <-c.binIn:
				io.Copy(c.conn, r)
			case <-c.done:
				return
			}
			c.ack <- struct{}{}
		}
	}()
}

func (c *DataConnManager) Close() {
	close(c.done)
}
