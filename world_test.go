package m2

import "testing"

func TestM2(t *testing.T) {
	expected := "hello world"
	got := World()
	if expected != got {
		t.Errorf("expected: %v\tgot %v\n", expected, got)
	}
}
