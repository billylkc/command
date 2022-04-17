/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/billylkc/command/command"
	"github.com/spf13/cobra"
)

var (
	fgOut     string // furniture out file, e.g. input.toml
	fgImpOnly bool   // Important furniture only, default false
)

// furnitureGenerateCmd represents the furnitureGenerate command
var furnitureGenerateCmd = &cobra.Command{
	Use:     "furnitureGenerate",
	Short:   "[g] Generate sample input file for furniture command.",
	Long:    "[g] Generate sample input file for furniture command.",
	Aliases: []string{"g"},
	Example: `
  command utility furniture generate --output furniture.toml
  command u f g -o furniture.toml
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// output
		f, err := os.OpenFile(fgOut, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		defer f.Close()
		if err != nil {
			return err
		}
		w := bufio.NewWriter(f)

		err = command.GenerateFurnitureExample(w, fgImpOnly)
		if err != nil {
			return err
		}
		fmt.Printf("Output funiture toml file - %s\n", fgOut)
		return nil
	},
}

func init() {
	furnitureCmd.AddCommand(furnitureGenerateCmd)
	furnitureGenerateCmd.Flags().StringVarP(&fgOut, "output", "o", "furniture.toml", "Output file in toml format")
	furnitureGenerateCmd.Flags().BoolVarP(&fgImpOnly, "impFlag", "i", false, "Important items only")
}
