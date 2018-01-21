package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func Test_ElementsByTagName(t *testing.T) {
	in := `<!doctype html>

<html lang="en">
<head>
  <meta charset="utf-8">
  <title>The HTML5 Herald</title>
  <link rel="stylesheet" href="css/styles.css?v=1.0">
</head>
<body>
  <main>
  <h1>hoge</h1>
  <h2>fuga</h2>
  <h3>piyo</h3>
  <div id="content">
    test test
    test test
  </div>
  <div id="content2"> </div>
  </main>
</body>
</html>`

	doc, _ := html.Parse(strings.NewReader(in))
	nodes := ElementsByTagName(doc, "div")
	if len(nodes) != 2 {
		t.Error("nodes must be 2")
	}
	var called int
	for _, attr := range nodes[1].Attr {
		if attr.Key == "id" {
			called += 1
			if attr.Val != "content" {
				t.Errorf("id is must content but %s", attr.Val)
			}
		}
	}
	if called != 1 {
		t.Errorf("id should be found in nodes")
	}
	called = 0
	for _, attr := range nodes[0].Attr {
		if attr.Key == "id" {
			called += 1
			if attr.Val != "content2" {
				t.Errorf("id is must content2 but %s", attr.Val)
			}
		}
	}
	if called != 1 {
		t.Errorf("id should be found in nodes")
	}
}

func Test_ElementsByTagName2(t *testing.T) {
	in := `<!doctype html>

<html lang="en">
<head>
  <meta charset="utf-8">
  <title>The HTML5 Herald</title>
  <link rel="stylesheet" href="css/styles.css?v=1.0">
</head>
<body>
  <main>
  <h1 class="test1" >hoge</h1>
  <h2 id="test2" >fuga</h2>
  <h3 data-id="testtest" >piyo</h3>
  <div id="content">
    test test
    test test
  </div>
  <div id="content2"> </div>
  </main>
</body>
</html>`

	doc, _ := html.Parse(strings.NewReader(in))
	nodes := ElementsByTagName(doc, "h1", "h2", "h3")
	if len(nodes) != 3 {
		t.Error("nodes must be 2")
	}

	var called1 int
	var called2 int
	var called3 int
	for _, node := range nodes {
		if node.Data == "h1" {
			for _, attr := range node.Attr {
				if attr.Key == "class" {
					called1 += 1
					if attr.Val != "test1" {
						t.Errorf("must be test1 but %s", attr.Val)
					}
				}
			}
		}
		if node.Data == "h2" {
			for _, attr := range node.Attr {
				if attr.Key == "id" {
					called2 += 1
					if attr.Val != "test2" {
						t.Errorf("must be test2 but %s", attr.Val)
					}
				}
			}
		}
		if node.Data == "h3" {
			for _, attr := range node.Attr {
				if attr.Key == "data-id" {
					called3 += 1
					if attr.Val != "testtest" {
						t.Errorf("must be testtest but %s", attr.Val)
					}
				}
			}
		}
	}

	if called1 != 1 || called2 != 1 || called3 != 1 {
		t.Errorf("must be called %d, %d, %d", called1, called2, called3)
	}
}
