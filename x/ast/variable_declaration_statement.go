package ast

import (
	"fmt"
	"strings"

	"github.com/bocha-io/garnetutils/x/converter"
	"github.com/bocha-io/garnetutils/x/utils"
	"github.com/buger/jsonparser"
)

const (
	VariableDeclarationStatement = "VariableDeclarationStatement"
)

func (a *ASTConverter) BytesToVariableDeclaration(value []byte) (name string, typeValue string, err error) {
	err = nil
	if string(value) == "null" {
		return "_", "_", nil
	}
	name, errInternal := jsonparser.GetString(value, "name")
	if errInternal != nil {
		return "", "", errInternal
	}
	typeNameObject, _, _, errInternal := jsonparser.Get(value, "typeName")
	if errInternal != nil {
		return "", "", errInternal
	}
	typeName, errInternal := a.processNodeType(typeNameObject)
	if errInternal != nil {
		return "", "", errInternal
	}

	return name, typeName, err
}

func (a *ASTConverter) processVariableDeclarationStatement(data []byte) (string, error) {
	declarations := []string{}
	_, err := jsonparser.ArrayEach(
		data,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			name, typeName, err := a.BytesToVariableDeclaration(value)
			if err != nil {
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
		isArray := ""

		initialValue, _, _, err := jsonparser.Get(data, "initialValue")
		if err != nil {
			// It has no initial value, it's just a var declaration
			val := ""
			for _, v := range declarations {
				splited := strings.SplitAfter(v, " ")
				if len(splited) == 2 {
					varType := utils.SolidityTypeToGolang(splited[0], converter.GetEnumKeys(a.Enums))
					val += "var " + splited[1] + " " + varType + "\n"
				}
			}
			return val, nil
		}

		// It has initial value
		value, err = a.processNodeType(initialValue)
		if err != nil {
			return "", err
		}

		ret := declarations[0]
		// remove type
		splited := strings.SplitAfter(ret, " ")
		if len(splited) == 2 {
			ret = splited[1]
		}

		// TODO: if we are creating a new array, we need to add the []type{} string, but if the function returns the array that's not needed
		// Check if there is a way to get from the ast if they are setting each position in the array
		// if len(strings.Split(splited[0], "]")) == 2 {
		// 	isArray = "[]" + utils.SolidityTypeToGolang(strings.Split(splited[0], "]")[1]) + "{"
		// }

		// if there is more than one declaration, it's a tuple
		if len(declarations) > 1 {
			ret = ""
			for k, v := range declarations {
				// remove type
				splited := strings.SplitAfter(v, " ")
				if len(splited) == 2 {
					v = splited[1]
				}

				ret += v
				if k != len(declarations)-1 {
					ret += ", "
				}
			}
			// ret += ")"
		}

		if isArray != "" {
			return ret + " := " + isArray + value + "}", nil
		}
		return ret + " := " + value, nil
	}

	return "", fmt.Errorf("no declarations in this block")
}
