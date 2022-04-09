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
	mIn  string // mermaid input file, e.g. input.mmd (mermaid)
	mOut string // mermaid output file e.g. result.html
)

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
		b, err := ioutil.ReadFile(mIn)
		if err != nil {
			return err
		}

		// output
		f, err := os.OpenFile(mOut, os.O_WRONLY|os.O_CREATE, 0644)
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
		fmt.Printf("Output Mermaid file - %s\n", mOut)

		return nil
	},
}

func init() {
	utilityCmd.AddCommand(mermaidCmd)
	mermaidCmd.Flags().StringVarP(&mIn, "input", "i", "input.mmd", "Input file in mermaid syntax format")
	mermaidCmd.Flags().StringVarP(&mOut, "output", "o", "result.html", "Output file in output format")
}
