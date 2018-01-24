package ex02

import "io"

type counter struct {
	writer io.Writer
	size   int64
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &counter{writer: w, size: 0}
	return c, &c.size
}

func (c *counter) Write(p []byte) (int, error) {
	size, err := c.writer.Write(p)
	if err != nil {
		return size, err
	}
	c.size += int64(size)
	return size, nil
}
