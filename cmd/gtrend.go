package cmd

import (
	"os"
	"strings"

	"github.com/billylkc/command/command"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var country string

// gtrendCmd represents the gtrend command
var gtrendCmd = &cobra.Command{
	Use:     "gtrend",
	Aliases: []string{"tr"},
	Short:   "[tr] Daily keywords from google trend.",
	Long:    `[tr] Daily keywords from google trend.`,
	Example: `
  news trend -c "hk"
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		country = strings.ToUpper(country)
		res, err := command.Gtrend(country)
		if err != nil {
			return err
		}

		// Print result
		rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Topic", "Traffic", "Title"})
		for _, r := range res {
			t.AppendRow(table.Row{r.Topic, r.Traffic, r.Title}, rowConfigAutoMerge)
		}
		t.SetColumnConfigs([]table.ColumnConfig{
			{Number: 1, AutoMerge: true},
			{Number: 2, AutoMerge: true},
			{Number: 3, WidthMax: 80},
		})
		// t.SetColumnConfigs([]table.ColumnConfig{
		//  {Number: 1, WidthMax: 64},
		// })
		t.AppendSeparator()
		t.Style().Options.SeparateRows = true
		t.Render()

		return nil
	},
}

func init() {
	newsCmd.AddCommand(gtrendCmd)
	gtrendCmd.Flags().StringVarP(&country, "country", "c", "", "Country of the trend.")
}
