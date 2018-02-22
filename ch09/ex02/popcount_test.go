package ex02

import "testing"

func TestPopCount(t *testing.T) {
	if i := PopCount(1); i != 1 {
		t.Errorf("must be 1 but %d", i)
	}

	if i := PopCount(5); i != 2 {
		t.Errorf("must be 2 but %d", i)
	}

	if i := PopCount(255); i != 8 {
		t.Errorf("must be 8 but %d", 8)
	}

	if i := PopCount(256); i != 1 {
		t.Errorf("must be 1 but %d", i)
	}

	if i := PopCount(18446744073709551615); i != 64 {
		t.Errorf("must be 64 but %d", i)
	}
}
