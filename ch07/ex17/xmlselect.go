package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	queries := parseQuery(os.Args[1:])
	run(os.Stdin, os.Stdout, queries)
}

type Elem struct {
	Name    string
	Classes []string
	Id      string
}

func parseToken(tok xml.StartElement) *Elem {
	elem := &Elem{Name: tok.Name.Local}
	for _, attr := range tok.Attr {
		if attr.Name.Local == "class" {
			elem.Classes = strings.Split(attr.Value, " ")
		} else if attr.Name.Local == "id" {
			elem.Id = attr.Value
		}
	}
	return elem
}

func Names(elems []*Elem) []string {
	var names []string
	for _, e := range elems {
		names = append(names, e.Name)
	}
	return names
}

type Query struct {
	Type  string
	Value string
}

func parseQuery(query []string) []Query {
	var queries []Query
	for _, q := range query {
		res := Query{}
		if strings.HasPrefix(q, ".") {
			res.Type = "class"
			res.Value = q[1:]
		} else if strings.HasPrefix(q, "#") {
			res.Type = "id"
			res.Value = q[1:]
		} else {
			res.Type = "tag"
			res.Value = q
		}
		queries = append(queries, res)
	}
	return queries
}

func (q *Query) match(e *Elem) bool {
	if q.Type == "tag" {
		return q.Value == e.Name
	}
	if q.Type == "class" {
		for _, class := range e.Classes {
			if q.Value == class {
				return true
			}
		}
	}
	if q.Type == "id" {
		return q.Type == e.Id
	}

	return false
}

func run(r io.Reader, w io.Writer, contains []Query) {
	dec := xml.NewDecoder(r)
	var stack []*Elem

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, parseToken(tok))
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, contains) {
				fmt.Fprintf(w, "%s: %s\n", strings.Join(Names(stack), " "), tok)
			}
		}
	}
}

func containsAll(x []*Elem, y []Query) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if y[0].match(x[0]) {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
