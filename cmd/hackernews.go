package cmd

import (
	"fmt"

	"github.com/billylkc/command/command"
	"github.com/billylkc/myutil"

	"github.com/spf13/cobra"
)

var all bool = false

// hackernewsCmd represents the hackernews command
var hackernewsCmd = &cobra.Command{
	Use:     "hackernews",
	Aliases: []string{"hn"},
	Short:   "[hn] Hackernews.",
	Long:    `[hn] Hackernews.`,
	Example: `
  command news hackernews
  command n hn
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var articles Articles

		res, err := command.GetHackerNews(all)
		if err != nil {
			return err
		}

		// Formatting
		for _, r := range res {
			var a Article
			a.Date = r.Updated.Format("2006-01-02")
			title := myutil.TextYellow(r.Title)
			link := myutil.TextGreen(fmt.Sprintf("%s", r.Link))
			abstract := myutil.BreakLongParagraph(r.Content, 100, 0)

			a.Content = fmt.Sprintf("%s\n\n%s\n\n%s", title, link, abstract)
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
	newsCmd.AddCommand(hackernewsCmd)
	hackernewsCmd.Flags().BoolVarP(&all, "all", "a", false, "All news. Default false.")
}
