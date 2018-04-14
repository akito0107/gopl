package ex07

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestDecoder_Decode(t *testing.T) {
	t.Run("string slice", func(t *testing.T) {
		input := `("Test1" "Test2" "Test3")
("Test1" "Test2" "Test3")
("Test1" "Test2" "Test3")
("Test1" "Test2" "Test3")`

		var called int
		expectCall := 4
		expect := []string{"Test1", "Test2", "Test3"}
		buf := bytes.NewBufferString(input)
		decoder := NewDecoder(buf)
		for {
			var in []string
			err := decoder.Decode(&in)
			if err == io.EOF {
				if called != expectCall {
					t.Errorf("must call %d, but %d", expectCall, called)
				}
				return
			} else if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(in, expect) {
				t.Errorf("actual: %+v, expect: %+v", in, expect)
			}
			called++
		}
	})
	t.Run("int slice", func(t *testing.T) {
		input := `(1 2 3 4 5)
(1 2 3 4 5)
(1 2 3 4 5)
(1 2 3 4 5)
(1 2 3 4 5)`

		var called int
		expectCall := 5
		expect := []int{1, 2, 3, 4, 5}
		buf := bytes.NewBufferString(input)
		decoder := NewDecoder(buf)
		for {
			var in []int
			err := decoder.Decode(&in)
			if err == io.EOF {
				if called != expectCall {
					t.Errorf("must call %d, but %d", expectCall, called)
				}
				return
			} else if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(in, expect) {
				t.Errorf("actual: %+v, expect: %+v", in, expect)
			}
			called++
		}
	})
	t.Run("struct", func(t *testing.T) {
		input := `((String "Test")(Integer 10))
		((String "Test")(Integer 10))
		((String "Test")(Integer 10))
		((String "Test")(Integer 10))`

		type testStruct struct {
			String  string
			Integer int
		}

		var called int
		expectCall := 4
		expect := testStruct{"Test", 10}
		buf := bytes.NewBufferString(input)
		decoder := NewDecoder(buf)
		for {
			var in testStruct
			err := decoder.Decode(&in)
			if err == io.EOF {
				if called != expectCall {
					t.Errorf("must call %d, but %d", expectCall, called)
				}
				return
			} else if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(in, expect) {
				t.Errorf("actual: %+v, expect: %+v", in, expect)
			}
			called++
		}
	})
}
