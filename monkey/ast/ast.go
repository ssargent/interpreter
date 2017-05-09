package ast

// Program represents the entire A-S-T
type Program struct {
	Statements []Statement
}

//TokenLiteral WTSE-1
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}
