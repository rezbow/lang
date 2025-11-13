package ast

type NodeType int

type Node interface {
	Equal(other Node) bool
	String() string
}

type AST struct {
	Root []Node
}
