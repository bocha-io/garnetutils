package converter

import (
	"fmt"
	"os"
	"strings"
)

func CreateTypesString() string {
	return `package garnethelpers

import "github.com/bocha-io/garnet/x/indexer/data"

type GameState struct {
	db    *data.Database
	world *data.World
}

func NewGameState(db *data.Database) *GameState {
	return &GameState{
		db:    db,
		world: db.GetDefaultWorld(),
	}
}
`
}

func CreateGettersString(tables []Table, c Converter) string {
	functionsString := ""
	// Getters
	for _, v := range tables {
		functionsString += fmt.Sprintf("\n%s", c.MultiValueTable(v.Key, v.Values))
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

func GenerateFiles(mainStruct string, mudConfig []byte, path string) error {
	if path == "" {
		path = "/Users/hanchon/devel/bocha-io/transpiler/x/garnethelpers/"
	}

	if path[len(path)-1] != '/' {
		path += "/"
	}

	// Convert to JSON
	jsonFile := MudConfigToJSON(mudConfig)
	// Tables
	tables := GetTablesFromJSON(jsonFile)

	c := Converter{mainStruct: mainStruct}

	gettersString := CreateGettersString(tables, c)
	if err := os.WriteFile(path+"getters.go", []byte(gettersString), 0644); err != nil {
		return err
	}

	eventsString := CreateEventsString(tables, c)
	if err := os.WriteFile(path+"setters.go", []byte(eventsString), 0644); err != nil {
		return err
	}

	if err := os.WriteFile(path+"types.go", []byte(CreateTypesString()), 0644); err != nil {
		return err
	}

	return nil
}
