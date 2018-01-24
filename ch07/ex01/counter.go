package ex01

import (
	"bufio"
	"bytes"
)

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	length := 0
	counter := 0
	scanner := bufio.NewScanner(bytes.NewReader(p))
	for scanner.Scan() {
		length += len([]byte(scanner.Text()))
		counter += 1
	}
	if err := scanner.Err(); err != nil {
		return length, err
	}
	*l += LineCounter(counter)
	return len(p), nil
}

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	length := 0
	counter := 0
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		length += len([]byte(scanner.Text()))
		counter += 1
	}
	if err := scanner.Err(); err != nil {
		return length, err
	}
	*w += WordCounter(counter)
	return len(p), nil
}
