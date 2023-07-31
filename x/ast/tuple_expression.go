package ast

import "github.com/buger/jsonparser"

const TupleExpression = "TupleExpression"

func processTupleExpression(data []byte) (string, error) {
	components := []string{}
	_, err := jsonparser.ArrayEach(
		data,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			val, err := processNodeType(value)
			if err != nil {
				return
			}
			components = append(components, val)
		},
		"components",
	)
	if err != nil {
		return "", nil
	}

	ret := ""
	for k, v := range components {
		ret += v
		if k != len(components)-1 {
			ret += ", "
		}
	}

	return ret, nil
}