package main

import (
	"bytes"
	"log"
	"testing"

	"strings"

	"github.com/andreyvit/diff"
	"golang.org/x/net/html"
)

func Test_ppElement(t *testing.T) {
	cases := []struct {
		in    *html.Node
		out   string
		depth int
		name  string
	}{
		{
			in: &html.Node{
				Data: "img",
				Type: html.ElementNode,
				Attr: []html.Attribute{
					{Key: "src", Val: "http://some.host.com/test.png"},
				},
			},
			out:   "<img src=\"http://some.host.com/test.png\" />\n",
			depth: 0,
			name:  "basic image parser",
		},
		{
			in: &html.Node{
				Data: "img",
				Type: html.ElementNode,
				Attr: []html.Attribute{
					{Key: "src", Val: "http://some.host.com/test.png"},
				},
			},
			out:   "  <img src=\"http://some.host.com/test.png\" />\n",
			depth: 2,
			name:  "depth check",
		},
		{
			in: &html.Node{
				Data: "div",
				Type: html.ElementNode,
				Attr: []html.Attribute{
					{Key: "class", Val: "test-class"},
					{Namespace: "ng", Key: "data", Val: "test-data"},
				},
				FirstChild: &html.Node{},
			},
			out:   "<div class=\"test-class\" ng:data=\"test-data\" >\n",
			depth: 0,
			name:  "child",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			act := new(bytes.Buffer)
			ppElement(act, c.in, c.depth)
			if bytes.Compare([]byte(c.out), act.Bytes()) != 0 {
				log.Println(diff.LineDiff(c.out, act.String()))
				t.Fatal("not matched")
			}
		})
	}
}

func Test_pp(t *testing.T) {
	in := `<!doctype html>

<html lang="en">
<head>
  <meta charset="utf-8">
  <title>The HTML5 Herald</title>
  <link rel="stylesheet" href="css/styles.css?v=1.0">
</head>
<body>
  <main>
<!-- this is commented node -->
<!--
multi line
-->
  <div id="content">
    test test
    test test
  </div>
  </main>
</body>
</html>`
	expected := `  <html lang="en" >
    <head >
      <meta charset="utf-8" />
      <title >
        The HTML5 Herald
      </title>
      <link rel="stylesheet" href="css/styles.css?v=1.0" />
    </head>
    <body >
      <main >
        <!--  this is commented node -->
        <!-
        multi line
        -->
        <div id="content" >
          test test
          test test
        </div>
      </main>
    </body>
  </html>
`

	act := new(bytes.Buffer)

	pp(strings.NewReader(in), act)

	if bytes.Compare([]byte(expected), act.Bytes()) != 0 {
		t.Log(diff.LineDiff(expected, act.String()))
		t.Fatal("not matched")
	}
}
