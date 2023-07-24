package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "garnetutils",
	Short: "Garnet cli util to generate go files using a MUD config",
	Long:  `To generate files use: garnetutils generate -i ./mud.config.ts -o ./out/`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
