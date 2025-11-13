package parser

// this file implements Recursive Descent Parsing
// for parsing expression with precedents

import (
	"fmt"
	"log"
	"strconv"

	"github.com/rezbow/lang/ast"
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

func (p *Parser) value(level PrecedentLevel) ast.Node {
	if level >= PrecedentLevelThree {
		t := p.next()
		switch t.T {
		case tokenizer.TokenNumber:
			return &ast.NodeNumber{N: toInt(t.Content)}
		case tokenizer.TokenLeftParen:
			return p.expr(PrecedentLevelOne)
		case tokenizer.TokenPlus:
			return p.value(level)
		default:
			p.err = fmt.Errorf("Unexpected token: %q", t.Content)
			return nil
		}
	} else {
		return p.expr(level)
	}
}

func (p *Parser) expr(level PrecedentLevel) ast.Node {
	left := p.value(level + 1)
	for p.peek() != tokenizer.TokenEOF && p.peek() != tokenizer.TokenRightParen && level == tokenPrecedent(p.peek()) {
		op := p.next()
		right := p.value(level + 1)
		switch op.T {
		case tokenizer.TokenPlus:
			left = &ast.NodeBinaryOperator{
				Left:  left,
				Right: right,
				Op:    "+",
			}
		case tokenizer.TokenMinus:
			left = &ast.NodeBinaryOperator{
				Left:  left,
				Right: right,
				Op:    "-",
			}
		case tokenizer.TokenMul:
			left = &ast.NodeBinaryOperator{
				Left:  left,
				Right: right,
				Op:    "*",
			}
		}
	}
	return left
}
