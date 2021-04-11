package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// stripLineNumCmd represents the stripLineNum command
var stripLineNumCmd = &cobra.Command{
	Use:     "stripLineNum",
	Aliases: []string{"sl"},
	Short:   "[sl] Strip line number.",
	Long:    `[sl] Strip line number.`,
	Example: `
  command utility stripLineNum
  command u sl
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		scanner := bufio.NewScanner(os.Stdin)

		// create a map of string with number
		m := make(map[string]bool)
		m = map[string]bool{
			"0": true,
			"1": true,
			"2": true,
			"3": true,
			"4": true,
			"5": true,
			"6": true,
			"7": true,
			"8": true,
			"9": true,
		}

		fmt.Println("Reading text - Until . ")
		var results []string
		for scanner.Scan() {
			line := scanner.Text()
			if line == "." {
				break
			}

			// only those lines not starting with numbers
			if len(line) > 0 && line != " " {
				char := string([]rune(line)[0])
				if _, ok := m[char]; !ok {
					if strings.TrimSpace(line) != "" {
						results = append(results, line)
					}
				}
			}
		}

		// Print result
		fmt.Println("============================")
		fmt.Println("Output:")
		fmt.Println("============================")
		for _, res := range results {
			if strings.TrimSpace(res) != "" {
				fmt.Println(res)
			}
		}
		return nil
	},
}

func init() {
	utilityCmd.AddCommand(stripLineNumCmd)
}
