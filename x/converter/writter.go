package converter

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func CreateTypesString(mainStruct string) string {
	return fmt.Sprintf(`package garnethelpers

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

var BlockchainConnection GameObject
`, mainStruct, mainStruct, mainStruct, mainStruct)
}

func CreateGettersString(tables []Table, c Converter) string {
	functionsString := ""
	// Getters
	for _, v := range tables {
		functionsString += fmt.Sprintf("\n%s", c.MultiValueTable(v.Key, v.Values, v.Singleton))
	}

	gettersFile := "package garnethelpers\n\nimport (\n"

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

	eventsFile := "package garnethelpers\n\nimport (\n"

	if strings.Contains(eventsString, "big.") {
		eventsFile += "\t\"math/big\"\n"
	}

	eventsFile += "\n\t\"github.com/bocha-io/garnet/x/indexer/data\"\n)"

	eventsFile += eventsString

	return strings.ReplaceAll(eventsFile, "    ", "\t")
}

func CreateHelpersString(tables []Table, c Converter) string {
	helpersString := ""
	// Events
	for _, v := range tables {
		helpersString += fmt.Sprintf("\n%s", CreateHelper(v.Key, v.Values, v.Singleton))
	}

	eventsFile := "package garnethelpers\n\nimport (\n\t\"strings\"\n\n\t\"github.com/bocha-io/garnet/x/indexer/data\"\n)\n\n"
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

	eventsString := CreateEventsString(tables, c)
	if err := os.WriteFile(path+"setters.go", []byte(eventsString), 0o600); err != nil {
		return err
	}

	helpers := CreateHelpersString(tables, c)
	if err := os.WriteFile(path+"helpers.go", []byte(helpers), 0o600); err != nil {
		return err
	}

	enumsString := "package garnethelpers\n\n"
	for _, v := range enums {
		for k, e := range v.Values {
			if k == 0 {
				enumsString += fmt.Sprintf("const (\n\t%s = iota\n", e)
			} else {
				enumsString += "\t" + e + "\n"
			}
			if k == len(v.Values)-1 {
				enumsString += ")\n"
			}
		}

	}
	if err := os.WriteFile(path+"enums.go", []byte(enumsString), 0o600); err != nil {
		return err
	}

	return os.WriteFile(path+"types.go", []byte(CreateTypesString(c.mainStruct)), 0o600)
}
