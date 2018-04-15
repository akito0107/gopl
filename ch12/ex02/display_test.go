package ex02

import (
	"testing"

	"bytes"
	"os"
)

type Cycle struct {
	Value int
	Tail  *Cycle
}

func TestDisplay(t *testing.T) {
	var buf bytes.Buffer
	writer = &buf
	defer func() {
		writer = os.Stdout
	}()

	var c Cycle
	c = Cycle{42, &c}

	Display("test", c)

	expect := `Display test (ex02.Cycle): 
test.Value = 42
(*test.Tail).Value = 42
(*(*test.Tail).Tail).Value = 42
(*(*(*test.Tail).Tail).Tail).Value = 42
(*(*(*(*test.Tail).Tail).Tail).Tail).Value = 42
(*(*(*(*(*test.Tail).Tail).Tail).Tail).Tail)... too deep`

	if act := buf.String(); act != expect {
		t.Errorf("expect \n%s\nbut \n%s\n", expect, act)
	}
}
