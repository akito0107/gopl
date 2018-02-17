package ex03

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"
	"time"
)

func incomingUrls() chan string {
	urls := []string{
		"http://golang.org",
		"http://godoc.org",
		"http://play.golang.org",
		"http://gopl.io",
		"http://golang.org",
		"http://godoc.org",
		"http://play.golang.org",
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

type URLMap struct {
	urls map[string]int
	lock sync.Mutex
}

func NewURLMap() *URLMap {
	return &URLMap{
		urls: make(map[string]int, 4),
	}
}

func (u *URLMap) Add(url string) {
	u.lock.Lock()
	defer u.lock.Unlock()

	u.urls[url] += 1
}

func TestMemo(t *testing.T) {
	u := NewURLMap()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u.Add(r.URL.String())
		fmt.Fprint(w, "got request")
	}))
	defer ts.Close()

	defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
	http.DefaultTransport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
		return url.Parse(ts.URL)
	}
	defer func() {
		http.DefaultTransport.(*http.Transport).Proxy = defaultProxy
	}()

	m := New(httpGetBody, nil)
	var n sync.WaitGroup
	for url := range incomingUrls() {
		n.Add(1)
		start := time.Now()
		done := make(chan struct{})
		go func(url string, done chan struct{}) {
			_, err := m.Get(url, done)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%s, %s\n", url, time.Since(start))
			n.Done()
		}(url, done)
	}
	n.Wait()

	for k, v := range u.urls {
		if v != 1 {
			t.Errorf("must be called once url %s, but %d\n", k, v)
		}
	}
}

func TestMemoCancel(t *testing.T) {
	u := NewURLMap()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.NewTicker(3 * time.Second)
		defer t.Stop()
		<-t.C
		fmt.Fprint(w, "got request")
	}))
	defer ts.Close()

	defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
	http.DefaultTransport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
		return url.Parse(ts.URL)
	}
	defer func() {
		http.DefaultTransport.(*http.Transport).Proxy = defaultProxy
	}()

	m := New(httpGetBody, nil)
	var n sync.WaitGroup
	for url := range incomingUrls() {
		n.Add(1)
		start := time.Now()
		done := make(chan struct{})
		go func(url string, done chan struct{}) {
			_, err := m.Get(url, done)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%s, %s\n", url, time.Since(start))
			n.Done()
		}(url, done)

		go func(done chan struct{}) {
			close(done)
		}(done)
	}
	n.Wait()
	for k, v := range u.urls {
		if v != 2 {
			t.Errorf("must be called uncached url %s, but %d\n", k, v)
		}
	}
}
