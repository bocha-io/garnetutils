package ast

import (
	"fmt"

	"github.com/buger/jsonparser"
)

const ParameterList = "ParameterList"

func processParameterList(data []byte) (string, error) {
	parameters := []string{}
	_, err := jsonparser.ArrayEach(
		data,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			name, err := jsonparser.GetString(value, "name")
			if err != nil {
				return
			}
			typeName, err := jsonparser.GetString(value, "typeName", "name")
			if err != nil {
				return
			}
			parameters = append(parameters, fmt.Sprintf("%s %s", name, typeName))
		},
		"parameters",
	)
	if err != nil {
		return "", nil
	}

	ret := ""
	for k, v := range parameters {
		ret += v
		if k != len(parameters)-1 {
			ret += ", "
		}
	}
	return ret, nil
}