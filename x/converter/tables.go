package converter

import (
	"fmt"
	"strings"

	"github.com/buger/jsonparser"
)

func GetTablesFromJSON(mudConfigJSON []byte) []Table {
	tables := []Table{}
	err := jsonparser.ObjectEach(
		mudConfigJSON,
		func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			if dataType == jsonparser.String {
				tables = append(tables, Table{
					Key: string(key),
					Values: []Field{
						{
							Key:  defaultKeyForField,
							Type: string(value),
						},
					},
				})
			} else if dataType == jsonparser.Object {
				table := Table{
					Key:    string(key),
					Values: []Field{},
				}
				err := jsonparser.ObjectEach(value, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
					table.Values = append(table.Values, Field{Key: string(key), Type: string(value)})
					return nil
				}, schemaName)
				tables = append(tables, table)
				if err != nil {
					return err
				}
			}
			return nil
		},
		configName,
		tablesName,
	)
	if err != nil {
		panic(fmt.Sprintf("could not create table data: %s", err.Error()))
	}

	return tables
}

func GetEnumsFromJSON(mudConfigJSON []byte) []Enum {
	enums := []Enum{}
	err := jsonparser.ObjectEach(
		mudConfigJSON,
		func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			e := Enum{
				Key:    string(key),
				Values: []string{},
			}

			if dataType == jsonparser.Array {
				_, err := jsonparser.ArrayEach(
					mudConfigJSON,
					func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
						e.Values = append(e.Values, string(value))
					},
					configName,
					enumsName,
					string(key),
				)
				if err != nil {
					return err
				}
			}
			enums = append(enums, e)
			return nil
		},
		configName,
		enumsName,
	)

	// If there is no enums in the ts file, the ObjectEach will return an error, we ignore that error
	if err != nil && !strings.Contains(err.Error(), "Key path not found") {
		panic(fmt.Sprintf("could not create enum data: %s", err.Error()))
	}

	return enums
}
