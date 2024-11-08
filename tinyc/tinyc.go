package tinyc

import "fmt"

type Program struct {
	Statements []Statement
}

func (p Program) Generate() string {
	var result string
	for _, stmt := range p.Statements {
		result += stmt.Generate() + "\n"
	}
	return result
}

type Statement interface {
	Generate() string
}

type ExprStatement struct {
	Expression Expression
}

func (es ExprStatement) Generate() string {
	return es.Expression.Generate() + ";"
}

type IfStatement struct {
	Condition Expression
	ThenStmt  Statement
	ElseStmt  Statement
}

func (ifs IfStatement) Generate() string {
	result := fmt.Sprintf("if (%s) %s", ifs.Condition.Generate(), ifs.ThenStmt.Generate())
	if ifs.ElseStmt != nil {
		result += " else " + ifs.ElseStmt.Generate()
	}
	return result
}

type WhileStatement struct {
	Condition Expression
	Body      Statement
}

func (ws WhileStatement) Generate() string {
	return fmt.Sprintf("while (%s) %s", ws.Condition.Generate(), ws.Body.Generate())
}

type Expression interface {
	Generate() string
}

type AssignExpr struct {
	ID    string
	Right Expression
}

func (ae AssignExpr) Generate() string {
	return fmt.Sprintf("%s = %s", ae.ID, ae.Right.Generate())
}

type TestExpr struct {
	Left  Expression
	Op    string // Set this to "<" if it’s a comparison, or leave it empty otherwise
	Right Expression
}

func (te TestExpr) Generate() string {
	if te.Op != "" {
		return fmt.Sprintf("%s %s %s", te.Left.Generate(), te.Op, te.Right.Generate())
	}
	return te.Left.Generate()
}

type SumExpr struct {
	Left  Expression // Left side of the sum or single term if `Op` is empty
	Op    string     // "+" or "-" for addition or subtraction, empty for a single term
	Right Expression
}

func (se SumExpr) Generate() string {
	if se.Op == "" {
		return se.Left.Generate()
	}
	return fmt.Sprintf("(%s %s %s)", se.Left.Generate(), se.Op, se.Right.Generate())
}

type TermExpr struct {
	ID        string
	Integer   string
	ParenExpr Expression
}

func (te TermExpr) Generate() string {
	if te.ID != "" {
		return te.ID
	} else if te.Integer != "" {
		return te.Integer
	} else if te.ParenExpr != nil {
		return fmt.Sprintf("(%s)", te.ParenExpr.Generate())
	}
	return ""
}
