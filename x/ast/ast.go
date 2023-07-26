package ast

import (
	"fmt"

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
			fmt.Println(string(symbolName))
			symbols = append(symbols, string(symbolName))

		},
		"symbolAliases",
	)
	if err != nil {
		return SymbolImport{}, err
	}

	return SymbolImport{path: absolutePath, symbols: symbols}, nil
}

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
			definition = v
		}
	}

	_ = definition

	for _, v := range imports {
		fmt.Println(v)
	}

	return err
}
