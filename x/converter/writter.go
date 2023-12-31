package converter

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func CreateTypesString(mainStruct string) string {
	return fmt.Sprintf(`// File autogenerated with garnetutils. DO NOT EDIT

package garnethelpers

import "github.com/bocha-io/garnet/x/indexer/data"

type %s struct {
	db     *data.Database
	world  *data.World
    active bool
}

func New%s(db *data.Database) *%s {
	return &%s{
		db:     db,
		world:  db.GetDefaultWorld(),
        active: true,
	}
}
`, mainStruct, mainStruct, mainStruct, mainStruct)
}

func CreateTablesString(tables []Table) string {
	ret := "// File autogenerated with garnetutils. DO NOT EDIT\n\npackage garnethelpers\n\nconst (\n"
	// Tables
	for _, v := range tables {
		ret += fmt.Sprintf("%sTableName = \"%s\"\n", v.Key, v.Key)
	}
	ret += ")"

	return strings.ReplaceAll(ret, "    ", "\t")
}

func CreateGettersString(tables []Table, c Converter) string {
	functionsString := ""
	// Getters
	for _, v := range tables {
		functionsString += fmt.Sprintf("\n%s", c.MultiValueTable(v.Key, v.Values, v.Singleton))
		functionsString += fmt.Sprintf("\n%s", c.GetRows(v.Key, v.Values))
	}

	gettersFile := "// File autogenerated with garnetutils. DO NOT EDIT\n\npackage garnethelpers\n\nimport (\n"

	if strings.Contains(functionsString, "fmt") {
		gettersFile += "\t\"fmt\"\n"
	}

	if strings.Contains(functionsString, "strconv") {
		gettersFile += "\t\"strconv\"\n"
	}

	if strings.Contains(functionsString, "strings") {
		gettersFile += "\t\"strings\"\n"
	}

	gettersFile += "\n\t\"github.com/bocha-io/garnet/x/indexer/data\"\n)"

	gettersFile += functionsString

	return strings.ReplaceAll(gettersFile, "    ", "\t")
}

func CreateEventsString(tables []Table, c Converter) string {
	eventsString := ""
	// Events
	for _, v := range tables {
		eventsString += fmt.Sprintf("\n%s", c.CreateEventFunction(v.Key, v.Values))
	}

	eventsFile := "// File autogenerated with garnetutils. DO NOT EDIT\n\npackage garnethelpers\n\nimport (\n"

	if strings.Contains(eventsString, "big.") {
		eventsFile += "\t\"math/big\"\n\n"
	}
	eventsFile += "\t\"github.com/ethereum/go-ethereum/common/hexutil\"\n"
	eventsFile += "\t\"github.com/bocha-io/garnet/x/indexer/data\"\n)"

	eventsFile += `
func BytesEventFromString(val string) []byte{
    ret, err := hexutil.Decode(val)
    if err != nil {
        panic(err.Error())
    }
    return ret
}
`

	eventsFile += eventsString

	return strings.ReplaceAll(eventsFile, "    ", "\t")
}

func CreateHelpersString(tables []Table, enums []Enum) string {
	helpersString := ""
	// Events
	for _, v := range tables {
		helpersString += fmt.Sprintf("\n%s", CreateHelper(v.Key, v.Values, v.Singleton, enums))
	}

	eventsFile := "// File autogenerated with garnetutils. DO NOT EDIT\n\npackage garnethelpers\n\nimport (\n\t\"strings\"\n\n\t\"github.com/bocha-io/garnet/x/indexer/data\"\n)\n\n"
	eventsFile += "//nolint govet\nconst EmptyBytes = string(int64(0))\n"
	eventsFile += CreateHelperStruct()

	eventsFile += helpersString

	return strings.ReplaceAll(eventsFile, "    ", "\t")
}

func GenerateFiles(mainStruct string, mudConfig []byte, path string) error {
	if path == "" {
		path = "/Users/hanchon/devel/bocha-io/garnetutils/x/garnethelpers/"
	}

	if path[len(path)-1] != '/' {
		path += "/"
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Convert to JSON
	jsonFile := MudConfigToJSON(mudConfig)
	// Tables
	tables := GetTablesFromJSON(jsonFile)
	// Enums
	enums := GetEnumsFromJSON(jsonFile)

	c := Converter{mainStruct: mainStruct}

	gettersString := CreateGettersString(tables, c)
	if err := os.WriteFile(path+"getters.go", []byte(gettersString), 0o600); err != nil {
		return err
	}

	tablesString := CreateTablesString(tables)
	if err := os.WriteFile(path+"tables.go", []byte(tablesString), 0o600); err != nil {
		return err
	}

	eventsString := CreateEventsString(tables, c)
	if err := os.WriteFile(path+"setters.go", []byte(eventsString), 0o600); err != nil {
		return err
	}

	helpers := CreateHelpersString(tables, enums)
	if err := os.WriteFile(path+"helpers.go", []byte(helpers), 0o600); err != nil {
		return err
	}

	enumsString := "// File autogenerated with garnetutils. DO NOT EDIT\n\npackage garnethelpers\n\n"
	for _, v := range enums {
		for k, e := range v.Values {
			if k == 0 {
				enumsString += fmt.Sprintf("const (\n\t%s = iota\n", e)
			} else {
				enumsString += "\t" + e + "\n"
			}
			if k == len(v.Values)-1 {
				enumsString += ")\n\n"
			}
		}
	}
	if err := os.WriteFile(path+"enums.go", []byte(enumsString), 0o600); err != nil {
		return err
	}

	return os.WriteFile(path+"types.go", []byte(CreateTypesString(c.mainStruct)), 0o600)
}
