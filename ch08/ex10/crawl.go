package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

const debug = true

type linkSet struct {
	depth int
	url   string
}

var done = make(chan struct{})
var tokens = make(chan struct{}, 20)
var limit = flag.Int("depth", 2, "crawl depth (default 2)")

func main() {
	flag.Parse()
	worklist := make(chan []linkSet)

	var n int

	go func() {
		os.Stdin.Read(make([]byte, 1))
		log.Println("cancelled")
		close(done)
		close(worklist)
	}()

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
		select {
		case list := <-worklist:
			for _, link := range list {
				if link.depth > *limit {
					continue
				}
				if !seen[link.url] {
					seen[link.url] = true
					n++
					// wg.Add(1)
					// go crawl(wg, link, worklist)
					go crawl(link, worklist)
				}
			}
		case <-done:
			for range worklist {
			}
			if debug {
				panic("just dubug")
			}
		}
	}
}

// func crawl(wg *sync.WaitGroup, link linkSet, workChan chan []linkSet) {
func crawl(link linkSet, workChan chan []linkSet) {
	// defer wg.Done()
	log.Printf("depth: %d url: %s \n", link.depth, link.url)

	if cancelled() {
		return
	}
	tokens <- struct{}{}
	defer func() {
		<-tokens
	}()

	depth := link.depth
	list, err := extract(link.url)
	if err != nil {
		log.Printf("error occured in %s, err: %+v", link.url, err)
		return
	}

	var urls []linkSet
	for _, url := range list {
		urls = append(urls, linkSet{
			url:   url,
			depth: depth + 1,
		})
	}
	if !cancelled() {
		workChan <- urls
	}
}

func extract(url string) ([]string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if cancelled() {
		return nil, err
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
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}

	forEachNode(doc, visitNode, nil)
	return links, nil
}

type parseFunc func(n *html.Node)

func forEachNode(n *html.Node, pre, post parseFunc) *html.Node {
	var stack []*html.Node
	stack = append(stack, n)

	for len(stack) > 0 {
		n = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pre != nil {
			pre(n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			stack = append(stack, c)
		}
		if post != nil {
			post(n)
		}
	}
	return nil
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
