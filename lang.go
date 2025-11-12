package lang

import (
	"log"
	"strconv"
)

type PrecedentLevel int

const (
	PrecedentLevelOne   PrecedentLevel = iota // + and -
	PrecedentLevelTwo                         // * and /
	PrecedentLevelThree                       // unary and () and number
)

func tokenPrecedent(t TokenType) PrecedentLevel {
	switch t {
	case TokenPlus, TokenMinus:
		return PrecedentLevelOne
	case TokenMul:
		return PrecedentLevelTwo
	default:
		return PrecedentLevelThree
	}
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Println(err)
	}
	return n
}

type Evaluator struct {
	tokens []Token
	cursor int
}

func (e *Evaluator) value(level PrecedentLevel) int {
	if level >= PrecedentLevelThree {
		t := e.next()
		switch t.T {
		case TokenNumber:
			return toInt(t.Content)
		case TokenLeftParen:
			return e.expr(PrecedentLevelOne)
		default:
			return 0
		}
	} else {
		return e.expr(level)
	}
}

func (e *Evaluator) expr(level PrecedentLevel) int {
	v := e.value(level + 1)
	for e.peek() != TokenEOF && e.peek() != TokenRightParen && level == tokenPrecedent(e.peek()) {
		op := e.next()
		right := e.value(level + 1)
		switch op.T {
		case TokenPlus:
			v = v + right
		case TokenMinus:
			v = v - right
		case TokenMul:
			v = v * right
		}
	}
	return v
}

func (e *Evaluator) Eval() int {
	return e.expr(PrecedentLevelOne)
}

func (e *Evaluator) peek() TokenType {
	return e.tokens[e.cursor].T
}
func (e *Evaluator) next() Token {
	t := e.tokens[e.cursor]
	if t.T != TokenEOF {
		e.cursor++
	}
	return t
}

func Run(expr string) int {
	tokens := tokenize(expr)
	E := Evaluator{tokens: tokens}
	return E.Eval()
}
