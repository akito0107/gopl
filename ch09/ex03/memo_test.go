package memo5

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func incomingUrls() chan string {
	urls := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"http://gopl.io",
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"http://gopl.io",
	}

	u := make(chan string)

	go func() {
		for _, url := range urls {
			u <- url
		}
		close(u)
	}()

	return u
}

func TestMemo1(t *testing.T) {
	m := New(httpGetBody, nil)
	var n sync.WaitGroup
	for url := range incomingUrls() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url, nil)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}
