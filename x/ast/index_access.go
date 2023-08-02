package ast

import (
	"fmt"

	"github.com/buger/jsonparser"
)

const IndexAccess = "IndexAccess"

func (a *ASTConverter) processIndexAccess(data []byte) (string, error) {
	// Base
	baseExpression, _, _, err := jsonparser.Get(data, "baseExpression")
	if err != nil {
		return "", err
	}

	base, err := a.processNodeType(baseExpression)
	if err != nil {
		return "", err
	}

	// Index
	indexExpression, _, _, err := jsonparser.Get(data, "indexExpression")
	if err != nil {
		return "", err
	}

	value, err := a.processNodeType(indexExpression)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`%s[%s]`, base, value), nil
}
