package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("must be passed more 1 locale")
	}
	locales := map[string]string{}
	for _, c := range args {
		locale := strings.Split(c, "=")
		if len(locale) != 2 {
			log.Fatalf("invalid locale sequence %s\n", c)
		}
		locales[locale[0]] = locale[1]
	}

	wg := &sync.WaitGroup{}

	for k, v := range locales {
		conn, err := net.Dial("tcp", v)
		if err != nil {
			log.Fatal(err)
		}
		wg.Add(1)
		go handleConn(k, wg, conn)
	}
	wg.Wait()
}

func handleConn(name string, wg *sync.WaitGroup, c net.Conn) {
	defer func() {
		wg.Done()
		c.Close()
	}()
	writer := &TimeWriter{
		name:  name,
		inner: os.Stdout,
	}
	for {
		mustCopy(writer, c)
	}
}

type TimeWriter struct {
	name  string
	inner io.Writer
}

func (t *TimeWriter) Write(p []byte) (n int, err error) {
	t.inner.Write([]byte(fmt.Sprintf("%s: ", t.name)))
	return t.inner.Write(p)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
