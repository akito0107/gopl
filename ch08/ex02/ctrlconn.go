package main

import (
	"bufio"
	"fmt"
	"net"
)

var users = map[string]string{
	"user": "12345",
}

type CtrlConnManager struct {
	conn net.Conn
	in   chan string
	ack  chan struct{}
	out  chan string
	done chan struct{}
}

func NewCtrlConnManager(conn net.Conn) *CtrlConnManager {
	return &CtrlConnManager{
		conn: conn,
		in:   make(chan string),
		ack:  make(chan struct{}),
		out:  make(chan string),
		done: make(chan struct{}),
	}
}

func (c *CtrlConnManager) SendMessage(code int, mes string) {
	c.Send(fmt.Sprintf("%d %s\n", code, mes))
}

func (c *CtrlConnManager) Send(mes string) {
	c.in <- mes
	<-c.ack
}

func (c *CtrlConnManager) Recv() string {
	return <-c.out
}

func (c *CtrlConnManager) Run() {
	go func() {
		input := bufio.NewScanner(c.conn)
		for input.Scan() {
			c.out <- input.Text()
		}
		c.Close()
	}()
	go func() {
		defer c.conn.Close()
		for {
			select {
			case mes := <-c.in:
				fmt.Fprintf(c.conn, mes)
			case <-c.done:
				return
			}
			c.ack <- struct{}{}
		}
	}()
}

func (c *CtrlConnManager) Close() {
	close(c.done)
}
