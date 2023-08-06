package ast

import "github.com/buger/jsonparser"

const ForStatement = "ForStatement"

func (a *Converter) processForStatement(data []byte) (string, error) {
	ret := "for "
	// initializationExpression
	initializationExpressionObject, _, _, err := jsonparser.Get(data, "initializationExpression")
	if err != nil {
		return "", err
	}
	initExpression, err := a.processNodeType(initializationExpressionObject)
	if err != nil {
		return "", err
	}

	ret += initExpression + "; "

	// condition
	conditionObject, _, _, err := jsonparser.Get(data, "condition")
	if err != nil {
		return "", err
	}
	condition, err := a.processNodeType(conditionObject)
	if err != nil {
		return "", err
	}

	ret += condition + "; "

	// loopExpression
	loopExpressionObject, _, _, err := jsonparser.Get(data, "loopExpression")
	if err != nil {
		return "", err
	}
	loopExpression, err := a.processNodeType(loopExpressionObject)
	if err != nil {
		return "", err
	}

	ret += loopExpression + " {"

	statements, err := a.getStatementsFromBody(data)
	if err != nil {
		return "", err
	}

	ret += statements

	ret += "\n}"

	return ret, err
}
