package ex05

import (
	"bytes"
	"io"
	"testing"
)

func TestLimitReader(t *testing.T) {
	buf := bytes.NewReader([]byte("abc"))
	l := LimitReader(buf, 2)
	buf1 := make([]byte, 1)
	_, err := l.Read(buf1)
	if err != nil {
		t.Fatal("must be no error")
	}
	buf2 := make([]byte, 1)
	_, err = l.Read(buf2)
	if err != io.EOF {
		t.Fatal("must return io.EOF")
	}
}
