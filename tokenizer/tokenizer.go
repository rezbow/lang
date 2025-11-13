package tokenizer

import (
	"strings"
	"unicode"
)

type TokenType int

type Token struct {
	T       TokenType
	Content string
}

const (
	TokenNumber TokenType = iota
	TokenPlus
	TokenMinus
	TokenMul
	TokenLeftParen
	TokenRightParen
	TokenIdentifier
	TokenAssign
	TokenEOF
)

var (
	EOF = Token{T: TokenEOF, Content: "EOF"}
)

func skipSpace(expr string, cursor int) int {
	for ; cursor < len(expr) && unicode.IsSpace(rune(expr[cursor])); cursor++ {
	}
	return cursor
}

func identifier(expr string, cursor int) (Token, int) {
	var stringBuilder strings.Builder
	for ; cursor < len(expr) && unicode.IsLetter(rune(expr[cursor])); cursor++ {
		stringBuilder.WriteByte(expr[cursor])
	}
	return Token{
		T:       TokenIdentifier,
		Content: stringBuilder.String(),
	}, cursor
}

func number(expr string, cursor int) (Token, int) {
	var stringBuilder strings.Builder
	for ; cursor < len(expr) && unicode.IsDigit(rune(expr[cursor])); cursor++ {
		stringBuilder.WriteByte(expr[cursor])
	}
	return Token{
		T:       TokenNumber,
		Content: stringBuilder.String(),
	}, cursor
}

func Tokenize(expr string) []Token {
	var tokens []Token
	for cursor := 0; cursor < len(expr); {
		cursor = skipSpace(expr, cursor)
		if cursor >= len(expr) {
			break
		}
		ch := rune(expr[cursor])
		switch {
		case unicode.IsDigit(ch):
			var t Token
			t, cursor = number(expr, cursor)
			tokens = append(tokens, t)
		case unicode.IsLetter(ch):
			var t Token
			t, cursor = identifier(expr, cursor)
			tokens = append(tokens, t)
		case ch == '+':
			tokens = append(tokens, Token{T: TokenPlus, Content: "+"})
			cursor++
		case ch == '-':
			tokens = append(tokens, Token{T: TokenMinus, Content: "-"})
			cursor++
		case ch == '*':
			tokens = append(tokens, Token{T: TokenMul, Content: "*"})
			cursor++
		case ch == '(':
			tokens = append(tokens, Token{T: TokenLeftParen, Content: "("})
			cursor++
		case ch == ')':
			tokens = append(tokens, Token{T: TokenRightParen, Content: ")"})
			cursor++
		case ch == '=':
			tokens = append(tokens, Token{T: TokenAssign, Content: "="})
			cursor++
		default:
			cursor++
		}
	}
	tokens = append(tokens, EOF)
	return tokens
}
