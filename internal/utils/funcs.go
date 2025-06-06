package utils

// IsOperator takes a rune and returns a boolean indicating whether the rune is an operator.
func IsOperator(char rune) bool {
	switch char {
	case '+', '-', '*', '/', '^':
		return true
	default:
		return false
	}
}

// GetAssociativity returns the associativity of an operator.
// For operators '+', '-', '*', and '/', it returns "left".
// For operator '^', it returns "right".
// If the operator is not recognized, it returns an empty string.
func GetAssociativity(op rune) string {
	switch op {
	case '+', '-', '*', '/':
		return "left"
	case '^':
		return "right"
	default:
		return ""
	}
}

// GetPrecedence returns the precedence of an operator.
// For operators '+', '-', it returns 1.
// For operators '*', '/', it returns 2.
// For operator '^', it returns 3.
// If the operator is not recognized, it returns 0.
func GetPrecedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	case '^':
		return 3
	default:
		return 0
	}
}

// GetOperatorRes pops two numbers from the result stack and applies the
// given operator function to them, returning the result.
func GetOperatorRes(resStack *Stack[float64], funcOp func(float64, float64) float64) float64 {
	right := resStack.Pop()
	left := resStack.Pop()
	return funcOp(left, right)
}
