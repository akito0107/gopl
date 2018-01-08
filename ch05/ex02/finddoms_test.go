package main

import (
	"bytes"
	"testing"

	"golang.org/x/net/html"
)

func Test_visit(t *testing.T) {
	in := `<HTML><HEAD><meta http-equiv="content-type" content="text/html;charset=utf-8">
<TITLE>302 Moved</TITLE></HEAD><BODY>
<H1>302 Moved</H1>
The document has moved
<A HREF="http://www.google.co.jp/?gfe_rd=cr&amp;dcr=0&amp;ei=ySZTWvSUI6H98Ae_nrqQDA">here</A>.
</BODY></HTML>`
	stdin := bytes.NewBufferString(in)
	d, _ := html.Parse(stdin)
	out := visit(map[string]int{}, d)

	if out["html"] != 1 {
		t.Errorf("html tag must be 1 but %d", out["html"])
	}

	if out["h1"] != 1 {
		t.Errorf("h1 tag must be 1 but %d", out["html"])
	}

	if out["a"] != 1 {
		t.Errorf("a tag must be 1 but %d", out["html"])
	}
}
