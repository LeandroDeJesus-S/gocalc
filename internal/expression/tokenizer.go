package expression

import (
	"unicode"
	"strings"
	"github.com/LeandroDeJesus-S/gocalc/internal/utils"
)

// Tokenize takes a mathematical expression and returns a slice of Tokens, where
// each Token represents a number or an operator. The numbers are represented as
// strings, and the operators are represented as runes. The precedence of each
// operator is also stored in the Token. The function handles parentheses and
// spaces correctly.
func Tokenize(expression string) []utils.Token {
	tokens := []utils.Token{}
	num := strings.Builder{}

	for _, char := range expression {
		switch {
		case unicode.IsSpace(char):
			continue

		case unicode.IsDigit(char) || char == '.':
			num.WriteRune(char)
			continue

		default:
			var t utils.Token
			if num.Len() > 0 {
				t = utils.Token{
					Value:      num.String(),
					Type:       "number",
					Precedence: 0,
				}
				tokens = append(tokens, t)
				num.Reset()
			}

			if utils.IsOperator(char) {
				t = utils.Token{
					Value:         string(char),
					Type:          "operator",
					Precedence:    utils.GetPrecedence(char),
					Associativity: utils.GetAssociativity(char),
				}

			}

			if char == '(' {
				t = utils.Token{
					Value:      string(char),
					Type:       "left_paren",
					Precedence: 0,
				}
			}

			if char == ')' {
				t = utils.Token{
					Value:      string(char),
					Type:       "right_paren",
					Precedence: 0,
				}
			}
			tokens = append(tokens, t)
			continue
		}
	}
	if num.Len() > 0 {
		t := utils.Token{
			Value:      num.String(),
			Type:       "number",
			Precedence: 0,
		}
		tokens = append(tokens, t)
		num.Reset()
	}
	return tokens
}
