package operations

import "testing"

// Mul returns the product of two float64 numbers.
func TestMul(t *testing.T) {
	A := 1.0
	B := 2.0
	expected := 2.0
	actual := Mul(A, B)
	if actual != expected {
		t.Errorf("Mul(%f, %f) = %f, want %f", A, B, actual, expected)
	}
}
