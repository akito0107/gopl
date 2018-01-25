package main

import (
	"strings"
	"testing"
)

func Test_parse(t *testing.T) {
	in := `<div class="root">
  <div class="class">
    <h2 id="id">test</h2>
  </div>
</div>`
	r := strings.NewReader(in)
	elem := parse(r)

	if act := len(elem.Children); act != 1 {
		t.Errorf("children must be 1 but actually: %d", act)
	}

	ch, ok := elem.Children[0].(*Element)
	if !ok {
		t.Fatal("first ch is must be Element")
	}
	if len(ch.Attr) != 1 || string(ch.Attr[0].Name.Local) != "class" || ch.Attr[0].Value != "class" {
		t.Errorf("parse failed  %+v", ch)
	}

}
