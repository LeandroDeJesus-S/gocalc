package utils

import (
	"testing"
)

func TestIsOperatorShouldReturnTrueIfCharIsAnOperator(t *testing.T) {
	operators := []rune{'+', '-', '*', '/', '^'}
	for _, op := range operators {
		if !IsOperator(op) {
			t.Errorf("Expected %c to be an operator", op)
		}
	}
}

func TestIsOperatorShouldReturnFalseIfCharIsNotAnOperator(t *testing.T) {
	nonOperators := []rune{'a', '1', '!', '@'}
	for _, nonOp := range nonOperators {
		if IsOperator(nonOp) {
			t.Errorf("Expected %c to not be an operator", nonOp)
		}
	}
}

func TestGetAssociativityShouldReturnCorrectAssociativity(t *testing.T) {
	associativityTests := map[rune]string{
		'+': "left",
		'-': "left",
		'*': "left",
		'/': "left",
		'^': "right",
		'%': "",
	}

	for op, expected := range associativityTests {
		if result := GetAssociativity(op); result != expected {
			t.Errorf("Expected associativity of %c to be %s, got %s", op, expected, result)
		}
	}
}

func TestGetPrecedenceShouldReturnCorrectPrecedence(t *testing.T) {
	precedenceTests := map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
		'^': 3,
		'%': 0, // Assuming % is not defined in the precedence
	}

	for op, expected := range precedenceTests {
		if result := GetPrecedence(op); result != expected {
			t.Errorf("Expected precedence of %c to be %d, got %d", op, expected, result)
		}
	}
}

func TestGetOperatorResShouldReturnCorrectResult(t *testing.T) {
	resStack := &Stack[float64]{}
	resStack.Push(3)
	resStack.Push(4)

	result := GetOperatorRes(resStack, func(left, right float64) float64 {
		return left + right
	})

	if result != 7 {
		t.Errorf("Expected result to be 7, got %f", result)
	}
}
