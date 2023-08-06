package ast

import (
	"fmt"

	"github.com/bocha-io/garnetutils/x/converter"
	"github.com/bocha-io/garnetutils/x/utils"
	"github.com/buger/jsonparser"
)

const StructDefinition = "StructDefinition"

type member struct {
	name      string
	typeValue string
}

func (a *Converter) processStructDefinition(data []byte) (string, error) {
	members := []member{}
	_, err := jsonparser.ArrayEach(
		data,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			name, typeValue, errProcess := a.BytesToVariableDeclaration(value)
			if errProcess != nil {
				return
			}
			members = append(members, member{name: name, typeValue: typeValue})
		},
		"members",
	)
	if err != nil {
		return "", err
	}

	structName, err := jsonparser.GetString(data, "name")
	if err != nil {
		return "", err
	}

	ret := fmt.Sprintf("type %s struct {\n", structName)
	constructorParams := "("
	initValues := ""
	for k, v := range members {
		ret += fmt.Sprintf("%s %s\n", v.name, utils.SolidityTypeToGolang(v.typeValue, converter.GetEnumKeys(a.Enums)))
		constructorParams += fmt.Sprintf("%s %s", v.name, utils.SolidityTypeToGolang(v.typeValue, converter.GetEnumKeys(a.Enums)))
		initValues += "temp." + v.name + "=" + v.name + "\n"
		if k != len(members)-1 {
			constructorParams += ", "
		}
	}
	constructorParams += ")"
	ret += "}\nfunc New" + structName + constructorParams + " " + structName + " {\ntemp:=" + structName + "{}\n" + initValues + "return temp\n}\n"

	return ret, nil
}
