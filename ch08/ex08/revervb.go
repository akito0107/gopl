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
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	wg.Done()
}

func handleConn(c net.Conn) {
	conn := c.(*net.TCPConn)
	defer conn.CloseRead()
	input := bufio.NewScanner(conn)

	wg := &sync.WaitGroup{}
	ack := make(chan struct{})

	go func(conn *net.TCPConn, wg *sync.WaitGroup, ack chan struct{}) {
		for input.Scan() {
			wg.Add(1)
			ack <- struct{}{}
			go echo(conn, wg, input.Text(), 1*time.Second)
		}
	}(conn, wg, ack)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	wd := false

HANDLER_MAIN:
	for {
		select {
		case <-ticker.C:
			log.Printf("ticker and flag : %t\n", wd)
			if wd {
				wd = false
				continue
			}
			break HANDLER_MAIN
		case <-ack:
			wd = true
			continue
		default:
			continue
		}
	}
	log.Println("no message received close channel")

	wg.Wait()
	conn.CloseWrite()
}
