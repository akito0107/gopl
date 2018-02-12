package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"

	c := client{
		board: ch,
		name:  who,
	}
	entering <- c
	input := bufio.NewScanner(conn)

	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- c
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprint(conn, msg)
	}
}
