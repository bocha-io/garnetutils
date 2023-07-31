package ast

import "github.com/buger/jsonparser"

const IfStatement = "IfStatement"

func processIfStatement(data []byte) (string, error) {
	ret := "if "
	// condition

	conditionObject, _, _, err := jsonparser.Get(data, "condition")
	if err != nil {
		return "", err
	}
	condition, err := processNodeType(conditionObject)

	ret += condition + " {\n"

	// true
	trueBodyObject, _, _, err := jsonparser.Get(data, "trueBody")
	if err != nil {
		return "", err
	}
	trueBody, err := processNodeType(trueBodyObject)
	ret += trueBody
	ret += "\n}"

	// false
	falseBodyObject, dataType, _, _ := jsonparser.Get(data, "falseBody")
	if dataType != jsonparser.NotExist {
		ret += " else {\n"
		falseBody, err := processNodeType(falseBodyObject)
		if err != nil {
			return "", err
		}
		ret += falseBody + "\n}"

	}

	return ret, err
}
