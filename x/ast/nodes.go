package ast

import (
	"fmt"

	"github.com/buger/jsonparser"
)

const (
	VariableDeclarationStatement = "VariableDeclarationStatement"
	BinaryOperation              = "BinaryOperation"
	Return                       = "Return"
	Block                        = "Block"
	ContractDefinition           = "ContractDefinition"

	FunctionDefinition = "FunctionDefinition"
	ParameterList      = "ParameterList"

	VariableDeclaration = "VariableDeclaration"
	ElementaryTypeName  = "ElementaryTypeName"
)

const (
	Identifier = "Identifier"
)

// Operators
const (
	OperatorAnd = "&&"
	OperatorGE  = ">="
	OperatorAdd = "+"
)

func getNodeType(data []byte) (string, error) {
	return jsonparser.GetString(data, "nodeType")
}

func processVariableDeclarationStatement(data []byte) (string, error) {
	// This only supports one var at the time
	declarations := []string{}
	_, err := jsonparser.ArrayEach(
		data,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			// isConstant, err := jsonparser.GetBoolean(value, "constant")
			name, err := jsonparser.GetString(value, "name")
			typeName, err := jsonparser.GetString(value, "typeName", "name")
			declarations = append(declarations, fmt.Sprintf("var %s %s", typeName, name))
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

		ret := ""
		for k, v := range declarations {
			ret += v + " = " + value
			if k != len(declarations)-1 {
				ret += "\n"
			}
		}
		return ret, nil
	}

	return "", fmt.Errorf("no declarations in this block")
}

func processBranches(data []byte) (string, string, error) {
	leftExpression, _, _, err := jsonparser.Get(data, "leftExpression")
	if err != nil {
		return "", "", err
	}
	leftside, err := processNodeType(leftExpression)
	if err != nil {
		return "", "", err
	}

	rightExpression, _, _, err := jsonparser.Get(data, "rightExpression")
	if err != nil {
		return "", "", err
	}
	rightSide, err := processNodeType(rightExpression)
	return leftside, rightSide, err

}

func processBinaryOperation(data []byte) (string, error) {
	operator, err := jsonparser.GetString(data, "operator")
	if err != nil {
		return "", err
	}

	switch operator {
	case OperatorAnd:
		left, right, err := processBranches(data)
		if err != nil {
			return "", err
		}
		return left + " " + operator + " " + right, nil

	case OperatorGE:
		left, right, err := processBranches(data)
		if err != nil {
			return "", err
		}
		return left + " " + operator + " " + right, nil

	case OperatorAdd:
		left, right, err := processBranches(data)
		if err != nil {
			return "", err
		}
		return left + " " + operator + " " + right, nil
	}

	return "", nil

}

func processIdentifier(data []byte) (string, error) {
	return jsonparser.GetString(data, "name")
}

func processReturn(data []byte) (string, error) {
	expression, _, _, err := jsonparser.Get(data, "expression")
	if err != nil {
		return "", err
	}

	expresionValue, err := processNodeType(expression)
	if err != nil {
		return "", err
	}

	return "return " + expresionValue, nil
}

func processBlock(data []byte) (string, error) {
	statements := []string{}
	_, err := jsonparser.ArrayEach(
		data,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			statement, err := processNodeType(value)
			if err != nil {
				return
			}
			statements = append(statements, statement)
		},
		"statements",
	)
	if err != nil {
		return "", err
	}

	ret := ""
	for k, v := range statements {
		ret += v
		if k != len(statements)-1 {
			ret += "\n"
		}
	}
	return ret, nil
}

func processContractDefinition(data []byte) (string, error) {
	ret := ""
	_, err := jsonparser.ArrayEach(
		data,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			nodeString, err := processNodeType(value)
			if err != nil {
				return
			}
			ret = fmt.Sprintf("%s\n%s", ret, nodeString)
		},
		"nodes",
	)
	if err != nil {
		return "", err
	}
	return ret, nil
}

func processFunctionDefinition(data []byte) (string, error) {
	functionName, err := jsonparser.GetString(data, "name")
	if err != nil {
		return "", err
	}
	ret := "func " + functionName

	// Parameters
	parameters, _, _, err := jsonparser.Get(data, "parameters")
	if err != nil {
		return "", err
	}

	parametersString, err := processNodeType(parameters)
	if err != nil {
		return "", err
	}
	ret += " (" + parametersString + ") "

	// Returns
	returns, _, _, err := jsonparser.Get(data, "returnParameters")
	if err != nil {
		return "", err
	}

	returnsString, err := processNodeType(returns)
	if err != nil {
		return "", err
	}
	ret += " (" + returnsString + ") {"

	// Function body
	body, _, _, err := jsonparser.Get(data, "body")
	if err != nil {
		return "", err
	}

	_, err = jsonparser.ArrayEach(
		body,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			nodeString, err := processNodeType(value)
			if err != nil {
				return
			}
			ret = fmt.Sprintf("%s\n%s", ret, nodeString)
		},
		"statements",
	)
	if err != nil {
		return "", nil
	}

	// Close function
	ret += "\n}"
	return ret, nil
}

func processParameterList(data []byte) (string, error) {
	parameters := []string{}

	_, err := jsonparser.ArrayEach(
		data,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			// isConstant, err := jsonparser.GetBoolean(value, "constant")
			name, err := jsonparser.GetString(value, "name")
			if err != nil {
				return
			}
			typeName, err := jsonparser.GetString(value, "typeName", "name")
			if err != nil {
				return
			}
			// typeName, _, _, err := jsonparser.Get(data, "typeName")
			// if err != nil {
			// 	return "", err
			// }
			//
			// typeValue, err := processNodeType(typeName)
			// if err != nil {
			// 	return "", err
			// }
			parameters = append(parameters, fmt.Sprintf("%s %s", name, typeName))
		},
		"parameters",
	)
	if err != nil {
		return "", nil
	}

	ret := ""
	for k, v := range parameters {
		ret += v
		if k != len(parameters)-1 {
			ret += ", "
		}
	}
	return ret, nil

}

// func processElementaryTypeName(data []byte) (string, error) {
// 	return jsonparser.GetString(data, "name")
// }

func processNodeType(data []byte) (string, error) {
	// fmt.Println("processing node type", string(data))
	nodeType, err := getNodeType(data)
	if err != nil {
		return "", err
	}

	switch nodeType {
	case VariableDeclarationStatement:
		return processVariableDeclarationStatement(data)
	case BinaryOperation:
		return processBinaryOperation(data)
	case Identifier:
		return processIdentifier(data)
	case Return:
		return processReturn(data)
	case Block:
		return processBlock(data)
	case ContractDefinition:
		return processContractDefinition(data)
	case FunctionDefinition:
		return processFunctionDefinition(data)
	case ParameterList:
		return processParameterList(data)
	// case VariableDeclaration:
	// 	return processVariableDeclaration(data)
	// case ElementaryTypeName:
	// 	return processElementaryTypeName(data)

	default:
		fmt.Println(nodeType)
	}

	return "", nil
}
