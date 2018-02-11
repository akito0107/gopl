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
		handleConn(conn)
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
	for input.Scan() {
		wg.Add(1)
		go echo(conn, wg, input.Text(), 1*time.Second)
	}
	wg.Wait()
	conn.CloseWrite()
}
