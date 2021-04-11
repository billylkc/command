package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var cheatsheetCmd = &cobra.Command{
	Use:     "cheatsheet",
	Short:   "[c] Getting some cheatsheet from cht.sh",
	Long:    `[c] Getting some cheatsheet from cht.sh`,
	Aliases: []string{"ch"},
	Example: `
  command utility cheatsheet python dataframe
  command u ch python dataframe
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		endpoints := "http://cht.sh"
		lang := args[0]
		args = args[1:]
		commands := fmt.Sprintf("curl %s/%s/%s", endpoints, lang, strings.Join(args, "+"))
		fmt.Println(commands)

		// Execute command
		err, out, _ := shellOut(commands)
		if err != nil {
			log.Printf("error: %v\n", err)
		}

		fmt.Println(out)

		return nil
	},
}

func init() {
	utilityCmd.AddCommand(cheatsheetCmd)
}

// shell output to bash
func shellOut(command string) (error, string, string) {
	const shell = "bash"

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(shell, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}
