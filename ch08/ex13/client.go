package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	board chan<- string
	name  string
}

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

	var timeout chan struct{}

	ticker := func(who string) {
		t := time.NewTicker(5 * time.Minute)
		defer t.Stop()
		select {
		case <-timeout:
			return
		case <-t.C:
			log.Printf("client: %s, timeout", who)
			conn.Close()
		}
	}

	go ticker(who)
	for input.Scan() {
		go func() {
			timeout <- struct{}{}
		}()
		messages <- who + ": " + input.Text()
		go ticker(who)
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
