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

func processBranches(data []byte) (string, string, error) {
	leftExpression, _, _, err := jsonparser.Get(data, "leftExpression")
	if err != nil {
		return "", "", err
	}
	leftside, err := processNodeType(leftExpression)
	if err != nil {
		return "", "", err
	}

	rightExpression, _, _, err := jsonparser.Get(data, "rightExpression")
	if err != nil {
		return "", "", err
	}
	rightSide, err := processNodeType(rightExpression)
	return leftside, rightSide, err

}

func processBinaryOperation(data []byte) (string, error) {
	operator, err := jsonparser.GetString(data, "operator")
	if err != nil {
		return "", err
	}

	switch operator {
	// Maybe we need to support another type of operator in the future
	default:
		// it has left and right side
		left, right, err := processBranches(data)
		if err != nil {
			return "", err
		}
		return left + " " + operator + " " + right, nil
	}
}
