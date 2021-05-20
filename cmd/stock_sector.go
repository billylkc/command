package cmd

import (
	"fmt"

	"github.com/billylkc/command/command"
	"github.com/spf13/cobra"
)

// sectorCmd represents the hk command
var sectorCmd = &cobra.Command{
	Use:     "sector",
	Aliases: []string{"s"},
	Short:   "[s] Extract sector performance.",
	Long:    `[s] Extract sector performance`,
	Example: `
  command stock sector
  command s sector
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := command.GetSectorOveriew()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(PrettyPrint(res))
		fmt.Println(len(res))
		return nil
	},
}

func init() {
	stockCmd.AddCommand(sectorCmd)
}
