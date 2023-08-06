package ast

import (
	"fmt"

	"github.com/buger/jsonparser"
)

const (
	UserDefinedTypeName = "UserDefinedTypeName"
	ElementaryTypeName  = "ElementaryTypeName"
	ArrayTypeName       = "ArrayTypeName"
)

func (a *ASTConverter) processUserDefinedTypeName(data []byte) (string, error) {
	return jsonparser.GetString(data, "pathNode", "name")
}

func (a *ASTConverter) processElementaryTypeName(data []byte) (string, error) {
	return jsonparser.GetString(data, "name")
}

func (a *ASTConverter) processArrayTypeName(data []byte) (string, error) {
	baseType, _, _, err := jsonparser.Get(data, "baseType")
	if err != nil {
		return "", err
	}
	typeString, err := a.processNodeType(baseType)
	if err != nil {
		return "", err
	}

	length, err := jsonparser.GetString(data, "length", "value")
	if err != nil {
		return fmt.Sprintf("[]%s", typeString), nil
	}

	return fmt.Sprintf("[%s]%s", length, typeString), nil
}
