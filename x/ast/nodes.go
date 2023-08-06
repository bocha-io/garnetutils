package ast

import (
	"github.com/buger/jsonparser"
)

func (a *Converter) processNodeType(data []byte) (string, error) {
	nodeType, err := jsonparser.GetString(data, "nodeType")
	if err != nil {
		return "", err
	}

	switch nodeType {
	case VariableDeclarationStatement:
		return a.processVariableDeclarationStatement(data)
	case BinaryOperation:
		return a.processBinaryOperation(data)
	case Identifier:
		return a.processIdentifier(data)
	case Return:
		return a.processReturn(data)
	case Block:
		return a.processBlock(data)
	case ContractDefinition:
		return a.processContractDefinition(data)
	case FunctionDefinition:
		return a.processFunctionDefinition(data)
	case ParameterList:
		return a.processParameterList(data)
	case FunctionCall:
		return a.processFunctionCall(data)
	case Literal:
		return a.processLiteral(data)
	case MemberAccess:
		return a.processMemberAccess(data)
	case IfStatement:
		return a.processIfStatement(data)
	case ExpressionStatement:
		return a.processExpressionStatement(data)
	case Assignment:
		return a.processAssignment(data)
	case TupleExpression:
		return a.processTupleExpression(data)
	case UnaryOperation:
		return a.processUnaryOperation(data)
	case IndexAccess:
		return a.processIndexAccess(data)
	case ElementaryTypeName:
		return a.processElementaryTypeName(data)
	case UserDefinedTypeName:
		return a.processUserDefinedTypeName(data)
	case ArrayTypeName:
		return a.processArrayTypeName(data)
	case ForStatement:
		return a.processForStatement(data)
	case Continue:
		return a.processContinue(data)
	case StructDefinition:
		return a.processStructDefinition(data)
	default:
		panic(nodeType + " not registered")
	}
}
