package main

import (
	"io"
	"os/exec"
	"fmt"
)

type writer struct {
	out io.Writer
	cmd *exec.Cmd
	cmdin io.WriteCloser
	cmdout io.ReadCloser
}

func NewWriter(out io.Writer) io.WriteCloser {
	cmd := exec.Command("bzip2")
	w := &writer{out, cmd, nil, nil}
	return w
}

func (w *writer) Write(data []byte) (b int, err error) {
	cmdin, err := w.cmd.StdinPipe()
	if err != nil {
		return 0, err
	}
	w.cmdin = cmdin

	cmdout, err := w.cmd.StdoutPipe()
	if err != nil {
		return 0, err
	}
	w.cmdout = cmdout
	defer func() {
		fmt.Println("----------------------")
		w.cmdout.
		_, err = io.Copy(w.out, w.cmdout)
		fmt.Println("----------------------")
	}()

	if err := w.cmd.Start(); err != nil {
		return 0, err
	}
	return cmdin.Write(data)
}


func (w *writer) Close() (err error) {
	if w.cmdin != nil {
		err = w.cmdin.Close()
	}
	if w.cmdout != nil {
		err = w.cmdout.Close()
	}

	return nil
}

