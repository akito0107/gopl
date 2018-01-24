package main

import "io"

type reader struct {
	body   string
	cursor int64
}

func StringReader(str string) io.Reader {
	return &reader{
		body:   str,
		cursor: 0,
	}
}

func (r *reader) Read(p []byte) (int, error) {
	if r.cursor >= int64(len(r.body)) {
		return 0, io.EOF
	}
	n := copy(p, r.body[r.cursor:])
	r.cursor += int64(n)
	return n, nil
}
