package ast

import (
	"bytes"
	"strings"

	"github.com/ssargent/interpreter/monkey/token"
)

// FunctionLiteral is WTSE-1
type FunctionLiteral struct {
	Token      token.Token // the 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

//TokenLiteral is wtse-1
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }

//String
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}
