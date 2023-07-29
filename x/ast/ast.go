package ast

import (
	"fmt"
	"strings"

	"github.com/buger/jsonparser"
)

const (
	pragmaDirective    = "PragmaDirective"
	importDirective    = "ImportDirective"
	contractDefinition = "ContractDefinition"
)

type SymbolImport struct {
	path    string
	symbols []string
}

func getNodes(abiData []byte) ([][]byte, error) {
	nodes := [][]byte{}
	if _, err := jsonparser.ArrayEach(
		abiData,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			nodes = append(nodes, value)
		},
		"ast", "nodes",
	); err != nil {
		return [][]byte{}, err
	}

	return nodes, nil
}

func processImport(symbolData []byte) (SymbolImport, error) {
	absolutePath, err := jsonparser.GetString(symbolData, "absolutePath")
	if err != nil {
		return SymbolImport{}, err
	}

	symbols := []string{}

	_, err = jsonparser.ArrayEach(
		symbolData,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			symbolName, err := jsonparser.GetString(value, "foreign", "name")
			if err != nil {
				panic(err)
			}
			symbols = append(symbols, string(symbolName))

		},
		"symbolAliases",
	)
	if err != nil {
		return SymbolImport{}, err
	}

	return SymbolImport{path: absolutePath, symbols: symbols}, nil
}

func generateGoImports(symbols []SymbolImport) string {
	ret := ""
	for _, v := range symbols {
		if strings.Contains(v.path, "node_modules") {
			// ignore node modules imports
			continue
		}

		if strings.Contains(v.path, "tables") {
			continue
		}

		if strings.Contains(v.path, "types") {
			continue
		}

		if strings.Contains(v.path, "src") {
			// TODO: import the function from that file
			continue
		}

	}
	return ret
}

func processDeclarations(data []byte) ([]string, error) {
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
	return declarations, err
}

// const (
// 	VariableDeclarationStatement = "VariableDeclarationStatement"
// 	BinaryOperation              = "BinaryOperation"
// )
//
// const (
// 	OperatorAnd = "&&"
// )

func processInitialValue(data []byte) string {
	val, err := processNodeType(data)
	if err != nil {
		panic(err)
	}
	return val
	// val, err := jsonparser.GetString(data, "nodeType")
	// if err != nil {
	// 	return
	// }
	// fmt.Println("nodetype", string(val))
	// if val == VariableDeclarationStatement {
	// 	initialValue, _, _, err := jsonparser.Get(data, "initialValue")
	// 	if err != nil {
	// 		return
	// 	}
	//
	// 	val, err := jsonparser.GetString(initialValue, "nodeType")
	// 	if err != nil {
	// 		return
	// 	}
	// 	if string(val) == BinaryOperation {
	// 		val, err := jsonparser.GetString(initialValue, "operator")
	// 		if err != nil {
	// 			return
	// 		}
	// 		if string(val) == OperatorAnd {
	// 			leftExpression, _, _, err := jsonparser.Get(initialValue, "leftExpression")
	// 			if err != nil {
	// 				return
	// 			}
	// 			val, err := jsonparser.GetString(leftExpression, "nodeType")
	// 			if err != nil {
	// 				return
	// 			}
	// 			fmt.Println(val)
	//
	// 			rightExpression, _, _, err := jsonparser.Get(initialValue, "rightExpression")
	// 			if err != nil {
	// 				return
	// 			}
	// 			val, err = jsonparser.GetString(rightExpression, "nodeType")
	// 			if err != nil {
	// 				return
	// 			}
	// 			fmt.Println(val)
	//
	// 		}
	// 	}
	//
	// leftExpressionValue, _, _, err := jsonparser.Get(data, "leftExpression")
	// if err != nil {
	// 	return
	// }
	// val, err := jsonparser.GetString(leftExpressionValue, "nodeType")
	// fmt.Println(val)

	// READ FIRST THE nodeType to check if there is left and right or only left

	// data,
	// func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
	// leftExpressionValue, _, _, _ := jsonparser.Get(initialValue, "leftExpression")
	// fmt.Printf("left value: %s", string(leftExpressionValue))

	// 	},
	// 	"initialValue",
	// )
	// if err != nil {
	// 	panic(err.Error())
	// }

}

// func processStatements(data []byte) {
// 	_, err := jsonparser.ArrayEach(
// 		data,
// 		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
// 			fmt.Println(processNodeType(value))
// 			// declarations, err := processDeclarations(value)
// 			// // TODO: maybe loop here
// 			// if err == nil && len(declarations) != 0 {
// 			// 	fmt.Println(declarations[0] + " " + processInitialValue(value))
// 			// } else {
// 			//              // No declaration
// 			//
// 			//
// 			// }
//
// 		},
// 		"statements",
// 	)
// 	if err != nil {
// 		panic(err.Error())
// 	}
//
// }

// func processContractDefinition(data []byte) error {
// 	_, err := jsonparser.ArrayEach(
// 		data,
// 		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
// 			nodeType, err := jsonparser.GetString(value, "nodeType")
// 			if err != nil {
// 				panic(err)
// 			}
// 			switch string(nodeType) {
// 			case FunctionDefinition:
// 				body, _, _, _ := jsonparser.Get(value, "body")
// 				// fmt.Println(string(body))
// 				processStatements(body)
// 			}
// 			fmt.Println(string(nodeType))
//
// 		},
// 		"nodes",
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func ProcessAST(data []byte) error {
	imports := []SymbolImport{}
	definition := []byte{}
	nodes, err := getNodes(data)
	if err != nil {
		return err
	}

	for _, v := range nodes {
		value, err := jsonparser.GetString(v, "nodeType")
		if err != nil {
			return err
		}

		switch value {
		case importDirective:
			importData, err := processImport(v)
			if err != nil {
				return err
			}
			imports = append(imports, importData)

		case contractDefinition:
			// processContractDefinition(v)
			a, err := processNodeType(v)
			if err != nil {
				return err
			}
			fmt.Println("----")
			fmt.Println(a)
			definition = v
		}
	}

	_ = definition

	// for _, v := range imports {
	// 	fmt.Println(v)
	// }

	return nil
}
