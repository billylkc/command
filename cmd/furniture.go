/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/billylkc/command/command"
	"github.com/spf13/cobra"
)

var (
	fIn  string // furniture input file, e.g. input.toml
	fOut string // furniture output file e.g. result.html
)

// furnitureCmd represents the furniture command
var furnitureCmd = &cobra.Command{
	Use:     "furniture",
	Short:   "[f] Furniture",
	Long:    "[f] Furniture",
	Aliases: []string{"f"},
	Example: `
  command utility furniture --input furniture.toml --output result.html
  command u f -i output/furniture.toml -o output/furniture.html
`,

	RunE: func(cmd *cobra.Command, args []string) error {

		// Read file content from input file
		b, err := ioutil.ReadFile(fIn)
		if err != nil {
			return err
		}

		// output
		f, err := os.OpenFile(fOut, os.O_WRONLY|os.O_CREATE, 0644)
		defer f.Close()
		if err != nil {
			return err
		}
		w := bufio.NewWriter(f)

		err = command.ParseFurniture(b, w)
		if err != nil {
			return err
		}

		fmt.Printf("Output result file - %s\n", fOut)

		return nil
	},
}

func init() {
	utilityCmd.AddCommand(furnitureCmd)
	furnitureCmd.Flags().StringVarP(&fIn, "input", "i", "input.toml", "Input file in mermaid syntax format")
	furnitureCmd.Flags().StringVarP(&fOut, "output", "o", "result.html", "Output file in output format")
}
