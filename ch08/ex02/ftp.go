package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

var port = flag.Int("port", 8021, "listen port (default 8000)")

func main() {
	flag.Parse()

	s := &FTPServer{
		host: "localhost",
		port: *port,
		done: make(chan struct{}),
	}

	s.Run()
}

type FTPServer struct {
	host string
	port int
	done chan struct{}
}

func (s *FTPServer) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("port %d listen \n", s.port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
		go func() {
			<-s.done
			return
		}()
	}
}

func (s *FTPServer) Close() {
	close(s.done)
}

func handleConn(c net.Conn) {
	defer c.Close()

	log.Println("connected")
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	log.Println(exPath)

	done := make(chan struct{})
	s := NewSessionController(c, exPath, done)

	if err := s.Login(); err != nil {
		s.SendError(err)
		return
	}

	for {
		select {
		case mes := <-s.WaitCtrl():
			log.Printf(">> incomming message %s \n", mes)
			commands := strings.Split(mes, " ")
			if err := s.Handle(commands[0], strings.Join(commands[1:], " ")); err != nil {
				log.Printf("connection closed with error %v", err)
			}
		case <-s.done:
			log.Println("done received")
			return
		}
	}
}
