package converter

import "fmt"

func (c *Converter) CreateEventFunction(tableName string, fields []Field) string {
	returnValues := ""
	for _, v := range fields {
		goType := int64Type
		switch v.Type {
		case bytes32Type:
			goType = "[]byte"
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

	firstLine := fmt.Sprintf(`
func create%sEvent(ID string, %s) data.MudEvent {
    return data.MudEvent{
        Table: "%s",
        Key: ID,
        Fields: []data.Field{`, tableName, returnValues, tableName)

	fieldsEvents := "\n"
	for _, v := range fields {
		// Uint, Int and Enums will return int64 in go
		dataString := fmt.Sprintf(`data.UintField{Data: *big.NewInt(%s)}`, v.Key)
		switch v.Type {
		case bytes32Type:
			dataString = fmt.Sprintf(`data.NewBytesField(%s)}`, v.Key)
		case boolType:
			dataString = fmt.Sprintf(`data.BoolField{Data: %s}`, v.Key)
		}
		fieldsEvents += fmt.Sprintf("            {Key: \"%s\", Data: %s},\n", v.Key, dataString)
	}
	fieldsEvents += "        },\n    },"

	return fmt.Sprintf("%s%s\n}", firstLine, fieldsEvents)
}
