package expression

import (
	"github.com/LeandroDeJesus-S/gocalc/internal/utils"
	"testing"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name    string
		postFix utils.Queue[utils.Token]
		want    float64
		wantErr bool
	}{
		{
			name: "Simple addition",
			postFix: utils.Queue[utils.Token]{
				Values: []utils.Token{
					{Value: "2", Type: "number"},
					{Value: "3", Type: "number"},
					{Value: "+", Type: "operator"},
				},
			},
			want: 5,
		},
		{
			name: "Simple subtraction",
			postFix: utils.Queue[utils.Token]{
				Values: []utils.Token{
					{Value: "5", Type: "number"},
					{Value: "2", Type: "number"},
					{Value: "-", Type: "operator"},
				},
			},
			want: 3,
		},
		{
			name: "Simple multiplication",
			postFix: utils.Queue[utils.Token]{
				Values: []utils.Token{
					{Value: "4", Type: "number"},
					{Value: "5", Type: "number"},
					{Value: "*", Type: "operator"},
				},
			},
			want: 20,
		},
		{
			name: "Simple division",
			postFix: utils.Queue[utils.Token]{
				Values: []utils.Token{
					{Value: "10", Type: "number"},
					{Value: "2", Type: "number"},
					{Value: "/", Type: "operator"},
				},
			},
			want: 5,
		},
		{
			name: "Simple exponentiation",
			postFix: utils.Queue[utils.Token]{
				Values: []utils.Token{
					{Value: "2", Type: "number"},
					{Value: "3", Type: "number"},
					{Value: "^", Type: "operator"},
				},
			},
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantErr {
						t.Errorf("Evaluate() panicked: %v", r)
					}
					return
				}
				if tt.wantErr {
					t.Errorf(
						"Evaluate(name=%s, postFix=%v, want=%v, wantErr=%v) did not panic",
						tt.name,
						tt.postFix,
						tt.want,
						tt.wantErr,
					)
				}
			}()

			got := Evaluate(tt.postFix)
			if got != tt.want {
				t.Errorf("Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
