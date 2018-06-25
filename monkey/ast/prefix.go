package ast

import (
	"bytes"

	"github.com/ssargent/interpreter/monkey/token"
)

// PrefixExpression WTSE-1
type PrefixExpression struct {
	Token    token.Token // the prefix token, eg !
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

//TokenLiteral WTSE-1
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
