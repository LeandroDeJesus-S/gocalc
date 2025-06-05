package operations

func Pow(base, exponent float64) float64 {
	if exponent == 0 {
		return 1
	}
	if exponent < 0 {
		return 1 / Pow(base, -exponent)
	}
	result := 1.0
	for range int(exponent) {
		result *= base
	}
	return result
}
