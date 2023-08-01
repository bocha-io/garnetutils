package converter

import "fmt"

type Converter struct {
	mainStruct string
}

func (c Converter) SingleValueString(tableName string) string {
	return fmt.Sprintf(`func (g *%s) get%s(rowID string) (data.Field, string, error) {
	return data.GetRowFromIDUsingString(g.db, g.world, rowID, "%s")
}`, c.mainStruct, tableName, tableName)
}

func (c Converter) SingleValueInt(tableName string) string {
	return fmt.Sprintf(`func (g *%s) get%s(key string) (int64, error) {
	return data.GetInt64UsingString(g.db, g.world, key, "%s")
}`, c.mainStruct, tableName, tableName)
}

func processFieldsForGetter(fields []Field) (string, string, []string) {
	errorReturn := ""
	returnValues := "("
	goFields := []string{}
	for _, v := range fields {
		// Uint, Int and Enums will return int64 in go
		goType := int64Type
		var tempReturn string
		switch v.Type {
		case bytes32Type:
			goType = stringType
			tempReturn = "\"\""
		case boolType:
			goType = boolType
			tempReturn = "false"
		default:
			tempReturn = "0"
		}

		// Function return types
		if returnValues == "(" {
			returnValues = fmt.Sprintf("%s%s", returnValues, goType)
		} else {
			returnValues = fmt.Sprintf("%s, %s", returnValues, goType)
		}

		// Empty values to return errors
		if errorReturn == "" {
			errorReturn = tempReturn
		} else {
			errorReturn = fmt.Sprintf("%s, %s", errorReturn, tempReturn)
		}

		// List to iterate and get fields
		goFields = append(goFields, goType)
	}
	returnValues += ", error)"
	return errorReturn, returnValues, goFields
}

func (c Converter) MultiValueTable(tableName string, fields []Field) string {
	errorReturn, returnValues, goFields := processFieldsForGetter(fields)

	firstLine := fmt.Sprintf(
		`func (g *%s) get%s(key string) %s {`,
		c.mainStruct,
		tableName,
		returnValues,
	)
	getValues := fmt.Sprintf(`
    fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "%s")
    if err != nil {
        return %s, err
    }`,
		tableName, errorReturn)

	checkLenght := fmt.Sprintf(`
    if len(fields) != %d {
        return %s, fmt.Errorf("invalid amount of fields")
    }`,
		len(fields), errorReturn)

	getters := ""
	validReturn := ""
	for k, v := range goFields {
		switch v {
		case int64Type:
			getters = fmt.Sprintf(`%s
    field%d, err := strconv.ParseInt(fields[%d].Data.String(), 10, 32)
    if err != nil {
        return %s, err
    }`, getters, k, k, errorReturn)

		case boolType:
			getters = fmt.Sprintf(`%s
    field%d := fields[%d].Data.String() == "true"`,
				getters, k, k)

		case stringType:
			getters = fmt.Sprintf(`%s
    field%d := strings.ReplaceAll(fields[%d].Data.String(), "\"", "")`,
				getters, k, k)
		}

		if validReturn == "" {
			validReturn = fmt.Sprintf("field%d", k)
		} else {
			validReturn = fmt.Sprintf("%s, field%d", validReturn, k)
		}

	}
	validReturn = fmt.Sprintf("    return %s, nil\n}", validReturn)

	return fmt.Sprintf(`
%s
%s
%s
%s
%s`,
		firstLine, getValues, checkLenght, getters, validReturn)
}
