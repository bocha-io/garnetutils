package converter

import (
	"fmt"
)

type Converter struct {
	mainStruct string
}

func (c Converter) SingleValueString(tableName string) string {
	return fmt.Sprintf(`func (g *%s) Get%s(rowID string) (data.Field, string, error) {
	return data.GetRowFromIDUsingString(g.db, g.world, rowID, "%s")
}`, c.mainStruct, tableName, tableName)
}

func (c Converter) SingleValueInt(tableName string) string {
	return fmt.Sprintf(`func (g *%s) Get%s(key string) (int64, error) {
	return data.GetInt64UsingString(g.db, g.world, key, "%s")
}`, c.mainStruct, tableName, tableName)
}

func processFieldsForGetter(fields []Field) (string, string, []string) {
	errorReturn := ""
	returnValues := "("
	goFields := []string{}
	for _, v := range fields {
		// Uint, Int and Enums will return int64 in go
		goType := Int64Type
		var tempReturn string
		switch v.Type {
		case Bytes32Type:
			goType = StringType
			tempReturn = "\"\""
		case BoolType:
			goType = BoolType
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

func (c Converter) createProcessFieldFunction(
	tableName string,
	returnValues string,
	fields []Field,
	errorReturn string,
	goFields []string,
) string {
	checkLenght := fmt.Sprintf(`if len(fields) != %d {
        return %s, fmt.Errorf("invalid amount of fields")
    }`,
		len(fields), errorReturn)

	getters := ""
	validReturn := ""
	for k, v := range goFields {
		switch v {
		case Int64Type:
			getters = fmt.Sprintf(`%s
    field%d, err := strconv.ParseInt(fields[%d].Data.String(), 10, 32)
    if err != nil {
        return %s, err
    }`, getters, k, k, errorReturn)

		case BoolType:
			getters = fmt.Sprintf(`%s
    field%d := fields[%d].Data.String() == "true"`,
				getters, k, k)

		case StringType:
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

	return fmt.Sprintf(`func (g *%s) ProcessFields%s(fields []data.Field) %s {
%s
%s
%s
`, c.mainStruct, tableName, returnValues, checkLenght, getters, validReturn)
}

func (c Converter) MultiValueTable(tableName string, fields []Field, singleton bool) string {
	errorReturn, returnValues, goFields := processFieldsForGetter(fields)
	args := "(key string)"
	key := ""
	if singleton {
		args = "()"
		key = "\n    key := \"\""
	}

	processFunction := c.createProcessFieldFunction(
		tableName,
		returnValues,
		fields,
		errorReturn,
		goFields,
	)

	firstLine := fmt.Sprintf(
		`func (g *%s) Get%s%s %s {%s`,
		c.mainStruct,
		tableName,
		args,
		returnValues,
		key,
	)
	getValues := fmt.Sprintf(
		`    fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "%s")
    if err != nil {
        return %s, err
    }`,
		tableName,
		errorReturn,
	)

	return fmt.Sprintf(`
%s
%s
%s
	return g.ProcessFields%s(fields)
}
`,
		processFunction, firstLine, getValues, tableName)
}

func (c Converter) GetRows(tableName string, fields []Field) string {
	_, _, goFields := processFieldsForGetter(fields)

	args := ""
	for k, v := range goFields {
		args += fmt.Sprintf("arg%d %s", k, v)
		if k != len(goFields)-1 {
			args += ", "
		}
	}

	zeroLine := fmt.Sprintf(
		`func (g %s) GetAllRows%s() map[string][]data.Field{
	table := g.world.GetTableByName("%s")
	return g.db.GetRows(table)
}
`,
		c.mainStruct,
		tableName,
		tableName,
	)

	firstLine := fmt.Sprintf(
		`func (g %s) GetRows%s(%s) []string{
    rows := g.GetAllRows%s()
 	for k, fields := range rows {`,
		c.mainStruct,
		tableName,
		args,
		tableName,
	)

	checkLenght := fmt.Sprintf(`        if len(fields) != %d {
            continue
        }`,
		len(fields))

	getters := ""
	for k, v := range goFields {
		switch v {
		case Int64Type:
			getters = fmt.Sprintf(`%s
        field%d, err := strconv.ParseInt(fields[%d].Data.String(), 10, 32)
        if err != nil {
            continue
        }
        if field%d != arg%d {
            continue
        }`, getters, k, k, k, k)

		case BoolType:
			getters = fmt.Sprintf(`%s
        field%d := fields[%d].Data.String() == "true"
        if field%d != arg%d {
            continue
        }`, getters, k, k, k, k)

		case StringType:
			getters = fmt.Sprintf(`%s
        field%d := strings.ReplaceAll(fields[%d].Data.String(), "\"", "")
        if field%d != arg%d {
            continue
        }`, getters, k, k, k, k)
		}
	}
	getters += ("\n        return []string{k}\n    }")

	return fmt.Sprintf(`
%s
%s
%s
%s
    return []string{}
}
`,
		zeroLine, firstLine, checkLenght, getters)
}
