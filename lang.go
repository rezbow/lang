package lang

import (
	"fmt"
	"log"
	"strconv"

	"github.com/rezbow/lang/tokenizer"
)

type PrecedentLevel int

const (
	PrecedentLevelOne   PrecedentLevel = iota // + and -
	PrecedentLevelTwo                         // * and /
	PrecedentLevelThree                       // unary and () and number
)

func tokenPrecedent(t tokenizer.TokenType) PrecedentLevel {
	switch t {
	case tokenizer.TokenPlus, tokenizer.TokenMinus:
		return PrecedentLevelOne
	case tokenizer.TokenMul:
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
	tokens []tokenizer.Token
	cursor int
	err    error
}

func (e *Evaluator) value(level PrecedentLevel) int {
	if level >= PrecedentLevelThree {
		t := e.next()
		switch t.T {
		case tokenizer.TokenNumber:
			return toInt(t.Content)
		case tokenizer.TokenLeftParen:
			return e.expr(PrecedentLevelOne)
		case tokenizer.TokenMinus:
			return -(e.value(level))
		case tokenizer.TokenPlus:
			return e.value(level)
		default:
			e.err = fmt.Errorf("Unexpected token: %q", t.Content)
			return 0
		}
	} else {
		return e.expr(level)
	}
}

func (e *Evaluator) expr(level PrecedentLevel) int {
	isRightParen := func() bool {
		is := e.peek() == tokenizer.TokenRightParen
		if is {
			e.next()
		}
		return is
	}
	v := e.value(level + 1)
	for e.peek() != tokenizer.TokenEOF && !isRightParen() && level == tokenPrecedent(e.peek()) {
		op := e.next()
		right := e.value(level + 1)
		switch op.T {
		case tokenizer.TokenPlus:
			v = v + right
		case tokenizer.TokenMinus:
			v = v - right
		case tokenizer.TokenMul:
			v = v * right
		}
	}
	return v
}

func (e *Evaluator) Eval() int {
	return e.expr(PrecedentLevelOne)
}

func (e *Evaluator) peek() tokenizer.TokenType {
	return e.tokens[e.cursor].T
}
func (e *Evaluator) next() tokenizer.Token {
	t := e.tokens[e.cursor]
	if t.T != tokenizer.TokenEOF {
		e.cursor++
	}
	return t
}

func Run(expr string) (int, error) {
	tokens := tokenizer.Tokenize(expr)
	E := Evaluator{tokens: tokens}
	return E.Eval(), E.err
}
