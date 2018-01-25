package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestQuery_parse(t *testing.T) {
	t.Run("single values", func(t *testing.T) {
		in := []string{"div", ".class", "#id"}
		q := parseQuery(in)
		if q[0].Type != "tag" || q[0].Value != "div" {
			t.Errorf("parse failed: %+v", q)
		}
		if q[1].Type != "class" || q[1].Value != "class" {
			t.Errorf("parse failed: %+v", q)
		}
		if q[2].Type != "id" || q[2].Value != "id" {
			t.Errorf("parse failed: %+v", q)
		}
	})
}

func TestQuery_match(t *testing.T) {
	cases := []struct {
		name  string
		query *Query
		elem  *Elem
		want  bool
	}{
		{
			name: "class match",
			query: &Query{
				Type:  "class",
				Value: "class",
			},
			elem: &Elem{
				Classes: []string{
					"class",
					"test",
				},
			},
			want: true,
		},
		{
			name: "tag match",
			query: &Query{
				Type:  "tag",
				Value: "div",
			},
			elem: &Elem{
				Name: "div",
			},
			want: true,
		},
		{
			name: "id match",
			query: &Query{
				Type:  "id",
				Value: "id",
			},
			elem: &Elem{
				Id: "id",
			},
			want: true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if act := c.query.match(c.elem); act != c.want {
				t.Errorf("want %s, but %s", c.want, act)
			}
		})
	}
}

func Test_run(t *testing.T) {
	in := `
<div>
  <div class="class">
    <h2 id="id">test</h2>
  </div>
</div>
`
	r := strings.NewReader(in)
	queries := parseQuery([]string{
		"div", ".class", "#id",
	})
	var buf []byte
	w := bytes.NewBuffer(buf)
	run(r, w, queries)

	want := "div div h2: test\n"
	if act := w.String(); act != want {
		t.Errorf("must be [%s] but [%s]", want, act)
	}
}
