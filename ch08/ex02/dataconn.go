package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

type DataConnOpener interface {
	Open(host string, port int) (net.Conn, error)
}

type dataConnOpener struct{}

func DefaultDataConnOpener() DataConnOpener {
	return &dataConnOpener{}
}

func (dataConnOpener) Open(host string, port int) (net.Conn, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("connecting to: %s \n", addr)
	return net.Dial("tcp", addr)
}

type DataConnManager struct {
	conn      net.Conn
	binIn     chan io.Reader
	ack       chan struct{}
	done      chan struct{}
	errorChan chan error
}

func NewDataConnManager(conn net.Conn) *DataConnManager {
	return &DataConnManager{
		conn:      conn,
		binIn:     make(chan io.Reader),
		ack:       make(chan struct{}),
		done:      make(chan struct{}),
		errorChan: make(chan error),
	}
}

func (c *DataConnManager) SendBin(r io.Reader) error {
	c.binIn <- r
	select {
	case <-c.ack:
		return nil
	case e := <-c.errorChan:
		return e
	}
}

func (c *DataConnManager) CopyFromConn(w io.Writer) error {
	_, err := io.Copy(w, c.conn)
	if err != nil {
		return err
	}
	return nil
}

func (c *DataConnManager) Run() {
	go func() {
		defer c.conn.Close()
		for {
			select {
			case r := <-c.binIn:
				io.Copy(c.conn, r)
				c.ack <- struct{}{}
			case <-c.done:
				return
			}
		}
	}()
}

func (c *DataConnManager) Close() {
	close(c.done)
}
