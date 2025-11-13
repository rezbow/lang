package parser

import (
	"github.com/rezbow/lang/ast"
	"github.com/rezbow/lang/tokenizer"
)

type Parser struct {
	tokens []tokenizer.Token
	cursor int
	err    error
}

func (p *Parser) peek() tokenizer.TokenType {
	return p.tokens[p.cursor].T
}
func (p *Parser) next() tokenizer.Token {
	t := p.tokens[p.cursor]
	if t.T != tokenizer.TokenEOF {
		p.cursor++
	}
	return t
}

func (p *Parser) Parse() *ast.AST {
	program := &ast.AST{}
	for p.peek() != tokenizer.TokenEOF {
		program.Root = append(program.Root, p.expr(PrecedentLevelOne))
	}
	return program
}

func Parse(source string) *ast.AST {
	p := &Parser{
		tokens: tokenizer.Tokenize(source),
	}
	return p.Parse()
}

/*
 */
