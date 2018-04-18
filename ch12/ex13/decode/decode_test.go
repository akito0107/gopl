package decode

import (
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {

	type Movie struct {
		Title    string  `sexpr:"title"`
		Subtitle string  `sexpr:"subtitle"`
		Year     int
		Float    float64 `sexpr:"float"`
		Color    bool
		Sequel   *string `sexpr:"sequel"`
	}

	input := `((title "Dr. Strangelove")
 (subtitle "How I Learned to Stop Worrying and Love the Bomb")
 (Year 1964)
 (float 12.2345)
 (Color true) 
 (sequel nil))`

	expect := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Float:    12.2345,
		Color:    true,
	}

	var m Movie
	if err := Unmarshal([]byte(input), &m); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(m, expect) {
		t.Errorf("Must be same value expect: %+v\n, actual: %+v", expect, m)
	}

}
