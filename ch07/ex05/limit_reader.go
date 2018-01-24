package ex05

import "io"

type limitReader struct {
	inner  io.Reader
	cursor int64
	limit  int64
}

func LimitReader(r io.Reader, limit int64) io.Reader {
	return &limitReader{
		inner:  r,
		cursor: 0,
		limit:  limit,
	}
}

func (l *limitReader) Read(p []byte) (int, error) {
	n, err := l.inner.Read(p)
	l.cursor += int64(n)
	if l.limit <= l.cursor {
		return int(l.cursor), io.EOF
	}
	return n, err
}
