package ast

import "github.com/buger/jsonparser"

const Return = "Return"

func (a *Converter) processReturn(data []byte) (string, error) {
	expression, _, _, err := jsonparser.Get(data, "expression")
	if err != nil {
		return "", err
	}

	expresionValue, err := a.processNodeType(expression)
	if err != nil {
		return "", err
	}

	return "return " + expresionValue, nil
}
