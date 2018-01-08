package main

import "golang.org/x/net/html"

var depth int

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, startElement(id), nil)
}

type parseFunc func(n *html.Node) bool

func forEachNode(n *html.Node, pre, post parseFunc) *html.Node {
	var stack []*html.Node
	stack = append(stack, n)

	for len(stack) > 0 {
		n = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pre != nil {
			if !pre(n) {
				return n
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			stack = append(stack, c)
		}
		if post != nil {
			if !post(n) {
				return n
			}
		}
	}
	return nil
}

func startElement(id string) parseFunc {
	return func(n *html.Node) bool {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return false
			}
		}
		return true
	}
}
