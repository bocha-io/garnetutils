package converter

import (
	"fmt"
	"strings"
)

func CreateHelper(tableName string, fields []Field) string {
	ret := fmt.Sprintf(`type %s struct {
}

`, tableName)

	_, returnValues, _ := processFieldsForGetter(fields)
	returnValues = strings.Replace(returnValues, ", error", "", 1)

	getValues := ""
	for k := range fields {
		getValues += fmt.Sprintf("field%d", k)
		if len(fields)-1 != k {
			getValues += ", "
		}
	}

	ret += fmt.Sprintf(`func (%s) get(key string) %s {
    if !BlockchainConnection.active {
        panic("game object is not active")
    }
    %s, err := BlockchainConnection.get%s(key)
    if err != nil {
        panic("value not found")
    }
    return %s
}
`, tableName, returnValues, getValues, tableName, getValues)

	params := createSettersReturnsValues(fields)
	args := ""
	for k, v := range fields {
		args += v.Key
		if len(fields)-1 != k {
			args += ", "
		}
	}

	ret += fmt.Sprintf(`
func (%s) set(ID string, %s) data.MudEvent {
    return Create%sEvent(ID, %s)
}
`, tableName, params, tableName, args)

	return ret
}
