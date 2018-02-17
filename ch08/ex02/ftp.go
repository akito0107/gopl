package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var port = flag.Int("port", 8021, "listen port (default 8000)")

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
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

	s := NewSession(c, exPath)
	handleLogin(s)

	for {
		mes := s.RecvCtrl()
		log.Printf(">> incomming message %s \n", mes)
		commands := strings.Split(mes, " ")
		handle(s, commands[0], strings.Join(commands[1:], " "))
	}
}

func handleLogin(s *Session) {
	s.SendCtrl(ReadyForUser, "my go ftp server ready")

	userseq := strings.Split(s.RecvCtrl(), " ")
	if !strings.EqualFold(userseq[0], "USER") {
		s.SendCtrl(SyntaxError, "Invalid Sequence.\n")
		return
	}

	pass := users[userseq[1]]
	s.SendCtrl(NeedPassword, "Send Password.")

	passseq := strings.Split(s.RecvCtrl(), " ")
	if !strings.EqualFold(passseq[0], "PASS") {
		s.SendCtrl(SyntaxError, "Invalid Sequence")
		return
	}

	if pass != passseq[1] {
		s.SendCtrl(NotLoggedIn, "Authentication failed")
		return
	}
	s.SendCtrl(UserLoggedIn, "Login Successful")
}

func handle(s *Session, command string, arg string) {
	switch command {
	case "SYST":
		s.SendCtrl(SystemType, "UNIX Type: L8")
	case "FEAT":
		s.SendCtrl(SystemStatusReply, "End.")
	case "PWD":
		s.SendCtrl(Created, fmt.Sprintf("\"%s\" is the current directory.", s.CurrentPath()))
		//case "PASV":
		//	u := UserSession{host: "127.0.0.1"}
		//	u.Listen()
		//	w.Printf(EnteringPassiveMode, "Entering Passive Mode (%s)\n", u.FormatNetwork())
	case "PORT":
		network := strings.Split(arg, ",")
		host := strings.Join(network[0:4], ".")

		base, err := strconv.Atoi(network[4])
		if err != nil {
			log.Fatal(err)
		}
		p, err := strconv.Atoi(network[5])
		if err != nil {
			log.Fatal(err)
		}
		port := base*256 + p
		s.OpenDataConn(host, port)
		s.SendCtrl(OK, "PORT command successful")
	case "LIST":
		s.SendCtrl(TransferStarting, "start.")
		s.Ls()
		s.SendCtrl(CloseDataConnection, "Transfer complete")
	case "CWD":
		s.Cd(arg)
		s.SendCtrl(RequestedCompleted, fmt.Sprintf("%s is a current directory.", s.CurrentPath()))
	default:
		s.SendCtrl(NotImplemented, "Not Implemented")
	}
}
