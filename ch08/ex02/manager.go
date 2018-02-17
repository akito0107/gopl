package main

import (
	"bufio"
	"fmt"
	"net"
)

var users = map[string]string{
	"user": "12345",
}

type ConnManager struct {
	conn net.Conn
	in   chan string
	out  chan string
}

func NewConnManager(conn net.Conn) *ConnManager {
	return &ConnManager{
		conn: conn,
		in:   make(chan string),
		out:  make(chan string),
	}
}

func (c *ConnManager) SendMessage(code int, mes string) {
	c.Send(fmt.Sprintf("%d %s\n", code, mes))
}

func (c *ConnManager) Send(mes string) {
	c.in <- mes
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
		for {
			select {
			case mes := <-c.in:
				fmt.Fprintf(c.conn, mes)
			}
		}
	}()
}
