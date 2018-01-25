package main

import (
	"fmt"
	"log"
	"os"

	"io"

	"strings"

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

func visit(texts []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		texts = append(texts, text)
	}

	if c := n.FirstChild; c != nil {
		if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
			return texts
		} else {
			texts = visit(texts, c)
		}
	}

	if c := n.NextSibling; c != nil {
		texts = visit(texts, c)
	}
	return texts
}
