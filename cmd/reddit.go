package cmd

import (
	"fmt"
	"strings"

	"github.com/billylkc/command/command"
	"github.com/billylkc/myutil"
	"github.com/spf13/cobra"
)

// redditCmd represents the reddit command
var redditCmd = &cobra.Command{
	Use:     "reddit [subreddit]",
	Aliases: []string{"r"},
	Short:   "[r] Reddit.",
	Long:    `[r] Reddit.`,
	Example: `
  command news reddit wallstreetbets
  command n r wallstreetbets
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var articles []Article
		var s string
		if len(args) > 0 {
			s = strings.Join(args, " ")
		} else {
			return fmt.Errorf("not enough argumetns\n")
		}
		res, err := command.Reddit(s)
		if err != nil {
			return err
		}

		// Formatting
		for _, r := range res {
			var a Article
			a.Date = r.Updated.Format("2006-01-02")
			title := myutil.TextYellow(myutil.BreakLongStr(r.Title, 100, 0))
			link := myutil.TextGreen(fmt.Sprintf("%s", r.Link))
			abstract := myutil.BreakLongParagraph(r.Content, 100, 0)

			a.Content = fmt.Sprintf("%s\n\n%s\n\n%s", title, link, abstract)
			articles = append(articles, a)
		}

		headers := []string{"Date", "Content"}
		ignores := []string{""}
		data := myutil.InterfaceSlice(articles)
		err = myutil.PrintTable(data, headers, ignores, 1)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	newsCmd.AddCommand(redditCmd)
}
