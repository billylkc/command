package cmd

import (
	"fmt"

	"github.com/billylkc/command/command"
	"github.com/spf13/cobra"
)

var twitterCmd = &cobra.Command{
	Use:     "twitter [hashtags]",
	Aliases: []string{"t"},
	Short:   "[t] twitter.",
	Long:    `[t] twitter.`,
	Example: `
  command news twitter donaldtrump
  command n t donaldtrump
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("No Hashtags \n")
		}
		tags := getHashTags(args...)
		command.ListenTweets(tags...)
		return nil
	},
}

func init() {
	newsCmd.AddCommand(twitterCmd)
}

// convert args with hashtags
func getHashTags(args ...string) []string {
	var tags []string
	for _, t := range args {
		t = fmt.Sprintf("#%s", t)
		tags = append(tags, t)
	}
	return tags
}
