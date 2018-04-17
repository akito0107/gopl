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
}

func NewWriter(out io.Writer) io.WriteCloser {
	cmd := exec.Command("bzip2")
	cmdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Stdout = out
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	return &writer{out, cmd, cmdin}
}

func (w *writer) Write(data []byte) (int, error) {
	return w.cmdin.Write(data)
}

func (w *writer) Close() (err error) {
	err = w.cmdin.Close()

	return
}
