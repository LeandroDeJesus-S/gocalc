package expression

import (
	"testing"

	"github.com/LeandroDeJesus-S/gocalc/internal/utils"
)

func TestShuntingYardShouldReturnCorrectQueue(t *testing.T) {
	expression := "3 + 4 * 2 / (1 - 5) ^ 2 ^ 3"
	expectedQueue := []utils.Token{
		{Value: "3", Type: "number", Precedence: 0, Associativity: ""},
		{Value: "4", Type: "number", Precedence: 0, Associativity: ""},
		{Value: "2", Type: "number", Precedence: 0, Associativity: ""},
		{Value: "*", Type: "operator", Precedence: 2, Associativity: "left"},
		{Value: "1", Type: "number", Precedence: 0, Associativity: ""},
		{Value: "5", Type: "number", Precedence: 0, Associativity: ""},
		{Value: "-", Type: "operator", Precedence: 1, Associativity: "left"},
		{Value: "2", Type: "number", Precedence: 0, Associativity: ""},
		{Value: "3", Type: "number", Precedence: 0, Associativity: ""},
		{Value: "^", Type: "operator", Precedence: 3, Associativity: "right"},
		{Value: "^", Type: "operator", Precedence: 3, Associativity: "right"},
		{Value: "/", Type: "operator", Precedence: 2, Associativity: "left"},
		{Value: "+", Type: "operator", Precedence: 1, Associativity: "left"},
	}
	tokens := Tokenize(expression)
	queue := ShuntingYard(tokens)

	if queue.Len() != len(expectedQueue) {
		t.Errorf("Expected queue length %d, got %d", len(expectedQueue), queue.Len())
	}
	for i, token := range queue.Values {
		if token != expectedQueue[i] {
			t.Errorf("Expected token %d to be %v, got %v", i, expectedQueue[i], token)
		}
	}
}
