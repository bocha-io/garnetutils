package ast

import (
	"github.com/buger/jsonparser"
)

func processNodeType(data []byte) (string, error) {
	nodeType, err := jsonparser.GetString(data, "nodeType")
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
	default:
		panic(nodeType + " not registered")
	}
}
