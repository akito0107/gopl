package main

import (
	"io"
	"log"
	"os/exec"
)

type writer struct {
	out    io.Writer
	cmd    *exec.Cmd
	cmdin  io.WriteCloser
	cmdout io.ReadCloser
	done   chan struct{}
}

func NewWriter(out io.Writer) io.WriteCloser {
	cmd := exec.Command("bzip2")
	cmdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(out, cmdout)
		close(done)
	}()
	cmd.Start()
	return &writer{out, cmd, cmdin, cmdout, done}
}

func (w *writer) Write(data []byte) (int, error) {
	return w.cmdin.Write(data)
}

func (w *writer) Close() (err error) {
	w.cmdin.Close()
	<-w.done
	w.cmdout.Close()
	return nil
}
