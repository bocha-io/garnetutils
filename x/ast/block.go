package ast

import "github.com/buger/jsonparser"

const Block = "Block"

func (a *ASTConverter) processBlock(data []byte) (string, error) {
	statements := []string{}
	_, err := jsonparser.ArrayEach(
		data,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			statement, errProcess := a.processNodeType(value)
			if errProcess != nil {
				return
			}
			statements = append(statements, statement)
		},
		"statements",
	)
	if err != nil {
		return "", err
	}

	ret := ""
	for k, v := range statements {
		ret += v
		if k != len(statements)-1 {
			ret += "\n"
		}
	}
	return ret, nil
}
