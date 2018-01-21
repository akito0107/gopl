package main

import (
	"fmt"

	"io"

	"strings"

	"golang.org/x/net/html"
)

var depth int

func pp(r io.Reader, w io.Writer) error {
	doc, err := html.Parse(r)
	if err != nil {
		return err
	}
	forEachNode(doc, startElement(w), endElement(w))
	return nil
}

type parseFunc func(n *html.Node)

func forEachNode(n *html.Node, pre, post parseFunc) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func startElement(w io.Writer) parseFunc {
	return func(n *html.Node) {
		if n.Type == html.ElementNode || n.Type == html.CommentNode || n.Type == html.TextNode {
			depth++
		}
		ppElement(w, n, depth*2)
	}
}

func endElement(w io.Writer) parseFunc {
	return func(n *html.Node) {
		if n.Type == html.ElementNode || n.Type == html.CommentNode || n.Type == html.TextNode {
			if n.FirstChild != nil {
				fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
			}
			depth--
		}
	}
}

func ppElement(w io.Writer, n *html.Node, depth int) {
	switch n.Type {
	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		text = strings.TrimRight(text, "\n")
		if text != "" {
			for _, s := range strings.Split(text, "\n") {
				fmt.Fprintf(w, "%*s%s\n", depth, "", strings.TrimSpace(s))
			}
		}
		break
	case html.CommentNode:
		comment := n.Data
		if strings.Contains(comment, "\n") {
			fmt.Fprintf(w, "%*s<!-\n", depth, "")
			for _, s := range strings.Split(comment, "\n") {
				s = strings.TrimSpace(s)
				if s != "" {
					fmt.Fprintf(w, "%*s%s\n", depth, "", strings.TrimSpace(s))
				}
			}
			fmt.Fprintf(w, "%*s-->\n", depth, "")
		} else {
			fmt.Fprintf(w, "%*s<!-- %s-->\n", depth, "", n.Data)
		}
		break
	case html.ElementNode:
		fmt.Fprintf(w, "%*s<%s ", depth, "", n.Data)
		for _, attr := range n.Attr {
			if attr.Namespace != "" {
				fmt.Fprintf(w, "%s:", attr.Namespace)
			}
			fmt.Fprintf(w, "%s=\"%s\" ", attr.Key, attr.Val)
		}
		if n.FirstChild == nil {
			fmt.Fprint(w, "/>\n")
		} else {
			fmt.Fprint(w, ">\n")
		}
		break
	}
}
