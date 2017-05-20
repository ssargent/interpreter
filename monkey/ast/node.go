package ast

// Node represents a node of the A-S-T WTSE-1
type Node interface {
	TokenLiteral() string
	String() string
}
