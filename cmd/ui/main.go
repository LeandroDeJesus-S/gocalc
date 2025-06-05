package main

import (
	"fmt"
	"github.com/LeandroDeJesus-S/gocalc/internal/expression"
)

func main() {
	exp := "3 + 4"

	tokens := expression.Tokenize(exp)
	postFix := expression.ShuntingYard(tokens)
	result := expression.Evaluate(postFix)
	fmt.Println("Result:", result)
}
