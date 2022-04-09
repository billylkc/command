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

var inputFile string
var outputFile string

// mermaidCmd represents the mermaid command
var mermaidCmd = &cobra.Command{
	Use:     "mermaid",
	Short:   "[m] Create mermaid html file with the input mermaid syntax.",
	Long:    `[m] Create mermaid html file with the input mermaid syntax.`,
	Aliases: []string{"m"},
	Example: `
  command utility mermaid --input input.mmd --output result.html
  command u m -i input.mmd -o result.html
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// Read file content from input file
		b, err := ioutil.ReadFile(inputFile)
		if err != nil {
			return err
		}

		// output
		f, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE, 0644)
		defer f.Close()
		if err != nil {
			return err
		}
		w := bufio.NewWriter(f)

		// Parse file
		err = command.ParseMermaid(b, w)
		if err != nil {
			return err
		}
		fmt.Printf("Output Mermaid file - %s\n", outputFile)

		return nil
	},
}

func init() {
	utilityCmd.AddCommand(mermaidCmd)
	mermaidCmd.Flags().StringVarP(&inputFile, "input", "i", "input.mmd", "Input file in mermaid syntax format")
	mermaidCmd.Flags().StringVarP(&outputFile, "output", "o", "result.html", "Output file in output format")
}
