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
	for d, c := range visit(map[string]int{}, doc) {
		fmt.Fprintf(w, "dom: %s, %d \n", d, c)
	}
	return nil
}

func visit(doms map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		doms[n.Data]++
	}
	if c := n.FirstChild; c != nil {
		doms = visit(doms, c)
	}
	if c := n.NextSibling; c != nil {
		doms = visit(doms, c)
	}
	return doms
}
