package operations

import (
	"testing"
)

func TestPowShowReturnValidResult(t *testing.T) {
	bases := []float64{2, 3, 4}
	exponents := []float64{2, 3, -1}
	expectedResults := []float64{4, 27, 0.25}

	for i, b := range bases {
		result := Pow(b, exponents[i])
		if result != expectedResults[i] {
			t.Errorf("Pow(%f, %f) = %f; want %f", b, exponents[i], result, expectedResults[i])
		}
	}
}
