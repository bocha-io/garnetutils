package ast

import "github.com/buger/jsonparser"

const Identifier = "Identifier"

func (a *Converter) processIdentifier(data []byte) (string, error) {
	return jsonparser.GetString(data, "name")
}
