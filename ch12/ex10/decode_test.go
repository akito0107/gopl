package ex10

import (
	"log"
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {

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

func TestUnmarshal_Interface(t *testing.T) {

	type Interface struct {
		Int    interface{}
		Int8   interface{}
		Int16  interface{}
		Int32  interface{}
		Int64  interface{}
		Uint8  interface{}
		Uint16 interface{}
		Uint32 interface{}
		Uint64 interface{}
		String interface{}
	}

	input := `((Int("int" 10))
 (Int8("int8" 10))
 (Int16("int16" 10))
 (Int32("int32" 10))
 (Int64("int64" 10))
 (Uint8("uint8" 10))
 (Uint16("uint16" 10))
 (Uint32("uint32" 10))
 (Uint64("uint64" 10))
 (String("string" "string")))`

	expect := Interface{
		Int:    int(10),
		Int8:   int8(10),
		Int16:  int16(10),
		Int32:  int32(10),
		Int64:  int64(10),
		Uint8:  uint8(10),
		Uint16: uint16(10),
		Uint32: uint32(10),
		Uint64: uint64(10),
		String: "string",
	}

	b, err := Marshal(expect)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(string(b))

	var i Interface

	if err := Unmarshal([]byte(input), &i); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(i, expect) {
		t.Errorf("Must be same value expect: %+v\n, actual: %+v", expect, i)
	}

}
