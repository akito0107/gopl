package main

import (
	"fmt"

	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}

	outline(doc)
}

type parseFunc func(n *html.Node)

func outline(doc *html.Node) *html.Node {
	var depth int

	startElement := func() parseFunc {
		return func(n *html.Node) {
			if n.Type == html.ElementNode {
				fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
				depth++
			}
		}
	}

	endElement := func() parseFunc {
		return func(n *html.Node) {
			if n.Type == html.ElementNode {
				depth--
				fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			}
		}
	}

	return forEachNode(doc, startElement(), endElement())
}

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
