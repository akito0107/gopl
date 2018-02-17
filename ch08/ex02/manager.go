package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

var users = map[string]string{
	"user": "12345",
}

type ConnManager struct {
	conn  net.Conn
	in    chan string
	binIn chan io.Reader
	out   chan string
	done  chan struct{}
}

func NewConnManager(conn net.Conn) *ConnManager {
	return &ConnManager{
		conn:  conn,
		in:    make(chan string),
		binIn: make(chan io.Reader),
		out:   make(chan string),
		done:  make(chan struct{}),
	}
}

func (c *ConnManager) SendMessage(code int, mes string) {
	c.Send(fmt.Sprintf("%d %s\n", code, mes))
}

func (c *ConnManager) Send(mes string) {
	c.in <- mes
}

func (c *ConnManager) SendBin(r io.Reader) {
	c.binIn <- r
}

func (c *ConnManager) Recv() string {
	return <-c.out
}

func (c *ConnManager) Run() {
	go func() {
		input := bufio.NewScanner(c.conn)
		for input.Scan() {
			c.out <- input.Text()
		}
	}()
	go func() {
		defer c.conn.Close()
		for {
			select {
			case mes := <-c.in:
				fmt.Fprintf(c.conn, mes)
			case r := <-c.binIn:
				io.Copy(c.conn, r)
			case <-c.done:
				return
			}
		}
	}()
}

func (c *ConnManager) Close() {
	close(c.done)
}
