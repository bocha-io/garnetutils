package ast

import "github.com/buger/jsonparser"

const Assignment = "Assignment"

func processAssignment(data []byte) (string, error) {
	operator, err := jsonparser.GetString(data, "operator")
	if err != nil {
		return "", err
	}

	leftExpression, _, _, err := jsonparser.Get(data, "leftHandSide")
	if err != nil {
		return "", err
	}
	leftSide, err := processNodeType(leftExpression)
	if err != nil {
		return "", err
	}

	rightExpression, _, _, err := jsonparser.Get(data, "rightHandSide")
	if err != nil {
		return "", err
	}
	rightSide, err := processNodeType(rightExpression)
	if err != nil {
		return "", err
	}

	return leftSide + " " + operator + " " + rightSide, nil
}
