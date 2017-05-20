package ast

import (
	"bytes"

	"github.com/ssargent/interpreter/monkey/token"
)

//ReturnStatement WTSE-1
type ReturnStatement struct {
	Token       token.Token // the return Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral WTSE-1
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

//String WTSE-1
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}
