package ast

import (
	"fmt"

	"github.com/bocha-io/garnetutils/x/utils"
	"github.com/buger/jsonparser"
)

const ParameterList = "ParameterList"

func (a *ASTConverter) processParameterList(data []byte) (string, error) {
	parameters := []string{}
	_, err := jsonparser.ArrayEach(
		data,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			name, errInternal := jsonparser.GetString(value, "name")
			if errInternal != nil {
				return
			}
			typeNameObject, _, _, errInternal := jsonparser.Get(value, "typeName")
			if errInternal != nil {
				return
			}
			typeName, errInternal := a.processNodeType(typeNameObject)
			if errInternal != nil {
				return
			}

			parameters = append(parameters, fmt.Sprintf("%s %s", name, utils.SolidityTypeToGolang(typeName)))
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
