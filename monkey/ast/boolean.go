package ast

import "github.com/ssargent/interpreter/monkey/token"

//Boolean wtse-1
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

//TokenLiteral WTSE-1
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }

//WTSE-1
func (b *Boolean) String() string { return b.Token.Literal }
