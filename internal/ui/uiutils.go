package ui

import (
	"fmt"
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/LeandroDeJesus-S/gocalc/internal/expression"
	"github.com/LeandroDeJesus-S/gocalc/internal/utils"
)

// UpdateDisplay manages the display text based on the button pressed.
// It handles numbers, operators, parentheses, and the decimal point.
// It also ensures that the display text is valid according to the rules of arithmetic expressions.
// The function returns the updated display text.
// If the input is invalid, it returns the current text unchanged.
func UpdateDisplay(label string, currentText string, parentheses *utils.Stack[string]) string {
	numberPattern := regexp.MustCompile(`^\d+(\.\d+)?$`)
	if len(currentText) == 0 {
		currentText = "0"
	}

	switch label {
	case ".":
		if !numberPattern.MatchString(string(currentText[len(currentText)-1])) {
			return currentText
		}
		return currentText + label

	case ")":
		if parentheses.Len() <= 0 {
			return currentText
		}
		lastChar := currentText[len(currentText)-1]
		if lastChar == ')' || numberPattern.MatchString(string(lastChar)) {
			parentheses.Pop()
			return currentText + label
		}

	case "(":
		parentheses.Push(currentText)
		if currentText == "0" {
			return label
		}
		if utils.IsOperator(rune(currentText[len(currentText)-1])) {
			return currentText + label
		}

	case "*", "/", "+", "-", "^":
		if numberPattern.MatchString(string(currentText[len(currentText)-1])) || currentText[len(currentText)-1] == ')' {
			return currentText + label
		}
		return currentText

	default:
		if currentText == "0" {
			return label
		}
		return currentText + label
	}

	return currentText
}

// ResolveExpression takes a mathematical expression in string format and returns 
// its evaluated result as a float64. It tokenizes the expression, converts it to 
// postfix notation using the Shunting-yard algorithm, and then evaluates the result.
func ResolveExpression(exp string) float64 {
	tokens := expression.Tokenize(exp)
	postFix := expression.ShuntingYard(tokens)
	result := expression.Evaluate(postFix)
	return result
}


// GetButton returns a new Button with the given label. When the button is
// clicked, its callback function calls UpdateDisplay to update the display's
// text based on the button's label and the current state of the parentheses
// stack. It then sets the display's text to the result of the update.
func GetButton(label string, display *widget.Entry, parentheses *utils.Stack[string]) fyne.CanvasObject {
	return widget.NewButton(label, func() {
		display.SetText(UpdateDisplay(label, display.Text, parentheses))
	})
}

// GetClearButton returns a new Button with the label "C". When the button is
// clicked, its callback function resets the display's text to "0" and clears the
// parentheses stack.
func GetClearButton(display *widget.Entry, parentheses *utils.Stack[string]) fyne.CanvasObject {
	return widget.NewButton("C", func() {
		display.SetText("0")
		*parentheses = utils.Stack[string]{}
	})
}

// GetEqualButton returns a new Button with the label "=". When the button is
// clicked, its callback function resolves the expression in the display's text
// using ResolveExpression and sets the display's text to the result. It also
// clears the parentheses stack.
func GetEqualButton(display *widget.Entry, parentheses *utils.Stack[string]) fyne.CanvasObject {
	return widget.NewButton("=", func() {
		str_exp := display.Text
		result := ResolveExpression(str_exp)
		display.SetText(fmt.Sprint(result))
		*parentheses = utils.Stack[string]{}
	})
}