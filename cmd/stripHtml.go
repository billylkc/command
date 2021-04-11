package cmd

import (
	"bufio"
	"fmt"
	"os"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/spf13/cobra"
)

// stripHtmlCmd represents the stripHtml command
var stripHtmlCmd = &cobra.Command{
	Use:     "stripHtml",
	Aliases: []string{"shtml"},
	Short:   "[sh] Strip HTML.",
	Long:    `[sh] Strip line numberHTML.`,
	Example: `
  command utility stripHtml
  command u shtml
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Reading text - Until . ")
		scanner := bufio.NewScanner(os.Stdin)
		res := ""
		for scanner.Scan() {
			line := scanner.Text()
			if line == "." {
				break
			}
			res += line + "\n"
		}

		fmt.Println("============================")
		fmt.Printf("\nOutput:\n")
		fmt.Println("============================")
		fmt.Println(strip.StripTags(res))

		return nil
	},
}

func init() {
	utilityCmd.AddCommand(stripHtmlCmd)
}
