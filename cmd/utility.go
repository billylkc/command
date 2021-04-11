package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// utilCmd represents the util command
var utilityCmd = &cobra.Command{
	Use:     "utility",
	Short:   "[u] A collection of utility functions.",
	Long:    `[u] A collection of utility functions.`,
	Aliases: []string{"u"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(utilityCmd)
}
