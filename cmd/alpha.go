/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bocha-io/garnetutils/x/ast"
	"github.com/bocha-io/garnetutils/x/converter"
	"github.com/spf13/cobra"
)

// alphaCmd represents the build command
var alphaCmd = &cobra.Command{
	Use:   "alpha",
	Short: "It generates helpers and predictions for your MUD contracts",
	Long: `This command is in alpha state, it was created to support the Eternal Legends
    contracts.
    There are some limitations:
    You can not use Conditionals in solidity, you must use If statements.
    You can not use init an array in the same line, you need to declare it and set position by position.
    Your functions must have unique names.
    Your structs must be declared inside the contract definition.


    This command will be improved in the future to remove the limitations.
    `,
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

		// Enums
		jsonFile := converter.MudConfigToJSON(mudConfigFile)
		enums := converter.GetEnumsFromJSON(jsonFile)
		ast.ProcessAllSolidityFiles(input, filepath.Join(input, "src"), output, enums)

		// Try to fmt the files if the user has gofmt installed
		_, _ = exec.Command("gofmt", "-w", output).Output()
	},
}

func init() {
	rootCmd.AddCommand(alphaCmd)

	alphaCmd.Flags().
		StringP("input", "i", "./", "Input: Path to your mud contracts (where the mud.config.ts file is located)")
	alphaCmd.Flags().
		StringP("output", "o", "/tmp/garnetgenerated/", "Output: Path to the output folder.")
}
