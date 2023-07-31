package ast

import "github.com/buger/jsonparser"

const MemberAccess = "MemberAccess"

func processMemberAccess(data []byte) (string, error) {
	member, err := jsonparser.GetString(data, "memberName")
	if err != nil {
		return "", err
	}

	expressionObject, _, _, err := jsonparser.Get(data, "expression")
	if err != nil {
		return "", err
	}

	expression, err := processNodeType(expressionObject)
	if err != nil {
		return "", err
	}

	return expression + "." + member, nil
}
