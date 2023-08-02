package ast

import "github.com/buger/jsonparser"

const Identifier = "Identifier"

func (a *ASTConverter) processIdentifier(data []byte) (string, error) {
	return jsonparser.GetString(data, "name")
}
