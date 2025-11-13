package tokenizer

import (
	"slices"
	"testing"
)

func TestTokenizer(t *testing.T) {
	t.Run("arithmetics tokens", func(t *testing.T) {
		expr := "(1 + 20) * (238 - 99)"
		want := []Token{
			{T: TokenLeftParen, Content: "("},
			{T: TokenNumber, Content: "1"},
			{T: TokenPlus, Content: "+"},
			{T: TokenNumber, Content: "20"},
			{T: TokenRightParen, Content: ")"},
			{T: TokenMul, Content: "*"},
			{T: TokenLeftParen, Content: "("},
			{T: TokenNumber, Content: "238"},
			{T: TokenMinus, Content: "-"},
			{T: TokenNumber, Content: "99"},
			{T: TokenRightParen, Content: ")"},
			EOF,
		}
		got := Tokenize(expr)
		if !slices.Equal(got, want) {
			t.Errorf("expr: %q -> got %v, wanted %v", expr, got, want)
		}
	})
	t.Run("vairable assignment tokens", func(t *testing.T) {
		statement := "x = 10"
		want := []Token{
			{T: TokenIdentifier, Content: "x"},
			{T: TokenAssign, Content: "="},
			{T: TokenNumber, Content: "10"},
			EOF,
		}
		got := Tokenize(statement)
		if !slices.Equal(got, want) {
			t.Errorf("statement: %q -> got %v, wanted %v", statement, got, want)
		}
	})
}
