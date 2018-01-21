package main

import "testing"

func Test_min(t *testing.T) {
	if i := min(0, 1, 2, 3, 4, 5); i != 0 {
		t.Errorf("min value must 0 but %d", i)
	}
	if i := min(-1, 0); i != -1 {
		t.Errorf("min value must -1 but %d", i)
	}
}

func Test_max(t *testing.T) {
	if i := max(0, 1, 2, 3, 4, 5); i != 5 {
		t.Errorf("min value must 5 but %d", i)
	}
	if i := max(-1, 0); i != 0 {
		t.Errorf("min value must 0 but %d", i)
	}
}
