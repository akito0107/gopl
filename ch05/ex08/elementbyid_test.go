package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestElementByID(t *testing.T) {
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
	doc, _ := html.Parse(strings.NewReader(in))
	if ElementByID(doc, "content") == nil {
		t.Errorf("content must be found")
	}
	if ElementByID(doc, "content2") != nil {
		t.Errorf("content2 must be not found")
	}
}
