package bzip

import (
	"io"
	"os/exec"
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

func (w *writer) Write(data []byte) (int, error) {
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

	if err := w.cmd.Start(); err != nil {
		return 0, err
	}
	return cmdin.Write(data)
}


func (w *writer) Close() error {
}

