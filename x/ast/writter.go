package ast

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bocha-io/garnetutils/x/converter"
)

// // Attack file
// attack, err := os.ReadFile(filepath.Join(input, "out", "AttackSystem.sol", "AttackSystem.json"))
//
//	if err != nil {
//	    fmt.Printf("error opening the config: %s\n", err.Error())
//	    return ""
//	}
//
// // Convert to JSON
// jsonFile := converter.MudConfigToJSON(mudConfigFile)
// // Enums
// enums := converter.GetEnumsFromJSON(jsonFile)
//
//	if output[len(output)-1] != '/' {
//	    output += "/"
//	}
func ProcessSolidityFiles(basePath string, fileName string, outputFolder string, enums []converter.Enum) error {
	path := filepath.Join(basePath, "out", fileName+".sol")
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			splited := strings.Split(file.Name(), ".")
			if len(splited) == 2 && splited[1] == "json" {
				if err := ProcessSolidityFile(filepath.Join(path, file.Name()), splited[0], outputFolder, enums); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func ProcessSolidityFile(path string, fileName string, outputFolder string, enums []converter.Enum) error {
	// content, err := OpenSolidityFile(path)
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	generated, err := GenerateGoFileFromSolidy(content, enums)
	if err != nil {
		return err
	}

	return SaveGoFile(outputFolder, fileName, generated)
}

func OpenSolidityFile(basePath string, fileName string) ([]byte, error) {
	return os.ReadFile(filepath.Join(basePath, "out", fileName+".sol", fileName+".json"))
}

func SaveGoFile(outputFolder string, fileName string, fileContent string) error {
	return os.WriteFile(filepath.Join(outputFolder, fileName+".go"), []byte(fileContent), 0o600)
}

func GenerateGoFileFromSolidy(file []byte, enums []converter.Enum) (string, error) {
	astConvereter := NewASTConverter()
	astConvereter.Enums = enums

	val, err := astConvereter.ProcessAST(file)
	if err != nil {
		return "", fmt.Errorf("error generating ast: %s", err.Error())
	}

	val = "package garnethelpers\n\n" + val
	// Replace the getkeyswithvalue module
	quotesRegex := regexp.MustCompile(`p\.get(Keys)WithValue\(([A-Za-z]+)TableId, p\.[A-Za-z]+\(([A-Za-z0-9, ]+)\)\)`)
	val = quotesRegex.ReplaceAllString(val, "p.$2$1($3)")

	return val, nil
}

func ProcessAllSolidityFiles(basePath string, currentPath string, destination string, enums []converter.Enum) {
	files, err := os.ReadDir(currentPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			if file.Name() == "tables" || file.Name() == "codegen" || file.Name() == "world" {
				continue
			}
			ProcessAllSolidityFiles(basePath, filepath.Join(currentPath, file.Name()), destination, enums)
		} else {
			// Ignore interfaces
			if file.Name()[0] == 'I' {
				continue
			}
			fmt.Println(file.Name())
			filename := strings.Split(file.Name(), ".")[0]
			fmt.Println(filename)
			if filename == "addressToEntityKey" {
				continue
			}
			err = ProcessSolidityFiles(basePath, filename, destination, enums)
			if err != nil {
				panic(err)
			}
		}
	}
}
