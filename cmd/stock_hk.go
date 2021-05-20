package cmd

import (
	"fmt"

	"github.com/billylkc/command/command"
	"github.com/spf13/cobra"
)

// hkStockCmd represents the hkStock command
var hkStockCmd = &cobra.Command{
	Use:     "hkStock",
	Aliases: []string{"hk"},
	Short:   "Get Hong Kong stock price.",
	Long:    `Get Hong Kong stock price.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		quandl := command.NewQuandl()
		// quandl.Option(SetLimit(10))
		// quandl.Option(SetEndDate("2021-01-02"))
		// quandl.Option(SetOrder("desc"))

		result, err := quandl.GetHistoricalPrice(5)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(PrettyPrint(result))
		return nil
	},
}

func init() {
	stockCmd.AddCommand(hkStockCmd)
}
