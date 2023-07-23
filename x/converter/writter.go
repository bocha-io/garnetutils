package converter

import (
	"fmt"
	"os"
	"strings"
)

func GenerateFiles(mainStruct string, mudConfig []byte, path string) []string {
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

	gettersFile = strings.ReplaceAll(gettersFile, "    ", "\t")
	_ = os.WriteFile(path+"getters.go", []byte(gettersFile), 0644)

	fmt.Println(gettersFile)

	return []string{""}
}
