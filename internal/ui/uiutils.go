package ui

import (
	"fmt"
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/LeandroDeJesus-S/gocalc/internal/expression"
	"github.com/LeandroDeJesus-S/gocalc/internal/utils"
)

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

func ResolveExpression(exp string) float64 {
	tokens := expression.Tokenize(exp)
	postFix := expression.ShuntingYard(tokens)
	result := expression.Evaluate(postFix)
	return result
}


func GetButton(label string, display *widget.Entry, parentheses *utils.Stack[string]) fyne.CanvasObject {
	return widget.NewButton(label, func() {
		display.SetText(UpdateDisplay(label, display.Text, parentheses))
	})
}

func GetClearButton(display *widget.Entry, parentheses *utils.Stack[string]) fyne.CanvasObject {
	return widget.NewButton("C", func() {
		display.SetText("0")
		*parentheses = utils.Stack[string]{}
	})
}

func GetEqualButton(display *widget.Entry, parentheses *utils.Stack[string]) fyne.CanvasObject {
	return widget.NewButton("=", func() {
		str_exp := display.Text
		result := ResolveExpression(str_exp)
		display.SetText(fmt.Sprint(result))
		*parentheses = utils.Stack[string]{}
	})
}