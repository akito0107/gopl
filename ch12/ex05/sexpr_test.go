package ex05

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func Test_encode(t *testing.T) {

	strangelove := Movie{
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

	j, err := Marshal(strangelove)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(j))
	m := &Movie{}
	err = json.Unmarshal(j, m)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(strangelove, *m) {
		t.Errorf("must be same: expect:\n %+v, actual:\n %+v", strangelove, m)
	}

}
