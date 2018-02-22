package ex05

import (
	"io"
	"testing"
	"strings"
	"log"
)

func TestLimitReader(t *testing.T) {
	for _, c := range []struct {
		name string
		source string
		limit int64
		bufSize int
	} {
		{
			name: "basic",
			source: "basicsource", // 11
			limit: 5,
			bufSize: 10,
		},
		{
			name: "exceeds buffer",
			source: "basicsource", // 11
			limit: 20,
			bufSize: 10,
		},
		{
			name: "min buf size",
			source: "basicsource", // 11
			limit: 20,
			bufSize: 4,
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			r1 := strings.NewReader(c.source)
			r2 := strings.NewReader(c.source)

			b1 := make([]byte, c.bufSize)
			b2 := make([]byte, c.bufSize)

			l1 := LimitReader(r1, c.limit)
			l2 := io.LimitReader(r2, c.limit)

			n1, err1 := l1.Read(b1)
			n2, err2 := l2.Read(b2)

			log.Printf("%s bytes: %d, %d \n", c.name, n1, n2)
			log.Printf("%s error: %v, %v \n", c.name, err1, err2)

			if n1 != n2 {
				t.Errorf("must be same, %d, %d", n1, n2)
			}

			if err1 != err2 {
				t.Errorf("must be same error, %v, %v", err1, err2)
			}
		})
	}
}

func TestLimitReaderDoubleRead(t *testing.T) {
	for _, c := range []struct {
		name string
		source string
		limit int64
		bufSize1 int
		bufSize2 int
	} {
		{
			name: "basic",
			source: "source-source", // 11
			limit: 15,
			bufSize1: 5,
			bufSize2: 10,
		},
		{
			name: "exceeds buffer",
			source: "source-source", // 11
			limit: 20,
			bufSize1: 20,
			bufSize2: 10,
		},
		{
			name: "min buf size",
			source: "source-source", // 11
			limit: 20,
			bufSize1: 2,
			bufSize2: 2,
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			r1 := strings.NewReader(c.source)
			r2 := strings.NewReader(c.source)

			b11 := make([]byte, c.bufSize1)
			b21 := make([]byte, c.bufSize1)

			b12 := make([]byte, c.bufSize2)
			b22 := make([]byte, c.bufSize2)

			l1 := LimitReader(r1, c.limit)
			l2 := io.LimitReader(r2, c.limit)

			n1, err1 := l1.Read(b11)
			n2, err2 := l2.Read(b21)

			log.Printf("%s bytes: %d, %d \n", c.name, n1, n2)
			log.Printf("%s error: %v, %v \n", c.name, err1, err2)

			if n1 != n2 {
				t.Errorf("must be same, %d, %d", n1, n2)
			}

			if err1 != err2 {
				t.Errorf("must be same error, %v, %v", err1, err2)

			}

			n1, err1 = l1.Read(b12)
			n2, err2 = l2.Read(b22)

		    log.Printf("%s bytes: %d, %d \n", c.name, n1, n2)
			log.Printf("%s error: %v, %v \n", c.name, err1, err2)

			if n1 != n2 {
				t.Errorf("must be same, %d, %d", n1, n2)
			}

			if err1 != err2 {
				t.Errorf("must be same error, %v, %v", err1, err2)
			}

		})
	}
}
