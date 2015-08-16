package runner

import "testing"

func TestRoot(t *testing.T) {
	r := root()
	if r != "." {
		t.Error("root is not .")
	}
}
