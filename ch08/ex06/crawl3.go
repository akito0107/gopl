package main

import (
	"flag"
	"log"
	"os"

	"github.com/akito0107/gopl/ch05/links"
)

type linkSet struct {
	depth int
	url   string
}

var tokens = make(chan struct{}, 20)

var limit = flag.Int("depth", 2, "crawl depth (default 2)")

func main() {
	flag.Parse()
	worklist := make(chan []linkSet)
	var n int

	n++
	go func() {
		var list []linkSet
		for _, u := range os.Args[1:] {
			l := linkSet{
				depth: 1,
				url:   u,
			}
			list = append(list, l)
		}
		worklist <- list
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if link.depth > *limit {
				continue
			}
			if !seen[link.url] {
				seen[link.url] = true
				n++
				go func(link linkSet) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(link linkSet) []linkSet {
	log.Printf("depth: %d url: %s \n", link.depth, link.url)
	tokens <- struct{}{}
	depth := link.depth
	list, err := links.Extract(link.url)
	<-tokens
	if err != nil {
		log.Printf("error occured in %s, err: %+v", link.url, err)
	}
	var urls []linkSet

	for _, url := range list {
		urls = append(urls, linkSet{
			url:   url,
			depth: depth + 1,
		})
	}
	return urls
}
