package cmd

import (
	"fmt"

	"github.com/billylkc/command/command"
	"github.com/billylkc/myutil"

	. "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var all bool = false

// hackernewsCmd represents the hackernews command
var hackernewsCmd = &cobra.Command{
	Use:     "hackernews",
	Aliases: []string{"hn"},
	Short:   "[hn] Hackernews.",
	Long:    `[hn] Hackernews.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		type Article struct {
			Date    string
			Content string
		}
		var articles []Article

		res, err := command.GetHackerNews(all)
		if err != nil {
			return err
		}

		// Formatting
		for _, r := range res {
			var a Article
			a.Date = r.Date
			title := Sprintf(Yellow(r.Title))
			link := Sprintf(Green(fmt.Sprintf("%s", r.Link)))
			abstract := myutil.BreakLongParagraph(r.Abstract, 100, 0)

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
	newsCmd.AddCommand(hackernewsCmd)
	hackernewsCmd.Flags().BoolVarP(&all, "all", "a", false, "All news. Default false.")
}
