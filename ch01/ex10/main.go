package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, "/tmp", ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, cacheroot string, ch chan<- string) {
	cachepath := filepath.Join(cacheroot, fmt.Sprintf("%x", sha256.Sum256([]byte(url))))
	if fi, err := os.Stat(cachepath); !os.IsNotExist(err) {
		start := time.Now()
		secs := time.Since(start).Seconds()
		if err != nil {
			ch <- fmt.Sprint(err)
			return
		}
		ch <- fmt.Sprintf("%.2fs %7d %s (from cache)", secs, fi.Size(), url)
	}

	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	out, err := os.Create(cachepath)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(out, resp.Body)
	defer resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
