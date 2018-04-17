package main

import (
	"bytes"
	"testing"

	"io"

	"github.com/akito0107/gopl/ch13/bzip/bzip"
)

func TestWriter(t *testing.T) {
	src := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, 
sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut 
aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse 
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, 
sunt in culpa qui officia deserunt mollit anim id est laborum.`

	var out1 bytes.Buffer
	var out2 bytes.Buffer
	buf1 := bytes.NewBufferString(src)
	buf2 := bytes.NewBufferString(src)

	w1 := NewWriter(&out1)
	io.Copy(w1, buf1)
	w1.Close()

	w2 := bzip.NewWriter(&out2)
	io.Copy(w2, buf2)
	w2.Close()

	if bytes.Compare(out1.Bytes(), out2.Bytes()) != 0 {
		t.Errorf("must be same value")
	}
}
