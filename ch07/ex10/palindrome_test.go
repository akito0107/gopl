package ex10

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		name string
		in   stringSorter
		out  bool
	}{
		{
			name: "is palindrome",
			in:   []string{"A", "L", "A", "X", "A", "L", "A"},
			out:  true,
		},
		{
			name: "not palindrome",
			in:   []string{"t", "o", "m", "a", "t", "o"},
			out:  false,
		},
		{
			name: "multibyte",
			in:   []string{"と", "ま", "と"},
			out:  true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if act := IsPalindrome(c.in); act != c.out {
				t.Errorf("must %s, but %s", c.out, act)
			}
		})
	}
}

type stringSorter []string

func (x stringSorter) Len() int           { return len(x) }
func (x stringSorter) Less(i, j int) bool { return x[i] < x[j] }
func (x stringSorter) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
