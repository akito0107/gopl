package ex16

import "testing"

func Test_Join(t *testing.T) {
	if out := Join(",", "test", "test1", "test2"); out != "test,test1,test2" {
		t.Errorf("must be test,test1,test2, but actual: %s\n", out)
	}
	if out := Join("", "test", "test1", "test2"); out != "testtest1test2" {
		t.Errorf("must be testtest1test2, but actual: %s\n", out)
	}
	if out := Join(","); out != "" {
		t.Errorf("must be blank, but actual: %s\n", out)
	}
}
