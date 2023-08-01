package ast

import (
	"fmt"

	"github.com/buger/jsonparser"
)

const FunctionCall = "FunctionCall"

func processFunctionCall(data []byte) (string, error) {
	kind, err := jsonparser.GetString(data, "kind")
	if err != nil {
		return "", err
	}

	if kind == "typeConversion" {
		// nodeType: ElementaryTypeNameExpression
		// arguments
		arguments := []string{}
		_, err := jsonparser.ArrayEach(
			data,
			func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				argument, errProcess := processNodeType(value)
				if errProcess != nil {
					return
				}
				arguments = append(arguments, argument)
			},
			"arguments",
		)
		if err != nil {
			return "", nil
		}

		ret := "("
		for k, v := range arguments {
			ret += v
			if k != len(arguments)-1 {
				ret += ", "
			}
		}
		ret += ")"

		// expression
		val, err := jsonparser.GetString(data, "expression", "nodeType")
		if err != nil {
			return "", err
		}

		funcType := ""
		if val == "ElementaryTypeNameExpression" {
			funcType, err = jsonparser.GetString(data, "expression", "typeName", "name")
			if err != nil {
				return "", err
			}
		}
		return funcType + ret, nil

	} else if kind == "functionCall" {
		// arguments
		arguments := []string{}
		_, err := jsonparser.ArrayEach(
			data,
			func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				argument, errProcess := processNodeType(value)
				if errProcess != nil {
					return
				}
				arguments = append(arguments, argument)
			},
			"arguments",
		)
		if err != nil {
			return "", err
		}

		ret := "("
		if len(arguments) > 0 {
			for k, v := range arguments {
				ret += v
				if k != len(arguments)-1 {
					ret += ", "
				}
			}
		}
		ret += ")"

		// expression
		expressionObject, _, _, err := jsonparser.Get(data, "expression")
		if err != nil {
			return "", err
		}
		expression, err := processNodeType(expressionObject)
		if err != nil {
			return "", err
		}
		if expression == "require" {
			if len(arguments) != 2 {
				return "", fmt.Errorf("invalid arguments for require")
			}

			return fmt.Sprintf(`if !(%s) {
    panic(%s)
 }`, arguments[0], arguments[1]), nil
		}
		return expression + ret, nil
	}

	return "", fmt.Errorf("%s function kind not processed", kind)
}
