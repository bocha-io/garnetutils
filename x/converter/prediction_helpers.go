package converter

import (
	"fmt"
	"strings"
)

const PredictionObject = "Prediction"

func CreateHelperStruct() string {
	return fmt.Sprintf(`type %s struct {
    events []data.MudEvent
}

func (%s) addressToEntityKey(address string) string{
        return strings.Replace(address, "0x", "0x000000000000000000000000", 1)
}
`, PredictionObject, PredictionObject)
	// TODO: add NewPredictionObject function
}

func CreateHelper(tableName string, fields []Field, sigleton bool) string {
	ret := ""
	_, returnValues, _ := processFieldsForGetter(fields)
	returnValues = strings.Replace(returnValues, ", error", "", 1)

	getValues := ""
	for k := range fields {
		getValues += fmt.Sprintf("field%d", k)
		if len(fields)-1 != k {
			getValues += ", "
		}
	}

	argsGetter := "key string"
	key := "key"
	if sigleton {
		argsGetter = ""
		key = ""
	}

	ret += fmt.Sprintf(`func (%s) %sGet(%s) %s {
    if !BlockchainConnection.active {
        panic("game object is not active")
    }
    %s, err := BlockchainConnection.get%s(%s)
    if err != nil {
        panic("value not found")
    }
    return %s
}
`, PredictionObject, tableName, argsGetter, returnValues, getValues, tableName, key, getValues)

	params := createSettersReturnsValues(fields)
	args := ""
	for k, v := range fields {
		args += v.Key
		if len(fields)-1 != k {
			args += ", "
		}
	}

	ret += fmt.Sprintf(`
func (p %s) %sSet(ID string, %s) {
    p.events = append(p.events, Create%sEvent(ID, %s))
}
`, PredictionObject, tableName, params, tableName, args)

	ret += fmt.Sprintf(`
func (p %s) %sDeleteRecord(ID string) {
    p.events = append(p.events, Delete%sEvent(ID))
}
`, PredictionObject, tableName, tableName)

	ret += fmt.Sprintf(`
func (p %s) getKeys%s(ID string) {
    p.events = append(p.events, Delete%sEvent(ID))
}
`, PredictionObject, tableName, tableName)

	return ret
}
