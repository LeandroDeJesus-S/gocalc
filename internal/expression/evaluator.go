package expression

import (
	"fmt"
	"strconv"

	"github.com/LeandroDeJesus-S/gocalc/internal/operations"
	"github.com/LeandroDeJesus-S/gocalc/internal/utils"
)

// Evaluate takes a postfix expression represented as a queue of Tokens and
// evaluates it. The expression is evaluated by iterating over the queue and
// performing the corresponding operation when an operator is encountered. The
// result of the evaluation is returned as a float64.
func Evaluate(postFix utils.Queue[utils.Token]) float64 {
	resultStack := utils.Stack[float64]{}
	for _, token := range postFix.Values {
		if token.Type == "number" {
			num, err := strconv.ParseFloat(token.Value, 64)
			if err != nil {
				panic(fmt.Sprintf("Error parsing number '%s': %v", token.Value, err))
			}
			resultStack.Push(num)
			
		} else if token.Type == "operator" {
			var res float64
			switch token.Value {
			case "+":
				res = utils.GetOperatorRes(&resultStack, operations.Sum)
			case "-":
				res = utils.GetOperatorRes(&resultStack, operations.Sub)
			case "*":
				res = utils.GetOperatorRes(&resultStack, operations.Mul)
			case "/":
				res = utils.GetOperatorRes(&resultStack, operations.Div)
			case "^":
				res = utils.GetOperatorRes(&resultStack, operations.Pow)
			}
			resultStack.Push(res)
		}
	}
	return resultStack.Pop()
}
