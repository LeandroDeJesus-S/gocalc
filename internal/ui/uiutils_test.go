package ui

import (
	"testing"

	"github.com/LeandroDeJesus-S/gocalc/internal/utils"
)

func TestUpdateDisplay(t *testing.T) {
	tests := []struct {
		name        string
		label       string
		currentText string
		want        string
		setupParen  func(*utils.Stack[string])
	}{
		{
			name:        "Append number to non-zero",
			label:       "5",
			currentText: "12",
			want:        "125",
			setupParen:  func(p *utils.Stack[string]) {},
		},
		{
			name:        "Replace zero with number",
			label:       "7",
			currentText: "0",
			want:        "7",
			setupParen:  func(p *utils.Stack[string]) {},
		},
		{
			name:        "Do not append dot after operator",
			label:       ".",
			currentText: "+",
			want:        "+",
			setupParen:  func(p *utils.Stack[string]) {},
		},
		{
			name:        "Append dot after number",
			label:       ".",
			currentText: "5",
			want:        "5.",
			setupParen:  func(p *utils.Stack[string]) {},
		},
		{
			name:        "Append operator after number",
			label:       "+",
			currentText: "5",
			want:        "5+",
			setupParen:  func(p *utils.Stack[string]) {},
		},
		{
			name:        "Do not append operator after operator",
			label:       "+",
			currentText: "+",
			want:        "+",
			setupParen:  func(p *utils.Stack[string]) {},
		},
		{
			name:        "Append left parenthesis",
			label:       "(",
			currentText: "0",
			want:        "(",
			setupParen:  func(p *utils.Stack[string]) {},
		},
		{
			name:        "Append right parenthesis with stack",
			label:       ")",
			currentText: "5",
			want:        "5)",
			setupParen: func(p *utils.Stack[string]) {
				p.Push("test")
			},
		},
		{
			name:        "No append right parenthesis if stack empty",
			label:       ")",
			currentText: "5",
			want:        "5",
			setupParen:  func(p *utils.Stack[string]) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var parentheses utils.Stack[string]
			tt.setupParen(&parentheses)
			got := UpdateDisplay(tt.label, tt.currentText, &parentheses)
			if got != tt.want {
				t.Errorf("UpdateDisplay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolveExpression(t *testing.T) {
	tests := []struct {
		expression string
		want       float64
	}{
		{"2+3", 5},
		{"2*3", 6},
		{"6/2", 3},
		{"2^3", 8},
		{"(2+3)*4", 20},
	}

	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			got := ResolveExpression(tt.expression)
			if got != tt.want {
				t.Errorf("ResolveExpression(%v) = %v, want %v", tt.expression, got, tt.want)
			}
		})
	}
}
