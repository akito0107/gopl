package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var users = map[string]string{
	"user": "12345",
}

type CtrlConnManager struct {
	conn      net.Conn
	in        chan string
	ack       chan struct{}
	out       chan string
	done      chan struct{}
	errorChan chan error
}

func NewCtrlConnManager(conn net.Conn) *CtrlConnManager {
	return &CtrlConnManager{
		conn:      conn,
		in:        make(chan string),
		ack:       make(chan struct{}),
		out:       make(chan string),
		done:      make(chan struct{}),
		errorChan: make(chan error),
	}
}

func (c *CtrlConnManager) SendMessage(code int, mes string) error {
	return c.Send(fmt.Sprintf("%d %s\n", code, mes))
}

func (c *CtrlConnManager) Send(mes string) error {
	c.in <- mes
	select {
	case <-c.ack:
		return nil
	case e := <-c.errorChan:
		return e
	}
}

func (c *CtrlConnManager) Recv() string {
	return <-c.out
}

func (c *CtrlConnManager) Run() {
	go func() {
		defer c.Close()
		input := bufio.NewScanner(c.conn)
		for input.Scan() {
			c.out <- input.Text()
		}
		log.Println("Going to closing")
	}()
	go func() {
		defer c.conn.Close()
		for {
			select {
			case mes := <-c.in:
				fmt.Fprintf(c.conn, mes)
				c.ack <- struct{}{}
			case <-c.done:
				return
			}
		}
	}()
}

func (c *CtrlConnManager) Close() {
	close(c.done)
}
