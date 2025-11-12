package lang

import (
	"slices"
	"testing"
)

func TestTokenizer(t *testing.T) {
	expr := "1 + 20 * (238 - 99)"
	want := []Token{
		{T: TokenNumber, Content: "1"},
		{T: TokenPlus, Content: "+"},
		{T: TokenNumber, Content: "20"},
		{T: TokenMul, Content: "*"},
		{T: TokenLeftParen, Content: "("},
		{T: TokenNumber, Content: "238"},
		{T: TokenMinus, Content: "-"},
		{T: TokenNumber, Content: "99"},
		{T: TokenRightParen, Content: ")"},
		EOF,
	}
	got := tokenize(expr)
	if !slices.Equal(got, want) {
		t.Errorf("expr: %q -> got %v, wanted %v", expr, got, want)
	}
}
