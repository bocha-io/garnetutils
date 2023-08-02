package ast

import (
	"fmt"

	"github.com/buger/jsonparser"
)

const Literal = "Literal"

func (a *ASTConverter) processLiteral(data []byte) (string, error) {
	kind, err := jsonparser.GetString(data, "kind")
	if err != nil {
		return "", err
	}
	if kind == "number" || kind == "bool" {
		return jsonparser.GetString(data, "value")
	}

	if kind == "string" {
		val, err := jsonparser.GetString(data, "value")
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`"%s"`, val), nil
	}

	return "", fmt.Errorf("%s literal not parsed", err)
}
