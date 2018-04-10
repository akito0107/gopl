package format_test

import (
	"testing"

	"github.com/akito0107/gopl/ch12/ex01"
)

type key struct {
	k string
}

func TestDisplay(t *testing.T) {
	m := map[key]struct{}{}
	k := key{"test"}
	m[k] = struct{}{}
	format.Display("test", m)
}
