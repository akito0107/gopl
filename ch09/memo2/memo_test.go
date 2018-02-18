package memo2

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestMemo1(t *testing.T) {
	m := New(httpGetBody)
	var n sync.WaitGroup
	incomingUrls := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"http://gopl.io",
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"http://gopl.io",
	}
	for _, url := range incomingUrls {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
		n.Wait()
	}
}
