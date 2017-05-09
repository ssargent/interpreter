package ast

// Expression represents an Expression in the A-S-T WTSE-1
type Expression interface {
	Node
	expressionNode()
}
