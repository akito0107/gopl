package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	con := conn.(*net.TCPConn)

	go handleReader(con, done)
	mustCopy(con, os.Stdin)
	con.CloseWrite()
	<-done
}

func handleReader(conn *net.TCPConn, done chan struct{}) {
	defer conn.CloseRead()
	io.Copy(os.Stdout, conn)
	log.Println("done")
	done <- struct{}{}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
