package ast

import "github.com/buger/jsonparser"

const BinaryOperation = "BinaryOperation"

// Operators
const (
	OperatorAnd = "&&"
	OperatorGE  = ">="
	OperatorLE  = "<="
	OperatorAdd = "+"
	OperatorSub = "-"
	OperatorMul = "*"
	OperatorL   = "<"
	OperatorG   = ">"
	OperatorNE  = "!="
	OperatorE   = "=="
)

func (a *ASTConverter) processBranches(data []byte) (string, string, error) {
	leftExpression, _, _, err := jsonparser.Get(data, "leftExpression")
	if err != nil {
		return "", "", err
	}
	leftside, err := a.processNodeType(leftExpression)
	if err != nil {
		return "", "", err
	}

	rightExpression, _, _, err := jsonparser.Get(data, "rightExpression")
	if err != nil {
		return "", "", err
	}
	rightSide, err := a.processNodeType(rightExpression)
	return leftside, rightSide, err
}

func (a *ASTConverter) processBinaryOperation(data []byte) (string, error) {
	operator, err := jsonparser.GetString(data, "operator")
	if err != nil {
		return "", err
	}

	left, right, err := a.processBranches(data)
	if err != nil {
		return "", err
	}
	return left + " " + operator + " " + right, nil
}
