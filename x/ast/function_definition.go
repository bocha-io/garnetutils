package ast

import (
	"fmt"
	"strings"

	"github.com/bocha-io/garnetutils/x/converter"
	"github.com/buger/jsonparser"
)

const FunctionDefinition = "FunctionDefinition"

func (a *ASTConverter) processFunctionDefinition(data []byte) (string, error) {
	functionName, err := jsonparser.GetString(data, "name")
	if err != nil {
		return "", err
	}
	functionHeader := "func " + "(p " + converter.PredictionObject + ") " + functionName

	// Parameters
	parameters, _, _, err := jsonparser.Get(data, "parameters")
	if err != nil {
		return "", err
	}

	parametersString, err := a.processNodeType(parameters)
	if err != nil {
		return "", err
	}
	functionParameters := " (" + parametersString + ") "

	// Returns
	returns, _, _, err := jsonparser.Get(data, "returnParameters")
	if err != nil {
		return "", err
	}

	returnsString, err := a.processNodeType(returns)
	if err != nil {
		return "", err
	}

	functionReturns := " {"
	if returnsString != "" {
		functionReturns = " (" + returnsString + ") {"
	}

	// Function body
	body, _, _, err := jsonparser.Get(data, "body")
	if err != nil {
		return "", err
	}

	statements := ""
	_, err = jsonparser.ArrayEach(
		body,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			nodeString, errProcess := a.processNodeType(value)
			if errProcess != nil {
				return
			}
			statements = fmt.Sprintf("%s\n%s", statements, nodeString)
		},
		"statements",
	)
	if err != nil {
		return "", nil
	}

	fixedStatements := strings.ReplaceAll(statements, "p._msgSender()", "senderAddress")
	if statements != fixedStatements {
		functionParameters = strings.Replace(functionParameters, ")", ", senderAddress string)", 1)
	}

	// statements
	return fmt.Sprintf("%s%s%s\n%s\n}", functionHeader, functionParameters, functionReturns, fixedStatements), nil
}
