package main

import "golang.org/x/net/html"

type parseFunc func(n *html.Node) []*html.Node

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	pre := func(n *html.Node) []*html.Node {
		var result []*html.Node
		if contains(n.Data, name) {
			result = append(result, n)
		}
		return result
	}
	return forEachNode(doc, pre)
}

func contains(target string, strs []string) bool {
	for _, s := range strs {
		if s == target {
			return true
		}
	}
	return false
}

func forEachNode(n *html.Node, pre parseFunc) []*html.Node {
	var stack []*html.Node
	stack = append(stack, n)
	var result []*html.Node

	for len(stack) > 0 {
		n = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, pre(n)...)
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			stack = append(stack, c)
		}
	}
	return result
}
