package operations

import "testing"

func TestDiv(t *testing.T) {
	A := 1.0
	B := 2.0
	expected := 0.5
	actual := Div(A, B)
	if actual != expected {
		t.Errorf("Div(%f, %f) = %f, want %f", A, B, actual, expected)
	}
}
