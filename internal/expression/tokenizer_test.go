package expression

import (
	"testing"

	"github.com/LeandroDeJesus-S/gocalc/internal/utils"
)

func TestTokenizeShouldReturnCorrectTokens(t *testing.T) {
	expression := "3.5 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3"
	expectedTokens := []utils.Token{
		{Value: "3.5", Type: "number", Precedence: 0, Associativity: ""},
		{Value: "+", Type: "operator", Precedence: 1, Associativity: "left"},
		{Value: "4", Type: "number", Precedence: 0, Associativity: ""},
		{Value: "*", Type: "operator", Precedence: 2, Associativity: "left"},
		{Value: "2", Type: "number", Precedence: 0, Associativity: ""},
		{Value: "/", Type: "operator", Precedence: 2, Associativity: "left"},
		{Value: "(", Type: "left_paren", Precedence: 0, Associativity: ""},
		{Value: "1", Type: "number", Precedence: 0, Associativity: ""},
		{Value: "-", Type: "operator", Precedence: 1, Associativity: "left"},
		{Value: "5", Type: "number", Precedence: 0, Associativity: ""},
		{Value: ")", Type: "right_paren", Precedence: 0, Associativity: ""},
		{Value: "^", Type: "operator", Precedence: 3, Associativity: "right"},
		{Value: "2", Type: "number", Precedence: 0, Associativity: ""},
		{Value: "^", Type: "operator", Precedence: 3, Associativity: "right"},
		{Value: "3", Type: "number", Precedence: 0, Associativity: ""},
	}

	tokens := Tokenize(expression)
	if len(tokens) != len(expectedTokens) {
		t.Errorf("Expected %d tokens, got %d", len(expectedTokens), len(tokens))
	}

	for i, token := range tokens {
		if token != expectedTokens[i] {
			t.Errorf("Expected token %d to be %v, got %v", i, expectedTokens[i], token)
		}
	}
}
