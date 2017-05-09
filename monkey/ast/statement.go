package ast

// Statement represents a Statement in the A-S-T WTSE-1
type Statement interface {
	Node
	statementNode()
}
