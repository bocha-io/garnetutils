/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/bocha-io/garnetutils/x/ast"
	"github.com/bocha-io/garnetutils/x/converter"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		input, err := cmd.Flags().GetString("input")
		if err != nil {
			fmt.Println("error reading input flag")
			return
		}

		mudConfigFile, err := os.ReadFile(filepath.Join(input, "mud.config.ts"))
		if err != nil {
			fmt.Printf("error opening the config: %s\n", err.Error())
			return
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println("error reading output flag")
			return
		}

		err = converter.GenerateFiles("GameObject", mudConfigFile, output)
		if err != nil {
			fmt.Printf("error generating files: %s", err.Error())
		}

		// Attack file
		attack, err := os.ReadFile(filepath.Join(input, "out", "AttackSystem.sol", "AttackSystem.json"))
		if err != nil {
			fmt.Printf("error opening the config: %s\n", err.Error())
			return
		}

		// Convert to JSON
		jsonFile := converter.MudConfigToJSON(mudConfigFile)
		// Enums
		enums := converter.GetEnumsFromJSON(jsonFile)
		astConvereter := ast.NewASTConverter()
		astConvereter.Enums = enums

		val, err := astConvereter.ProcessAST(attack)
		if err != nil {
			fmt.Printf("error generating ast: %s", err.Error())
		}

		if output[len(output)-1] != '/' {
			output += "/"
		}

		val = "package garnethelpers\n\n" + val

		// Replace the getkeyswithvalue module
		quotesRegex := regexp.MustCompile(`p\.get(Keys)WithValue\(([A-Za-z]+)TableId, p\.[A-Za-z]+\(([A-Za-z0-9, ]+)\)\)`)
		val = quotesRegex.ReplaceAllString(val, "p.$2$1($3)")

		if err := os.WriteFile(output+"attack.go", []byte(val), 0o600); err != nil {
			return
		}

		// LibCover
		libcover, err := os.ReadFile(filepath.Join(input, "out", "LibCover.sol", "LibCover.json"))
		if err != nil {
			fmt.Printf("error opening the config: %s\n", err.Error())
			return
		}

		val, err = astConvereter.ProcessAST(libcover)
		if err != nil {
			fmt.Printf("error generating ast: %s", err.Error())
		}

		if output[len(output)-1] != '/' {
			output += "/"
		}

		val = "package garnethelpers\n\n" + val
		if err := os.WriteFile(output+"libcover.go", []byte(val), 0o600); err != nil {
			return
		}

		// endMatch
		endMatch, err := os.ReadFile(filepath.Join(input, "out", "endMatch.sol", "endMatch.json"))
		if err != nil {
			fmt.Printf("error opening the config: %s\n", err.Error())
			return
		}

		val, err = astConvereter.ProcessAST(endMatch)
		if err != nil {
			fmt.Printf("error generating ast: %s", err.Error())
		}

		if output[len(output)-1] != '/' {
			output += "/"
		}

		val = "package garnethelpers\n\n" + val
		if err := os.WriteFile(output+"endmatch.go", []byte(val), 0o600); err != nil {
			return
		}

	},
}

func ReadFiles(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		return
	}

	for _, file := range files {
		if file.IsDir() {
			ReadFiles(filepath.Join(path, file.Name()))
		} else {
			fmt.Println(file.Name())
		}
	}
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().
		StringP("input", "i", "./", "Input: Path to your mud contracts (where the mud.config.ts file is located)")
	buildCmd.Flags().
		StringP("output", "o", "/tmp/garnetgenerated/", "Output: Path to the output folder.")
}
