package ast

import (
	"fmt"
	"strings"

	"github.com/bocha-io/garnetutils/x/converter"
	"github.com/buger/jsonparser"
)

const FunctionDefinition = "FunctionDefinition"

func (a *Converter) getStatementsFromBody(data []byte) (string, error) {
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
			if statements == "" {
				statements = nodeString
			} else {
				statements = fmt.Sprintf("%s\n%s", statements, nodeString)
			}
		},
		"statements",
	)
	if err != nil {
		return "", err
	}

	return statements, nil
}

func (a *Converter) processFunctionDefinition(data []byte) (string, error) {
	functionName, err := jsonparser.GetString(data, "name")
	if err != nil {
		return "", err
	}
	functionHeader := "func " + "(p *" + converter.PredictionObject + ") " + functionName

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

	statements, err := a.getStatementsFromBody(data)
	if err != nil {
		return "", err
	}

	fixedStatements := strings.ReplaceAll(statements, "p._msgSender()", "senderAddress")
	if statements != fixedStatements {
		if functionParameters == " () " {
			functionParameters = strings.Replace(
				functionParameters,
				")",
				"senderAddress string)",
				1,
			)
		} else {
			functionParameters = strings.Replace(functionParameters, ")", ", senderAddress string)", 1)
		}
	}

	// statements
	return fmt.Sprintf(
		"%s%s%s\n%s\n}\n",
		functionHeader,
		functionParameters,
		functionReturns,
		fixedStatements,
	), nil
}
