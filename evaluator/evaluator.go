package evaluator

import (
	"bangu/ast"
	"bangu/object"
)

func Eval(node ast.Node) object.Object {
	switch n := node.(type) {
	// Statements
	case *ast.Program:
		return evalStatements(n.Statements)
	case *ast.ExpressionStatement:
		return Eval(n.Expression)
	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: n.Value}

	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt)
	}

	return result
}
