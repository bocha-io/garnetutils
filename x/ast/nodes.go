package ast

import (
	"fmt"

	"github.com/buger/jsonparser"
)

const (
	// VariableDeclaration = "VariableDeclaration"
	// ElementaryTypeName = "ElementaryTypeName"

	IfStatement = "IfStatement"

	ExpressionStatement = "ExpressionStatement"
	Assignment          = "Assignment"

	TupleExpression = "TupleExpression"
)

const (
	Identifier = "Identifier"
	Literal    = "Literal"
)

const (
	UnaryOperation = "UnaryOperation"
)

func getNodeType(data []byte) (string, error) {
	return jsonparser.GetString(data, "nodeType")
}

func processIdentifier(data []byte) (string, error) {
	return jsonparser.GetString(data, "name")
}

// func processElementaryTypeName(data []byte) (string, error) {
// 	return jsonparser.GetString(data, "name")
// }

func processLiteral(data []byte) (string, error) {
	kind, err := jsonparser.GetString(data, "kind")
	if err != nil {
		return "", err
	}
	if kind == "number" {
		return jsonparser.GetString(data, "value")
	}
	return "", fmt.Errorf("%s literal not parsed", err)
}

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

func processExpressionStatement(data []byte) (string, error) {
	expression, _, _, err := jsonparser.Get(data, "expression")
	if err != nil {
		return "", err
	}

	expresionValue, err := processNodeType(expression)
	if err != nil {
		return "", err
	}

	return expresionValue, nil
}

func processAssignment(data []byte) (string, error) {
	operator, err := jsonparser.GetString(data, "operator")
	if err != nil {
		return "", err
	}

	leftExpression, _, _, err := jsonparser.Get(data, "leftHandSide")
	if err != nil {
		return "", err
	}
	leftSide, err := processNodeType(leftExpression)
	if err != nil {
		return "", err
	}

	rightExpression, _, _, err := jsonparser.Get(data, "rightHandSide")
	if err != nil {
		return "", err
	}
	rightSide, err := processNodeType(rightExpression)
	return leftSide + " " + operator + " " + rightSide, nil
}

func processTupleExpression(data []byte) (string, error) {
	components := []string{}
	_, err := jsonparser.ArrayEach(
		data,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			val, err := processNodeType(value)
			if err != nil {
				return
			}
			components = append(components, val)
		},
		"components",
	)
	if err != nil {
		return "", nil
	}

	// ret := "("
	ret := ""
	for k, v := range components {
		ret += v
		if k != len(components)-1 {
			ret += ", "
		}
	}
	// ret += ")"

	return ret, nil
}

func processUnaryOperation(data []byte) (string, error) {
	operator, err := jsonparser.GetString(data, "operator")
	if err != nil {
		return "", err
	}

	subExpressionObject, _, _, err := jsonparser.Get(data, "subExpression")
	if err != nil {
		return "", err
	}
	subExpression, err := processNodeType(subExpressionObject)
	if err != nil {
		return "", err
	}

	switch operator {
	case "!":
		return operator + subExpression, nil
	case "--":
		return subExpression + operator, nil
	case "++":
		return subExpression + operator, nil
	}

	return operator + subExpression, nil
}

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
	case FunctionCall:
		return processFunctionCall(data)
	case Literal:
		return processLiteral(data)
	case MemberAccess:
		return processMemberAccess(data)
	case IfStatement:
		return processIfStatement(data)
	case ExpressionStatement:
		return processExpressionStatement(data)
	case Assignment:
		return processAssignment(data)

	case TupleExpression:
		return processTupleExpression(data)

	case UnaryOperation:
		return processUnaryOperation(data)

	// case VariableDeclaration:
	// 	return processVariableDeclaration(data)
	// case ElementaryTypeName:
	// 	return processElementaryTypeName(data)

	default:
		fmt.Println(nodeType)
	}

	return "", nil
}
