package sexpr

import (
	"bytes"
	"reflect"
	"testing"
)

func TestDecoder_Decode(t *testing.T) {

	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}

	input := `((Title "Dr. Strangelove")
 (Subtitle "How I Learned to Stop Worrying and Love the Bomb")
 (Year 1964)
 (Color nil)
 (Actor (("Dr. Strangelove" "Peter Sellers")
         ("Grp. Capt. Lionel Mandrake" "Peter Sellers")
         ("Pres. Merkin Muffley" "Peter Sellers")
         ("Gen. Buck Turgidson" "George C. Scott")
         ("Brig. Gen. Jack D. Ripper" "Sterling Hayden")
         ("Maj. T.J. \"King\" Kong" "Slim Pickens")))
 (Oscars ("Best Actor (Nomin.)"
		  "Best Adapted Screenplay (Nomin.)"
		  "Best Director (Nomin.)"
          "Best Picture (Nomin.)"))
 (Sequel nil))`

	expect := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	buf := bytes.NewBufferString(input)
	decoder := NewDecoder(buf)
	var m Movie
	if err := decoder.Decode(&m); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(m, expect) {
		t.Errorf("Must be same value expect: %+v\n, actual: %+v", expect, m)
	}

}
