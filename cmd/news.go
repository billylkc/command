package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// newsCmd represents the news command
var newsCmd = &cobra.Command{
	Use:     "news",
	Short:   "[n] A collection of different news sources.",
	Long:    `[n] A collection of different news sources.`,
	Aliases: []string{"n"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(newsCmd)
}
