package main

import (
	"fmt"
	"log"
	"os"

	"io"

	"golang.org/x/net/html"
)

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func run(r io.Reader, w io.Writer) error {
	doc, err := html.Parse(r)
	if err != nil {
		return err
	}
	for _, link := range visit(nil, doc) {
		fmt.Fprintln(w, link)
	}
	return nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		if n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
		if n.Data == "img" || n.Data == "script" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}
	}
	if c := n.FirstChild; c != nil {
		links = visit(links, c)
	}
	if c := n.NextSibling; c != nil {
		links = visit(links, c)
	}
	return links
}
