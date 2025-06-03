package operations

import "testing"

func TestSub(t *testing.T) {
	A := 1.0
	B := 2.0
	expected := -1.0
	actual := Sub(A, B)
	if actual != expected {
		t.Errorf("Sub(%f, %f) = %f, want %f", A, B, actual, expected)
	}
}
