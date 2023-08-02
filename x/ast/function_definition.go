package ast

import (
	"fmt"

	"github.com/buger/jsonparser"
)

const FunctionDefinition = "FunctionDefinition"

func (a *ASTConverter) processFunctionDefinition(data []byte) (string, error) {
	functionName, err := jsonparser.GetString(data, "name")
	if err != nil {
		return "", err
	}
	ret := "func " + functionName

	// Parameters
	parameters, _, _, err := jsonparser.Get(data, "parameters")
	if err != nil {
		return "", err
	}

	parametersString, err := a.processNodeType(parameters)
	if err != nil {
		return "", err
	}
	ret += " (" + parametersString + ") "

	// Returns
	returns, _, _, err := jsonparser.Get(data, "returnParameters")
	if err != nil {
		return "", err
	}

	returnsString, err := a.processNodeType(returns)
	if err != nil {
		return "", err
	}

	if returnsString != "" {
		ret += " (" + returnsString + ")"
	}
	ret += " {"

	// Function body
	body, _, _, err := jsonparser.Get(data, "body")
	if err != nil {
		return "", err
	}

	_, err = jsonparser.ArrayEach(
		body,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			nodeString, errProcess := a.processNodeType(value)
			if errProcess != nil {
				return
			}
			ret = fmt.Sprintf("%s\n%s", ret, nodeString)
		},
		"statements",
	)
	if err != nil {
		return "", nil
	}

	// Close function
	ret += "\n}"
	return ret, nil
}
