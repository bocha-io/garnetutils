package ast

import "github.com/buger/jsonparser"

const ExpressionStatement = "ExpressionStatement"

func (a *Converter) processExpressionStatement(data []byte) (string, error) {
	expression, _, _, err := jsonparser.Get(data, "expression")
	if err != nil {
		return "", err
	}

	expresionValue, err := a.processNodeType(expression)
	if err != nil {
		return "", err
	}

	return expresionValue, nil
}
