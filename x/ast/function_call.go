package ast

import (
	"fmt"
	"strings"

	"github.com/buger/jsonparser"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const FunctionCall = "FunctionCall"

func (a *Converter) processFunctionCall(data []byte) (string, error) {
	kind, err := jsonparser.GetString(data, "kind")
	if err != nil {
		return "", err
	}

	switch kind {
	case "typeConversion":
		// nodeType: ElementaryTypeNameExpression
		// arguments
		arguments := []string{}
		_, err := jsonparser.ArrayEach(
			data,
			func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				argument, errProcess := a.processNodeType(value)
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
			if funcType == "bytes32" {
				funcType = "string"
			}
			if funcType == "uint32" {
				funcType = "int64"
			}
		}
		return funcType + ret, nil

	case "functionCall":
		// arguments
		arguments := []string{}
		_, err := jsonparser.ArrayEach(
			data,
			func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				argument, errProcess := a.processNodeType(value)
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
		expression, err := a.processNodeType(expressionObject)
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

		isMUDTable := false
		// Update the expression is it's using a MUD table
		for _, v := range a.imports {
			if isMUDTable {
				break
			}
			if strings.Contains(v.path, "tables") {
				for _, symbolName := range v.symbols {
					splited := strings.Split(expression, ".")
					if len(splited) == 2 {
						if splited[0] == symbolName {
							caser := cases.Title(language.English)
							expression = "p." + splited[0] + caser.String(splited[1])
							isMUDTable = true
							break
						}
					}
				}
			} else {
				for _, v := range strings.Split(v.path, "/") {
					// Remove extension from the file
					v = strings.Split(v, ".")[0]

					splited := strings.Split(expression, ".")
					if len(splited) == 2 {
						if splited[0] == v {
							expression = "p." + splited[1]
							isMUDTable = true
							break
						}
					}
				}
			}
		}

		if !isMUDTable {
			expression = "p." + expression
		}

		return expression + ret, nil
	case "structConstructorCall":
		// expression
		expression, _, _, err := jsonparser.Get(data, "expression")
		if err != nil {
			return "", err
		}

		expresionValue, err := a.processNodeType(expression)
		if err != nil {
			return "", err
		}

		ret := "New" + expresionValue

		// arguments
		arguments := []string{}
		_, err = jsonparser.ArrayEach(
			data,
			func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				argument, errProcess := a.processNodeType(value)
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

		ret += "("
		for k, v := range arguments {
			ret += v
			if k != len(arguments)-1 {
				ret += ", "
			}
		}
		ret += ")"
		return ret, nil
	default:
		return "", fmt.Errorf("%s function kind not processed", kind)
	}
}
