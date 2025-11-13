package ast

import "fmt"

type NodeBinaryOperator struct {
	Left  Node
	Right Node
	Op    string
}

func (n *NodeBinaryOperator) Equal(other Node) bool {
	otherNode, ok := other.(*NodeBinaryOperator)
	if !ok {
		return false
	}
	if n.Left.Equal(otherNode.Left) && n.Right.Equal(otherNode.Right) && otherNode.Op == n.Op {
		return true
	}
	return false
}

func (n *NodeBinaryOperator) String() string {
	return fmt.Sprintf("NodeBinaryOperator[Left:%q, Right:%q, OP: %q]", n.Left.String(), n.Right.String(), n.Op)
}

type NodeNumber struct {
	N int
}

func (n *NodeNumber) Equal(other Node) bool {
	otherNode, ok := other.(*NodeNumber)
	if !ok {
		return false
	}
	if n.N == otherNode.N {
		return true
	}
	return false
}

func (n *NodeNumber) String() string {
	return fmt.Sprintf("NodeNumber[Value:%d]", n.N)
}
