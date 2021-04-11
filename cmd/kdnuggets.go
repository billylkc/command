package cmd

import (
	"fmt"
	"time"

	"github.com/billylkc/command/command"
	"github.com/billylkc/myutil"
	"github.com/spf13/cobra"
)

// kdnuggetsCmd represents the kdnuggets command
var kdnuggetsCmd = &cobra.Command{
	Use:     "kdnuggets",
	Aliases: []string{"kd"},
	Short:   "[kd] KDnuggets",
	Long:    `[kd] KDnuggets`,
	Example: `
  command news kdnuggets
  command n kd
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var articles Articles

		err := myutil.HandleDateArgs(&date, &nrecords, 1, args...)
		if err != nil {
			return err
		}

		d, err := myutil.ParseDateInput(date, "m")
		if err != nil {
			return err
		}

		res, err := command.KDNuggets(d, nrecords)
		if err != nil {
			return err
		}

		// Formatting
		for _, r := range res {
			var a Article
			a.Date = r.Updated.Format("2006-01-02")
			title := myutil.BreakLongStr(r.Title, 60, 80)
			title = myutil.TextYellow(myutil.HandleCamalCase(title))
			link := myutil.TextGreen(fmt.Sprintf("%s", r.Link))
			tags := myutil.TextMagenta(fmt.Sprintf("[%s]", r.Tags))
			abstract := myutil.BreakLongParagraph(r.Content, 100, 0)

			a.Content = fmt.Sprintf("%s\n\n%s\n%s\n\n%s", title, link, tags, abstract)
			articles = append(articles, a)
		}
		err = articles.PrintTable()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	today := time.Now().Format("2006-01-02")
	newsCmd.AddCommand(kdnuggetsCmd)
	kdnuggetsCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
