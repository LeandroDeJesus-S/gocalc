package operations

import "testing"

func TestSum(t *testing.T) {
	A := 1.0
	B := 2.0
	expected := 3.0
	actual := Sum(A, B)
	if actual != expected {
		t.Errorf("Sum(%f, %f) = %f, want %f", A, B, actual, expected)
	}
}
