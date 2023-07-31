package ast

import "github.com/buger/jsonparser"

const Identifier = "Identifier"

func processIdentifier(data []byte) (string, error) {
	return jsonparser.GetString(data, "name")
}
