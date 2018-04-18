package encode

import (
	"testing"

	"github.com/andreyvit/diff"
)

type Movie struct {
	Title    string            `sexpr:"title"`
	Subtitle string            `sexpr:"subtitle"`
	Year     int
	Color    bool              `sexpr:"color"`
	Actor    map[string]string `sexpr:"actor"`
	Oscars   []string          `sexpr:"oscars"`
	Sequel   *string           `sexpr:"sequel"`
}

func TestMarshal(t *testing.T) {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	expect := `((title "Dr. Strangelove")
 (subtitle "How I Learned to Stop Worrying and Love the Bomb")
 (Year 1964)
 (oscars ("Best Actor (Nomin.)"
          "Best Adapted Screenplay (Nomin.)"
          "Best Director (Nomin.)"
          "Best Picture (Nomin.)"))
)`
	var zeroint int
	var zeromap map[struct{}]struct{}
	var zeroslice []int
	type teststruct struct {
		Foo *int
		bar *string
	}
	zerostruct1 := teststruct{}
	i := 12
	var s string
	zerostruct2 := teststruct{
		Foo: &i,
		bar: &s,
	}
	cases := []struct {
		name string
		in   interface{}
		out  string
	}{
		{
			name: "strangelove",
			in:   strangelove,
			out:  expect,
		},
		{
			name: "zero int",
			in:   zeroint,
			out:  "",
		},
		{
			name: "zero map",
			in:   zeromap,
			out:  "",
		},
		{
			name: "zeroslice",
			in:   zeroslice,
			out:  "",
		},
		{
			name: "zerostruct1",
			in:   zerostruct1,
			out:  "",
		},
		{
			name: "zerostruct2",
			in:   zerostruct2,
			out: `((Foo 12)
)`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			b, err := Marshal(c.in)
			if err != nil {
				t.Fatal(err)
			}
			if act := string(b); act != c.out {
				t.Errorf(diff.LineDiff(act, c.out))
			}
		})
	}

}
