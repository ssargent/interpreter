package ast

import "github.com/ssargent/interpreter/monkey/token"

// LetStatement is a statemnt that is a let statement WTSE-1
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral is a function
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
