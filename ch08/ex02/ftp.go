package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
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

type ResponseWriter struct {
	w io.Writer
}

func (w *ResponseWriter) Printf(code int, format string, arg ...interface{}) {
	f := fmt.Sprintf("%d %s", code, format)
	if len(arg) == 0 {
		log.Printf(f)
		fmt.Fprintf(w.w, f)
	} else {
		log.Printf(f, arg)
		log.Printf(f, arg)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	log.Println("connected")

	s := NewSession(c)
	handleLogin(s)

	for input.Scan() {
		t := input.Text()
		log.Printf(">> incomming message %s \n", t)
		commands := strings.Split(t, " ")
		handle(commands[0], strings.Join(commands[1:], " "), w)
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

type UserSession struct {
	host string
	port int
	conn net.Conn
}

// passive mode
func (s *UserSession) Listen() {
	l, err := net.Listen("tcp", s.host+":0")
	if err != nil {
		log.Fatal(err)
	}
	networks := strings.Split(l.Addr().String(), ":")
	s.port, err = strconv.Atoi(networks[1])
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}
		s.conn = conn
		input := bufio.NewScanner(conn)
		for input.Scan() {
			t := input.Text()
			log.Printf(">>> incomming passive message %s \n", t)
		}
	}()
}

func (u *UserSession) Connect(mes <-chan string) {
	addr := fmt.Sprintf("%s:%d", u.host, u.port)
	log.Printf("connecting to: %s \n", addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connected: to %s\n", addr)
	u.conn = conn

	go func() {
		defer conn.Close()
		go func() {
			input := bufio.NewScanner(conn)
			for input.Scan() {
				t := input.Text()
				log.Printf(">>> incomming message %s \n", t)
			}
		}()
		message := <-mes
		log.Println("--------")
		fmt.Fprintf(conn, "message: %s\n", message)
	}()
}

func (s *UserSession) FormatNetwork() string {
	hs := strings.Replace(s.host, ".", ",", 0)
	return fmt.Sprintf("%s,%d", hs, s.port)
}

func handle(command string, arg string, w *ResponseWriter) {
	switch command {
	case "SYST":
		w.Printf(SystemType, "UNIX Type: L8\n")
	case "FEAT":
		w.Printf(SystemStatusReply, "End.\n")
	case "PWD":
		w.Printf(Created, "\"/\" is the current directory.\n")
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

		u := UserSession{
			host: host,
			port: port,
		}
		m := make(chan string)
		u.Connect(m)
		w.Printf(OK, "OK.\n")

		go func() {
			m <- "lslslsls"
		}()

	default:
		w.Printf(NotImplemented, "Not Implemented\n")
	}
}
