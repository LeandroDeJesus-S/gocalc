package utils

func IsOperator(char rune) bool {
	switch char {
	case '+', '-', '*', '/', '^':
		return true
	default:
		return false
	}
}

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

func GetOperatorRes(resStack *Stack[float64], funcOp func(float64, float64) float64) float64 {
	right := resStack.Pop()
	left := resStack.Pop()
	return funcOp(left, right)
}
