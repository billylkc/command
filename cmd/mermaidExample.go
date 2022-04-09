/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/billylkc/command/command"
	"github.com/spf13/cobra"
)

// mermaidExampleCmd represents the example command
var mermaidExampleCmd = &cobra.Command{
	Use:     "example",
	Short:   "[e] Display useful mermaid syntax for different types.",
	Long:    "[e] Display useful mermaid syntax for different types.",
	Aliases: []string{"e"},
	RunE: func(cmd *cobra.Command, args []string) error {
		_ = command.PrintMermaidExample()
		return nil
	},
}

func init() {
	mermaidCmd.AddCommand(mermaidExampleCmd)

}
