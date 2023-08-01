package ast

import (
	"fmt"

	"github.com/buger/jsonparser"
)

const (
	VariableDeclarationStatement = "VariableDeclarationStatement"
)

func processVariableDeclarationStatement(data []byte) (string, error) {
	// This only supports one var at the time
	declarations := []string{}
	_, err := jsonparser.ArrayEach(
		data,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			name, errInternal := jsonparser.GetString(value, "name")
			if errInternal != nil {
				return
			}
			typeNameObject, _, _, errInternal := jsonparser.Get(value, "typeName")
			if errInternal != nil {
				return
			}
			typeName, errInternal := processNodeType(typeNameObject)
			if errInternal != nil {
				return
			}
			declarations = append(declarations, fmt.Sprintf("%s %s", typeName, name))
		},
		"declarations",
	)
	if err != nil {
		return "", nil
	}

	value := ""
	if len(declarations) != 0 {
		initialValue, _, _, err := jsonparser.Get(data, "initialValue")
		if err != nil {
			return "", err
		}

		value, err = processNodeType(initialValue)
		if err != nil {
			return "", err
		}

		ret := declarations[0]
		// if there is more than one declaration, it's a tuple
		if len(declarations) > 1 {
			ret = "("
			for k, v := range declarations {
				ret += v
				if k != len(declarations)-1 {
					ret += ", "
				}
			}
			ret += ")"
		}

		ret += " := " + value
		return ret, nil
	}

	return "", fmt.Errorf("no declarations in this block")
}
