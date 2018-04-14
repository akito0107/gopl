package ex10

import (
	"reflect"
	"testing"
)

func TestDecoder_Decode(t *testing.T) {

	type Movie struct {
		Title, Subtitle string
		Year            int
		Float           float64
		Color           bool
		Sequel          *string
	}

	input := `((Title "Dr. Strangelove")
 (Subtitle "How I Learned to Stop Worrying and Love the Bomb")
 (Year 1964)
 (Float 12.2345)
 (Color true) 
 (Sequel nil))`

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
