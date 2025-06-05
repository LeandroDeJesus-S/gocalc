package expression

import "github.com/LeandroDeJesus-S/gocalc/internal/utils"


// ShuntingYard implements the Shunting-yard algorithm, which is used to parse a
// mathematical expression specified in infix notation and convert it into a
// Reverse Polish notation (RPN) expression. 
func ShuntingYard(tokens []utils.Token) utils.Queue[utils.Token] {
	queue := utils.Queue[utils.Token]{}
	stack := utils.Stack[utils.Token]{}

	for _, token := range tokens {
		if token.Type == "number" {
			queue.Enqueue(token)

		} else if token.Type == "operator" {
			stackIsNotEmpty := stack.Len() > 0
			isStackTopOperator := stack.Top().Type == "operator"
			hasStackTopGreatestPrecedence := stack.Top().Precedence > token.Precedence
			hasStackTopSamePrecedence := stack.Top().Precedence == token.Precedence
			hasTokenLeftAssociativity := token.Associativity == "left"

			for stackIsNotEmpty && isStackTopOperator && (hasStackTopGreatestPrecedence || (hasStackTopSamePrecedence && hasTokenLeftAssociativity)) {
				queue.Enqueue(stack.Pop())

				stackIsNotEmpty = stack.Len() > 0
				isStackTopOperator = stack.Top().Type == "operator"
				hasStackTopGreatestPrecedence = stack.Top().Precedence > token.Precedence
				hasStackTopSamePrecedence = stack.Top().Precedence == token.Precedence
				hasTokenLeftAssociativity = token.Associativity == "left"
			}
			stack.Push(token)

		} else if token.Type == "left_paren" {
			stack.Push(token)

		} else if token.Type == "right_paren" {
			for stack.Len() > 0 && stack.Top().Type != "left_paren" {
				queue.Enqueue(stack.Pop())
			}
			stack.Pop()
		}
	}

	for stack.Len() > 0 {
		queue.Enqueue(stack.Pop())
	}

	return queue
}
