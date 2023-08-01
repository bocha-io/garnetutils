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

func processUserDefinedTypeName(data []byte) (string, error) {
	return jsonparser.GetString(data, "pathNode", "name")
}

func processElementaryTypeName(data []byte) (string, error) {
	return jsonparser.GetString(data, "name")
}

func processArrayTypeName(data []byte) (string, error) {
	baseType, _, _, err := jsonparser.Get(data, "baseType")
	if err != nil {
		return "", err
	}
	typeString, err := processNodeType(baseType)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("[]%s", typeString), nil
}
