package cmd

import (
	"fmt"
	"os"

	"github.com/bocha-io/garnetutils/x/converter"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the getters, events and types go files based on your MUD config",
	Long: `To generate files use:
    garnetutils generate -i ./mud.config.ts -o /tmp/garnetgenerated/

    -i param is your mud config path
    -o param is your destination folder
`,

	Run: func(cmd *cobra.Command, args []string) {
		input, err := cmd.Flags().GetString("input")
		if err != nil {
			fmt.Println("error reading input flag")
			return
		}

		mudConfigFile, err := os.ReadFile(input)
		if err != nil {
			fmt.Printf("error opening the file: %s\n", err.Error())
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

		fmt.Printf("files generated at %s\n", output)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().
		StringP("input", "i", "./mud.config.ts", "Input: Path to your mud.config.ts file.")
	generateCmd.Flags().
		StringP("output", "o", "/tmp/garnetgenerated/", "Output: Path to the output folder.")
}
