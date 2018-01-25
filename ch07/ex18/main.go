package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node interface{}
type CharData string
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func parse(r io.Reader) *Element {
	dec := xml.NewDecoder(r)
	var stack []*Element
	var root *Element

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "parse failded: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			element := &Element{Type: tok.Name, Attr: tok.Attr}
			if len(stack) == 0 {
				root = element
			} else {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, element)
			}
			stack = append(stack, element)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			ch := strings.TrimSpace(string(tok))
			if ch != "" {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, CharData(string(tok)))
			}
		}
	}
	return root
}
