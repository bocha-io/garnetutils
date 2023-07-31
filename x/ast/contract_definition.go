package ast

import (
	"fmt"

	"github.com/buger/jsonparser"
)

const ContractDefinition = "ContractDefinition"

func processContractDefinition(data []byte) (string, error) {
	ret := ""
	_, err := jsonparser.ArrayEach(
		data,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			nodeString, err := processNodeType(value)
			if err != nil {
				return
			}
			ret = fmt.Sprintf("%s\n%s", ret, nodeString)
		},
		"nodes",
	)
	if err != nil {
		return "", err
	}
	return ret, nil
}
