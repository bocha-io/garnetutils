package ast

import "github.com/buger/jsonparser"

const UnaryOperation = "UnaryOperation"

func (a *Converter) processUnaryOperation(data []byte) (string, error) {
	operator, err := jsonparser.GetString(data, "operator")
	if err != nil {
		return "", err
	}

	subExpressionObject, _, _, err := jsonparser.Get(data, "subExpression")
	if err != nil {
		return "", err
	}
	subExpression, err := a.processNodeType(subExpressionObject)
	if err != nil {
		return "", err
	}

	switch operator {
	case "!":
		return operator + subExpression, nil
	case "--":
		return subExpression + operator, nil
	case "++":
		return subExpression + operator, nil
	}

	return operator + subExpression, nil
}
