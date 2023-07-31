package ast

import (
	"fmt"

	"github.com/buger/jsonparser"
)

const Literal = "Literal"

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
