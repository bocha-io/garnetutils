package converter

import "fmt"

func createSettersReturnsValues(fields []Field) string {
	returnValues := ""
	for _, v := range fields {
		goType := int64Type
		switch v.Type {
		case bytes32Type:
			goType = "string"
		case boolType:
			goType = boolType
		}
		// Function return types
		if returnValues == "" {
			returnValues = fmt.Sprintf("%s %s", v.Key, goType)
		} else {
			returnValues = fmt.Sprintf("%s, %s %s", returnValues, v.Key, goType)
		}
	}
	return returnValues
}

func (c *Converter) CreateEventFunction(tableName string, fields []Field) string {
	returnValues := createSettersReturnsValues(fields)

	firstLine := fmt.Sprintf(`
func Create%sEvent(ID string, %s) data.MudEvent {
    return data.MudEvent{
        Table: "%s",
        Key:   ID,
        Fields: []data.Field{`, tableName, returnValues, tableName)

	fieldsEvents := "\n"
	for _, v := range fields {
		// Uint, Int and Enums will return int64 in go
		dataString := fmt.Sprintf(`data.UintField{Data: *big.NewInt(%s)}`, v.Key)
		switch v.Type {
		case bytes32Type:
			dataString = fmt.Sprintf(`data.NewBytesField(%s)`, v.Key)
		case boolType:
			dataString = fmt.Sprintf(`data.BoolField{Data: %s}`, v.Key)
		}
		fieldsEvents += fmt.Sprintf("            {Key: \"%s\", Data: %s},\n", v.Key, dataString)
	}
	fieldsEvents += "        },\n    }"

	setter := fmt.Sprintf("%s%s\n}", firstLine, fieldsEvents)

	remover := fmt.Sprintf(`

func Delete%sEvent(ID string) data.MudEvent {
    return data.MudEvent{
        Table:  "%s",
        Key:    ID,
        Fields: []data.Field{},
    }
}
`, tableName, tableName)

	return setter + remover
}
