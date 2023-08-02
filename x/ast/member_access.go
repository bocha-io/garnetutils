package ast

import "github.com/buger/jsonparser"

const MemberAccess = "MemberAccess"

func (a *ASTConverter) processMemberAccess(data []byte) (string, error) {
	member, err := jsonparser.GetString(data, "memberName")
	if err != nil {
		return "", err
	}

	expressionObject, _, _, err := jsonparser.Get(data, "expression")
	if err != nil {
		return "", err
	}

	expression, err := a.processNodeType(expressionObject)
	if err != nil {
		return "", err
	}

	// Remove enum name if it's a class defined by MUD
	for _, v := range a.Enums {
		if expression == v.Key {
			return member, nil
		}
	}

	return expression + "." + member, nil
}
