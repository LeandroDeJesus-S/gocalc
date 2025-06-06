package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	
	"github.com/LeandroDeJesus-S/gocalc/internal/utils"
)


func Home() {
	a := app.New()
	w := a.NewWindow("Home")
	w.Resize(fyne.NewSize(280, w.Canvas().Size().Height))

	display := widget.NewEntry()
	display.SetText("0")

	var parentheses utils.Stack[string]
	isUpdating := false

	display.OnChanged = func(s string) {
		if isUpdating {
			return
		}

		isUpdating = true
		if s == "" {
			display.SetText("0")
		} else {
			lastChar := s[len(s)-1:]
			newText := UpdateDisplay(lastChar, s[:len(s)-1], &parentheses)
			display.SetText(newText)
		}
		isUpdating = false
	}
	display.OnSubmitted = func(s string) {
		result := ResolveExpression(s)
		display.SetText(fmt.Sprint(result))
		parentheses = utils.Stack[string]{}
	}

	content := container.New(
		layout.NewGridLayout(4),
		GetClearButton(display, &parentheses),
		GetButton("(", display, &parentheses),
		GetButton(")", display, &parentheses),
		GetButton("^", display, &parentheses),

		GetButton("7", display, &parentheses),
		GetButton("8", display, &parentheses),
		GetButton("9", display, &parentheses),
		GetButton("/", display, &parentheses),

		GetButton("4", display, &parentheses),
		GetButton("5", display, &parentheses),
		GetButton("6", display, &parentheses),
		GetButton("*", display, &parentheses),

		GetButton("1", display, &parentheses),
		GetButton("2", display, &parentheses),
		GetButton("3", display, &parentheses),
		GetButton("-", display, &parentheses),

		GetButton("0", display, &parentheses),
		GetButton(".", display, &parentheses),
		GetEqualButton(display, &parentheses),
		GetButton("+", display, &parentheses),
	)
	content = container.New(
		layout.NewVBoxLayout(),
		display,
		content,
	)

	w.SetContent(content)

	w.ShowAndRun()
}
