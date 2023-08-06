package ast

import (
	"github.com/bocha-io/garnetutils/x/converter"
	"github.com/buger/jsonparser"
)

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

func (a *Converter) processBranches(data []byte) (string, string, error) {
	leftExpression, _, _, err := jsonparser.Get(data, "leftExpression")
	if err != nil {
		return "", "", err
	}
	leftside, err := a.processNodeType(leftExpression)
	if err != nil {
		return "", "", err
	}

	// special case, leftExpression is bytes32 and rightExpression is number
	isSpecialCase := false
	leftType, err := jsonparser.GetString(data, "leftExpression", "typeDescriptions", "typeString")
	if err == nil && leftType == converter.Bytes32Type {
		isSpecialCase = true
	}

	rightExpression, _, _, err := jsonparser.Get(data, "rightExpression")
	if err != nil {
		return "", "", err
	}
	rightSide, err := a.processNodeType(rightExpression)
	if isSpecialCase {
		kind, err := jsonparser.GetString(data, "rightExpression", "kind")
		if err == nil && kind == "number" {
			rightSide = "string(" + rightSide + ")"
		}
	}

	return leftside, rightSide, err
}

func (a *Converter) processBinaryOperation(data []byte) (string, error) {
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
