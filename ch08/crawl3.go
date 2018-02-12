package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akito0107/gopl/ch05/links"
)

type linkSet struct {
	depth int
	url   string
}

func main() {
	worklist := make(chan []linkSet)
	unseenLinks := make(chan linkSet)

	go func() {
		for _, u := range os.Args[1:] {
			l := linkSet{
				depth: 1,
				url:   u,
			}
			worklist <- l
		}
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
