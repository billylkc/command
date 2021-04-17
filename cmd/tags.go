package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/billylkc/myutil"
	"github.com/spf13/cobra"
)

// tagsCmd represents the tags command
var gitTagCmd = &cobra.Command{
	Use:     "tag",
	Aliases: []string{"gt"},
	Short:   "[gt] Git tag next minor version.",
	Long:    `[gt] Git tag next minor version.`,
	Example: `
  command utility tag
  command u gt
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err, out, _ := myutil.Shell("git status")
		if err != nil {
			return err
		}

		fmt.Printf("out: %+v\n", out)
		b := strings.Contains(out, "nothing to commit, working tree clean")
		if !b { // Need commit
			fmt.Errorf("Need to commit first. \n")
			os.Exit(0)
		}
		fmt.Println("Continue")

		return nil
	},
}

func init() {
	utilityCmd.AddCommand(gitTagCmd)
}
