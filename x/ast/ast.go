package ast

import (
	"strings"

	"github.com/bocha-io/garnetutils/x/converter"
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
			symbolName, errParser := jsonparser.GetString(value, "foreign", "name")
			if errParser != nil {
				// TODO: pass this error to the rest to the outside function
				return
			}
			symbols = append(symbols, symbolName)
		},
		"symbolAliases",
	)
	if err != nil {
		return SymbolImport{}, err
	}

	return SymbolImport{path: absolutePath, symbols: symbols}, nil
}

func GenerateGoImports(symbols []SymbolImport) string {
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

type ASTConverter struct {
	imports []SymbolImport
	Enums   []converter.Enum
}

func NewASTConverter() *ASTConverter {
	return &ASTConverter{
		imports: []SymbolImport{},
		Enums:   []converter.Enum{},
	}
}

func (a *ASTConverter) ProcessAST(data []byte) (string, error) {
	definition := []byte{}
	nodes, err := getNodes(data)
	if err != nil {
		return "", err
	}

	ret := ""

	for _, v := range nodes {
		value, err := jsonparser.GetString(v, "nodeType")
		if err != nil {
			return "", err
		}

		switch value {
		case importDirective:
			importData, err := processImport(v)
			if err != nil {
				return "", err
			}
			a.imports = append(a.imports, importData)

		case contractDefinition:
			a, err := a.processNodeType(v)
			if err != nil {
				return "", err
			}
			// fmt.Println("----")
			// fmt.Println(a)
			ret += a + "\n"
			definition = v

		// helpers files where functions are outside a contract
		case FunctionDefinition:
			a, err := a.processNodeType(v)
			if err != nil {
				return "", err
			}
			ret += a + "\n"
		}
	}

	_ = definition

	return ret, nil
}
