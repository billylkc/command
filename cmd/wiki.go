package cmd

import (
	"fmt"
	"os"

	"github.com/billylkc/command/command"
	"github.com/billylkc/command/util"
	"github.com/spf13/cobra"
)

// wikiCmd represents the wiki command
var wikiCmd = &cobra.Command{
	Use:     "wiki [query]",
	Short:   "[w] Quick summary from the wiki page.",
	Long:    `[w] Quick summary from the wiki page.`,
	Aliases: []string{"w"},
	Example: `
  command wiki logistic regression
  command w regression
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		q := util.ParseQuery(args...)
		result, err := command.GetWiki(q)
		if err != nil {
			return err
		}

		fmt.Printf("%s (%s)\n\n", result.Title, result.Link)
		fmt.Println(result.Summary)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(wikiCmd)
}
