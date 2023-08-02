package ast

import "github.com/buger/jsonparser"

const Assignment = "Assignment"

func (a *ASTConverter) processAssignment(data []byte) (string, error) {
	operator, err := jsonparser.GetString(data, "operator")
	if err != nil {
		return "", err
	}

	leftExpression, _, _, err := jsonparser.Get(data, "leftHandSide")
	if err != nil {
		return "", err
	}
	leftSide, err := a.processNodeType(leftExpression)
	if err != nil {
		return "", err
	}

	rightExpression, _, _, err := jsonparser.Get(data, "rightHandSide")
	if err != nil {
		return "", err
	}
	rightSide, err := a.processNodeType(rightExpression)
	if err != nil {
		return "", err
	}

	return leftSide + " " + operator + " " + rightSide, nil
}
