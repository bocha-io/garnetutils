package converter

import (
	"fmt"
	"strings"

	"github.com/bocha-io/garnetutils/x/utils"
)

const PredictionObject = "Prediction"

func CreateHelperStruct() string {
	return fmt.Sprintf(`type %s struct {
    Events               []data.MudEvent
    BlockchainConnection %s

}

func New%s(db *data.Database) *%s {
	return &%s{
        Events:                []data.MudEvent{},
		BlockchainConnection:  *New%s(db),
	}
}

func (%s) addressToEntityKey(address string) string{
        return strings.Replace(address, "0x", "0x000000000000000000000000", 1)
}
`, PredictionObject, "GameObject", PredictionObject, PredictionObject, PredictionObject, "GameObject", PredictionObject)
}

func CreateHelper(tableName string, fields []Field, sigleton bool, enums []Enum) string {
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

	getFromPrediction := fmt.Sprintf(`
    var temp *data.MudEvent = nil
	for k, v := range p.Events {
		if v.Table == "%s" && v.Key == %s {
			temp = &p.Events[k]
		}
	}

	if temp != nil {
		%s, _ := p.BlockchainConnection.ProcessFields%s(temp.Fields)
		return %s
	}
    `, tableName, key, getValues, tableName, getValues)

	if sigleton {
		argsGetter = ""
		key = ""
		getFromPrediction = ""
	}

	ret += fmt.Sprintf(`func (p *%s) %sGet(%s) %s {
    %s
    if !p.BlockchainConnection.active {
        panic("game object is not active")
    }
    %s, _ := p.BlockchainConnection.Get%s(%s)
    return %s
}
`, PredictionObject, tableName, argsGetter, returnValues, getFromPrediction, getValues, tableName, key, getValues)

	params := createSettersReturnsValues(fields)
	args := ""
	keysArgs := ""
	for k, v := range fields {
		args += v.Key
		keysArgs += v.Key + " " + utils.SolidityTypeToGolang(v.Type, GetEnumKeys(enums))
		if len(fields)-1 != k {
			args += ", "
			keysArgs += ", "
		}
	}

	ret += fmt.Sprintf(`
func (p *%s) %sSet(ID string, %s) {
    p.Events = append(p.Events, Create%sEvent(ID, %s))
}
`, PredictionObject, tableName, params, tableName, args)

	ret += fmt.Sprintf(`
func (p *%s) %sDeleterecord(ID string) {
    p.Events = append(p.Events, Delete%sEvent(ID))
}
`, PredictionObject, tableName, tableName)

	// TODO: read keys using the prediction values

	ret += fmt.Sprintf(`
func (p *%s) %sKeys(%s) []string {
    if !p.BlockchainConnection.active {
        panic("game object is not active")
    }
    return p.BlockchainConnection.GetRows%s(%s)
}
`, PredictionObject, tableName, keysArgs, tableName, args)

	return ret
}
