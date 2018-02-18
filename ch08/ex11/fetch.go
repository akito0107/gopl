package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

var done = make(chan struct{})

func main() {
	resChan := make(chan struct{}, len(os.Args[1:]))
	wg := &sync.WaitGroup{}
	for _, url := range os.Args[1:] {
		wg.Add(1)
		go fetcher(wg, url, resChan)
	}

	<-resChan
	close(done)
	wg.Wait()

	return
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func fetcher(wg *sync.WaitGroup, url string, resChan chan<- struct{}) {
	defer wg.Done()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("newRequest Failed: %s, error: %v", url, err)
		return
	}

	if cancelled() {
		return
	}

	cancelChan := make(chan struct{})
	req.Cancel = cancelChan

	go func() {
		select {
		case <-done:
			close(cancelChan)
		}
	}()

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Printf("fetch: %v\n", err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Printf("fetch: reading %s: %v\n", url, err)
	}
	log.Printf("%s \n", b)
	resChan <- struct{}{}
}
