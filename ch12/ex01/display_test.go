package ex01

import (
	"testing"

	"bytes"
	"os"
)

type key struct {
	Name string
	id   int
}

func TestDisplayStruct(t *testing.T) {
	var buf bytes.Buffer
	writer = &buf
	defer func() {
		writer = os.Stdout
	}()

	m := map[key]bool{}
	k1 := key{"test", 1}
	k2 := key{"test2", 2}
	m[k1] = false
	m[k2] = true
	Display("test", m)

	expect := `Display test (map[ex01.key]bool): 
test[ex01.key{Name: "test", id: 1}] = false
test[ex01.key{Name: "test2", id: 2}] = true
`

	if act := buf.String(); act != expect {
		t.Errorf("expect %s, but %s", expect, act)
	}
}

func TestDisplayArray(t *testing.T) {
	var buf bytes.Buffer
	writer = &buf
	defer func() {
		writer = os.Stdout
	}()

	m := map[[4]int]bool{}
	k1 := [4]int{1, 2, 3, 4}
	k2 := [4]int{5, 6, 7, 8}
	m[k1] = false
	m[k2] = true
	Display("test", m)

	expect := `Display test (map[[4]int]bool): 
test[[4]int{1, 2, 3, 4}] = false
test[[4]int{5, 6, 7, 8}] = true
`

	if act := buf.String(); act != expect {
		t.Errorf("expect %s, but %s", expect, act)
	}
}
