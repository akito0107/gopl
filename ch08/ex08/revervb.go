package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, wg *sync.WaitGroup, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	wg.Done()
}

func timer(ack chan struct{}, c net.Conn) {
	for {
		select {
		case <-time.After(10 * time.Second):
			c.Close()
			return
		case _, ok := <-ack:
			if !ok {
				return
			}
		}
	}
}

func handleConn(c net.Conn) {
	conn := c.(*net.TCPConn)
	defer conn.CloseRead()
	input := bufio.NewScanner(conn)

	wg := &sync.WaitGroup{}
	ack := make(chan struct{})

	go timer(ack, conn)
	for input.Scan() {
		wg.Add(1)
		ack <- struct{}{}
		go echo(conn, wg, input.Text(), 1*time.Second)
	}
	log.Println("no message received close channel")
	close(ack)

	wg.Wait()
	conn.CloseWrite()
}
