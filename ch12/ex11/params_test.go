package ex11

import (
	"testing"
)

func TestPack(t *testing.T) {
	data := struct {
		String string `http:"string"`
		Int    int    `http:"int"`
		Bool   bool   `http:"bool"`
		Slice  []int  `http:"slice"`
	}{
		"test",
		123,
		true,
		[]int{1, 2, 3, 4, 5},
	}
	expect := "string=test&int=123&bool=true&slice=[1,2,3,4,5]"

	s, err := Pack(&data)
	if err != nil {
		t.Fatal(err)
	}
	if s != expect {
		t.Errorf("must be %s but %s", expect, s)
	}
}
